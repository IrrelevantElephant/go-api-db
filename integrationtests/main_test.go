package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
)

type album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func Test_GetAll_ReturnsSuccess(t *testing.T) {
	url := fmt.Sprintf("%s/albums", getUrl())
	fmt.Println(url)
	response, err := http.Get(url)

	if err != nil {
		t.Fatalf("Failed to execute GET request: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}
}

func Test_Post_ReturnsLocation(t *testing.T) {
	url := fmt.Sprintf("%s/albums", getUrl())
	fmt.Println(url)

	var myalbum = album{
		Title:  "my new title",
		Artist: "an artist",
		Price:  21.22,
	}

	jsonValue, _ := json.Marshal(myalbum)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to execute POST request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, but got %d", http.StatusCreated, response.StatusCode)
	}

	location := response.Header.Get("location")
	if location == "" {
		t.Fatal("Location header is empty")
	}

	getResponse, err := http.Get(location)
	if err != nil {
		t.Fatalf("Failed to execute GET request to %s: %v", location, err)
	}
	defer getResponse.Body.Close()

	if getResponse.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, getResponse.StatusCode)
	}

	var returnedAlbum album
	if err := json.NewDecoder(getResponse.Body).Decode(&returnedAlbum); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if !reflect.DeepEqual(returnedAlbum, myalbum) {
		t.Fatalf("Expected album %+v, but got %+v", myalbum, returnedAlbum)
	}
}

func getUrl() string {
	return os.Getenv("SUT_URL")
}
