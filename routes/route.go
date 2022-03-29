package routes

import (
	"database/sql"
	"simple-upload-file/controllers"
)

func InitializeRoutes(db *sql.DB) {
	// Membuat sebuah group router
	// routerGroup := r.Group("/api/v1")
	controllers.NewUserController(db)

}
