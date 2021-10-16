package auth

import (
	"arinuryadi/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateTextAuth(fileName string, text []string) {
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

func CreateJsonAuth(fileName string, text []string) {
	auth := config.AuthLog{}
	var output []config.AuthLog

	for _, each_ln := range text {
		data := strings.Split(each_ln, ":")

		detail1 := strings.Split(data[2], " ")

		auth.Date = data[0] + "-" + data[1] + "-" + detail1[0]
		auth.Server = detail1[1]
		auth.IdentityName = strings.TrimSpace(detail1[2])
		auth.Desc = strings.TrimLeft(strings.Join(data[3:], ""), " ")

		output = append(output, auth)
	}

	json, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
		return
	}

	unescapeJson, _ := config.UnescapeUnicodeCharactersInJSON(json)

	_ = ioutil.WriteFile(fileName, unescapeJson, 0644)
}
