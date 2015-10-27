// Package time contains the work-horse for processing all requests.
//
// This package will instantiate useable objects, and include methods
//  for database interaction and processing.
package time

import (
	"go/build"
	"net/http"
	"path"

	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/gorilla/mux"
)

type Time struct {
	Id int `json:"id"`
}

// New will provision a fresh Time object.
// Time provides the brunt of the shortoftime work, dealing with
//  all data and file system efforts.
func New() *Time {
	t := new(Time)

	return t
}

// GetMainFile will fetch the path of the actual main HTML file
func (t *Time) GetMainFile() string {
	return path.Join(getPath(), "/templates/layout.html")
}

// GetResource will fetch the path of a specific resource file
func (t *Time) GetResource(r *http.Request) string {
	// of course, this needs filtering
	return path.Join(getPath(), "/templates/"+mux.Vars(r)["resource"])
}

func getPath() string {
	return path.Join(build.Default.GOPATH, "src/github.com/CerealBoy/shortoftime")
}

/* func (t *Time) Setup() {

  // schema
  Table: item
    identifier
	title
	content
	owner
	created
	updated
	state
  Table: tags
    identifier
	colour
	title
	display
  Table: user
    identifier
	name
	display
	email
	created
	logged
	oauth
  Table: logs
    identifier
	file
	date
	level
	message

} */
