package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/qevo/go-char/char/gen"
	"github.com/qevo/go-char/char/sub"
	"github.com/qevo/go-shiftidx/shiftidx"
)

type Alphabet struct {
	Code string
	UC   string
	LC   string
}

func (alphabet *Alphabet) Set(lang string) *Alphabet {
	alphabet.Code = lang
	return alphabet
}

func (alphabet *Alphabet) String() string {
	return alphabet.Code
}

var lang string
var exclude string
var shift int
var text string
var langMaps map[string]Alphabet

func init() {
	langMaps = make(map[string]Alphabet)
	langMaps["en"] = Alphabet{
		Code: "en",
		UC:   "[A-Z]",
		LC:   "[a-z]",
	}

	flag.StringVar(&lang, "lang", "en", "the user provided language code")
	flag.StringVar(&exclude, "skip", "[[:space:]]", "the user provided character exclusion list")
	flag.IntVar(&shift, "n", 3, "the number of positions to shift the alphabet")
	flag.StringVar(&text, "text", "Hello World!", "the text you want to encode")
	flag.Parse()
}

func main() {
	langMap := langMaps[lang]

	var uc, lc string
	var e error = nil

	if uc, e = gen.Regex(langMap.UC); e != nil {
		fmt.Println(e.Error())
		return
	}

	if lc, e = gen.Regex(langMap.LC); e != nil {
		fmt.Println(e.Error())
		return
	}

	ucR := strings.Join(shiftidx.Rotate[string](strings.Split(uc, ""), shift), "")
	lcR := strings.Join(shiftidx.Rotate[string](strings.Split(lc, ""), shift), "")

	charShifter, _ := sub.Create(
		[]string{uc, lc}, // standard alphabet
		[]string{ucR, lcR}, // shifted alphabet
		[]string{`[[:space:]]`}, // excludes
	)

	s, _ := charShifter.Do(text)

	fmt.Println("")
	fmt.Println(text)
	fmt.Println("")
	fmt.Println(s)
	fmt.Println("")
}
