package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger inicializa el logger global de Zap
func InitLogger() {
	// Configuración base
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // Formato de tiempo legible
	config.EncodeLevel = zapcore.CapitalLevelEncoder // Niveles en mayúsculas (INFO, ERROR, etc)

	// Crear el encoder (JSON para producción, Consola para desarrollo)
	var encoder zapcore.Encoder
	env := os.Getenv("APP_ENV")
	if env == "production" {
		encoder = zapcore.NewJSONEncoder(config)
	} else {
		// Para desarrollo, usamos colores y formato de consola
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(config)
	}

	// Definir el nivel de log (Info por defecto)
	logLevel := zap.InfoLevel
	if env != "production" {
		logLevel = zap.DebugLevel
	}

	// Crear el core del logger que escribe en la salida estándar (consola)
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	// Construir el logger
	Log = zap.New(core, zap.AddCaller())
}

// Sync asegura que todos los logs en buffer se escriban antes de cerrar la app
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}