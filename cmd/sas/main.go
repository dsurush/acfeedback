package main

import (
	"acfeedback/cmd/sas/app"
	"acfeedback/pkg/core/services"
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9998", "Server port")
	dsn  = flag.String("dsn", "postgres://username:password@localhost:5501/app", "Postgres DSN")
)

func main() {
	fmt.Print()
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	router := httprouter.New()
	pool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		log.Printf("%e", err)
	}
	svc := services.NewFeedbackSvc(pool)
	server := app.NewMainServer(pool, router, svc)
	server.Start()
	panic(http.ListenAndServe(addr, server))
}
