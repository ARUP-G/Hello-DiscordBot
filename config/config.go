package config

import (
	"encoding/json" //to use unmersal
	"fmt"
	"io/ioutil" // read the file
)

var (
	Token     string
	BotPrefix string

	//config is a pointer to config struct
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	//Reading config file

	//ReadFile reads the file named by filename and returns the contents. A successful call returns err == nil, not err == EOF.
	// here file stores the data as byte format
	file, err := ioutil.ReadFile("./config.json")

	ErrorCheck(err)

	//printing the file is console
	fmt.Println(string(file))

	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v(config 'here').
	// If v(config) is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	err = json.Unmarshal(file, &config)
	ErrorCheck(err)

	//golag readable values
	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil

}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
