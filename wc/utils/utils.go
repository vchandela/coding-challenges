package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FileCnt(input io.Reader, splitType string) int {
	s := bufio.NewScanner(input)

	switch splitType {
	case "-c":
		s.Split(bufio.ScanBytes)
	case "-w":
		s.Split(bufio.ScanWords)
	case "-m":
		s.Split(bufio.ScanRunes)
	}

	sum := 0
	for s.Scan() {
		sum += 1
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error in reading file:", err)
	}

	return sum
}

func FileAllCnt(fileName string, splitType string) int {
	f, err := os.Open("./" + fileName)
	check(err)
	defer f.Close()

	return FileCnt(f, splitType)
}
