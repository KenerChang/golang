package util

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/KenerChang/golang/logger"
	"github.com/KenerChang/golang/rest"
	"github.com/gorilla/mux"
	"net/http"
)

// SetRoutes recives routes info and return a mux router
func SetRoutes(routes []rest.Route) *mux.Router {
	apiRoutes := mux.NewRouter()
	for _, route := range routes {
		for _, endpoint := range route.Endpoints {
			path := fmt.Sprintf("/api/v%d/%s", endpoint.Version, route.Name)
			if endpoint.Path != "" {
				path = path + "/" + endpoint.Path
			}

			logger.Info.Printf("set path: %s", path)

			apiRoutes.
				HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					defer func() {
						if err := recover(); err != nil {
							errMsg := fmt.Sprintf("%#v", err)
							WriteWithStatus(w, errMsg, http.StatusInternalServerError)
							logger.Error.Println(r, "Panic error:", errMsg)
						}
					}()

					endpoint.Callback(w, r)
				}).
				Methods(endpoint.Method)
		}

		// init module
		if route.InitFunc != nil {
			route.InitFunc()
		}
	}

	return apiRoutes
}

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
