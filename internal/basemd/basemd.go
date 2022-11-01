package basemd

import (
	"github.com/rwxrob/scan"
)

func ScanWS(s *scan.R) bool {
	b, r, lp := s.P, s.R, s.LP
	var found bool
	for s.Scan() {
		switch s.R {
		case ' ', '\r', '\n':
			found = true
			continue
		}
		break
	}
	if !found {
		s.P, s.R, s.LP = b, r, lp
	}
	return found
}

func ScanEOL(s *scan.R) bool {
	b, r, lp := s.P, s.R, s.LP
	var found bool
	s.Scan()
	switch s.R {
	case '\n':
		found = true
	case '\r':
		if s.Scan() && s.R == '\n' {
			found = true
		}
	}
	if !found {
		s.P, s.R, s.LP = b, r, lp
	}
	return found
}

func ScanEOB(s *scan.R) bool {
	b, r, lp := s.P, s.R, s.LP
	var found bool
TOP:
	{
		switch {
		case s.End():
			found = true
		case ScanEOL(s) && ScanEOL(s):
			found = true
			goto TOP
		case ScanEOL(s) && s.End():
			found = true
		case ScanWS(s):
			goto TOP
		}
	}
	if !found {
		s.P, s.R, s.LP = b, r, lp
	}
	return found
}
