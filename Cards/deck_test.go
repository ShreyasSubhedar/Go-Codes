package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 , but got %v", len(d))

	}
	if d[0] != "A of ♠" {
		t.Errorf("Expected card to be A of ♠ , but got %v", d[0])
	}
	if d[len(d)-1] != "🦁 of ♣" {
		t.Errorf("Expected card to be 🦁 of ♣ , but got %v", d[len(d)-1])
	}
}
