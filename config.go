package querybuilder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadConfig(filepath string) (Config, error) {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)

	fmt.Printf("%s", data)

	return config, nil

}
