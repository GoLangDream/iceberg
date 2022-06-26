package log_test

import (
	"github.com/GoLangDream/iceberg/log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

var _ = Describe("Env", func() {

	It("日志默认的头信息包含iceberg提示 和 当前时间", func() {
		Expect(log.Prefix()).To(
			MatchRegexp(`\[iceberg \d{2}:\d{2}:\d{2}\]`),
		)
	})

	It("Info 能正确的带上前缀输出内容", func() {
		hook := test.NewGlobal()
		log.Info("abc test")

		Expect(hook.LastEntry().Level).To(Equal(logrus.InfoLevel))
		output, _ := hook.LastEntry().Bytes()
		Expect(string(output)).To(
			MatchRegexp(`.*\[iceberg \d{2}:\d{2}:\d{2}\].* abc test`),
		)

		hook.Reset()
		Expect(hook.LastEntry()).To(BeNil())
	})

	It("Infof 能正确的带上前缀输出内容", func() {
		hook := test.NewGlobal()
		log.Infof("abc test %s", "ddd")

		Expect(hook.LastEntry().Level).To(Equal(logrus.InfoLevel))
		output, _ := hook.LastEntry().Bytes()
		Expect(string(output)).To(
			MatchRegexp(`.*\[iceberg \d{2}:\d{2}:\d{2}\].* abc test ddd`),
		)

		hook.Reset()
		Expect(hook.LastEntry()).To(BeNil())
	})

})
