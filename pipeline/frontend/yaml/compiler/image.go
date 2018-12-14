package compiler

import (
	"github.com/docker/docker/reference"
	"strings"

	dref "github.com/docker/distribution/reference"
)

// trimImage returns the short image name without tag.
func trimImage(name string) string {
	ref, err := reference.ParseNamed(name)
	if err != nil && err != dref.ErrNameNotCanonical {
		return name
	}
	return reference.TrimNamed(ref).String()
}

// expandImage returns the fully qualified image name.
func expandImage(name string) string {
	ref, err := reference.ParseNamed(name)
	if err != nil && err != dref.ErrNameNotCanonical {
		return name
	}
	ret := reference.WithDefaultTag(ref).String()
	if dref.Domain(ref) == "" {
		if strings.Count(ret, "/") == 0 {
			ret = "docker.io/library/" + ret
		}
		if strings.Count(ret, "/") == 1 {
			ret = "docker.io/" + ret
		}
	}

	return ret
}

// matchImage returns true if the image name matches
// an image in the list. Note the image tag is not used
// in the matching logic.
func matchImage(from string, to ...string) bool {
	from = trimImage(from)
	for _, match := range to {
		if from == trimImage(match) {
			return true
		}
	}
	return false
}

// matchHostname returns true if the image hostname
// matches the specified hostname.
func matchHostname(image, hostname string) bool {
	ref, err := reference.ParseNamed(image)
	if err != nil && err != dref.ErrNameNotCanonical {
		return false
	}
	return ref.Hostname() == hostname
}
