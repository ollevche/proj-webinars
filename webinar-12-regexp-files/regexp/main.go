package main

import (
	"fmt"
	"regexp"
)

func main() {
	var matched bool

	matched, _ = regexp.MatchString(`hello`, "hello hello")
	fmt.Println(matched)

	re := regexp.MustCompile(`hello`)
	res := re.FindAllString(`hello hello he`, -1)
	fmt.Println(res)

	matched, _ = regexp.MatchString(`^hello$`, "hello ")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`^hello`, "hello ")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`hello$`, "hello ")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`\bcat`, "cat caterpillar")

	re = regexp.MustCompile(`\bcat`)
	res = re.FindAllString(`cat caterpillar`, -1)
	fmt.Println(res)

	re = regexp.MustCompile(`\Bcat`)
	res = re.FindAllString(`bobcat caterpillar`, -1)
	fmt.Println(res)

	matched, _ = regexp.MatchString(`\d`, "9")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`^GO\s\d\.\d\d$`, "GO 1.20")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`a+b=c`, "a+b=c")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`cats*`, "cat")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`^A{1}G{1,3}A{1,}!{0,2}$`, "AGGAA!!")
	fmt.Println(matched)

	re, _ = regexp.Compile(`<.*?>`)
	res = re.FindAllString("<p><b>Golang</b> <i>VS</i> <b>Python</b></p>", -1)
	fmt.Println(res)

	matched, _ = regexp.MatchString(`hello|bye`, "bye")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`[be]ye`, "eye")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`^[haHA]+$`, "hahah")
	fmt.Println(matched)

	matched, _ = regexp.MatchString(`^[haHA]+|[goGO]*$`, "")
	fmt.Println(matched)

	re = regexp.MustCompile(`[goGo]*\s(1\.\d+)`)
	ress := re.FindAllStringSubmatch("Go 1.20", -1)
	fmt.Println(ress)

	re, _ = regexp.Compile(`(\d{4})-(\d{2})-(\d{2})`)
	ress = re.FindAllStringSubmatch("Now is 2021-01-14", -1)
	fmt.Println(ress)
}
