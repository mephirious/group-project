package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 🔐 Секретный ключ для подписи токенов (лучше хранить в .env)
var jwtSecret = []byte("super-secret-key")

// Claims структура с пользовательскими данными для токена
type Claims struct {
	CustomerID int `json:"customer_id"`
	jwt.RegisteredClaims
}

// GenerateToken создает JWT-токен для пользователя
func GenerateToken(customerID int) (string, error) {
	// ⏳ Срок действия токена: 24 часа
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создаём claims (полезные данные в токене)
	claims := Claims{
		CustomerID: customerID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Создаем токен с алгоритмом HS256 (HMAC-SHA256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	return token.SignedString(jwtSecret)
}

// ParseToken проверяет и декодирует JWT, возвращая customerID
func ParseToken(tokenString string) (int, error) {
	// Разбираем токен и проверяем подпись
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи (должен быть HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	// Декодируем claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.CustomerID, nil
	}

	return 0, errors.New("неверный токен")
}
