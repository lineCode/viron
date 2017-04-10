// example-go JavaScript Client Example
//
// Code generated by goagen v1.1.0-dirty, DO NOT EDIT.
//
// Command:
// $ goagen
// --design=github.com/cam-inc/example-go/design
// --out=$(GOPATH)/src/github.com/cam-inc/example-go/public
// --version=v1.1.0-dirty

package js

import (
	"github.com/goadesign/goa"
)

// MountController mounts the JavaScript example controller under "/js".
// This is just an example, not the best way to do this. A better way would be to specify a file
// server using the Files DSL in the design.
// Use --noexample to prevent this file from being generated.
func MountController(service *goa.Service) {
	// Serve static files under js
	service.ServeFiles("/js/*filepath", "/Users/okada_n/workspace/dmc/example.go/src/github.com/cam-inc/example-go/public/js")
	service.LogInfo("mount", "ctrl", "JS", "action", "ServeFiles", "route", "GET /js/*")
}
