package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Post data to server
func Post(url string, headers map[string]string, data string) (string, error) {

	client := http.Client{}

	payload := []byte(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return "", errors.New("Http response code not valid [" + strconv.Itoa(resp.StatusCode) + "]")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Put data to server
func Put(url string, headers map[string]string, data string) (string, error) {

	client := http.Client{}

	payload := []byte(data)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return "", errors.New("Http response code not valid [" + strconv.Itoa(resp.StatusCode) + "]")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Get data from server
func Get(url string, headers map[string]string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return "", errors.New("Http response code not valid [" + strconv.Itoa(resp.StatusCode) + "]")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Head data header from server
func Head(url string, headers map[string]string) (*http.Header, error) {
	client := http.Client{}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, errors.New("Http response code not valid [" + strconv.Itoa(resp.StatusCode) + "]")
	}

	return &resp.Header, nil
}
