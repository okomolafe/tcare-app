package httputil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//MakeGetRequest makes a get request to a specified URL and unmarshals the response into the provided castTo object
func MakeGetRequest(url string, castTo interface{})  error {
	resp, err := http.Get(url)
	if err != nil {
		return  err
	}

	body := resp.Body
	defer func() {
		if body != nil {
			body.Close()
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return  err
	}

	err = json.Unmarshal(data, &castTo)	
	return err
}