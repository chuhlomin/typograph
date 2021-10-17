package typograph

import "regexp"

type NBSP struct {
}

func (nbsp *NBSP) Process(in []byte) []byte {
	afterShortWord := regexp.MustCompile(`\s([A-Za-zА-Яа-я]{1,2}) `)
	in = afterShortWord.ReplaceAll(in, []byte(" ${1}\u00a0"))

	return in
}
