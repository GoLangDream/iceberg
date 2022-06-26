package environment_test

import (
	"github.com/GoLangDream/iceberg/initializers"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEnvironment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Environment包测试")
}

var _ = BeforeSuite(func() {
	initializers.Init()
})
