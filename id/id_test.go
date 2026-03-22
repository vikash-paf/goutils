package id

import (
	"fmt"
	"regexp"
	"testing"
)

func TestUUID(t *testing.T) {
	u1 := UUID()
	u2 := UUID()

	if u1 == u2 {
		t.Error("Expected UUIDs to be unique")
	}

	// Basic V4 UUID regex check
	pattern := `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	matched, _ := regexp.MatchString(pattern, u1)
	if !matched {
		t.Errorf("UUID %s did not match V4 pattern", u1)
	}
}

func TestRandomString(t *testing.T) {
	s1 := RandomString(10)
	s2 := RandomString(10)

	if len(s1) != 10 {
		t.Errorf("Expected length 10, got %d", len(s1))
	}
	if s1 == s2 {
		t.Error("Expected random strings to be unique")
	}
}

func ExampleUUID() {
	uuid := UUID()
	fmt.Println(len(uuid))
	// Output: 36
}
