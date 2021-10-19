package typograph

import "regexp"

type Space struct {
}

func (space *Space) Process(in []byte) (out []byte) {
	comma := regexp.MustCompile(`([\.|,|!|;|\?])([^\sa-z»›\)\.…!\?])`)
	return comma.ReplaceAll(in, []byte("$1 $2"))
}
