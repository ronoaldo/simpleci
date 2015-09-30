// Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>
// Licensed under the Apache License, Version 2.0

package builder

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"ronoaldo.gopkg.net/aetools/vmproxy"
)

var (
	builderPod = `version: v1
kind: Pod
metadata:
  name: builder
spec:
  containers:
    - name: builder
      image: gcr.io/ronoaldoconsulting/builder-default:latest
      imagePullPolicy: Always
      ports:
        - containerPort: 8080
          hostPort: 80
          protocol: TCP
  restartPolicy: Always
  dnsPolicy: Default`

	builderVM = &vmproxy.VM{
		Path: "/api/builder/run",
		Instance: vmproxy.Instance{
			Name:        "builder-default",
			Zone:        "us-central1-f",
			Image:       vmproxy.ResourcePrefix + "/google-containers/global/images/container-vm-v20150806",
			MachineType: "n1-standard-1",
			Metadata: map[string]string{
				"google-container-manifest": builderPod,
			},
		},
	}
)

func init() {
	http.HandleFunc("/_ah/start", AhStart)
	http.HandleFunc("/_ah/stop", AhStop)
	http.Handle("/api/builder/run", builderVM)
}

// AhStart attempts to prevent a 503 error when servicing a loading request.
func AhStart(w http.ResponseWriter, r *http.Request) {
	log.Debugf(appengine.NewContext(r), "[vmproxy]: new instance started.")
}

// AhStop terminates all backend instance virtual machines.
func AhStop(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	log.Debugf(c, "[vmproxy]: terminating instances ...")
	if err := builderVM.Stop(c); err != nil {
		log.Warningf(c, "[vmproxy]: Error terminating echo vm: %v", err)
	}
	log.Debugf(c, "[vmproxy]: instance termination completed.")
}
