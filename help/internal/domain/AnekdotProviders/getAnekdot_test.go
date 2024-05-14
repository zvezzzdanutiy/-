package AnekdotProviders

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetAnekdot_Success(t *testing.T) {
	domain := New(
		&http.Client{})

	ctx := context.Background()
	category := "test"
	anekdot, err := domain.GetAnekdot(ctx, category)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if anekdot == "" {
		t.Errorf("Expected non-empty anekdot, got empty string")
	}
}

//write unit-test function for GenerateURL function

func TestGenerateURL(t *testing.T) {
	d := &Domain{}

	expectedAPIKey := "7c72cdbe4f22a675723cfccb9669cbe31d7c1aed2abf8ba146f7fa24426b920a"
	expectedPID := "s157zmtnncbm4m15b40k"

	// Установим текущее время для формирования URL
	currentTime := time.Now().Unix()
	expectedUTS := fmt.Sprintf("%d", currentTime)

	expectedStr := fmt.Sprintf("pid=%s&method=getRandItem&uts=%s", expectedPID, expectedUTS)
	expectedHash := GetMD5Hash(expectedStr + expectedAPIKey)
	expectedURL := "http://anecdotica.ru/api?" + expectedStr + "&hash=" + expectedHash

	generatedURL := d.GenerateURL()

	if generatedURL != expectedURL {
		t.Errorf("GenerateURL() returned incorrect URL. Expected: %s, Got: %s", expectedURL, generatedURL)
	}
}
