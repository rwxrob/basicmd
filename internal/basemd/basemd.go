package basemd

import (
	"github.com/rwxrob/scan"
)

func MatchWS(s *scan.R) int {
	switch {
	case s.Is("\r"):
		return 1
	case s.Is("\n"):
		return 1
	case s.Is(" "):
		return 1
	default:
		return -1
	}
}

func MatchEOL(s *scan.R) int {
	switch {
	case s.Peek("\r\n"):
		return 2
	case s.Peek("\n"):
		return 1
	default:
		return -1
	}
}

func MatchEOF(s *scan.R) int {
	if s.P == len(s.B) {
		return 0
	}
	return -1
}

func MatchEOB(s *scan.R) int {
	var b int
	m := s.P
TOP:
	{
		if r := MatchEOL(s); r >= 0 {
			s.P += r
			if rr := MatchEOL(s); rr >= 0 {
				s.P = m
				return r + rr
			}
			b += r
			s.P += r
			goto TOP
		}
		if r := MatchEOF(s); r >= 0 {
			s.P = m
			return b
		}
		if r := MatchWS(s); r >= 0 {
			b += r
			s.P += r
			goto TOP
		}
	}
	s.P = m
	return -1
}
