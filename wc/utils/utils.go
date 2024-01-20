package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	fileNameErr = errors.New(`entered file name doesn't follow naming regex: \.txt$`)
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

func ProcessFile(fileName string, splitType string) int {
	match, err := regexp.MatchString(`\.txt$`, fileName)
	check(err)

	if !match {
		check(fileNameErr)
	}

	f, err := os.Open("./" + fileName)
	check(err)
	defer f.Close()

	return FileCnt(f, splitType)
}

func ReadStdinAndStore() {
	f, err := os.Create("temp.txt")
	check(err)

	_, err = io.Copy(f, os.Stdin)
	check(err)
}

func ProcessStdin(splitType string) int {
	var num int

	switch splitType {
	case "-c":
		num = FileCnt(os.Stdin, "-c")
	case "-l":
		num = FileCnt(os.Stdin, "-l")
	case "-w":
		num = FileCnt(os.Stdin, "-w")
	case "-m":
		num = FileCnt(os.Stdin, "-m")
	}
	return num
}
