// Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>
// Licensed under the Apache License, Version 2.0

package www

import "net/http"

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/bitbucket/webook", bitbucketWebHook)
}

// index serves the index page for the user interface.
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
