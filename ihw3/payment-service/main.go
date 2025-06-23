package main

import (
	"context"
	"github.com/gorilla/mux"
	swag "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	_ "payment-service/docs"
	"payment-service/handlers"
	"payment-service/infrastructure/kafka"
	"payment-service/infrastructure/storage"
	"payment-service/pkg/postgres"
	"payment-service/services"
	"time"
)

// @title payment-service
// @version 1.0

// @host localhost:8002
// @BasePath /payment/
func main() {
	lg := log.New(os.Stdout, "payment-service ", log.LstdFlags)

	lg.Println("starting db...")
	db, err := postgres.Init()
	if err != nil {
		lg.Fatal(err)
	}
	accounts, err := storage.NewAccountDB(db)
	if err != nil {
		lg.Fatal(err)
	}

	lg.Println("starting inbox...")
	inbox, err := kafka.NewInbox(db)
	if err != nil {
		lg.Fatal(err)
	}

	lg.Println("starting outbox...")
	outbox, err := kafka.NewOutbox(db)
	if err != nil {
		lg.Fatal(err)
	}
	lg.Println("starting manager...")
	manager := storage.NewDBManager(db)

	lg.Println("starting broker...")
	broker := kafka.NewKafka()
	defer broker.Close()

	lg.Println("starting workers...")
	services.NewInboxWorker(broker, inbox, lg).Start(context.Background(), 2*time.Second)
	services.NewOutboxWorker(broker, outbox, lg).Start(context.Background(), 2*time.Second)
	services.NewPaymentWorker(accounts, inbox, outbox, manager, lg).StartPaying(context.Background(), 2*time.Second)

	s := services.NewAccountService(accounts, lg)

	lg.Println("tuning server...")

	h := handlers.NewHandler(s, lg)
	r := mux.NewRouter()

	sr := r.PathPrefix("/payment/account/").Subrouter()
	sr.Methods(http.MethodPost).Path("/create").HandlerFunc(h.CreateAccount)
	sr.Methods(http.MethodGet).Path("/get/{user_id}").HandlerFunc(h.GetAccount)
	sr.Methods(http.MethodGet).Path("/get").HandlerFunc(h.AllAccounts)
	sr.Methods(http.MethodPatch).Path("/update/{user_id}").HandlerFunc(h.UpdateBalance)

	r.PathPrefix("/docs/").Handler(swag.WrapHandler)

	log.Println("Сервер запущен на :8002")
	log.Fatal(http.ListenAndServe(":8002", r))

}
