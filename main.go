package main

import (
	"arinuryadi/file/log/apache"
	"arinuryadi/file/log/dpkg"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	filename := strings.Split(os.Args[1], "/")
	fileName := strings.Split(filename[len(filename)-1], ".")

	fmt.Println(fileName)
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

	if len(os.Args) == 2 {
		if filename[3] == "apache2" {
			apache.CreateText(fileName[0]+".txt", text)
		}

		if fileName[0] == "dpkg" {
			dpkg.CreateTextDpkg(fileName[0]+".txt", text)
		}

		return
	}

	if len(os.Args) >= 3 {
		if os.Args[2] == "-t" && os.Args[3] == "text" {
			if filename[3] == "apache2" {
				if len(os.Args) == 5 {
					if os.Args[4] == "-o" {
						if len(os.Args) == 6 {
							apache.CreateText(os.Args[5], text)
							return
						} else {
							log.Fatal("The folder is not set yet")
						}
					}
				} else {
					apache.CreateText(fileName[0]+".txt", text)
				}
			}

			if fileName[0] == "dpkg" {
				if len(os.Args) == 5 || len(os.Args) == 6 {
					if os.Args[4] == "-o" {
						if len(os.Args) == 6 {
							dpkg.CreateTextDpkg(os.Args[5], text)
							return
						} else {
							log.Fatal("The folder is not set yet")
						}
					}
				} else {
					dpkg.CreateTextDpkg(fileName[0]+".txt", text)
				}
			}
		}

		if os.Args[2] == "-t" && os.Args[3] == "json" {
			if filename[3] == "apache2" {
				if len(os.Args) == 5 || len(os.Args) == 6 {
					if os.Args[4] == "-o" {
						if len(os.Args) == 6 {
							apache.CreateJson(os.Args[5], text)
							return
						} else {
							log.Fatal("The folder is not set yet")
						}
					}
				} else {
					apache.CreateJson(fileName[0]+".json", text)
				}
			}

			if fileName[0] == "dpkg" {
				if len(os.Args) == 5 || len(os.Args) == 6 {
					if os.Args[4] == "-o" {
						if len(os.Args) == 6 {
							dpkg.CreateJson(os.Args[5], text)
							return
						} else {
							log.Fatal("The folder is not set yet")
						}
					}
				} else {
					dpkg.CreateJson(fileName[0]+".json", text)
				}
			}
		}

		if os.Args[2] == "-o" {
			if len(os.Args) == 4 {
				if filename[3] == "apache2" {
					apache.CreateText(os.Args[3], text)
					return
				}

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
