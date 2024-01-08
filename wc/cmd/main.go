package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"wc/utils"
)

func main() {
	switch len(os.Args) {
	case 2:
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
			switch os.Args[1] {
			case "-c":
				num_bytes := utils.FileCnt(os.Stdin, "-c")
				fmt.Println(num_bytes)
			case "-l":
				num_lines := utils.FileCnt(os.Stdin, "-l")
				fmt.Println(num_lines)
			case "-w":
				num_words := utils.FileCnt(os.Stdin, "-w")
				fmt.Println(num_words)
			case "-m":
				num_runes := utils.FileCnt(os.Stdin, "-m")
				fmt.Println(num_runes)
			}
		} else {
			match, err := regexp.MatchString(`\.txt$`, os.Args[1])
			if err != nil {
				log.Fatal(err)
			}

			if match {
				num_lines := utils.FileAllCnt(os.Args[1], "-l")
				num_words := utils.FileAllCnt(os.Args[1], "-w")
				num_bytes := utils.FileAllCnt(os.Args[1], "-c")
				fmt.Println(num_lines, num_words, num_bytes, os.Args[1])
			} else {
				fmt.Fprintln(os.Stderr, `Entered file name doesn't follow naming regex: \.txt$`)
			}
		}

	case 3:
		byteCountPtr := flag.String("c", "", "outputs the number of bytes in a file")
		lineCountPtr := flag.String("l", "", "outputs the number of lines in a file")
		wordCountPtr := flag.String("w", "", "outputs the number of words in a file")
		runeCountPtr := flag.String("m", "", "outputs the number of characters/runes (UTF-8) in a file")

		flag.Parse()

		switch os.Args[1] {
		case "-c":
			num_bytes := utils.FileAllCnt(*byteCountPtr, "-c")
			fmt.Println(num_bytes, *byteCountPtr)
		case "-l":
			num_lines := utils.FileAllCnt(*lineCountPtr, "-l")
			fmt.Println(num_lines, *lineCountPtr)
		case "-w":
			num_words := utils.FileAllCnt(*wordCountPtr, "-w")
			fmt.Println(num_words, *wordCountPtr)
		case "-m":
			num_runes := utils.FileAllCnt(*runeCountPtr, "-m")
			fmt.Println(num_runes, *runeCountPtr)
		}
	}

}
