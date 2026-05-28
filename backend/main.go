package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"backend/api"
	"backend/configs"
	"backend/db"
)

func main() {

	// MySQL Configuration
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Database Connection
	db, err := db.NewMySQLStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}

	// Check Database Connection
	initStorage(db)

	// Create API Server
	server := api.NewAPIServer(
		fmt.Sprintf(":%s", configs.Envs.Port),
		db,
	)

	// Run Server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// Check Database Connectivity
func initStorage(db *sql.DB) {

	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
