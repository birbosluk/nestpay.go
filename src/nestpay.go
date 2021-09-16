package nestpay

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

var EndPoints map[string]string = map[string]string{
	"asseco":     "https://entegrasyon.asseco-see.com.tr/fim/api",
	"akbank":     "https://www.sanalakpos.com/fim/api",
	"isbank":     "https://spos.isbank.com.tr/fim/api",
	"ziraatbank": "https://sanalpos2.ziraatbank.com.tr/fim/api",
	"halkbank":   "https://sanalpos.halkbank.com.tr/fim/api",
	"finansbank": "https://www.fbwebpos.com/fim/api",
	"teb":        "https://sanalpos.teb.com.tr/fim/api",
}

var Currencies map[string]string = map[string]string{
	"TRY": "949",
	"YTL": "949",
	"TRL": "949",
	"TL":  "949",
	"USD": "840",
	"EUR": "978",
	"GBP": "826",
	"JPY": "392",
}

type API struct {
	Bank string
}

type Request struct {
	XMLName    xml.Name    `xml:"CC5Request,omitempty"`
	Username   interface{} `xml:"Name,omitempty"`
	Password   interface{} `xml:"Password,omitempty"`
	ClientId   interface{} `xml:"ClientId,omitempty"`
	OrderId    interface{} `xml:"OrderId,omitempty"`
	GroupId    interface{} `xml:"GroupId,omitempty"`
	TransId    interface{} `xml:"TransId,omitempty"`
	UserId     interface{} `xml:"UserId,omitempty"`
	IPAddress  interface{} `xml:"IPAddress,omitempty"`
	Email      interface{} `xml:"Email,omitempty"`
	Mode       interface{} `xml:"Mode,omitempty"`
	Type       interface{} `xml:"Type,omitempty"`
	Number     interface{} `xml:"Number,omitempty"`
	Expires    interface{} `xml:"Expires,omitempty"`
	Cvv2Val    interface{} `xml:"Cvv2Val,omitempty"`
	Total      interface{} `xml:"Total,omitempty"`
	Currency   interface{} `xml:"Currency,omitempty"`
	Instalment interface{} `xml:"Instalment,omitempty"`

	PayerTxnId              interface{} `xml:"PayerTxnId,omitempty"`
	PayerSecurityLevel      interface{} `xml:"PayerSecurityLevel,omitempty"`
	PayerAuthenticationCode interface{} `xml:"PayerAuthenticationCode,omitempty"`
	CardholderPresentCode   interface{} `xml:"CardholderPresentCode,omitempty"`

	BillTo        *To       `xml:"BillTo,omitempty"`
	ShipTo        *To       `xml:"ShipTo,omitempty"`
	PbOrder       *Pb       `xml:"PbOrder,omitempty"`
	OrderItemList *ItemList `xml:"OrderItemList,omitempty"`

	VersionInfo interface{} `xml:"VersionInfo,omitempty"`
}

type To struct {
	Name       interface{} `xml:"Name,omitempty"`
	Company    interface{} `xml:"Company,omitempty"`
	Street1    interface{} `xml:"Street1,omitempty"`
	Street2    interface{} `xml:"Street2,omitempty"`
	Street3    interface{} `xml:"Street3,omitempty"`
	City       interface{} `xml:"City,omitempty"`
	StateProv  interface{} `xml:"StateProv,omitempty"`
	PostalCode interface{} `xml:"PostalCode,omitempty"`
	Country    interface{} `xml:"Country,omitempty"`
	TelVoice   interface{} `xml:"TelVoice,omitempty"`
}

type Pb struct {
	OrderType              interface{} `xml:"OrderType,omitempty"`
	TotalNumberPayments    interface{} `xml:"TotalNumberPayments,omitempty"`
	OrderFrequencyCycle    interface{} `xml:"OrderFrequencyCycle,omitempty"`
	OrderFrequencyInterval interface{} `xml:"OrderFrequencyInterval,omitempty"`
	Desc                   interface{} `xml:"Desc,omitempty"`
	Price                  interface{} `xml:"Price,omitempty"`
	Total                  interface{} `xml:"Total,omitempty"`
}

type Item struct {
	Id          interface{} `xml:"Id,omitempty"`
	ItemNumber  interface{} `xml:"ItemNumber,omitempty"`
	ProductCode interface{} `xml:"ProductCode,omitempty"`
	Qty         interface{} `xml:"Qty,omitempty"`
	Desc        interface{} `xml:"Desc,omitempty"`
	Price       interface{} `xml:"Price,omitempty"`
	Total       interface{} `xml:"Total,omitempty"`
}

type ItemList struct {
	Items []*Item `xml:"OrderItem,omitempty"`
}

type Response struct {
	XMLName        xml.Name `xml:"CC5Response,omitempty"`
	OrderId        string   `xml:"OrderId,omitempty"`
	GroupId        string   `xml:"GroupId,omitempty"`
	TransId        string   `xml:"TransId,omitempty"`
	Response       string   `xml:"Response,omitempty"`
	AuthCode       string   `xml:"AuthCode,omitempty"`
	HostRefNum     string   `xml:"HostRefNum,omitempty"`
	ProcReturnCode string   `xml:"ProcReturnCode,omitempty"`
	ErrMsg         string   `xml:"ErrMsg,omitempty"`
}

func (api *API) Transaction(ctx context.Context, req *Request) (res Response) {
	postdata, err := xml.Marshal(req)
	if err != nil {
		log.Println(err)
		return res
	}
	request, err := http.NewRequestWithContext(ctx, EndPoints[api.Bank], "POST", strings.NewReader(xml.Header+string(postdata)))
	if err != nil {
		log.Println(err)
		return res
	}
	request.Header.Set("Content-Type", "text/xml; charset=utf-8")
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return res
	}
	defer response.Body.Close()
	decoder := xml.NewDecoder(response.Body)
	decoder.Decode(&res)
	return res
}
