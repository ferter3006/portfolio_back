package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"crypto/hmac"
	"crypto/sha256"

	"github.com/google/uuid"

	"new-test/config"
)

// LoadRedsysConfigFromDefault carga la configuración desde config/redsys.json
func LoadRedsysConfigFromDefault() (*config.RedsysConfig, error) {
	return config.LoadRedsysConfig("config/redsys.json")
}

type RedsysPaymentRequest struct {
	Amount      string
	Order       string
	MerchantURL string
	UrlOK       string
	UrlKO       string
	ProductDesc string
}

func GenerateOrderID() string {
	return uuid.New().String()[:12]
}

func createSignature(secret, data string) string {
	key, _ := base64.StdEncoding.DecodeString(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func RedsysPayHandler(cfg *config.RedsysConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID := GenerateOrderID()
		amount := "100" // 1.00 EUR (en céntimos)
		merchantURL := "http://localhost:8080/redsys/notify"
		urlOK := "http://localhost:8080/redsys/ok"
		urlKO := "http://localhost:8080/redsys/ko"
		productDesc := "Test Product"

		params := map[string]string{
			"DS_MERCHANT_AMOUNT":             amount,
			"DS_MERCHANT_ORDER":              orderID,
			"DS_MERCHANT_MERCHANTCODE":       cfg.MerchantCode,
			"DS_MERCHANT_CURRENCY":           cfg.Currency,
			"DS_MERCHANT_TRANSACTIONTYPE":    "0",
			"DS_MERCHANT_TERMINAL":           cfg.Terminal,
			"DS_MERCHANT_MERCHANTURL":        merchantURL,
			"DS_MERCHANT_URLOK":              urlOK,
			"DS_MERCHANT_URLKO":              urlKO,
			"DS_MERCHANT_PRODUCTDESCRIPTION": productDesc,
		}

		jsonParams, _ := json.Marshal(params)
		merchantParams := base64.StdEncoding.EncodeToString(jsonParams)
		signature := createSignature(cfg.SecretKey, merchantParams)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		form := `<form id="redsysForm" action="` + cfg.TestURL + `" method="POST">
			<input type="hidden" name="Ds_SignatureVersion" value="HMAC_SHA256_V1" />
			<input type="hidden" name="Ds_MerchantParameters" value="` + merchantParams + `" />
			<input type="hidden" name="Ds_Signature" value="` + signature + `" />
		</form>
		<script>document.getElementById('redsysForm').submit();</script>`
		tmpl, _ := template.New("form").Parse(form)
		tmpl.Execute(w, nil)

		fmt.Print("form generated")

	}
}

func RedsysNotifyHandler(cfg *config.RedsysConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Aquí se procesa la notificación de Redsys (IPN)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
