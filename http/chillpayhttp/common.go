package chillpayhttp

import (
	"bytes"
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

type ChillpayConfig struct {
	ChillpayMd5          string `env:"CHILLPAY_MD5,required"`
	ChillpayMerchantcode string `env:"CHILLPAY_MERCHANTCODE,required"`
	ChillpayApikey       string `env:"CHILLPAY_APIKEY,required"`
	ChillpayUrl          string `env:"CHILLPAY_URL,required"`
}

var (
	POST            = []byte(fasthttp.MethodPost)
	GET             = []byte(fasthttp.MethodGet)
	PUT             = []byte(fasthttp.MethodPut)
	PATCH           = []byte(fasthttp.MethodPatch)
	DELETE          = []byte(fasthttp.MethodDelete)
	ApplicationJSON = []byte("application/json")
)

type HTTPRequest interface {
	NewRequest(body []byte, method []byte, url string) (*fasthttp.Request, *fasthttp.Response)
}

type fastHTTP struct {
	chillpayMerchantCode string
	chillpayApiKey       string
}

func NewRequester(merchantCode string, apiKey string) *fastHTTP {
	return &fastHTTP{
		chillpayMerchantCode: merchantCode,
		chillpayApiKey:       apiKey,
	}
}

func (u *fastHTTP) NewRequest(body []byte, method []byte, url string) (*fasthttp.Request, *fasthttp.Response) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetBody(body)
	req.Header.SetMethodBytes(method)
	req.SetRequestURIBytes([]byte(url))

	req.Header.Set("CHILLPAY-MerchantCode", u.chillpayMerchantCode)
	req.Header.Set("CHILLPAY-ApiKey", u.chillpayApiKey)
	req.Header.Set("Accept", "*/*")

	return req, resp
}

func CreateQuery(inf interface{}) (string, error) {
	v, err := query.Values(inf)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func ErrorMessage(b []byte) (ErrorMessageResponse, error) {
	var serverResponse ErrorMessageResponse
	reader := bytes.NewBuffer(b)
	err := json.NewDecoder(reader).Decode(&serverResponse)
	if err != nil {
		return ErrorMessageResponse{}, err
	}
	return serverResponse, nil
}
