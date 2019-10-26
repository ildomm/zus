package utils

import (
	"net/http"
	"syreclabs.com/go/faker"
)

func GenerateRandomHttpRequest() *http.Request {
	hTTPRequest := new(http.Request)
	hTTPRequest.RemoteAddr = faker.Internet().IpV4Address()

	return hTTPRequest
}
