package main

import (
	"context"
	"encoding/xml"
	"fmt"

	nestpay "github.com/ozgur-soft/nestpay/src"
)

func main() {
	api := &nestpay.API{"asseco"} // "asseco","akbank","isbank","ziraatbank","halkbank","finansbank","teb"
	request := new(nestpay.Request)
	request.ClientId = "" // Müşteri No
	request.Username = "" // Kullanıcı adı
	request.Password = "" // Şifre
	// Ödeme
	request.Type = "Auth"
	request.Mode = "P"                           // TEST : "T" - PRODUCTION "P"
	request.IPAddress = "1.2.3.4"                // Müşteri IP adresi (zorunlu)
	request.Number = "4242424242424242"          // Kart numarası
	request.Expires = "02/20"                    // Son kullanma tarihi (Ay ve Yılın son 2 hanesi) AA/YY
	request.Cvv2Val = "000"                      // Cvv2 Kodu (kartın arka yüzündeki 3 haneli numara)
	request.Total = "1.00"                       // Satış tutarı
	request.Instalment = ""                      // Taksit sayısı
	request.Currency = nestpay.Currencies["TRY"] // Para birimi
	// Fatura
	request.BillTo = new(nestpay.To)
	request.BillTo.Name = ""     // Kart sahibi
	request.BillTo.TelVoice = "" // Telefon numarası
	// 3D (varsa)
	//request.PayerTxnId = ""
	//request.PayerSecurityLevel = ""
	//request.PayerAuthenticationCode = ""
	//request.CardholderPresentCode = ""
	response := api.Transaction(context.Background(), request)
	pretty, _ := xml.MarshalIndent(response, " ", " ")
	fmt.Println(string(pretty))
}
