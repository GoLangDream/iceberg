package web_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iceberg web模块测试\n")
}

func getBody(response *http.Response) string {
	body, _ := ioutil.ReadAll(response.Body)
	return string(body)
}

func parseBody(response *http.Response, v any) {
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, v)
}
