package abnf

import (
	"testing"
	"sort"
	"strings"
)



// RFC 5234 errata 2968 3076 applied
var ABNF_ABNF_5234 string = `rulelist = 1*(rule / (*WSP c-nl))
rule = rulename defined-as elements c-nl
rulename = ALPHA *(ALPHA / DIGIT / "-")
defined-as = *c-wsp ("=" / "=/") *c-wsp
elements = alternation *WSP
c-wsp = WSP / (c-nl WSP)
c-nl = comment / CRLF
comment = ";" *(WSP / VCHAR) CRLF
alternation = concatenation *(*c-wsp "/" *c-wsp concatenation)
concatenation = repetition *(1*c-wsp repetition)
repetition = [ repeat ] element
repeat = 1*DIGIT / (*DIGIT "*" *DIGIT)
element = rulename / group / option / char-val / num-val / prose-val
group = "(" *c-wsp alternation *c-wsp ")"
option = "[" *c-wsp alternation *c-wsp "]"
char-val = DQUOTE *(%x20-21 / %x23-7E) DQUOTE
num-val = "%" (bin-val / dec-val / hex-val)
bin-val = "b" 1*BIT [ 1*("." 1*BIT) / ("-" 1*BIT) ]
dec-val = "d" 1*DIGIT [ 1*("." 1*DIGIT) / ("-" 1*DIGIT) ]
hex-val = "x" 1*HEXDIG [ 1*("." 1*HEXDIG) / ("-" 1*HEXDIG) ]
prose-val = "<" *(%x20-3D / %x3F-7E) ">"
ALPHA = %x41-5A / %x61-7A
BIT = "0" / "1"
CHAR = %x01-7F
CR = %x0D
CRLF = CR LF
CTL = %x00-1F / %x7F
DIGIT = %x30-39
DQUOTE = %x22
HEXDIG = DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
HTAB = %x09
LF = %x0A
LWSP = *(WSP / CRLF WSP)
OCTET = %x00-FF
SP = %x20
VCHAR = %x21-7E
WSP = SP / HTAB`

var ABNF_ABNF_7405 string = `rulelist = 1*(rule / (*WSP c-nl))
rule = rulename defined-as elements c-nl
rulename = ALPHA *(ALPHA / DIGIT / "-")
defined-as = *c-wsp ("=" / "=/") *c-wsp
elements = alternation *WSP
c-wsp = WSP / (c-nl WSP)
c-nl = comment / CRLF
comment = ";" *(WSP / VCHAR) CRLF
alternation = concatenation *(*c-wsp "/" *c-wsp concatenation)
concatenation = repetition *(1*c-wsp repetition)
repetition = [ repeat ] element
repeat = 1*DIGIT / (*DIGIT "*" *DIGIT)
element = rulename / group / option / char-val / num-val / prose-val
group = "(" *c-wsp alternation *c-wsp ")"
option = "[" *c-wsp alternation *c-wsp "]"
char-val = case-insensitive-string / case-sensitive-string
case-insensitive-string = [ "%i" ] quoted-string
case-sensitive-string = "%s" quoted-string
quoted-string = DQUOTE *(%x20-21 / %x23-7E) DQUOTE
num-val = "%" (bin-val / dec-val / hex-val)
bin-val = "b" 1*BIT [ 1*("." 1*BIT) / ("-" 1*BIT) ]
dec-val = "d" 1*DIGIT [ 1*("." 1*DIGIT) / ("-" 1*DIGIT) ]
hex-val = "x" 1*HEXDIG [ 1*("." 1*HEXDIG) / ("-" 1*HEXDIG) ]
prose-val = "<" *(%x20-3D / %x3F-7E) ">"
ALPHA = %x41-5A / %x61-7A
BIT = "0" / "1"
CHAR = %x01-7F
CR = %x0D
CRLF = CR LF
CTL = %x00-1F / %x7F
DIGIT = %x30-39
DQUOTE = %x22
HEXDIG = DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
HTAB = %x09
LF = %x0A
LWSP = *(WSP / CRLF WSP)
OCTET = %x00-FF
SP = %x20
VCHAR = %x21-7E
WSP = SP / HTAB`



func Test_ABNF_5234(t *testing.T) {
	ret := strings.Split(ABNF_ABNF_5234,"\n")
	sort.Strings(ret)
	sorted_ABNF_ABNF := strings.Join(ret,"\r\n")

	tested_ABNF := ABNF5234().String()
	if tested_ABNF != sorted_ABNF_ABNF {
		t.Errorf("difference between \n---------\n%s\n---------\nand\n---------\n%s\n---------\n ", sorted_ABNF_ABNF, tested_ABNF )

	}
}

func Test_ABNF_ABNF_5234(t *testing.T) {
	ret := strings.Split(ABNF_ABNF_5234,"\n")
	sort.Strings(ret)
	sorted_ABNF_ABNF := []byte(strings.Join(ret,"\r\n"))

	tested_ABNF := ABNF5234().Valid(sorted_ABNF_ABNF)
	if !tested_ABNF {
		t.Errorf("errors found in \n---------\n%s\n---------\n", ABNF_ABNF_5234)
	}
}

func Test_ABNF_7405(t *testing.T) {
	ret := strings.Split(ABNF_ABNF_7405,"\n")
	sort.Strings(ret)
	sorted_ABNF_ABNF := strings.Join(ret,"\r\n")

	tested_ABNF := ABNF7405().String()
	if tested_ABNF != sorted_ABNF_ABNF {
		t.Errorf("difference between \n---------\n%s\n---------\nand\n---------\n%s\n---------\n ", sorted_ABNF_ABNF, tested_ABNF )

	}
}

func Test_ABNF_ABNF_7405(t *testing.T) {
	ret := strings.Split(ABNF_ABNF_7405,"\n")
	sort.Strings(ret)
	sorted_ABNF_ABNF := []byte(strings.Join(ret,"\r\n"))

	tested_ABNF := ABNF7405().Valid(sorted_ABNF_ABNF)
	if !tested_ABNF {
			t.Errorf("errors found in \n---------\n%s\n---------\n", ABNF_ABNF_7405)
	}
}


func Benchmark_ABNF5234_call(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ABNF5234()
	}
}

func Benchmark_ABNF5234_String(b *testing.B) {
	abnf := ABNF5234()
	for i := 0; i < b.N; i++ {
		abnf.String()
	}
}


func Benchmark_ABNF5234_Valid(b *testing.B) {
	ret := strings.Split(ABNF_ABNF_5234,"\n")
	sort.Strings(ret)
	sorted_ABNF_ABNF := []byte(strings.Join(ret,"\r\n"))

	abnf := ABNF5234()
	for i := 0; i < b.N; i++ {
		abnf.Valid(sorted_ABNF_ABNF)
	}
}
