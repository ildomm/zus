package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	tokens "github.com/ildomm/zus/daos/tokens"
	"github.com/ildomm/zus/models"
	. "github.com/ildomm/zus/restapi/operations/tokens"
)

func TokensCreateHandlerResponder(params CreateHashParams) middleware.Responder {
	if tokens.TokenExists(params.Token.Token) {
		return NewCreateHashDefault(409)
	}

	token := new(models.TokenInfo)
	token.Token = params.Token.Token

	if token.Token == "" {
		return NewCreateHashDefault(400 ).WithPayload(&models.GeneralError{Code: 400, Message: "Token required"})
	}
	token, err := tokens.Insert(*token)

	if err != nil {
		return NewCreateHashDefault(500 ).WithPayload(&models.GeneralError{Code: 500, Message: err.Message})
	}

	response := new(models.TokenResponse)
	response.Hash = token.Hash

	return NewCreateHashOK().WithPayload(response)
}

func GetHashesHandlerResponder(params GetHashesParams) middleware.Responder {
	hashes := tokens.All()
	if hashes == nil {
		return NewGetHashsDefault(404)
	}

	return NewGetHashesOK().WithPayload(hashes)
}

func GetHashHandlerResponder(params GetHashParams) middleware.Responder {
	hash := tokens.ByHash(params.ID)
	if hash == nil {
		return NewGetHashDefault(404)
	}

	return NewGetHashOK().WithPayload(hash)
}