package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"simple-upload-file/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Server struct {
	db *sql.DB
}

func InitServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

// Connection Database mysql
func ConnectDB() (*sql.DB, error) {
	dbEngine := viper.GetString("database.engine")
	dbUser := viper.GetString("database.user")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbPass := viper.GetString("database.pass")
	dbSchema := viper.GetString("database.schema")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbSchema)
	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := sql.Open(dbEngine, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Print(err)
		_, _ = fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	return db, err
}

func StartRouter() error {
	host := viper.GetString("host")
	port := viper.GetString("port")
	fmt.Printf("Server has been running on %s port: %s\n", host, port)
	address := fmt.Sprintf("%s:%s", host, port)
	log.Fatal(http.ListenAndServe(address, nil))
	return nil
}

func Run() {
	// Connect DB
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Create Router
	// router := CreateRouter()

	InitServer(db)

	// Routes
	routes.InitializeRoutes(db)

	// Run Router / Server
	errRun := StartRouter()
	if errRun != nil {
		log.Fatal(errRun.Error())
	}

}
