package main

import (
	"regexp"
	"strings"
)

type Symbols struct {
}

// Symbols
// -> → →, <- → ←
// (c) → ©, (tm) → ™, (r) → ®
// Adding ° to C and F
func (symbols *Symbols) Process(in []byte) []byte {
	r := strings.NewReplacer(
		"->", "→", "<-", "←",
		"(r)", "®", "(R)", "®",
	)
	in = []byte(r.Replace(string(in)))

	copyRegexp := regexp.MustCompile(`(?i)(copyright )?\((c|с)\)`)
	in = copyRegexp.ReplaceAll(in, []byte(`©`))

	tmRegexp := regexp.MustCompile(`(?i)(\(tm\))`)
	in = tmRegexp.ReplaceAll(in, []byte(`™`))

	cfRegexp := regexp.MustCompile(`(^|\s*[+-≈±−—–]*)(\d+[.,\d]*?) (C|F)([\W\s.,:!?"]+|$)`)
	in = cfRegexp.ReplaceAll(in, []byte(`$1$2 °$3$4`))

	return in
}
