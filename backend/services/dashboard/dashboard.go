package dashboard

import (
	"database/sql"
	"net/http"

	"backend/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(
	db *sql.DB,
) *Handler {

	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(
	router *mux.Router,
) {

	router.HandleFunc(
		"/dashboard/stats",
		h.handleDashboardStats,
	).Methods(http.MethodGet)
}

// =========================
// DASHBOARD STATS
// =========================

func (h *Handler) handleDashboardStats(
	w http.ResponseWriter,
	r *http.Request,
) {

	var students int
	var marks int
	var assignments int
	var notices int
	var placements int
	var queries int

	h.db.QueryRow(
		"SELECT COUNT(*) FROM students",
	).Scan(&students)

	h.db.QueryRow(
		"SELECT COUNT(*) FROM marks",
	).Scan(&marks)

	h.db.QueryRow(
		"SELECT COUNT(*) FROM assignments",
	).Scan(&assignments)

	h.db.QueryRow(
		"SELECT COUNT(*) FROM notices",
	).Scan(&notices)

	h.db.QueryRow(
		"SELECT COUNT(*) FROM placements",
	).Scan(&placements)

	h.db.QueryRow(
		"SELECT COUNT(*) FROM queries",
	).Scan(&queries)

	utils.WriteJSON(
		w,
		http.StatusOK,
		map[string]int{
			"students":   students,
			"marks":      marks,
			"assignments": assignments,
			"notices":    notices,
			"placements": placements,
			"queries":    queries,
		},
	)
}