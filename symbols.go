package main

import (
	"regexp"
	"strings"
)

// Symbols
// -> → →, <- → ←
// (c) → ©, (tm) → ™, (r) → ®
// Adding ° to C and F
func Symbols(in string) (out string) {
	r := strings.NewReplacer(
		"->", "→", "<-", "←",
		"(r)", "®", "(R)", "®",
	)
	in = r.Replace(in)

	b := []byte(in)
	copyRegexp := regexp.MustCompile(`(?i)(copyright )?\((c|с)\)`)
	b = copyRegexp.ReplaceAll(b, []byte(`©`))

	tmRegexp := regexp.MustCompile(`(?i)(\(tm\))`)
	b = tmRegexp.ReplaceAll(b, []byte(`™`))

	cfRegexp := regexp.MustCompile(`(^|\s*[+-≈±−—–]*)(\d+[.,\d]*?) (C|F)([\W\s.,:!?"]+|$)`)
	b = cfRegexp.ReplaceAll(b, []byte(`$1$2 °$3$4`))

	return string(b)
}
