package iland_test

import (
	"encoding/json"
	"github.com/ilanddev/golang-sdk"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

// Load API configuration properties from config.json file.
func loadConfig() *iland.Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}
	decoder := json.NewDecoder(file)
	configuration := iland.Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}
	return &configuration
}

var c = iland.NewClient(loadConfig())

// Test Get request.
func TestGet(t *testing.T) {
	_, err := c.Get("/user/testman?expand=orgs")
	if err != nil {
		t.Fatal(err)
	}
}

// Test Post request.
func TestPostGet(t *testing.T) {
	updateEmails := "\n{\"username\":\"testman\",\"emails\":[\"testman@iland.dev\"]}"
	x, err := c.Post("/user/testman/alert-emails", updateEmails)
	log.Println(x)
	if err != nil {
		t.Fatal(err)
	}
	afterUpdate, err := c.Get("/user/testman/alert-emails")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, updateEmails, afterUpdate)
}

// Test Put Request.
func TestPutGet(t *testing.T) {
	updatePrefs := "\n{\"pref_map\":{\"a\":\"1\"}}"
	_, err := c.Put("/user/testman/preferences", updatePrefs)
	if err != nil {
		t.Fatal(err)
	}
	afterUpdate, err := c.Get("/user/testman/preferences")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, updatePrefs, afterUpdate)
}

//Test Error handling.
func TestError404(t *testing.T) {
	_, err := c.Get("/user/fakeUserHereMan")
	if err == nil {
		t.Fatal("Should have gotten an API error")
	} else {
		log.Println("Error successfully caught: ", err)
	}
}

// Test a delete request.
func TestDelete404(t *testing.T) {
	_, err := c.Delete("/user/XXXXXXXXXXX")
	if err == nil {
		t.Fatal("Should have gotten an API error")
	} else {
		log.Println("Error succesfully caught: ", err)
	}
}
