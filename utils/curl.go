package utils

import (
	"github.com/idoubi/goz"
)

func CurlGet(url string, param map[string]interface{}, timeout float32) (string, error) {
	cli := goz.NewClient()
	resp, err := cli.Get(url, goz.Options{
		Query: param,
	})
	if err != nil {
		return "", err
	}
	body, _ := resp.GetBody()
	return string(body), nil
}

func CurlPostForm(url string, param map[string]interface{}, timeout float32) (interface{}, error) {
	cli := goz.NewClient()
	resp, err := cli.Post(url, goz.Options{
		Timeout: timeout,
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: param,
	})
	if err != nil {
		return nil, err
	}
	body, _ := resp.GetBody()
	return body, nil
}

func CurlPostJson(url string, param struct{}, timeout float32) (interface{}, error) {
	cli := goz.NewClient()
	resp, err := cli.Post(url, goz.Options{
		Timeout: timeout,
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		JSON: param,
	})
	if err != nil {
		return nil, err
	}
	body, _ := resp.GetBody()
	return body, nil
}
