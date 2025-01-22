package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"os"

	"github.com/FulgurCode/stitch/models"
)

func RazorPaymentVerification(order models.VerifyOrder) error {
	var secret = os.Getenv("RAZORPAY_SECRET")
	data := order.OrderId+ "|" + order.PaymentId

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(order.RazorpaySignature)) != 1 {
		return errors.New("Payment failed")
	} else {
		return nil
	}
}
