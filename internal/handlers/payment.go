package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klarqqs/dayfi_backend/internal/config"
	"github.com/klarqqs/dayfi_backend/internal/stellar"
)

type PaymentRequest struct {
	SourceSecret string `json:"source_secret"`
	Destination  string `json:"destination"`
	Amount       string `json:"amount"`
	AssetCode    string `json:"asset_code"`
}

func SendPayment(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PaymentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		txHash, err := stellar.SendPayment(cfg.StellarClient, req.SourceSecret, req.Destination, req.Amount, req.AssetCode, cfg.NetworkPassphrase)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Payment sent successfully",
			"tx_hash":  txHash,
			"explorer": "https://stellar.expert/explorer/testnet/tx/" + txHash,
		})
	}
}
