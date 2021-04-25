package models

import (
	"bytes"
	"strings"
)

func Replace1(template, match1, replace1 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)

	return res
}

func Replace2(template, match1, replace1, match2, replace2 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)

	return res
}

func Replace3(template, match1, replace1, match2, replace2, match3, replace3 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)

	return res
}

func Replace4(template, match1, replace1, match2, replace2, match3, replace3, match4, replace4 string) string {

	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)
	res = strings.ReplaceAll(res, match4, replace4)

	return res
}

func Replace5(template, match1, replace1, match2, replace2, match3, replace3, //
	match4, replace4, match5, replace5 string) string {
	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)
	res = strings.ReplaceAll(res, match4, replace4)
	res = strings.ReplaceAll(res, match5, replace5)

	return res
}
func Replace6(template, match1, replace1, match2, replace2, match3, replace3, //
	match4, replace4, match5, replace5, match6, replace6 string) string {
	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)
	res = strings.ReplaceAll(res, match4, replace4)
	res = strings.ReplaceAll(res, match5, replace5)
	res = strings.ReplaceAll(res, match6, replace6)

	return res
}

func Replace7(template, match1, replace1, match2, replace2, match3, replace3, //
	match4, replace4, match5, replace5, match6, replace6, match7, replace7 string) string {
	res := ""
	res = strings.ReplaceAll(template, match1, replace1)
	res = strings.ReplaceAll(res, match2, replace2)
	res = strings.ReplaceAll(res, match3, replace3)
	res = strings.ReplaceAll(res, match4, replace4)
	res = strings.ReplaceAll(res, match5, replace5)
	res = strings.ReplaceAll(res, match6, replace6)
	res = strings.ReplaceAll(res, match7, replace7)

	return res
}

func makeFirstLowerCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}
