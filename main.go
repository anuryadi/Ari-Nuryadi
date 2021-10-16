package main

import (
	"arinuryadi/file/log/apache"
	"arinuryadi/file/log/auth"
	"arinuryadi/file/log/dpkg"
	"arinuryadi/file/log/kern"
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

	if os.Args[1] == "-h" {
		fmt.Println("Usage : go run main.go <path file> <flag> <detail>")
		fmt.Println("Option :")
		fmt.Println("  -t \t\t\t Flag For Option to Convert the File to Plaintext (text) or JSON (json)")
		fmt.Println("  -o \t\t\t Flag For Choose Where He Will Put the Output File. [/home/$USER/text.txt], [/home/$USER/text.json]")
		fmt.Println("  -h \t\t\t Flag For Instructions for Use.")
		fmt.Println("")
		fmt.Println("List File :")
		fmt.Println("> /var/log/apache2/error.log")
		fmt.Println("> /var/log/dpkg.log")
		fmt.Println("> /var/log/kern.log")
		fmt.Println("> /var/log/auth.log")
		fmt.Println("")
		fmt.Println("Example :")
		fmt.Println("> go run main.go /var/log/auth.log \t\t\t\t\t Output File txt")
		fmt.Println("> go run main.go /var/log/auth.log -t text \t\t\t\t Output File txt")
		fmt.Println("> go run main.go /var/log/auth.log -t json \t\t\t\t Output File json")
		fmt.Println("> go run main.go /var/log/auth.log -t json -o /home/$USER/text.json \t Output File in Folder /home/$USER/ with Type Json and File Name text.json")
		fmt.Println("> go run main.go /var/log/auth.log -o /home/$USER/text.txt \t\t Output File In Folder /home/$USER/ with File Name text.txt")
		fmt.Println("> go run main.go -h \t\t\t\t\t\t\t Instructions for Use.")

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

	processCreateFileNoFlag(fileName[0], filename[3], text)

	if len(os.Args) >= 3 {
		processCreateFileWithFlagT(fileName[0], filename[3], text)

		processCreateFileWithFlagO(fileName[0], filename[3], text)
	}
}

func processCreateFileNoFlag(fileName, filename string, text []string) {
	if len(os.Args) == 2 {
		if filename == "apache2" {
			apache.CreateText(fileName+".txt", text)
			return
		}

		if fileName == "dpkg" {
			dpkg.CreateTextDpkg(fileName+".txt", text)
			return
		}

		if fileName == "kern" {
			kern.CreateTextKern(fileName+".txt", text)
			return
		}

		if fileName == "auth" {
			auth.CreateTextAuth(fileName+".txt", text)
			return
		}
	}
}

func processCreateFileWithFlagT(fileName, filename string, text []string) {

	if os.Args[2] == "-t" {
		// Validasi
		if len(os.Args) > 3 {
			// Pembuatan File With Flag -t Jenis File Text
			if os.Args[3] == "text" {
				if filename == "apache2" {
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
						apache.CreateText(fileName+".txt", text)
					}
				}

				if fileName == "dpkg" {
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
						dpkg.CreateTextDpkg(fileName+".txt", text)
					}
				}

				if fileName == "kern" {
					if len(os.Args) == 5 || len(os.Args) == 6 {
						if os.Args[4] == "-o" {
							if len(os.Args) == 6 {
								kern.CreateTextKern(os.Args[5], text)
								return
							} else {
								log.Fatal("The folder is not set yet")
							}
						}
					} else {
						kern.CreateTextKern(fileName+".txt", text)
					}
				}

				if fileName == "auth" {
					if len(os.Args) == 5 || len(os.Args) == 6 {
						if os.Args[4] == "-o" {
							if len(os.Args) == 6 {
								auth.CreateTextAuth(os.Args[5], text)
								return
							} else {
								log.Fatal("The folder is not set yet")
							}
						}
					} else {
						auth.CreateTextAuth(fileName+".txt", text)
					}
				}
			}

			// Pembuatan File With Flag -t Jenis File Json
			if os.Args[3] == "json" {
				if filename == "apache2" {
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
						apache.CreateJson(fileName+".json", text)
					}
				}

				if fileName == "dpkg" {
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
						dpkg.CreateJson(fileName+".json", text)
					}
				}

				if fileName == "kern" {
					if len(os.Args) == 5 || len(os.Args) == 6 {
						if os.Args[4] == "-o" {
							if len(os.Args) == 6 {
								kern.CreateJsonKern(os.Args[5], text)
								return
							} else {
								log.Fatal("The folder is not set yet")
							}
						}
					} else {
						kern.CreateJsonKern(fileName+".json", text)
					}
				}

				if fileName == "auth" {
					if len(os.Args) == 5 || len(os.Args) == 6 {
						if os.Args[4] == "-o" {
							if len(os.Args) == 6 {
								auth.CreateJsonAuth(os.Args[5], text)
								return
							} else {
								log.Fatal("The folder is not set yet")
							}
						}
					} else {
						auth.CreateJsonAuth(fileName+".json", text)
					}
				}
			}
		} else {
			log.Fatal("Type Text is not set yet")
		}
	}
}

func processCreateFileWithFlagO(fileName, filename string, text []string) {
	if os.Args[2] == "-o" {
		if len(os.Args) == 4 {
			if filename == "apache2" {
				apache.CreateText(os.Args[3], text)
				return
			}
		} else {
			log.Fatal("The folder is not set yet")
		}

		if len(os.Args) == 4 {
			if fileName == "dpkg" {
				dpkg.CreateTextDpkg(os.Args[3], text)
				return
			}
		} else {
			log.Fatal("The folder is not set yet")
		}

		if len(os.Args) == 4 {
			if fileName == "kern" {
				kern.CreateTextKern(os.Args[3], text)
				return
			}
		} else {
			log.Fatal("The folder is not set yet")
		}

		if len(os.Args) == 4 {
			if fileName == "auth" {
				auth.CreateTextAuth(os.Args[3], text)
				return
			}
		} else {
			log.Fatal("The folder is not set yet")
		}
	}
}
