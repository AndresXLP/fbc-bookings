package dto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"time"

	"fbc-bookings/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Payment struct {
	Event       string    `json:"event"`
	Data        Data      `json:"data"`
	Environment string    `json:"environment"`
	Signature   Signature `json:"signature"`
	Timestamp   int       `json:"timestamp"`
	SentAt      time.Time `json:"sent_at"`
}
type Transaction struct {
	Id                string `json:"id"`
	AmountInCents     int    `json:"amount_in_cents"`
	Reference         string `json:"reference"`
	CustomerEmail     string `json:"customer_email"`
	Currency          string `json:"currency"`
	PaymentMethodType string `json:"payment_method_type"`
	RedirectURL       string `json:"redirect_url"`
	Status            string `json:"status"`
	ShippingAddress   any    `json:"shipping_address"`
	PaymentLinkID     any    `json:"payment_link_id"`
	PaymentSourceID   any    `json:"payment_source_id"`
}
type Data struct {
	Transaction Transaction `json:"transaction"`
}
type Signature struct {
	Properties []string `json:"properties"`
	Checksum   string   `json:"checksum"`
}

func (p *Payment) GetCheckSum256() string {
	secret := config.Environments().Payments.Secret
	data := ""

	for _, property := range p.Signature.Properties {
		title := cases.Title(language.LatinAmericanSpanish)
		replacer := strings.NewReplacer("transaction.", "")
		fieldName := title.String(replacer.Replace(property))
		if strings.Contains(fieldName, "_") {
			words := strings.Split(fieldName, "_")
			fieldName = ""
			for _, word := range words {
				fieldName += title.String(word)
			}
		}

		fieldValue := reflect.ValueOf(p.Data.Transaction).FieldByName(fieldName).Interface()
		data += fmt.Sprintf("%v", fieldValue)
	}
	data += fmt.Sprintf("%d", p.Timestamp) + secret
	hash := sha256.New().Sum([]byte(data))

	return hex.EncodeToString(hash)
}
