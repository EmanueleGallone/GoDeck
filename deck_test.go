package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	const length = 16

	if len(deck) != length {
		t.Errorf("Expected deck length of %v but we got %v", length, len(deck))
	}
}