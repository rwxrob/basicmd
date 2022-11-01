package basemd_test

import (
	"fmt"

	"github.com/rwxrob/basicmd/internal/basemd"
	"github.com/rwxrob/scan"
)

func ExampleScanWS() {
	s := new(scan.R)
	//s.Trace++
	s.B = []byte(`1 `)
	fmt.Println(basemd.ScanWS(s))
	s.Print() // nothing advanced at all
	s.Scan()
	s.Print()
	fmt.Println(basemd.ScanWS(s))
	s.Print()
	// Output:
	// false
	// 0 '\x00' "1 "
	// 1 '1' " "
	// true
	// 2 ' ' ""
}

func ExampleScanEOL_carriage_Feed() {
	s := new(scan.R)
	s.B = []byte("\r\n\r\n")
	s.Print()
	fmt.Println(basemd.ScanEOL(s))
	s.Print()
	// Output:
	// 0 '\x00' "\r\n\r\n"
	// true
	// 2 '\n' "\r\n"
}

func ExampleScanEOL_feed() {
	s := new(scan.R)
	s.B = []byte("\n")
	s.Print()
	fmt.Println(basemd.ScanEOL(s))
	s.Print()
	// Output:
	// 0 '\x00' "\n"
	// true
	// 1 '\n' ""
}

func ExampleScanEOL_carriage_Not_Enough() {
	s := new(scan.R)
	s.B = []byte("\r")
	s.Print()
	fmt.Println(basemd.ScanEOL(s))
	s.Print()
	// Output:
	// 0 '\x00' "\r"
	// false
	// 0 '\x00' "\r"
}

func ExampleScanEOB_line_Returns() {
	s := new(scan.R)
	s.B = []byte("\n\n")
	fmt.Println(basemd.ScanEOB(s))
	s.Print()
	// Output:
	// true
	// 2 '\n' ""
}

func ExampleScanEOB_carriage_and_Line_Returns() {
	s := new(scan.R)
	s.B = []byte("\r\n\r\n")
	fmt.Println(basemd.ScanEOB(s))
	s.Print()
	// Output:
	// true
	// 4 '\n' ""
}

func ExampleScanEOB_odd_Returns() {
	s := new(scan.R)
	s.B = []byte("\r\n\n")
	fmt.Println(basemd.ScanEOB(s))
	s.Print()
	// Output:
	// true
	// 3 '\n' ""
}

func ExampleScanEOB_extra_WS() {
	s := new(scan.R)
	s.B = []byte("   \r\n\r\n\r\n")
	fmt.Println(basemd.ScanEOB(s))
	s.Print()
	// Output:
	// true
	// 9 '\n' ""
}
