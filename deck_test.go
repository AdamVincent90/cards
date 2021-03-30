package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected deck length of 52 but got ", len(d))
	}

	if d[0] != "Two of Hearts" {
		t.Error("Expected Two of Hearts, instead got ", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Error("Expected Ace of Spades, instead got ", d[len(d)-1])
	}
}

func TestWriteToFileAndReadFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.writeToFile("_decktesting")

	d2 := readFile("_decktesting")

	if len(d2) != 52 {
		t.Error("Length of file expected 51, instead got ", len(d2))
	}

	os.Remove("_decktesdting")
}
