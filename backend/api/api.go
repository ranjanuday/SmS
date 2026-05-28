package api

import (
	"database/sql"
	"log"
	"net/http"

	"backend/services/assignment"
	"backend/services/dashboard"
	"backend/services/marks"
	"backend/services/notices"
	"backend/services/placement"
	"backend/services/query"
	"backend/services/student"
	"backend/services/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// =========================
// CREATE API SERVER
// =========================

func NewAPIServer(
	addr string,
	db *sql.DB,
) *APIServer {

	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// =========================
// RUN SERVER
// =========================

func (s *APIServer) Run() error {

	// =========================
	// MAIN ROUTER
	// =========================

	router := mux.NewRouter()

	// =========================
	// API PREFIX
	// =========================

	subrouter := router.PathPrefix(
		"/sms",
	).Subrouter()

	// =========================
	// USER / AUTH MODULE
	// =========================

	userStore := user.NewStore(
		s.db,
	)

	userHandler := user.NewHandler(
		userStore,
	)

	userHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// STUDENT MODULE
	// =========================

	studentStore := student.NewStore(
		s.db,
	)

	studentHandler := student.NewHandler(
		studentStore,
	)

	studentHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// MARKS MODULE
	// =========================

	marksStore := marks.NewStore(
		s.db,
	)

	marksHandler := marks.NewHandler(
		marksStore,
	)

	marksHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// NOTICE MODULE
	// =========================

	noticeStore := notices.NewStore(
		s.db,
	)

	noticeHandler := notices.NewHandler(
		noticeStore,
	)

	noticeHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// PLACEMENT MODULE
	// =========================

	placementStore := placement.NewStore(
		s.db,
	)

	placementHandler := placement.NewHandler(
		placementStore,
	)

	placementHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// ASSIGNMENT MODULE
	// =========================

	assignmentStore := assignment.NewStore(
		s.db,
	)

	assignmentHandler := assignment.NewHandler(
		assignmentStore,
	)

	assignmentHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// QUERY MODULE
	// =========================

	queryStore := query.NewStore(
		s.db,
	)

	queryHandler := query.NewHandler(
		queryStore,
	)

	queryHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// STATIC FILES
	// =========================

	router.PathPrefix("/").Handler(
		http.FileServer(
			http.Dir("./static"),
		),
	)

	// =========================
	// CORS
	// =========================

	c := cors.New(
		cors.Options{
			AllowedOrigins: []string{
				"http://localhost:5173",
			},

			AllowedMethods: []string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			},

			AllowedHeaders: []string{
				"Content-Type",
				"Authorization",
			},

			AllowCredentials: true,
		},
	)

	handler := c.Handler(
		router,
	)
	dashboardHandler := dashboard.NewHandler(
		s.db,
	)

	dashboardHandler.RegisterRoutes(
		subrouter,
	)

	// =========================
	// START SERVER
	// =========================

	log.Println(
		"Server running on",
		s.addr,
	)

	return http.ListenAndServe(
		s.addr,
		handler,
	)
	// =========================
	// DASHBOARD MODULE
	// =========================

	
}
