package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log — глобальный логгер для сервиса
var Log = logrus.New()

func InitLogger() {
	// Устанавливаем формат логов
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}
