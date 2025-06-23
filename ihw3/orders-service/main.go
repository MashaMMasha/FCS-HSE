package main

import (
	"context"
	"github.com/gorilla/mux"
	swag "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "orders-service/docs"
	"orders-service/handlers"
	"orders-service/infrastructure/kafka"
	"orders-service/infrastructure/storage"
	"orders-service/pkg/postgres"
	"orders-service/services"
	"os"
	"time"
)

// @title Orders
// @version 1.0

// @host localhost:8001
// @BasePath /order/
func main() {
	lg := log.New(os.Stdout, "Orders ", log.LstdFlags)

	db, err := postgres.Init()
	if err != nil {
		lg.Fatal(err)
	}

	orders, err := storage.NewOrderDB(db)
	outbox, err := kafka.NewOutbox(db)
	trxManager := storage.NewDBManager(db)

	if err != nil {
		lg.Fatal(err)
	}

	broker := kafka.NewKafka()
	defer broker.Close()

	service.NewOutboxWorker(broker, outbox, lg).Start(context.Background(), 2*time.Second)
	service.NewStatusWorker(broker, orders, lg).Start(context.Background(), 2*time.Second)

	s := service.NewOrderService(orders, outbox, trxManager)

	h := handlers.NewHandler(s, lg)
	r := mux.NewRouter()

	sr := r.PathPrefix("/order/").Subrouter()

	sr.Methods(http.MethodPost).Path("/create/{user_id}").HandlerFunc(h.CreateOrder)
	sr.Methods(http.MethodGet).Path("/get/{id}").HandlerFunc(h.GetOrder)
	sr.Methods(http.MethodGet).Path("/get").HandlerFunc(h.AllOrders)

	r.PathPrefix("/docs/").Handler(swag.WrapHandler)

	log.Println("Сервер запущен на :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
