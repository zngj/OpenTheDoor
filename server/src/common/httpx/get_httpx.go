package httpx

import (
	"net/http"
	"github.com/carsonsx/log4g"
	"io/ioutil"
	"encoding/json"
)

func Get(url string, header map[string]string, v interface{}) (err error) {
	var resp *http.Response
	client := new(http.Client)
	log4g.Debug("get    url: %s", url)
	log4g.Debug("get header: %s", header)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log4g.Error(err)
		return
	}
	if header != nil && len(header) > 0 {
		for k,v := range header {
			req.Header.Add(k, v)
		}
	}
	resp, err = client.Do(req)
	if err != nil {
		log4g.Error(err)
		return
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log4g.Error(err)
		return
	}
	log4g.Debug(string(body))
	err = json.Unmarshal(body, v)
	if err != nil {
		log4g.Error(err)
	}
	return
}