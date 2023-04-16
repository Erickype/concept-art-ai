package user

import (
	"context"
	"strings"
	"testing"
)

// TestCreateAndGetUser - test that the shortened URL is stored and retrieved from database.
func TestCreateAndGetUser(t *testing.T) {
	testEmail := "test@email.com"
	sp := CreateRequest{Email: testEmail}
	resp, err := Create(context.Background(), &sp)
	if err != nil {
		t.Fatal(err)
	}
	wantEmail := testEmail
	if resp.Status != "Created" {
		t.Errorf("got %q, want %q", resp.Status, wantEmail)
	}

	firstEmail := wantEmail
	gotEmail, err := Get(context.Background(), firstEmail)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Compare(gotEmail.Email, firstEmail) != 0 {
		t.Errorf("got %v, want %v", gotEmail.Email, firstEmail)
	}
}
