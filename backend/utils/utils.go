package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// =========================
// WRITE JSON RESPONSE
// =========================

func WriteJSON(
	w http.ResponseWriter,
	status int,
	v any,
) error {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		status,
	)

	return json.NewEncoder(
		w,
	).Encode(v)
}

// =========================
// WRITE ERROR RESPONSE
// =========================

func WriteError(
	w http.ResponseWriter,
	status int,
	err error,
) {

	WriteJSON(
		w,
		status,
		map[string]string{
			"error": err.Error(),
		},
	)
}

// =========================
// PARSE JSON REQUEST
// =========================

func ParseJSON(
	r *http.Request,
	v any,
) error {

	if r.Body == nil {

		return fmt.Errorf(
			"missing request body",
		)
	}

	return json.NewDecoder(
		r.Body,
	).Decode(v)
}

// =========================
// GET JWT TOKEN
// =========================

func GetTokenFromRequest(
	r *http.Request,
) string {

	// =========================
	// AUTH HEADER
	// =========================

	tokenAuth := r.Header.Get(
		"Authorization",
	)

	// Bearer <token>
	if tokenAuth != "" {

		splitToken := strings.Split(
			tokenAuth,
			"Bearer ",
		)

		if len(splitToken) == 2 {

			return strings.TrimSpace(
				splitToken[1],
			)
		}

		return tokenAuth
	}

	// =========================
	// QUERY PARAM
	// =========================

	tokenQuery := r.URL.Query().Get(
		"token",
	)

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}