package file

import (
	"encoding/json"
	"github.com/carsonsx/log4g"
	"io/ioutil"
)

func LoadJsonConfig(filename string, v interface{}) {
	log4g.Info("loading config file")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log4g.Warn(err)
	} else {
		err = json.Unmarshal(data, v)
		if err != nil {
			log4g.Fatal(err)
		}
		log4g.Info("loaded %s", filename)
	}
}
