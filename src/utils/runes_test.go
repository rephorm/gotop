package utils

import "testing"

const (
	ELLIPSIS = "…"
)

func TestTruncateFront(t *testing.T) {
	tests := []struct {
		s      string
		w      int
		prefix string
		want   string
	}{
		{"", 0, ELLIPSIS, ""},
		{"", 1, ELLIPSIS, ""},
		{"", 10, ELLIPSIS, ""},

		{"abcdef", 0, ELLIPSIS, ELLIPSIS},
		{"abcdef", 1, ELLIPSIS, ELLIPSIS},
		{"abcdef", 2, ELLIPSIS, ELLIPSIS + "f"},
		{"abcdef", 5, ELLIPSIS, ELLIPSIS + "cdef"},
		{"abcdef", 6, ELLIPSIS, "abcdef"},
		{"abcdef", 10, ELLIPSIS, "abcdef"},

		{"abcdef", 0, "...", "..."},
		{"abcdef", 1, "...", "..."},
		{"abcdef", 3, "...", "..."},
		{"abcdef", 4, "...", "...f"},
		{"abcdef", 5, "...", "...ef"},
		{"abcdef", 6, "...", "abcdef"},
		{"abcdef", 10, "...", "abcdef"},

		{"｟full～width｠", 15, ".", "｟full～width｠"},
		{"｟full～width｠", 14, ".", ".full～width｠"},
		{"｟full～width｠", 13, ".", ".ull～width｠"},
		{"｟full～width｠", 10, ".", ".～width｠"},
		{"｟full～width｠", 9, ".", ".width｠"},
		{"｟full～width｠", 8, ".", ".width｠"},
		{"｟full～width｠", 3, ".", ".｠"},
		{"｟full～width｠", 2, ".", "."},
	}

	for _, test := range tests {
		if got := TruncateFront(test.s, test.w, test.prefix); got != test.want {
			t.Errorf("TruncateFront(%q, %d, %q) = %q; want %q", test.s, test.w, test.prefix, got, test.want)
		}
	}
}

func TestScrollStateMove(t *testing.T) {
	tests := []struct {
		name   string
		in     ScrollState
		runes  string
		offset int
		want   ScrollState
	}{
		{
			name:   "empty",
			in:     ScrollState{0, 0, 5},
			runes:  "",
			offset: 1,
			want:   ScrollState{0, 0, 5},
		},
		{
			name:   "scroll middle, move right",
			in:     ScrollState{3, 0, 4},
			runes:  "abcdef",
			offset: 1,
			want:   ScrollState{4, 1, 4},
		},
		{
			name:   "scroll end, move right",
			in:     ScrollState{3, 1, 4},
			runes:  "abcdef",
			offset: 1,
			want:   ScrollState{4, 1, 4},
		},
		{
			name:   "scroll end, move to end",
			in:     ScrollState{5, 2, 4},
			runes:  "abcdef",
			offset: 1,
			want:   ScrollState{6, 2, 4},
		},
		{
			name:   "move past end does nothing",
			in:     ScrollState{6, 2, 4},
			runes:  "abcdef",
			offset: 1,
			want:   ScrollState{6, 2, 4},
		},
		{
			name:   "move back in window",
			in:     ScrollState{3, 1, 4},
			runes:  "abcdef",
			offset: -1,
			want:   ScrollState{2, 1, 4},
		},
		{
			name:   "move back scroll back",
			in:     ScrollState{2, 2, 4},
			runes:  "abcdef",
			offset: -1,
			want:   ScrollState{1, 1, 4},
		},
		{
			name:   "backspace from end",
			in:     ScrollState{6, 4, 4},
			runes:  "abcdef",
			offset: -1,
			want:   ScrollState{5, 2, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.in.Move([]rune(test.runes), test.offset); got != test.want {
				t.Errorf("(%v).Move(%q, %d) = %v; want %v", test.in, test.runes, test.offset, got, test.want)
			}
		})
	}
}
