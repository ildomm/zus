package handlers_test

import (
	"github.com/ildomm/linx_challenge/config"
	"github.com/ildomm/linx_challenge/db"
	handlers "github.com/ildomm/linx_challenge/handlers/users"
	"github.com/ildomm/linx_challenge/restapi/operations/users"
	"github.com/ildomm/linx_challenge/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go/build"
	"os"
	"reflect"
	"testing"
)

func TestUsersHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Users Handler Suite")
}
var _ = Describe("Users Handler", func() {

	BeforeSuite(func() {
		os.Chdir(build.Default.GOPATH + "/src/github.com/ildomm/linx_challenge/")
		config.Setup()
		db.Setup()
	})

	BeforeEach(func() {
		db.CleanDatabases()
	})

	Context("User", func() {

		It("success, it has to create new user", func() {
			params := utils.GenerateRandomUserParams(nil)
			middlewareResult := handlers.CreateUserHandlerResponder(*params)

			payload := utils.GetUsersCreationPayload(middlewareResult)

			Expect(reflect.TypeOf(middlewareResult)).To(Equal(reflect.TypeOf(users.NewCreateUserOK())))
			Expect(payload).ToNot(BeNil())
			Expect(payload.ID).To(Equal(params.ID.ID))
			Expect(db.CountTable("users")).To(Equal(int(1)))
		})

		It("fail, it has to deny ", func() {
			id := "fixed"
			params := utils.GenerateRandomUserParams(&id)
			handlers.CreateUserHandlerResponder(*params)
			Expect(db.CountTable("users")).To(Equal(int(1)))

			middlewareResult := handlers.CreateUserHandlerResponder(*params)

			payload := utils.GetUsersCreationPayload(middlewareResult)

			Expect(reflect.TypeOf(middlewareResult)).To(Equal(reflect.TypeOf(users.NewCreateUserDefault(409))))
			Expect(payload).To(BeNil())
			Expect(db.CountTable("users")).To(Equal(int(1)))
		})

	})

	Context("Url", func() {

		It("success, it has to create new user + url", func() {
			id := "fixed"
			paramsU := utils.GenerateRandomUserParams(&id)
			handlers.CreateUserHandlerResponder(*paramsU)

			paramsR := utils.GenerateRandomUrlParams(id)
			middlewareResult := handlers.CreateURLHandlerResponder(*paramsR)

			payload := utils.GetUrlsCreationPayload(middlewareResult)

			Expect(reflect.TypeOf(middlewareResult)).To(Equal(reflect.TypeOf(users.NewCreateURLOK())))
			Expect(payload).ToNot(BeNil())
			Expect(payload.ShortURL).ToNot(BeNil())
			Expect(db.CountTable("users")).To(Equal(int(1)))
			Expect(db.CountTable("urls")).To(Equal(int(1)))
		})

		It("success, it has to delete user and url(s)", func() {
			id := "fixed"
			paramsU := utils.GenerateRandomUserParams(&id)
			handlers.CreateUserHandlerResponder(*paramsU)

			paramsR := utils.GenerateRandomUrlParams(id)
			handlers.CreateURLHandlerResponder(*paramsR)

			paramsD := utils.GenerateRandomUserDeleteParams(&id)
			handlers.DeleteUserHandlerResponder(*paramsD)

			Expect(db.CountTable("users")).To(Equal(int(0)))
			Expect(db.CountTable("urls")).To(Equal(int(0)))
		})

	})


})
