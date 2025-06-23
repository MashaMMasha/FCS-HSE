package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"payment-service/pkg/json"
	"payment-service/services"
)

type Handler struct {
	as services.AccountServicing
	lg *log.Logger
}

func NewHandler(as services.AccountServicing, lg *log.Logger) *Handler {
	return &Handler{as: as, lg: lg}
}

// @Title Get all accounts
// @Description Возвращает все счета
// @Tags Accounts Info
// @Produce json
// @Success 200 {array} models.Account
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router  /account/get [get]
func (h *Handler) AllAccounts(w http.ResponseWriter, r *http.Request) {
	h.lg.Println("Fetching all accounts")

	accounts, err := h.as.All()
	if err != nil {
		h.lg.Printf("Error fetching all accounts: %v", err)
		http.Error(w, "Failed to retrieve accounts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.ToJSON(accounts, w); err != nil {
		h.lg.Printf("Error encoding accounts to JSON: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	h.lg.Printf("Successfully returned %d accounts", len(accounts))
}

// @Title Get Account By ID
// @Description Получить счет по ID
// @Tags Accounts Info
// @Produce json
// @Param   user_id  path  string  true  "ID пользователя"
// @Success 200 {object} models.Account "Информация о счете"
// @Failure 400 {string} string "Неверный ID пользователя"
// @Failure 404 {string} string "Счет не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /account/get/{user_id} [get]
func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]

	if userID == "" {
		h.lg.Println("Empty user_id provided")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	h.lg.Printf("Fetching account for user ID: %s", userID)

	account, err := h.as.Get(userID)
	if err != nil {
		h.lg.Printf("Error fetching account for user %s: %v", userID, err)
		http.Error(w, "Failed to retrieve account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.ToJSON(account, w); err != nil {
		h.lg.Printf("Error encoding account to JSON: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	h.lg.Printf("Successfully returned account for user %s", userID)
}

type createReq struct {
	FullName string  `json:"full_name" example:"Sergey Videnin" validate:"required"`
	Balance  float64 `json:"balance" example:"1000000000.00" validate:"gte=0"`
} //@name CreateAccountRequest

// @Title Create Account
// @Description Создает новый аккаунт
// @Tags Accounts Manage
// @Accept json
// @Produce json
// @Param   account_data body createReq true  "Данные для создания аккаунта"
// @Success 201 {string} string "Аккаунт успешно создан"
// @Failure 400 {string} string "Неверные данные запроса"
// @Failure 409 {string} string "Аккаунт уже существует"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router  /account/create [post]
func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	req := &createReq{}
	if err := json.FromJSON(req, r.Body); err != nil {
		h.lg.Printf("Error parsing request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.FullName == "" {
		http.Error(w, "Full name is required", http.StatusBadRequest)
		return
	}

	if req.Balance < 0 {
		http.Error(w, "Balance cannot be negative", http.StatusBadRequest)
		return
	}

	h.lg.Printf("Creating account for new user with balance %.2f", req.Balance)
	id, err := h.as.Add(req.FullName, req.Balance)
	if err != nil {
		h.lg.Printf("Error creating account for user %s: %v", id, err)
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Account created successfully", "id": "` + id + `"}`))

	h.lg.Printf("Successfully created account for user %s", id)
}

type updateReq struct {
	Amount float64 `json:"amount" example:"100.00" validate:"required"`
} //@name UpdateAccountRequest

// @Title Update Account Balance
// @Description Изменить баланс счета
// @Tags Accounts Manage
// @Accept json
// @Produce json
// @Param   user_id  path  string  true  "ID пользователя"
// @Param   update_data body updateReq true  "Данные для обновления баланса"
// @Success 200 {string} string "Баланс успешно обновлен"
// @Failure 400 {string} string "Неверные данные запроса"
// @Failure 404 {string} string "Аккаунт не найден"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router  /account/update/{user_id} [patch]
func (h *Handler) UpdateBalance(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]

	if userID == "" {
		h.lg.Println("Empty user_id provided for balance update")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	req := &updateReq{}
	if err := json.FromJSON(req, r.Body); err != nil {
		h.lg.Printf("Error parsing request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Amount == 0 {
		http.Error(w, "Amount cannot be zero", http.StatusBadRequest)
		return
	}

	h.lg.Printf("Updating balance for user %s by amount %.2f", userID, req.Amount)

	if err := h.as.Update(userID, req.Amount); err != nil {
		h.lg.Printf("Error updating balance for user %s: %v", userID, err)
		http.Error(w, "Failed to update balance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Balance updated successfully"}`))

	h.lg.Printf("Successfully updated balance for user %s", userID)
}
