package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Abunyawa/back_game/service"
	xhttp "github.com/Abunyawa/back_game/xtransport/http"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	PORT     string
	SIGN_KEY string
	DB_URL   string
)

func readEnv() {
	godotenv.Load()
	DB_URL = os.Getenv("DATABASE_URL")
	PORT = os.Getenv("PORT")
	SIGN_KEY = os.Getenv("SIGN_KEY")
}

func main() {
	readEnv()

	db, err := sql.Open("pgx", DB_URL)
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
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", PORT), h)
	}()

	fmt.Println("exiting", <-errs)
}
