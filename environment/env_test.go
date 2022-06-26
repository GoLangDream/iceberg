package environment_test

import (
	"github.com/GoLangDream/iceberg/environment"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Env", func() {
	BeforeEach(func() {
		environment.Set(environment.Development)
	})

	When("当环境变量设置了ICEBERG_ENV的时候", func() {
		Context("设置为production", func() {
			It("iceberg就会是production环境", func() {
				err := os.Setenv("ICEBERG_ENV", "production")
				environment.Init()
				if err != nil {
					Fail("环境变量设置失败")
				}
				Expect(environment.IsProduction()).To(BeTrue())
			})
		})

		Context("设置为test, iceberg就会是test环境", func() {
			It("iceberg就会是test环境", func() {
				err := os.Setenv("ICEBERG_ENV", "test")
				environment.Init()
				if err != nil {
					Fail("环境变量设置失败")
				}
				Expect(environment.IsTest()).To(BeTrue())
			})
		})

		Context("设置为其他值的时候", func() {
			It("iceberg就会是development环境", func() {
				err := os.Setenv("ICEBERG_ENV", "abddesd")
				environment.Init()
				if err != nil {
					Fail("环境变量设置失败")
				}
				Expect(environment.IsDevelopment()).To(BeTrue())
			})
		})
	})

	It("默认是开发环境", func() {
		Expect(environment.Get()).To(Equal(environment.Development))
	})

	It("能准确的设置环境", func() {
		environment.Set(environment.Production)
		Expect(environment.Get()).To(Equal(environment.Production))
	})

	It("能正确判断环境", func() {
		environment.Set(environment.Test)

		Expect(environment.IsTest()).To(Equal(true))
		Expect(environment.IsNotTest()).To(Equal(false))

		Expect(environment.IsProduction()).To(Equal(false))
		Expect(environment.IsNotProduction()).To(Equal(true))

		Expect(environment.IsDevelopment()).To(Equal(false))
		Expect(environment.IsNotDevelopment()).To(Equal(true))
	})
})
