package main

import (
	"testing"
)

func TestCountryLookupForIran(t *testing.T) {
	country := countryLookup("5.160.0.0")
	if country != "Iran" {
		t.Error("Expected Iran, but got ", country)
	}
}

func TestCountryLookupForUs(t *testing.T) {
	country := countryLookup("50.160.16.241")
	if country != "United States" {
		t.Error("Expected United States, but got", country)
	}
}
