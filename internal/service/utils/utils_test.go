package utils

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestGenerateAlias(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	alias := GenerateAlias()
	if len(alias) != aliasLength {
		t.Errorf("expected alias length %d, but got %d", aliasLength, len(alias))
	}

	for _, char := range alias {
		if !strings.ContainsRune(string(chars), char) {
			t.Errorf("unexpected character in alias: %c", char)
		}
	}
}

func TestGenerateAliasMultipleTimes(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	alias1 := GenerateAlias()
	alias2 := GenerateAlias()

	if alias1 == alias2 {
		t.Error("expected two different aliases")
	}
}

func TestIsValid(t *testing.T) {
	validUrl := "http://localhost"
	invalidUrl := "alsjdf;sdk"

	if !IsValid(validUrl) {
		t.Error("expected true")
	}

	if IsValid(invalidUrl) {
		t.Error("expected false")
	}
}
