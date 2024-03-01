package usecase

import "testing"

func TestFetchStoredRSSLink(t *testing.T) {
	expected := "https://echo.labstack.com"
	actual := FetchStoredRSSLink()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
