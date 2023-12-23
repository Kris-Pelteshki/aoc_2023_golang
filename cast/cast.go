package cast

// Suite of casting functions to speed up solutions
// This is NOT idiomatic Go... but AOC isn't about that...

import (
	"fmt"
	"strconv"
)

// ToInt will case a given arg into an int type.
// Supported types are:
//   - string
func ToInt(arg interface{}) int {
	switch v := arg.(type) {
	case string:
		var val int
		var err error
		val, err = strconv.Atoi(v)
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
		return val
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
}

func ToInts(args []string) (vals []int) {
	for _, arg := range args {
		vals = append(vals, ToInt(arg))
	}
	return vals
}

// ToString will case a given arg into an int type.
// Supported types are:
//   - int
//   - byte
//   - rune
func ToString(arg interface{}) (str string) {
	switch v := arg.(type) {
	case int:
		str = strconv.Itoa(v)
	case byte:
		str = string(rune(v))
	case rune:
		str = string(v)
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

const (
	ASCIICodeCapA   = int('A') // 65
	ASCIICodeCapZ   = int('Z') // 65
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 97
)

// ToASCIICode returns the ascii code of a given input
func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}

	return asciiVal
}

// ASCIIIntToChar returns a one character string of the given int
func ASCIIIntToChar(code int) string {
	return string(rune(code))
}
