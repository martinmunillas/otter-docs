package envs

import (
	"os"
	"os/exec"
)

var Etag string

func init() {
	if os.Getenv("IGNORE_CACHE") != "true" {
		hash, _ := exec.Command("git", "rev-parse", "HEAD").Output()
		Etag = string(hash)
	}
}
