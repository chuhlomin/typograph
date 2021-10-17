package main

import "regexp"

type Space struct {
}

func (space *Space) Process(in []byte) (out []byte) {
	comma := regexp.MustCompile(`(\.|,|!|;|\?)([^).…!;?\s[\])»›])`)
	return comma.ReplaceAll(in, []byte("$1 $2"))
}
