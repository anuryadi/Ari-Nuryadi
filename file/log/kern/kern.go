package kern

import (
	"arinuryadi/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateTextKern(fileName string, text []string) {
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

func CreateJsonKern(fileName string, text []string) {
	kern := config.KernLog{}
	var output []config.KernLog

	for _, each_ln := range text {
		o := strings.ReplaceAll(each_ln, "[", "|")
		c := strings.ReplaceAll(o, "]", "|")
		//fmt.Println(c)
		data := strings.Split(c, "|")

		data_1 := strings.Split(data[0], " ")

		kern.Date = strings.Join(data_1[:3], " ")
		kern.ComputerName = data_1[3]
		kern.Kern = strings.ReplaceAll(data_1[4], ":", "")
		kern.IdentityNumber = strings.TrimSpace(data[1])

		if len(data) == 5 {
			kern.Desc = strings.TrimSpace(data[3]) + data[4]
		} else {
			kern.Desc = strings.TrimLeft(data[2], " ")

		}

		output = append(output, kern)
	}

	json, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
		return
	}

	unescapeJson, _ := config.UnescapeUnicodeCharactersInJSON(json)

	_ = ioutil.WriteFile(fileName, unescapeJson, 0644)
}
