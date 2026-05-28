package query

import (
	"database/sql"
	"fmt"
	"net/http"

	"backend/services/auth"
	"backend/types"
	"backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// =========================
// STORE
// =========================

type Store struct {
	db *sql.DB
}

func NewStore(
	db *sql.DB,
) *Store {

	return &Store{
		db: db,
	}
}

// =========================
// HANDLER
// =========================

type Handler struct {
	store *Store
}

func NewHandler(
	store *Store,
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
	// STUDENT CREATE QUERY
	// =========================

	router.HandleFunc(
		"/queries",
		h.handleCreateQuery,
	).Methods(http.MethodPost)

	// =========================
	// ADMIN VIEW QUERIES
	// =========================

	router.HandleFunc(
		"/queries",
		auth.AdminOnly(
			h.handleGetQueries,
		),
	).Methods(http.MethodGet)
}

// =========================
// CREATE QUERY
// =========================

func (h *Handler) handleCreateQuery(
	w http.ResponseWriter,
	r *http.Request,
) {

	var payload types.RaiseQueryPayload

	if err := utils.ParseJSON(
		r,
		&payload,
	); err != nil {

		utils.WriteError(
			w,
			http.StatusBadRequest,
			err,
		)

		return
	}

	if err := utils.Validate.Struct(
		payload,
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

	query := `
	INSERT INTO queries
	(student_id, subject, message, status)
	VALUES (?, ?, ?, ?)
	`

	result, err := h.store.db.Exec(
		query,
		payload.StudentID,
		payload.Subject,
		payload.Message,
		"pending",
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	id, _ := result.LastInsertId()

	utils.WriteJSON(
		w,
		http.StatusCreated,
		map[string]interface{}{
			"message": "query raised successfully",
			"id":      id,
		},
	)
}

// =========================
// GET QUERIES
// =========================

func (h *Handler) handleGetQueries(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.store.db.Query(
		"SELECT * FROM queries",
	)

	if err != nil {

		utils.WriteError(
			w,
			http.StatusInternalServerError,
			err,
		)

		return
	}

	defer rows.Close()

	var queries []types.Query

	for rows.Next() {

		var q types.Query

		err := rows.Scan(
			&q.ID,
			&q.StudentID,
			&q.Subject,
			&q.Message,
			&q.Status,
			&q.CreatedAt,
		)

		if err != nil {

			utils.WriteError(
				w,
				http.StatusInternalServerError,
				err,
			)

			return
		}

		queries = append(
			queries,
			q,
		)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		queries,
	)
}