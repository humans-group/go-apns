package apns

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestHttpClient(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "APNs Package", []Reporter{junitReporter})
}
