package chillpayhttp

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type ChillpayHTTP struct {
	serviceURL string
	md5        string
	apprequest HTTPRequest
}

func NewChillpayHTTP(config ChillpayConfig) ChillpayHTTP {
	apprequest := NewRequester(config.ChillpayMerchantcode, config.ChillpayApikey)
	return ChillpayHTTP{
		serviceURL: config.ChillpayUrl,
		apprequest: apprequest,
		md5:        config.ChillpayMd5,
	}
}

func (r *ChillpayHTTP) GetPaylinkGenerate(condition *PaylinkGenerateRequest) (response *PaylinkGenerateResponse, err error) {
	errFrom := func(err error) (*PaylinkGenerateResponse, error) {
		return nil, err
	}

	url := fmt.Sprintf("%v/api/v1/paylink/generate", r.serviceURL)

	logrus.Info("url:", url)

	condition.GenerateChecksum(r.md5)

	b, err := json.Marshal(&condition)
	if err != nil {
		return errFrom(err)
	}

	logrus.Info("body:", string(b))

	req, resp := r.apprequest.NewRequest(b, POST, url)
	req.Header.SetContentTypeBytes(ApplicationJSON)
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			logrus.Info("request: ", err)
			return errFrom(err)
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			errResponse, err := ErrorMessage(body)
			if err != nil {
				logrus.Info("response:", err)
				return response, err
			}
			return response, fmt.Errorf("%s", errResponse.Message)
		}

		reader := bytes.NewBuffer(body)
		err := json.NewDecoder(reader).Decode(&response)
		if err != nil {
			logrus.Info("decode:", err)
			return errFrom(err)
		}
	}
	return response, nil
}

func (r *ChillpayHTTP) GetPaylinkDetail(condition *PaylinkDetailsRequest) (response *PaylinkGenerateResponse, err error) {
	errFrom := func(err error) (*PaylinkGenerateResponse, error) {
		return nil, err
	}

	url := fmt.Sprintf("%v/api/v1/paylink/details", r.serviceURL)

	condition.GenerateChecksum(r.md5)

	b, err := json.Marshal(&condition)
	if err != nil {
		return errFrom(err)
	}

	req, resp := r.apprequest.NewRequest(b, POST, url)
	req.Header.SetContentTypeBytes(ApplicationJSON)
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			return errFrom(err)
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			errResponse, err := ErrorMessage(body)
			if err != nil {
				return response, err
			}
			return response, fmt.Errorf("%s", errResponse.Message)
		}

		reader := bytes.NewBuffer(body)
		err := json.NewDecoder(reader).Decode(&response)
		if err != nil {
			return errFrom(err)
		}
	}
	return response, nil
}

func (r *ChillpayHTTP) GetPaylinkClose(condition *PaylinkDetailsRequest) (response *PaylinkGenerateResponse, err error) {
	errFrom := func(err error) (*PaylinkGenerateResponse, error) {
		return nil, err
	}

	url := fmt.Sprintf("%v/api/v1/paylink/close", r.serviceURL)

	b, err := json.Marshal(&condition)
	if err != nil {
		return errFrom(err)
	}

	req, resp := r.apprequest.NewRequest(b, POST, url)
	req.Header.SetContentTypeBytes(ApplicationJSON)
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			return errFrom(err)
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			errResponse, err := ErrorMessage(body)
			if err != nil {
				return response, err
			}
			return response, fmt.Errorf("%s", errResponse.Message)
		}

		reader := bytes.NewBuffer(body)
		err := json.NewDecoder(reader).Decode(&response)
		if err != nil {
			return errFrom(err)
		}
	}
	return response, nil
}

func (r *ChillpayHTTP) GetDetailByTransctionID(condition *TransctionIDRequest) (response *DetailByTransctionResponse, err error) {
	errFrom := func(err error) (*DetailByTransctionResponse, error) {
		return nil, err
	}

	url := fmt.Sprintf("%v/api/v1/paylinktransaction/details", r.serviceURL)

	condition.GenerateChecksum(r.md5)

	b, err := json.Marshal(&condition)
	if err != nil {
		return errFrom(err)
	}

	req, resp := r.apprequest.NewRequest(b, POST, url)
	req.Header.SetContentTypeBytes(ApplicationJSON)
	{
		err := fasthttp.Do(req, resp)
		if err != nil {
			return errFrom(err)
		}
		fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)
	}

	{
		body := resp.Body()
		if resp.StatusCode() != fasthttp.StatusOK {
			errResponse, err := ErrorMessage(body)
			if err != nil {
				return response, err
			}
			return response, fmt.Errorf("%s", errResponse.Message)
		}

		reader := bytes.NewBuffer(body)
		err := json.NewDecoder(reader).Decode(&response)
		if err != nil {
			return errFrom(err)
		}
	}
	return response, nil
}
