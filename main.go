package main

import (
	"AssessmentTestBackendTelkom/file/log/apache"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	filename := strings.Split(os.Args[1], "/")
	fileName := strings.Split(filename[len(filename)-1], ".")

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	fmt.Println(len(os.Args))
	if len(os.Args) == 2 {
		apache.CreateText(fileName[0]+".txt", text)
		return
	}

	if len(os.Args) >= 3 {
		if os.Args[2] == "-t" && os.Args[3] == "text" {
			if filename[3] == "apache2" {
				if len(os.Args) == 6 {
					if os.Args[4] == "-o" {
						apache.CreateText(os.Args[5], text)
						return
					} else {
						log.Fatal("The folder is not set yet")
					}
				} else {
					apache.CreateText(fileName[0]+".txt", text)
					return
				}
			}
		}

		if os.Args[2] == "-t" && os.Args[3] == "json" {
			if len(os.Args) == 6 {

				if os.Args[4] == "-o" {
					apache.CreateJson(os.Args[5], text)
					return
				}

			} else {
				apache.CreateJson(fileName[0]+".json", text)
				return
			}
		}

		if os.Args[2] == "-o" {
			if len(os.Args) == 4 {
				apache.CreateText(os.Args[3], text)
				return
			} else {
				log.Fatal("The folder is not set yet")
			}
		}

		if os.Args[2] == "-h" {
			fmt.Println("Cara Kerja")
			return
		}
	}
}
