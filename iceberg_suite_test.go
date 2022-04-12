package iceberg_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIceberg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iceberg Suite")
}
