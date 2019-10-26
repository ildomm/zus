package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	. "github.com/ildomm/zus/restapi/operations/tokens"
)

func TokensCreateHandlerResponder(params CreateHashParams) middleware.Responder {
	return NewCreateHashOK()
}

func GetHashesHandlerResponder(params GetHashesParams) middleware.Responder {
	return NewCreateHashOK()
}

func GetHashHandlerResponder(params GetHashParams) middleware.Responder {
	return NewCreateHashOK()
}