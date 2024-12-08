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
	// create transaction with status: CREATED
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	// update transaction with status: PAID
	// call service product to reduce the stock by 1
	// add new item to the user inventory by calling inventory service
}
