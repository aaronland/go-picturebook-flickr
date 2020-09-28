package filter

import (
	"regexp"
)

var flickr_re *regexp.Regexp

func init() {
	flickr_re = regexp.MustCompile(`\d+_[a-z0-9]+_o\..*$`)
}
