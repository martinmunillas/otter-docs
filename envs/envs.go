package envs

import (
	_ "embed"
	"os"
)

//go:embed .etag
var Etag string

func init() {
	if os.Getenv("IGNORE_CACHE") == "true" {
		Etag = ""
	}
}
