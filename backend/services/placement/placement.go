package placement

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
		"/placements",
		auth.AdminOnly(
			h.handleCreatePlacement,
		),
	).Methods(http.MethodPost)

	// =========================
	// GET PLACEMENTS
	// =========================

	router.HandleFunc(
		"/placements",
		h.handleGetPlacements,
	).Methods(http.MethodGet)
}

// =========================
// CREATE PLACEMENT
// =========================

func (h *Handler) handleCreatePlacement(
	w http.ResponseWriter,
	r *http.Request,
) {

	var payload types.CreatePlacementPayload

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
	INSERT INTO placements
	(student_id, company_name, role, lpa)
	VALUES (?, ?, ?, ?)
	`

	result, err := h.store.db.Exec(
		query,
		payload.StudentID,
		payload.CompanyName,
		payload.Role,
		payload.LPA,
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
			"message": "placement created successfully",
			"id":      id,
		},
	)
}

// =========================
// GET PLACEMENTS
// =========================

func (h *Handler) handleGetPlacements(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.store.db.Query(
		"SELECT * FROM placements",
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

	var placements []types.Placement

	for rows.Next() {

		var placement types.Placement

		err := rows.Scan(
			&placement.ID,
			&placement.StudentID,
			&placement.CompanyName,
			&placement.Role,
			&placement.LPA,
			&placement.CreatedAt,
		)

		if err != nil {

			utils.WriteError(
				w,
				http.StatusInternalServerError,
				err,
			)

			return
		}

		placements = append(
			placements,
			placement,
		)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		placements,
	)
}