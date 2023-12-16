package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 6 {
		t.Errorf("Expected deck length of 6, but got %v", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card of Ace of Spade, but got %v", d[0])
	}

	if d[len(d)-1] != "Queen of Diamond" {
		t.Errorf("Expected last card of Queen of Diamond, but got %v", d[len(d)-1])
	}
}
