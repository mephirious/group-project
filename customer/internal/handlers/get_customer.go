package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mephirious/group-project/services/customer/internal/repository"
	"github.com/mephirious/group-project/services/customer/pkg/auth"
)

// CustomerHandler обработчик HTTP-запросов
type CustomerHandler struct {
	Repo *repository.CustomerRepository
}

// GetProfileHandler получает профиль клиента
func (h *CustomerHandler) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// 1️⃣ Проверяем заголовок Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Токен не найден", http.StatusUnauthorized)
		return
	}

	// 2️⃣ Извлекаем JWT-токен
	token := strings.TrimPrefix(authHeader, "Bearer ")
	customerID, err := auth.ParseToken(token)
	if err != nil {
		http.Error(w, "Неверный токен", http.StatusUnauthorized)
		return
	}

	// 3️⃣ Ищем пользователя в MongoDB
	customer, err := h.Repo.GetCustomerByID(customerID)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// 4️⃣ Возвращаем JSON-профиль
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
