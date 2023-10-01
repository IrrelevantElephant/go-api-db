package integrationtests

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Test_GetAll_ReturnsSuccess(t *testing.T) {
	url := fmt.Sprintf("%s/albums", getUrl())

	fmt.Println(url)

	response, err := http.Get(url)

	if err != nil {
		t.FailNow()
	}

	if response.StatusCode != http.StatusOK {
		t.Fail()
	}
}

func getUrl() string {
	return os.Getenv("SUT_URL")
}
