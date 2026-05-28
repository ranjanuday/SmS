package notices

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
		"/notices",
		auth.AdminOnly(
			h.handleCreateNotice,
		),
	).Methods(http.MethodPost)

	// =========================
	// GET NOTICES
	// =========================

	router.HandleFunc(
		"/notices",
		h.handleGetNotices,
	).Methods(http.MethodGet)
}

// =========================
// CREATE NOTICE
// =========================

func (h *Handler) handleCreateNotice(
	w http.ResponseWriter,
	r *http.Request,
) {

	var payload types.CreateNoticePayload

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
	INSERT INTO notices
	(title, message)
	VALUES (?, ?)
	`

	result, err := h.store.db.Exec(
		query,
		payload.Title,
		payload.Message,
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
			"message": "notice created successfully",
			"id":      id,
		},
	)
}

// =========================
// GET NOTICES
// =========================

func (h *Handler) handleGetNotices(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := h.store.db.Query(
		"SELECT * FROM notices",
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

	var notices []types.Notice

	for rows.Next() {

		var notice types.Notice

		err := rows.Scan(
			&notice.ID,
			&notice.Title,
			&notice.Message,
			&notice.CreatedAt,
		)

		if err != nil {

			utils.WriteError(
				w,
				http.StatusInternalServerError,
				err,
			)

			return
		}

		notices = append(
			notices,
			notice,
		)
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		notices,
	)
}