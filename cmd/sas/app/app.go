package app

import (
	"acfeedback/pkg/core/services"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MainServer struct {
	pool        *pgxpool.Pool
	router      *httprouter.Router
	feedbackSvc *services.FeedbackSvc
}

func NewMainServer(pool *pgxpool.Pool, router *httprouter.Router, feedbackSvc *services.FeedbackSvc) *MainServer {
	return &MainServer{pool: pool, router: router, feedbackSvc: feedbackSvc}
}



func (server *MainServer) Start() {
	err := server.feedbackSvc.DbInit()
	if err != nil {
		panic("server don't created")
	}
	//feedback := models.FeedBack{
	//	Id:           0,
	//	Feedback:     "bla bla",
	//	Who:          1,
	//	Whom:         2,
	//	FeedbackTime: "",
	//	Remove:       false,
	//}
	//err = server.feedbackSvc.Save(feedback)
	//if err != nil {
	//	log.Fatalf("wrong 39/app, %e", err)
	//}
	server.InitRouts()
}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// delegation////
	server.router.ServeHTTP(writer, request)
}

