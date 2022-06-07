package environment_test

import (
	"github.com/GoLangDream/iceberg/environment"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Env", func() {

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
