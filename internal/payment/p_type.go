package payment

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Payments interface {
	NewPayment(*gin.Context)
	Verify(*gin.Context)
}

type Payment struct {
	log slog.Logger
}
