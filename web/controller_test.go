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
})
