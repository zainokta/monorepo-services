package transaction

import (
	"transaction/internal/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg config.Config
}

func NewHandler(cfg config.Config) Handler {
	return Handler{
		cfg: cfg,
	}
}

func (h *Handler) GetTransaction(c *gin.Context) {
	
}

func (h *Handler) CreateTransaction(c *gin.Context) {

}

func (h *Handler) UpdateTransaction(c *gin.Context) {

}