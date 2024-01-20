package main

import (
	"flag"
	"fmt"
	"os"
	"wc/utils"
)

const (
	stdinTempFile = "test.txt"
)

func main() {
	switch len(os.Args) {
	case 1:
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
			// Read from stdin and store in a temporary file to reuse input.
			// Also, input can be large so file preferred over bytes.buffer
			utils.ReadStdinAndStore()
			num_lines := utils.ProcessFile(stdinTempFile, "-l")
			num_words := utils.ProcessFile(stdinTempFile, "-w")
			num_bytes := utils.ProcessFile(stdinTempFile, "-c")
			fmt.Println(num_lines, num_words, num_bytes)
		}
	case 2:
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
			splitType := os.Args[1]
			num := utils.ProcessStdin(splitType)
			fmt.Println(num)
		} else {
			fileName := os.Args[1]
			num_lines := utils.ProcessFile(fileName, "-l")
			num_words := utils.ProcessFile(fileName, "-w")
			num_bytes := utils.ProcessFile(fileName, "-c")
			fmt.Println(num_lines, num_words, num_bytes, fileName)
		}

	case 3:
		byteCountPtr := flag.String("c", "", "outputs the number of bytes in a file")
		lineCountPtr := flag.String("l", "", "outputs the number of lines in a file")
		wordCountPtr := flag.String("w", "", "outputs the number of words in a file")
		runeCountPtr := flag.String("m", "", "outputs the number of characters/runes (UTF-8) in a file")

		flag.Parse()

		switch os.Args[1] {
		case "-c":
			num_bytes := utils.ProcessFile(*byteCountPtr, "-c")
			fmt.Println(num_bytes, *byteCountPtr)
		case "-l":
			num_lines := utils.ProcessFile(*lineCountPtr, "-l")
			fmt.Println(num_lines, *lineCountPtr)
		case "-w":
			num_words := utils.ProcessFile(*wordCountPtr, "-w")
			fmt.Println(num_words, *wordCountPtr)
		case "-m":
			num_runes := utils.ProcessFile(*runeCountPtr, "-m")
			fmt.Println(num_runes, *runeCountPtr)
		}
	}

}
