package handlers

import "testing"

func TestBuildDefaultCompanyCredentials(t *testing.T) {
	username, password := buildDefaultCompanyCredentials(42)

	if username != "company_42" {
		t.Fatalf("expected username to be company_42, got %s", username)
	}

	if password != "company@42" {
		t.Fatalf("expected password to be company@42, got %s", password)
	}
}
