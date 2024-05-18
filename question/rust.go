package question

import _ "embed"

//go:embed rust/ok.png
var rustOK []byte

//go:embed rust/1.png
var rust1 []byte

//go:embed rust/2.png
var rust2 []byte

//go:embed rust/3.png
var rust3 []byte

//go:embed rust/4.png
var rust4 []byte

//go:embed rust/5.png
var rust5 []byte

//go:embed rust/6.png
var rust6 []byte

//go:embed rust/7.png
var rust7 []byte

//go:embed rust/ok2.png
var rustOK2 []byte

func rust() ([]byte, string, []int, error) {
	type q struct {
		img []byte
		ans []int
	}

	rusts := []q{
		{img: rustOK, ans: []int{}},
		{img: rustOK2, ans: []int{}},
		{img: rust1, ans: []int{0, 1}},
		{img: rust2, ans: []int{0, 1}},
		{img: rust3, ans: []int{3, 4}},
		{img: rust4, ans: []int{3}},
		{img: rust5, ans: []int{4}},
		{img: rust6, ans: []int{3}},
		{img: rust7, ans: []int{6}},
	}

	n := randint(0, len(rusts))
	return rusts[n].img, "bugs", rusts[n].ans, nil
}
