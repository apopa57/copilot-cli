// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var message = "hello world"

// HealthCheck just returns true if the service is up.
func HealthCheck(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("🚑 healthcheck ok!")
	w.WriteHeader(http.StatusOK)
}

// ServiceDiscoveryGet just returns true no matter what
func ServiceDiscoveryGet(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Printf("Get on ServiceDiscovery endpoint Succeeded with message %s\n", message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func main() {
	router := httprouter.New()
	router.GET("/service-discovery/", ServiceDiscoveryGet)

	// Health Check
	router.GET("/", HealthCheck)

	log.Fatal(http.ListenAndServe(":80", router))
}
