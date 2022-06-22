package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const e = '\\'

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	rs := []rune(s)
	if unicode.IsDigit(rs[0]) { // краевой случай цифра в начале строки
		return "", ErrInvalidString
	}
	lnR := len(rs)
	prev, prevIsEsc, err := getNext(rs, 0)
	if err != nil {
		return "", err
	}

	var res strings.Builder
	i := 1
	if prevIsEsc {
		i++
	}
	for ; i < lnR; i++ {

		cur, curIsEsc, err := getNext(rs, i)
		if err != nil {
			return "", err
		}

		if curIsEsc {
			i++
		}
		if unicode.IsDigit(cur) {
			if !curIsEsc && !prevIsEsc && unicode.IsDigit(prev) {
				return "", ErrInvalidString
			}
			if curIsEsc {
				res.WriteRune(prev)
			} else {
				cnt, _ := strconv.Atoi(string(cur))
				res.WriteString(strings.Repeat(string(prev), cnt))
			}
		} else if prevIsEsc || !unicode.IsDigit(prev) {
			res.WriteRune(prev)
		}
		prev, prevIsEsc = cur, curIsEsc
	}

	if !unicode.IsDigit(prev) || prevIsEsc {
		res.WriteRune(prev)
	}
	return res.String(), nil
}

func getNext(rs []rune, i int) (r rune, isEsc bool, err error) {
	if i > len(rs)-1 {
		err = ErrInvalidString
		return
	}
	if rs[i] == e {
		if i >= len(rs)-1 {
			err = ErrInvalidString
			return
		}
		r = rs[i+1]
		isEsc = true
	} else {
		r = rs[i]
	}
	return
}
