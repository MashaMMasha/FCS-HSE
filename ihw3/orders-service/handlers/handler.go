package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"orders-service/pkg/json"
	service "orders-service/services"
)

type Handler struct {
	s service.OrderServicing
	l *log.Logger
}

func NewHandler(s service.OrderServicing, l *log.Logger) *Handler {
	return &Handler{s: s, l: l}
}

func (h *Handler) AllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.s.All()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.ToJSON(orders, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Title Get Order By ID
// @Description Получить заказ по ID
// @Tags Order Info
// @Produce json
// @Param   id  path  string  true  "ID заказа"
// @Success 200 {object} models.Order "Информация о заказе"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /get/{id} [get]
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	order, err := h.s.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.ToJSON(order, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type createReq struct {
	Amount float64 `json:"amount" example:"5.00"`
	Descr  string  `json:"descr" example:"Coca-cola"`
} //@name CreateOrderRequest

// @Title Create Order
// @Description Оформляет заказ
// @Tags Order Manage
// @Param   user_id  path  string  true  "ID пользователя"
// @Param   order     body  createReq true  "Детали заказа"
// @Success 200
// @Failure 400
// @Failure 500
// @Router  /create/{user_id} [post]
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]

	req := &createReq{}
	err := json.FromJSON(req, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID, err := h.s.Add(userID, req.Amount, req.Descr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Account created successfully", "id": "` + orderID + `"}`))
}
