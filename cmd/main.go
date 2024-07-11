package main

import (
	"github.com/gorilla/mux"
	"github.com/mao360/notifications/pkg/delivery"
	"github.com/mao360/notifications/pkg/repo"
	"github.com/mao360/notifications/pkg/service"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	conn := "username=postgres host=db port=5432 dbname=postgres password=1234 sslmode=disable"
	pool, err := repo.ConnToDB(conn)
	if err != nil {
		sugar.Fatalf("can`t conn to db: %s", err.Error())
	}

	repository := repo.NewRepo(pool)
	appService := service.NewService(repository)
	h := delivery.NewHandler(appService, sugar)

	r := mux.NewRouter()
	r.HandleFunc("/reg", h.Registration).Methods("POST")
	r.HandleFunc("/auth", h.Authorization).Methods("POST")
	r.HandleFunc("/subscribe", h.Auth(h.Subscribe)).Methods("POST")
	r.HandleFunc("/unsubscribe", h.Auth(h.Unsubscribe)).Methods("DELETE")
	r.HandleFunc("/notification", h.Auth(h.GetNotification)).Methods("GET")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		sugar.Fatalf("can`t start server: %s", err.Error())
	}
}
