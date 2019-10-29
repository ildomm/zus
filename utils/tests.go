package utils

import (
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ildomm/zus/models"
	"github.com/ildomm/zus/restapi/operations/tokens"
	"net/http"
	"syreclabs.com/go/faker"
)

func GenerateRandomHttpRequest() *http.Request {
	hTTPRequest := new(http.Request)
	hTTPRequest.RemoteAddr = faker.Internet().IpV4Address()

	return hTTPRequest
}

func GetHashesPayload(middlewareResult middleware.Responder) []*models.TokenInfo {
	out, err := json.Marshal(middlewareResult)
	if err != nil {
		panic(err)
	}
	payload := new(tokens.GetHashesOK)
	if err := json.Unmarshal(out, &payload); err != nil {
		panic(err)
	}

	return payload.Payload
}

func GetCreateHashPayload(middlewareResult middleware.Responder) *models.TokenResponse {
	out, err := json.Marshal(middlewareResult)
	if err != nil {
		panic(err)
	}
	handlerOK := new(tokens.CreateHashOK)
	if err := json.Unmarshal(out, &handlerOK); err != nil {
		panic(err)
	}

	return handlerOK.Payload
}

func GetHashPayload(middlewareResult middleware.Responder) *models.TokenInfo {
	out, err := json.Marshal(middlewareResult)
	if err != nil {
		panic(err)
	}
	payload := new(tokens.GetHashOK)
	if err := json.Unmarshal(out, &payload); err != nil {
		panic(err)
	}

	return payload.Payload
}
