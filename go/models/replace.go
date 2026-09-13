package models

import (
	"strings"
)

// Replace replaces all occurrences of old/new string pairs in template.
// pairs must have an even length (old1, new1, old2, new2, ...).
func Replace(template string, pairs ...string) string {
	if len(pairs) == 0 {
		return template
	}
	return strings.NewReplacer(pairs...).Replace(template)
}

func Replace1(template, match1, replace1 string) string {
	return Replace(template, match1, replace1)
}

func Replace2(template, match1, replace1, match2, replace2 string) string {
	return Replace(template, match1, replace1, match2, replace2)
}

func Replace3(template, match1, replace1, match2, replace2, match3, replace3 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3)
}

func Replace4(template, match1, replace1, match2, replace2, match3, replace3, match4, replace4 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3, match4, replace4)
}

func Replace5(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5)
}

func Replace6(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6)
}

func Replace7(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7)
}

func Replace8(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8)
}

func Replace9(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9)
}

func Replace10(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
	match10, replace10 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
		match10, replace10)
}

func Replace11(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
	match10, replace10, match11, replace11 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
		match10, replace10, match11, replace11)
}

func Replace12(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
	match10, replace10, match11, replace11, match12, replace12 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
		match10, replace10, match11, replace11, match12, replace12)
}

func Replace13(template, match1, replace1, match2, replace2, match3, replace3,
	match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
	match10, replace10, match11, replace11, match12, replace12, match13, replace13 string) string {
	return Replace(template, match1, replace1, match2, replace2, match3, replace3,
		match4, replace4, match5, replace5, match6, replace6, match7, replace7, match8, replace8, match9, replace9,
		match10, replace10, match11, replace11, match12, replace12, match13, replace13)
}
