package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var Mask = regexp.MustCompile(`[\p{L}-]+`)

func Top10(s string) []string {
	f := make(map[string]int)
	w := Mask.FindAllString(s, -1)
	for _, v := range w {
		if v == "-" {
			continue
		}
		fv := strings.ToLower(v)
		f[fv]++
	}
	r := rankByWordCount(f)
	N := len(r)
	if N > 10 {
		N = 10
	}
	res := make([]string, 0, N)
	for i := 0; i < N; i++ {
		res = append(res, r[i].Key)
	}
	return res
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		switch strings.Compare(p[i].Key, p[j].Key) {
		case 1:
			return false
		case -1:
			return true
		}
	}
	return p[i].Value > p[j].Value
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
