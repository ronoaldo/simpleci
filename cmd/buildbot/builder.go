// Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>
// Licensed under the Apache License, Version 2.0

package main

import (
	"encoding/json"
	"net/http"
	"os"

	"ronoaldo.gopkg.net/simpleci/api"
)

var (
	logger = log.New(os.Stderr, "[buildbot] ", log.LstdFlags|log.llongfile)
)

func init() {
	http.HandleFunc("/", build)
}

func build(w http.ResponseWriter, r *http.Request) {
	req := &api.BuildRequest{}
	if err := json.NewDecoder(r).Decode(req); err != nil {
		errorf(w, "unable to parse build request: %v", err)
		return
	}

	if err := checkout(req); err != nil {
		errorf(w, "unable to checkout project: %v", err)
		return
	}

	if err := runBuildCommands(req); err != nil {
		errorf(w, "unable to run build coomands: %v", err)
		return
	}

	if err := collectBuildResults(req); err != nil {
		errorf(w, "unable to collect build results: %v", err)
		return
	}
}

func checkout(req buildReq) error {
	return nil
}

func runBuildCommands(req buildReq) error {
	return nil
}

func collectBuildResults(req buildReq) error {
	return nil
}

func errorf(w http.ResponseWriter, msg string, args ...interface{}) {
	logger.Printf(msg, args...)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
