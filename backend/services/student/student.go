package student

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
		"/students",
		auth.AdminOnly(
			h.handleCreateStudent,
		),
	).Methods(http.MethodPost)

	// =========================
	// GET STUDENTS
	// =========================

	router.HandleFunc(
		"/students",
		h.handleGetStudents,
	).Methods(http.MethodGet)
}

// =========================
// CREATE STUDENT
// =========================

func (h *Handler) handleCreateStudent(
	w http.ResponseWriter,
	r *http.Request,
) {

	var payload types.CreateStudentPayload

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
	INSERT INTO students
	(student_id, roll, name, department)
	VALUES (?, ?, ?, ?)
	`

	result, err := h.store.db.Exec(
		query,
		payload.StudentID,
		payload.Roll,
		payload.Name,
		payload.Department,
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
			"message": "student created successfully",
			"id":      id,
		},
	)
}

// =========================
// GET STUDENTS
// =========================

func (h *Handler) handleGetStudents(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.store.db.Query(
		"SELECT * FROM students",
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

	var students []types.Student

	for rows.Next() {

		var student types.Student

		err := rows.Scan(
			&student.ID,
			&student.UserID,
			&student.StudentID,
			&student.Roll,
			&student.Name,
			&student.Department,
			&student.CreatedAt,
		)

		if err != nil {

			utils.WriteError(
				w,
				http.StatusInternalServerError,
				err,
			)

			return
		}

		students = append(
			students,
			student,
		)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		students,
	)
}