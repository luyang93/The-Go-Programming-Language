// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"gitlab.com/luyang93/The-Go-Programming-Language/ch04/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func noMust() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/*
75 issues:
----------------------------------------
Number: 56733
User:   rolandshoemaker
Title:  encoding/json: add (*Decoder).SetLimit
Age:    95 days
----------------------------------------
Number: 48298
User:   dsnet
Title:  encoding/json: add Decoder.DisallowDuplicateFields
Age:    526 days
----------------------------------------
Number: 36225
User:   dsnet
Title:  encoding/json: the Decoder.Decode API lends itself to misuse
Age:    1156 days
----------------------------------------
Number: 42571
User:   dsnet
Title:  encoding/json: clarify Decoder.InputOffset semantics
Age:    827 days
----------------------------------------
Number: 11046
User:   kurin
Title:  encoding/json: Decoder internally buffers full input
Age:    2816 days
----------------------------------------
Number: 56332
User:   gansvv
Title:  encoding/json: clearer error message for boolean like prefix in
Age:    121 days
----------------------------------------
Number: 43716
User:   ggaaooppeenngg
Title:  encoding/json: increment byte counter when using decoder.Token
Age:    763 days
----------------------------------------
Number: 29035
User:   jaswdr
Title:  proposal: encoding/json: add error var to compare  the returned
Age:    1540 days
----------------------------------------
Number: 28923
User:   mvdan
Title:  encoding/json: speed up the decoding scanner
Age:    1548 days
----------------------------------------
Number: 34543
User:   maxatome
Title:  encoding/json: Unmarshal & json.(*Decoder).Token report differen
Age:    1241 days
----------------------------------------
Number: 48950
User:   AlexanderYastrebov
Title:  encoding/json: calculate correct SyntaxError.Offset in the strea
Age:    492 days
----------------------------------------
Number: 14750
User:   cyberphone
Title:  encoding/json: parser ignores the case of member names
Age:    2535 days
----------------------------------------
Number: 32779
User:   rsc
Title:  encoding/json: memoize strings during decode
Age:    1333 days
----------------------------------------
Number: 31701
User:   lr1980
Title:  encoding/json: second decode after error impossible
Age:    1393 days
----------------------------------------
Number: 40127
User:   rogpeppe
Title:  encoding/json: add Encoder.EncodeToken method
Age:    953 days
----------------------------------------
Number: 40128
User:   rogpeppe
Title:  proposal: encoding/json: garbage-free reading of tokens
Age:    953 days
----------------------------------------
Number: 33854
User:   Qhesz
Title:  encoding/json: unmarshal option to treat omitted fields as null
Age:    1271 days
----------------------------------------
Number: 40982
User:   Segflow
Title:  encoding/json: use different error type for unknown field if the
Age:    909 days
----------------------------------------
Number: 5901
User:   rsc
Title:  encoding/json: allow per-Encoder/per-Decoder registration of mar
Age:    3502 days
----------------------------------------
Number: 16212
User:   josharian
Title:  encoding/json: do all reflect work before decoding
Age:    2424 days
----------------------------------------
Number: 6647
User:   btracey
Title:  x/tools/cmd/godoc: display type kind of each named type
Age:    3404 days
----------------------------------------
Number: 41144
User:   alvaroaleman
Title:  encoding/json: Unmarshaler breaks DisallowUnknownFields
Age:    900 days
----------------------------------------
Number: 43513
User:   AlexanderYastrebov
Title:  encoding/json: add line number to SyntaxError
Age:    773 days
----------------------------------------
Number: 48277
User:   Windsooon
Title:  encoding/json: add an example for InputOffset() function
Age:    526 days
----------------------------------------
Number: 34564
User:   mdempsky
Title:  go/internal/gcimporter: single source of truth for decoder logic
Age:    1240 days
----------------------------------------
Number: 26946
User:   deuill
Title:  encoding/json: clarify what happens when unmarshaling into a non
Age:    1650 days
----------------------------------------
Number: 27735
User:   dsnet
Title:  encoding/json: incorrect usage of sync.Pool
Age:    1613 days
----------------------------------------
Number: 53178
User:   AsterDY
Title:  proposal: encoding/json: more performant implementation with JIT
Age:    261 days
----------------------------------------
Number: 56299
User:   dop251
Title:  encoding/json: Excessive allocations when using the Token() API
Age:    122 days
----------------------------------------
Number: 22752
User:   buyology
Title:  proposal: encoding/json: add access to the underlying data causi
Age:    1920 days
*/
