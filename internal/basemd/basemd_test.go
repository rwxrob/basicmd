package basemd_test

import (
	"fmt"

	"github.com/rwxrob/basicmd/internal/basemd"
	"github.com/rwxrob/scan"
)

func ExampleMatchWS() {
	s := new(scan.R)
	s.B = []byte(` `)
	fmt.Println(basemd.MatchWS(s))
	// Output:
	// 1
}

func ExampleMatchEOL() {
	s := new(scan.R)
	s.B = []byte("\r\n")
	fmt.Println(basemd.MatchEOL(s))
	s.P = 0
	s.B = []byte("\n")
	fmt.Println(basemd.MatchEOL(s))
	s.P = 0
	s.B = []byte("\r")
	fmt.Println(basemd.MatchEOL(s))
	// Output:
	// 2
	// 1
	// -1
}

func ExampleMatchEOF() {
	s := new(scan.R)
	s.B = []byte(` `)
	s.Scan()
	fmt.Println(basemd.MatchEOF(s))
	s.P = 0
	s.B = []byte(` `)
	fmt.Println(basemd.MatchEOF(s))
	// Output:
	// 0
	// -1
}

func ExampleMatchEOB_line_Returns() {
	s := new(scan.R)
	s.B = []byte("\n\n")
	fmt.Println(basemd.MatchEOB(s))
	// Output:
	// 2
}

func ExampleMatchEOB_carriage_and_Line_Returns() {
	s := new(scan.R)
	s.B = []byte("\r\n\r\n")
	fmt.Println(basemd.MatchEOB(s))
	// Output:
	// 4
}

func ExampleMatchEOB_odd_Returns() {
	s := new(scan.R)
	s.B = []byte("\r\n\n")
	fmt.Println(basemd.MatchEOB(s))
	// Output:
	// 3
}

func ExampleMatchEOB_extra_WS() {
	s := new(scan.R)
	s.B = []byte("   \r\n\r\n\r\n")
	fmt.Println(basemd.MatchEOB(s))
	// Output:
	// 7
}
