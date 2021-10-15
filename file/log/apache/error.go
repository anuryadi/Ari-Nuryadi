package apache

import (
	"AssessmentTestBackendTelkom/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func CreateJson(fileName string, text []string) {
	error := config.Error{}
	var output []config.Error

	if strings.ToLower(os.Args[3]) == "json" {
		for _, each_ln := range text {
			u := strings.ReplaceAll(each_ln, "[", "|")
			t := strings.ReplaceAll(u, "]", "|")
			//fmt.Printf("%s\n", t)

			s := strings.Split(t, "|")

			date := strings.Split(s[1], " ")

			value := date[0] + " " + date[1] + " " + date[2] + " " + date[3] + " " + date[4]
			time, _ := time.Parse(time.ANSIC, value)
			error.Date = time.String()

			event := strings.Split(s[3], ":")
			error.Notice = event[0]

			pidTid := strings.Split(s[5], ":")
			pid := strings.Split(pidTid[0], " ")
			error.Pid = pid[1]
			tid := strings.Split(pidTid[1], " ")
			error.Tid = tid[1]

			status_detail := strings.Split(s[6], ":")
			error.Status = strings.TrimSpace(status_detail[0])
			error.Detail = status_detail[1]

			output = append(output, error)
		}

		json, err := json.Marshal(output)
		if err != nil {
			fmt.Println(err)
			return
		}

		_ = ioutil.WriteFile(fileName, json, 0644)
	}
}

func CreateText(fileName string, text []string) {
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
