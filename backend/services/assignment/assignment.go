package assignment

import (
	"database/sql"
	"fmt"
	"net/http"

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

	router.HandleFunc(
		"/assignments",
		h.handleCreateAssignment,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/assignments",
		h.handleGetAssignments,
	).Methods(http.MethodGet)
}

// =========================
// CREATE ASSIGNMENT
// =========================

func (h *Handler) handleCreateAssignment(
	w http.ResponseWriter,
	r *http.Request,
) {

	var payload types.CreateAssignmentPayload

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
	INSERT INTO assignments
	(title, description, due_date)
	VALUES (?, ?, ?)
	`

	result, err := h.store.db.Exec(
		query,
		payload.Title,
		payload.Description,
		payload.DueDate,
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
			"message": "assignment created successfully",
			"id":      id,
		},
	)
}

// =========================
// GET ASSIGNMENTS
// =========================

func (h *Handler) handleGetAssignments(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.store.db.Query(
		"SELECT * FROM assignments",
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

	var assignments []types.Assignment

	for rows.Next() {

		var assignment types.Assignment

		err := rows.Scan(
			&assignment.ID,
			&assignment.Title,
			&assignment.Description,
			&assignment.DueDate,
			&assignment.CreatedAt,
		)

		if err != nil {

			utils.WriteError(
				w,
				http.StatusInternalServerError,
				err,
			)

			return
		}

		assignments = append(
			assignments,
			assignment,
		)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		assignments,
	)
}