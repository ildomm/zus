package handlers_test

import (
	"github.com/ildomm/zus/config"
	"github.com/ildomm/zus/database"
	handler "github.com/ildomm/zus/handlers/tokens"
	"github.com/ildomm/zus/restapi/operations/tokens"
	"github.com/ildomm/zus/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go/build"
	"os"
	"reflect"
	"syreclabs.com/go/faker"
	"testing"
)

func TestTokensHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tokens Handler Suite")
}

var _ = Describe("Tokens Handler", func() {

	BeforeSuite(func() {
		os.Chdir(build.Default.GOPATH + "/src/github.com/ildomm/zus/")
		config.Setup()
		db.Setup()
	})

	BeforeEach(func() {
		db.CleanDatabases()
	})

	Context("Initialization", func() {
		It("success, has to be empty", func() {
			Expect(db.CountTable("tokens")).To(Equal(int(0)))
		})
	})

	Context("Creating", func() {

		It("fail, has to return error", func() {
			paramsCreate := new(tokens.CreateHashParams)
			paramsCreate.HTTPRequest = utils.GenerateRandomHttpRequest()

			token := new(tokens.CreateHashBody)
			paramsCreate.Token = *token
			paramsCreate.Token.Token = ""

			middlewareResultCreation := handler.TokensCreateHandlerResponder(*paramsCreate)
			Expect(reflect.TypeOf(middlewareResultCreation)).
				To(Equal(reflect.TypeOf(tokens.NewCreateHashDefault(400))))

			Expect(db.CountTable("tokens")).To(Equal(int(0)))
		})

		It("success, has to create", func() {
			paramsCreate := new(tokens.CreateHashParams)
			paramsCreate.HTTPRequest = utils.GenerateRandomHttpRequest()

			token := new(tokens.CreateHashBody)
			paramsCreate.Token = *token
			paramsCreate.Token.Token = faker.Bitcoin().Address()

			middlewareResult := handler.TokensCreateHandlerResponder(*paramsCreate)
			Expect(reflect.TypeOf(middlewareResult)).To(Equal(reflect.TypeOf(tokens.NewCreateHashOK())))

			Expect(db.CountTable("tokens")).To(Equal(int(1)))
		})
	})

	Context("Retrieving entry", func() {

		It("fail, has to return empty", func() {
			params := new(tokens.GetHashParams)
			params.HTTPRequest = utils.GenerateRandomHttpRequest()
			params.ID = faker.Bitcoin().Address()
			middlewareResult := handler.GetHashHandlerResponder(*params)

			payload := utils.GetHashesPayload(middlewareResult)
			Expect(payload).To(BeNil())
		})

		It("success, has to return some", func() {
			paramsCreate := new(tokens.CreateHashParams)
			paramsCreate.HTTPRequest = utils.GenerateRandomHttpRequest()

			token := new(tokens.CreateHashBody)
			paramsCreate.Token = *token
			paramsCreate.Token.Token = faker.Bitcoin().Address()

			middlewareResultCreation := handler.TokensCreateHandlerResponder(*paramsCreate)
			payload := utils.GetCreateHashPayload(middlewareResultCreation)
			Expect(payload).ToNot(BeNil())

			paramsSearch := new(tokens.GetHashParams)
			paramsSearch.HTTPRequest = utils.GenerateRandomHttpRequest()
			paramsSearch.ID = payload.Hash
			middlewareResultSearch := handler.GetHashHandlerResponder(*paramsSearch)

			payloadS := utils.GetHashPayload(middlewareResultSearch)
			Expect(payloadS).ToNot(BeNil())
		})
	})

	Context("Retrieving list", func() {

		It("fail, has to return empty", func() {
			params := new(tokens.GetHashesParams)
			params.HTTPRequest = utils.GenerateRandomHttpRequest()
			middlewareResult := handler.GetHashesHandlerResponder(*params)

			payload := utils.GetHashesPayload(middlewareResult)
			Expect(payload).To(BeNil())
		})

		It("success, has to return some ", func() {
			paramsCreate := new(tokens.CreateHashParams)
			paramsCreate.HTTPRequest = utils.GenerateRandomHttpRequest()

			token := new(tokens.CreateHashBody)
			paramsCreate.Token = *token
			paramsCreate.Token.Token = faker.Bitcoin().Address()

			middlewareResultCreation := handler.TokensCreateHandlerResponder(*paramsCreate)
			payload := utils.GetCreateHashPayload(middlewareResultCreation)
			Expect(payload).ToNot(BeNil())

			paramsSearch := new(tokens.GetHashesParams)
			paramsSearch.HTTPRequest = utils.GenerateRandomHttpRequest()
			middlewareResultSearch := handler.GetHashesHandlerResponder(*paramsSearch)

			payloadS := utils.GetHashesPayload(middlewareResultSearch)
			Expect(payloadS).ToNot(BeNil())
		})
	})

})
