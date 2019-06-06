package utils

import (
	rw "github.com/mattn/go-runewidth"
)

func TruncateFront(s string, w int, prefix string) string {
	if rw.StringWidth(s) <= w {
		return s
	}
	r := []rune(s)
	pw := rw.StringWidth(prefix)
	w -= pw
	width := 0
	i := len(r) - 1
	for ; i >= 0; i-- {
		cw := rw.RuneWidth(r[i])
		width += cw
		if width > w {
			break
		}
	}
	return prefix + string(r[i+1:len(r)])
}

type ScrollState struct {
	CursorPos   int // in runes
	WindowPos   int // in runes
	WindowWidth int // in cells
}

func (s ScrollState) Move(runes []rune, offset int) ScrollState {
	s.CursorPos += offset
	if s.CursorPos < 0 {
		s.CursorPos = 0
	} else if s.CursorPos > len(runes) {
		s.CursorPos = len(runes)
	}

	if s.CursorPos < s.WindowPos {
		s.WindowPos = s.CursorPos
	} else {
		w := 0
		for i := s.WindowPos; i <= s.CursorPos; i++ {
			if i >= len(runes) {
				w += 0
			} else {
				w += rw.RuneWidth(runes[i])
			}
		}
		for w > s.WindowWidth && s.WindowPos < len(runes) {
			w -= rw.RuneWidth(runes[s.WindowPos])
			s.WindowPos += 1
		}
	}

	// right align if needed.
	if w := rw.StringWidth(string(runes[s.WindowPos:])); w < s.WindowWidth {
		for s.WindowPos > 0 {
			w += rw.RuneWidth(runes[s.WindowPos-1])
			if w > s.WindowWidth {
				break
			}
			s.WindowPos -= 1
		}
	}
	return s
}

func (s ScrollState) PreCursor(runes []rune) string {
	return string(runes[s.WindowPos:s.CursorPos])
}

func (s ScrollState) PostCursor(runes []rune) (string, bool) {
	if s.CursorPos >= len(runes)-1 {
		return "", false
	}

	start := s.CursorPos + 1
	end := start
	w := rw.StringWidth(string(runes[s.WindowPos:start]))
	for i := start; i < len(runes); i++ {
		w += rw.RuneWidth(runes[i])
		if w > s.WindowWidth {
			break
		}
		end += 1
	}
	return string(runes[start:end]), end < len(runes)
}
