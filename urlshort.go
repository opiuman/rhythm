package rhythm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func ShortUrl(longUrl, token string) (string, error) {
	req := fmt.Sprintf("https://api-ssl.bitly.com/v3/shorten?access_token=%s&longurl=%s&format=json", token, longUrl)
	resp, err := http.Get(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	var br BitlyResp

	unmarshalErr := json.Unmarshal(body, &br)

	if unmarshalErr != nil {
		return "", err
	}

	return br.Data.Url, nil
}
