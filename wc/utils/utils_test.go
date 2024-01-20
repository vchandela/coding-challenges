package utils

import "testing"

func TestBytes(t *testing.T) {
	fileName := "../test.txt"
	wantNumBytes := 342190
	numBytes := ProcessFile(fileName, "-c")
	if wantNumBytes != numBytes {
		t.Fatalf("Wanted %v, Got %v", wantNumBytes, numBytes)
	}
}

func TestLines(t *testing.T) {
	fileName := "../test.txt"
	wantNumLines := 7145
	numLines := ProcessFile(fileName, "-l")
	if wantNumLines != numLines {
		t.Fatalf("Wanted %v, Got %v", wantNumLines, numLines)
	}
}

func TestWords(t *testing.T) {
	fileName := "../test.txt"
	wantNumWords := 58164
	numWords := ProcessFile(fileName, "-w")
	if wantNumWords != numWords {
		t.Fatalf("Wanted %v, Got %v", wantNumWords, numWords)
	}
}

func TestChars(t *testing.T) {
	fileName := "../test.txt"
	wantNumChars := 339292
	numChars := ProcessFile(fileName, "-m")
	if wantNumChars != numChars {
		t.Fatalf("Wanted %v, Got %v", wantNumChars, numChars)
	}
}
