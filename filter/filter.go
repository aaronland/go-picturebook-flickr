package filter

import (
	"regexp"
)

var flickr_re *regexp.Regexp

func init() {
	flickr_re = regexp.MustCompile(`o_\.\.*$`)
}
