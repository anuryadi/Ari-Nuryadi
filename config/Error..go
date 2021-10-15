package config

type Error struct {
	//Day		string	`json:"day"`
	//Month	string	`json:"month"`
	//Date	string	`json:"date"`
	//Time	string	`json:"time"`
	//Year	string	`json:"year"`
	Date   string `json:"date"`
	Notice string `json:"notice"`
	Pid    string `json:"pid"`
	Tid    string `json:"tid"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}
