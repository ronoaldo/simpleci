// Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>
// Licensed under the Apache License, Version 2.0

package www

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ronoaldo.gopkg.net/simpleci/api"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

type bitbucketPushReq struct {
	Actor struct {
		Username    string `json:username`
		DisplayName string `json:display_name`
	}
	Repository struct {
		FullName string `json:full_name`
		SCM      string `json:scm`
	}
	Changes []struct {
		Type string `json:type`
		Name string `json:name`
	}
	Created bool
	Closed  bool
	Forced  bool
}

// bitbucketWebHook allows user to have a webook.
// TODO(ronoaldo): auth token.
func bitbucketWebHook(w http.ResponseWriter, r *http.Request) {
	webHook := &bitbucketPushReq{}
	if err := json.NewDecoder(r).Decode(&webHook); err != nil {
		log.Warningf(c, "[www] unable to enqueue task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !webHook.Created || len(webHook.Changes) < 1 {
		log.Debugf("[www] discharding webhook with created=false: %v", webHook)
		return
	}
	branchName := "master"
	if webHook.Repository.SCM == "hg" {
		branchName = "default"
	}
	for _, c := range webHook.Changes {
		if c.Type == "branch" {
			branchName = c.Name
			break
		}
	}
	req := &api.BuildRequest{
		Repository: &api.Repository{
			URL: fmt.Sprintf("ssh://%s@bitbucket.org/%s", webHook.Repository.SCM, webHook.Repository.FullName),
			SCM: webHook.Repository.SCM,
		},
		Branch: branchName,
	}

	b, err := json.Marshal(req)
	if err != nil {
		log.Warningf(c, "[www] error generating build payload: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task := taskqueue.Task{
		Path:    "/api/builder/run",
		Payload: b,
		Method:  "POST",
	}
	if err = taskqueue.Add(c, task); err != nil {
		log.Warningf(c, "[www] unable to enqueue task: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Debugf(c, "[www] new build task enqueued: %v", task)
}
