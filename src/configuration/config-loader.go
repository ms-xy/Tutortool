package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func loadJsonConfig(path string, v interface{}) {
	if fi, err := os.Stat(path); err != nil || fi.IsDir() {
		if err == nil && fi.IsDir() {
			err = errors.New("not a file")
		}
		panic("could not find config file '" + path + "':\n" + err.Error())
	} else {
		if data, err := ioutil.ReadFile(path); err != nil {
			panic("could not read config file '" + path + "':\n" + err.Error())
		} else {
			if err := json.Unmarshal(data, v); err != nil {
				panic("could not parse config file '" + path + "':\n" + err.Error())
			}
		}
	}
}
