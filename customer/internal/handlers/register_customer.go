package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mephirious/group-project/services/customer/internal/models"
)

// Регистрация нового пользователя
func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 💡 Добавить хеширование пароля и сохранение в БД (опущено)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}
