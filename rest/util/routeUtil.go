package util

import (
	"fmt"
	"github.com/KenerChang/golang/logger"
	"github.com/KenerChang/golang/rest"
	"github.com/gorilla/mux"
	"net/http"
)

// SetRoutes recives routes info and return a mux
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
