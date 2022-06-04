package iceberg_test

import (
	"github.com/GoLangDream/iceberg/environment"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIceberg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iceberg 测试")
}

var _ = BeforeSuite(func() {
	environment.Set(environment.Test)
})
