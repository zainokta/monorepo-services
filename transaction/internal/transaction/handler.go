package transaction

import (
	"net/http"
	"transaction/internal/config"
	"transaction/pkg/event"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type Handler struct {
	cfg config.Config
	ch  *amqp.Channel
}

func NewHandler(cfg config.Config, ch *amqp.Channel) Handler {
	return Handler{
		cfg: cfg,
		ch:  ch,
	}
}

func (h *Handler) GetTransaction(c *gin.Context) {

}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	err := StoreTransaction(c.Request.Context(), transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	event.Publisher(h.ch, "", []byte(""))

	c.JSON(http.StatusOK, gin.H{
		"message": "success creating transaction",
	})
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	// update transaction with status: PAID
	// call service product to reduce the stock by 1
	// add new item to the user inventory by calling inventory service

	// api driven atau event driven
}
