package util

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

// WriteJSON is a sugar function which handle response json
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	js, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
	return nil
}

// ReadJSON is a sugar function which decodes json obj from request
func ReadJSON(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(target)
	return err
}

func WriteWithStatus(w http.ResponseWriter, content string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(content))
}

func NewSkipSecureVerifyClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}
