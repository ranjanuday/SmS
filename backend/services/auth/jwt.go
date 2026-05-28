package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"backend/configs"
	"backend/types"
	"backend/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserKey contextKey = "userID"
	RoleKey contextKey = "role"
)

// =========================
// JWT MIDDLEWARE
// =========================

func WithJWTAuth(
	handlerFunc http.HandlerFunc,
	store types.SMSStore,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		tokenString := utils.GetTokenFromRequest(
			r,
		)

		token, err := validateJWT(
			tokenString,
		)

		if err != nil {

			log.Printf(
				"failed to validate token: %v",
				err,
			)

			permissionDenied(w)

			return
		}

		if !token.Valid {

			permissionDenied(w)

			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// =========================
		// USER ID
		// =========================

		str, ok := claims["userID"].(string)

		if !ok {

			permissionDenied(w)

			return
		}

		userID, err := strconv.Atoi(
			str,
		)

		if err != nil {

			permissionDenied(w)

			return
		}

		// =========================
		// ROLE
		// =========================

		role, ok := claims["role"].(string)

		if !ok {

			permissionDenied(w)

			return
		}

		// =========================
		// CHECK USER
		// =========================

		u, err := store.GetUserByID(
			userID,
		)

		if err != nil {

			permissionDenied(w)

			return
		}

		// =========================
		// ADD TO CONTEXT
		// =========================

		ctx := r.Context()

		ctx = context.WithValue(
			ctx,
			UserKey,
			u.ID,
		)

		ctx = context.WithValue(
			ctx,
			RoleKey,
			role,
		)

		r = r.WithContext(
			ctx,
		)

		handlerFunc(
			w,
			r,
		)
	}
}

// =========================
// ADMIN ONLY
// =========================

func AdminOnly(
	handlerFunc http.HandlerFunc,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		role, ok := r.Context().Value(
			RoleKey,
		).(string)

		if !ok {

			permissionDenied(w)

			return
		}

		if role != "admin" {

			utils.WriteError(
				w,
				http.StatusForbidden,
				fmt.Errorf(
					"admin access required",
				),
			)

			return
		}

		handlerFunc(
			w,
			r,
		)
	}
}

// =========================
// CREATE JWT
// =========================

func CreateJWT(
	secret []byte,
	userID int,
	role string,
) (string, error) {

	expiration := time.Second * time.Duration(
		configs.Envs.JWTExpirationInSeconds,
	)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": strconv.Itoa(
				userID,
			),
			"role": role,
			"expiresAt": time.Now().
				Add(expiration).
				Unix(),
		},
	)

	tokenString, err := token.SignedString(
		secret,
	)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// =========================
// VALIDATE JWT
// =========================

func validateJWT(
	tokenString string,
) (*jwt.Token, error) {

	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf(
					"unexpected signing method",
				)
			}

			return []byte(
				configs.Envs.JWTSecret,
			), nil
		},
	)
}

// =========================
// PERMISSION DENIED
// =========================

func permissionDenied(
	w http.ResponseWriter,
) {

	utils.WriteError(
		w,
		http.StatusForbidden,
		fmt.Errorf(
			"permission denied",
		),
	)
}

// =========================
// GET USER ID
// =========================

func GetUserIDFromContext(
	ctx context.Context,
) int {

	userID, ok := ctx.Value(
		UserKey,
	).(int)

	if !ok {
		return -1
	}

	return userID
}

// =========================
// GET ROLE
// =========================

func GetRoleFromContext(
	ctx context.Context,
) string {

	role, ok := ctx.Value(
		RoleKey,
	).(string)

	if !ok {
		return ""
	}

	return role
}