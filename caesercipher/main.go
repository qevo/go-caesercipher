/*
Caeser Cipher example

In cryptography, a Caesar cipher, also known as Caesar's cipher, the shift cipher,
Caesar's code or Caesar shift, is one of the simplest and most widely known encryption
techniques. It is a type of substitution cipher in which each letter in the plaintext
is replaced by a letter some fixed number of positions down the alphabet. For example,
with a left shift of 3, D would be replaced by A, E would become B, and so on. The
method is named after Julius Caesar, who used it in his private correspondence.
Source: https://en.wikipedia.org/wiki/Caesar_cipher
*/

// Package caesercipher is a tool for encoding and decoding text by
// changing the position of letters in the alphabet.
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/qevo/go-char/char/gen"
	"github.com/qevo/go-char/char/sub"
	"github.com/qevo/go-shiftidx/shiftidx"
)

var lang, skip, text string
var shift int
var LangMaps map[string]Alphabet

func init() {
	LangMaps = make(map[string]Alphabet)
	LangMaps["en"] = Alphabet{
		Code: "en",
		uc:   "[A-Z]",
		lc:   "[a-z]",
	}

	flag.StringVar(&lang, "lang", "en", "the user provided language code")
	flag.StringVar(&skip, "skip", "[[:space:]][[:punct:]]", "the user provided character exclusion list")
	flag.IntVar(&shift, "n", 3, "the number of positions to shift the alphabet")
	flag.StringVar(&text, "text", "Hello World!", "the text you want to encode")
	flag.Parse()
}

func main() {
	// parse the values above
	langMap := LangMaps[lang]
	uc := langMap.GetUC()
	lc := langMap.GetLC()

	// shift the alphabet
	ucR, lcR := langMap.Shift(shift)
	exclude := langMap.GetSkipped(skip)

	fmt.Println("INPUTS")
	fmt.Println("\t" + uc)
	fmt.Println("\t" + lc)
	fmt.Println("\t   ... => ", shift)
	fmt.Println("\t" + ucR)
	fmt.Println("\t" + lcR)
	fmt.Println("")
	fmt.Println("\tSkip:")
	fmt.Println("\t" + exclude)
	fmt.Println("----------------------------------------")

	// get the character shifter
	charShifter, e := sub.Create(
		[]string{uc, lc},   // standard alphabet
		[]string{ucR, lcR}, // shifted alphabet
		[]string{exclude},  // excluded from substitution
	)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	// get the shifted text
	shiftedText, e := charShifter.Do(text)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println("")

	// show the start text for reference
	fmt.Println("")
	fmt.Println(text)
	fmt.Println("")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// show the shifted text
	fmt.Println("")
	fmt.Println(shiftedText)
	fmt.Println("")

	fmt.Println("----------------------------------------")
}

type Alphabet struct {
	Code string
	uc   string
	lc   string
}

func (alphabet *Alphabet) Set(lang string) *Alphabet {
	alphabet.Code = lang
	return alphabet
}
func (alphabet *Alphabet) String() string {
	return alphabet.Code
}
func (alphabet *Alphabet) Shift(n int) (string, string) {
	ucR := strings.Join(shiftidx.Rotate[string](strings.Split(alphabet.GetUC(), ""), n), "")
	lcR := strings.Join(shiftidx.Rotate[string](strings.Split(alphabet.GetLC(), ""), n), "")
	return ucR, lcR
}
func (alphabet *Alphabet) GetUC() string {
	// expand regex to get the alphabet in upper case
	if uc, e := expandRE(alphabet.uc); e != nil {
		fmt.Println(e.Error())
		return ""
	} else {
		return uc
	}
}
func (alphabet *Alphabet) GetLC() string {
	// expand regex to get the alphabet in lower case
	if lc, e := expandRE(alphabet.lc); e != nil {
		fmt.Println(e.Error())
		return ""
	} else {
		return lc
	}
}
func (alphabet *Alphabet) GetSkipped(in string) string {
	// expand regex to get the skipped characters
	if out, e := expandRE(in); e != nil {
		fmt.Println(e.Error())
		return ""
	} else {
		return out
	}
}
func expandRE(in string) (string, error) {
	if out, e := gen.Regex(in); e != nil {
		fmt.Println(e.Error())
		return "", e
	} else {
		return out, nil
	}
}
