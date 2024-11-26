package envs

import (
	"os"
)

var Etag = os.Getenv("ETAG")

func init() {
	if os.Getenv("IGNORE_CACHE") == "true" {
		Etag = ""
	}
}
