package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Abunyawa/back_auth/service"
	xhttp "github.com/Abunyawa/back_auth/xtransport/http"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	DB_NAME  string
	DB_USER  string
	DB_PASS  string
	DB_HOST  string
	DB_PORT  string
	SIGN_KEY string
)

func readEnv() {
	godotenv.Load()
	DB_NAME = os.Getenv("POSTGRES_DATABASE")
	DB_USER = os.Getenv("POSTGRES_USER")
	DB_PASS = os.Getenv("POSTGRES_PASS")
	DB_HOST = os.Getenv("POSTGRES_HOST")
	DB_PORT = os.Getenv("POSTGRES_PORT")
	SIGN_KEY = os.Getenv("SIGN_KEY")
}

func databaseUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}
func main() {
	readEnv()

	db, err := sql.Open("pgx", databaseUrl())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var s service.Service
	{
		s = service.NewService(db, SIGN_KEY)
	}

	var h http.Handler
	{
		h = xhttp.MakeHTTPHandler(s)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Println("Listening")
		errs <- http.ListenAndServe(":8080", h)
	}()

	fmt.Println("exiting", <-errs)
}
