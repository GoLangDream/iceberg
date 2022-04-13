package web

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controller", func() {
	Context("getNamespace 函数", func() {
		It("当没有子目录的时候返回空", func() {
			namespace := getNamespace("test_project/web/controllers")
			Expect(namespace).To(Equal(""))
		})

		It("当没有子目录, 而且路径最后有斜杠的时候返回空, ", func() {
			namespace := getNamespace("test_project/web/controllers/")
			Expect(namespace).To(Equal(""))
		})

		It("当有一层目录的时候，返回对应的空间名", func() {
			namespace := getNamespace("test_project/web/controllers/home")
			Expect(namespace).To(Equal("home"))
		})

		It("当有一层目录, 而且已斜杠结尾的时候，返回对应的空间名", func() {
			namespace := getNamespace("test_project/web/controllers/home/")
			Expect(namespace).To(Equal("home"))
		})
	})

	Context("getName 函数", func() {
		It("HomeController 应该返回 home", func() {
			name := getName("HomeController")
			Expect(name).To(Equal("home"))
		})

		It("如果struct 名不包括Controller，就返回空字符串", func() {
			name := getName("HomeContr")
			Expect(name).To(Equal(""))
		})

		It("Homecontroller 当controller小写时，不符合规则，返回空字符串", func() {
			name := getName("Homecontroller")
			Expect(name).To(Equal(""))
		})
	})
})
