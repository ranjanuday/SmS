package user

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/configs"
	"backend/services/auth"
	"backend/types"
	"backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.SMSStore
}

func NewHandler(
	store types.SMSStore,
) *Handler {

	return &Handler{
		store: store,
	}
}

// =========================
// REGISTER ROUTES
// =========================

func (h *Handler) RegisterRoutes(
	router *mux.Router,
) {

	// =========================
	// AUTH ROUTES
	// =========================

	router.HandleFunc(
		"/login",
		h.handleLogin,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/register",
		h.handleRegister,
	).Methods(http.MethodPost)

	// =========================
	// USER ROUTES
	// =========================

	router.HandleFunc(
		"/users/{userID}",
		auth.WithJWTAuth(
			h.handleGetUser,
			h.store,
		),
	).Methods(http.MethodGet)
}

// =========================
// LOGIN
// =========================

func (h *Handler) handleLogin(
	w http.ResponseWriter,
	r *http.Request,
) {

	var user types.LoginUserPayload

	if err := utils.ParseJSON(
		r,
		&user,
	); err != nil {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)

		return
	}

	if err := utils.Validate.Struct(
		user,
	); err != nil {

		errors := err.(validator.ValidationErrors)

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"invalid payload: %v",
				errors,
			),
		)

		return
	}

	u, err := h.store.GetUserByEmail(
		user.Email,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"invalid email or password",
			),
		)

		return
	}

	if !auth.ComparePasswords(
		u.Password,
		[]byte(user.Password),
	) {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"invalid email or password",
			),
		)

		return
	}

	secret := []byte(
		configs.Envs.JWTSecret,
	)

	token, err := auth.CreateJWT(
		secret,
		u.ID,
		u.Role,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]interface{}{
			"message": "login successful",
			"token":  token,
			"role":   u.Role,
			"user":   u,
		},
	)
}

// =========================
// REGISTER
// =========================

func (h *Handler) handleRegister(
	w http.ResponseWriter,
	r *http.Request,
) {

	var user types.RegisterUserPayload

	if err := utils.ParseJSON(
		r,
		&user,
	); err != nil {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)

		return
	}

	if err := utils.Validate.Struct(
		user,
	); err != nil {

		errors := err.(validator.ValidationErrors)

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"invalid payload: %v",
				errors,
			),
		)

		return
	}

	// =========================
	// CHECK EXISTING USER
	// =========================

	_, err := h.store.GetUserByEmail(
		user.Email,
	)

	if err == nil {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"user already exists",
			),
		)

		return
	}

	// =========================
	// HASH PASSWORD
	// =========================

	hashedPassword, err := auth.HashPassword(
		user.Password,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	// =========================
	// CREATE USER
	// =========================

	newUser := types.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
		Role:     user.Role,
	}

	err = h.store.CreateUser(
		newUser,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	// =========================
	// GET CREATED USER
	// =========================

	u, err := h.store.GetUserByEmail(
		user.Email,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	// =========================
	// CREATE JWT
	// =========================

	secret := []byte(
		configs.Envs.JWTSecret,
	)

	token, err := auth.CreateJWT(
		secret,
		u.ID,
		u.Role,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	// =========================
	// RESPONSE
	// =========================

	utils.WriteJSON(
		w,
		http.StatusCreated,
		map[string]interface{}{
			"message": "user registered successfully",
			"token":   token,
			"role":    u.Role,
			"user":    u,
		},
	)
}

// =========================
// GET USER
// =========================

func (h *Handler) handleGetUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	vars := mux.Vars(r)

	str, ok := vars["userID"]

	if !ok {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"missing user ID",
			),
		)

		return
	}

	userID, err := strconv.Atoi(
		str,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Errorf(
				"invalid user ID",
			),
		)

		return
	}

	user, err := h.store.GetUserByID(
		userID,
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		user,
	)
}