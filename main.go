package main

import (
	"fmt"
	"io"
	"os"
)

const (
	COMPRESS   = "-c"
	DECOMPRESS = "-d"
	RLE        = "RLE"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <algo> <mode> <input_file> <output_file>\n", os.Args[0])
		fmt.Println("Mode: -c for compression, -d for decompression")
		return
	}

	inputFilePath := os.Args[3]
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Error: Cannot open the input file <%s>\n", inputFilePath)
		return
	}
	defer inputFile.Close()

	outputFilePath := os.Args[4]
	outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: Cannot open the input file <%s>\n", outputFilePath)
		return
	}
	defer outputFile.Close()

	algo := os.Args[1]
	mode := os.Args[2]

	if algo == RLE {
		if mode == COMPRESS {
			rleCompress(inputFile, outputFile)
		} else if mode == DECOMPRESS {
			rleDecompress(inputFile, outputFile)
		} else {
			fmt.Println("Invalid mode. Use -c for compression or -d for decompression.")
		}
	} else {
		fmt.Println("Unsupported algorithm. Currently supported: RLE")
	}
}

func rleCompress(input *os.File, output *os.File) {
	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for i := 0; i < len(content); {
		count := 1
		for j := i + 1; j < len(content) && content[i] == content[j]; j++ {
			count++
		}

		output.WriteString(fmt.Sprintf("%c%d", content[i], count))
		i += count
	}
}

func rleDecompress(input *os.File, output *os.File) {}
