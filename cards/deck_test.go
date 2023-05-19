package main

import "testing"

func TestNedwDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d.containsDuplicates() {
		t.Errorf("The new deck contains duplicated cards")
	}
}
