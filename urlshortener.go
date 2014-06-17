package rhythm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var defaultToken = "3435"

type BitlyResp struct {
	StatusCode int    `json:"status_code"`
	StatusTxt  string `json:"status_txt"`
	Data       struct {
		LongUrl    string `json:"long_url"`
		Url        string `json:"url"`
		Hash       string `json:"hash"`
		GlobalHash string `json:"global_hash"`
		NewHash    int    `json:"new_hash"`
	} `json:"data"`
}

func ShortUrl(longUrl string) (string, error) {
	req := "https://api-ssl.bitly.com/v3/shorten?access_token=" + defaultToken + "&longurl=" + longUrl + "&format=json"
	resp, err := http.Get(req)
	if err != nil {
		return "error", fmt.Errorf("connection failed: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "error", fmt.Errorf("file read failed: %s", err)
	}
	var br BitlyResp

	UnmarshalErr := json.Unmarshal(body, &br)

	if UnmarshalErr != nil {
		return "error", fmt.Errorf("unmarshal failed: %s", UnmarshalErr)
	}

	return br.Data.Url, nil
}
