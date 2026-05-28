package marks

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
	// ADMIN ONLY
	// =========================

	router.HandleFunc(
		"/marks",
		auth.AdminOnly(
			h.handleAddMarks,
		),
	).Methods(http.MethodPost)

	// =========================
	// GET MARKS
	// =========================

	router.HandleFunc(
		"/marks",
		h.handleGetMarks,
	).Methods(http.MethodGet)
}

// =========================
// ADD MARKS
// =========================

func (h *Handler) handleAddMarks(
	w http.ResponseWriter,
	r *http.Request,
) {

	var payload types.AddMarksPayload

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
	INSERT INTO marks
	(student_id, subject, marks, grade)
	VALUES (?, ?, ?, ?)
	`

	result, err := h.store.db.Exec(
		query,
		payload.StudentID,
		payload.Subject,
		payload.Marks,
		payload.Grade,
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
			"message": "marks added successfully",
			"id":      id,
		},
	)
}

// =========================
// GET MARKS
// =========================

func (h *Handler) handleGetMarks(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.store.db.Query(
		"SELECT * FROM marks",
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

	var marksList []types.Marks

	for rows.Next() {

		var marks types.Marks

		err := rows.Scan(
			&marks.ID,
			&marks.StudentID,
			&marks.Subject,
			&marks.Marks,
			&marks.Grade,
			&marks.CreatedAt,
		)

		if err != nil {

			utils.WriteError(
				w,
				http.StatusInternalServerError,
				err,
			)

			return
		}

		marksList = append(
			marksList,
			marks,
		)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		marksList,
	)
}