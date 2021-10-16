package dpkg

import (
	"arinuryadi/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateJson(fileName string, text []string) {
	dpkg := config.DpkgLog{}
	var output []config.DpkgLog

	for _, each_ln := range text {
		data := strings.Split(each_ln, " ")
		dpkg.Date = data[0]
		dpkg.Time = data[1]
		dpkg.Status = data[2]
		dpkg.Desc = strings.Join(data[3:], " ")

		output = append(output, dpkg)

	}

	json, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
		return
	}

	unescapeJson, _ := config.UnescapeUnicodeCharactersInJSON(json)

	_ = ioutil.WriteFile(fileName, unescapeJson, 0644)
}

func CreateTextDpkg(fileName string, text []string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, each_ln := range text {
		_, err2 := f.WriteString(each_ln + "\n")

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
