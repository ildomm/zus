package db_test

import (
	"github.com/ildomm/zus/config"
	"github.com/ildomm/zus/database"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go/build"
	"os"
	"testing"
)

func TestDatabase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Database Suite")
}

var _ = Describe("Connection", func() {
	BeforeSuite(func() {
		os.Chdir(build.Default.GOPATH + "/src/github.com/ildomm/zus/")
		config.Setup()
		db.Setup()
	})

	BeforeEach(func() {
	})

	Context("Evolve", func() {
		It("able to reset", func() {
			defer GinkgoRecover()
			db.CleanDatabases()
			Expect(db.CountTable("tokens")).To(Equal(int(0)))
		})
	})

	Context("Connect", func() {

		It("has DATABASE instance", func() {
			defer GinkgoRecover()
			Expect(db.Setup()).ToNot(BeNil())
		})

	})
})
