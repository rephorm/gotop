package termui

import (
	"image"
	"strings"
	"unicode/utf8"

	"github.com/cjbassi/gotop/src/utils"
	. "github.com/gizak/termui/v3"
	rw "github.com/mattn/go-runewidth"
)

const (
	ELLIPSIS = "…"
	CURSOR   = " "
)

/*
Entry implements a text entry widget.

|----------------------|
Label: Value
Label: Long value trunc…
Label: [               ]   First cell reserved for scroll indicator
         ^
Label: [ Short value   ]   Cursor moves along while typing.
                    ^
Label: […lue truncated ]   Final cell can contain cursor when at end of value.
                      ^

Label: […lue truncated ]   Moving cursor back
                   ^
Label: […lue truncated ]   Moving cursor back more
         ^
Label: […ong value tru…]   Moving cursor back even more scrolls.
         ^
Label: […ong value tru…]   Moving cursor forward again maintains croll.
                     ^

Label: […ng value trun…]   Moving cursor further forward scrolls.
                     ^
Label: […alue truncate…]   Ditto
                     ^
Label: […lue truncated ]   Ditto, but indictor removed when entire value visible
                     ^
Label: […lue truncated ]   When at end of value, cursore moves into final cell.
                      ^
*/

type Entry struct {
	Block

	Style Style

	Label          string
	Value          []rune
	ShowWhenEmpty  bool
	UpdateCallback func(string)

	scroll  utils.ScrollState
	editing bool
}

func (self *Entry) SetEditing(editing bool) {
	self.editing = editing
	if editing {
		w := self.Dx() - rw.StringWidth(self.Label) - 5
		self.scroll = utils.ScrollState{WindowWidth: w}.Move(self.Value, len(self.Value))
	}
}

func (self *Entry) Insert(r rune) {
	self.Value = append(self.Value, 0)
	i := self.scroll.CursorPos
	copy(self.Value[i+1:], self.Value[i:])
	self.Value[i] = r
	self.scroll = self.scroll.Move(self.Value, 1)
}

func (self *Entry) Backspace() {
	i := self.scroll.CursorPos
	if i == 0 {
		return
	}
	copy(self.Value[i-1:], self.Value[i:])
	self.Value = self.Value[:len(self.Value)-1]
	self.scroll = self.scroll.Move(self.Value, -1)
}

func (self *Entry) Delete() {
	i := self.scroll.CursorPos
	if i >= len(self.Value) {
		return
	}
	copy(self.Value[i:], self.Value[i+1:])
	self.Value = self.Value[:len(self.Value)-1]
	self.scroll = self.scroll.Move(self.Value, 0)
}

func (self *Entry) update() {
	if self.UpdateCallback != nil {
		self.UpdateCallback(string(self.Value))
	}
}

// HandleEvent handles input events if the entry is being edited.
// Returns true if the event was handled.
func (self *Entry) HandleEvent(e Event) bool {
	if !self.editing {
		return false
	}
	if utf8.RuneCountInString(e.ID) == 1 {
		self.Insert([]rune(e.ID)[0])
		self.update()
		return true
	}
	switch e.ID {
	case "<C-c>", "<Escape>":
		self.Value = []rune{}
		self.editing = false
		self.update()
	case "<Enter>":
		self.editing = false
	case "<Backspace>":
		self.Backspace()
		self.update()
	case "<Delete>":
		self.Delete()
		self.update()
	case "<Space>":
		self.Insert(' ')
		self.update()
	case "<Left>":
		self.scroll = self.scroll.Move(self.Value, -1)
	case "<Right>":
		self.scroll = self.scroll.Move(self.Value, 1)
	default:
		return false
	}
	return true
}

func (self *Entry) Draw(buf *Buffer) {
	if len(self.Value) == 0 && !self.editing && !self.ShowWhenEmpty {
		return
	}

	style := self.Style
	label := self.Label
	if self.editing {
		label += "["
		style = NewStyle(style.Fg, style.Bg, ModifierBold)
	}
	cursorStyle := NewStyle(style.Bg, style.Fg, ModifierClear)

	p := image.Pt(self.Min.X, self.Min.Y)
	buf.SetString(label, style, p)
	p.X += rw.StringWidth(label)

	if !self.editing {
		val := utils.TruncateFront(string(self.Value), self.Max.X-p.X-1, ELLIPSIS)
		buf.SetString(val, style, p)
		p.X += 1
		buf.SetString(" ", style, p)
		return
	}

	tail := "] "

	indicator := " "
	if self.scroll.WindowPos > 0 {
		indicator = ELLIPSIS
	}
	buf.SetString(indicator, self.Style, p)
	p.X += rw.StringWidth(indicator)

	preCursor := self.scroll.PreCursor(self.Value)
	buf.SetString(preCursor, self.Style, p)
	p.X += rw.StringWidth(preCursor)

	cursor := CURSOR
	if self.scroll.CursorPos < len(self.Value) {
		cursor = string(self.Value[self.scroll.CursorPos])
	}
	buf.SetString(cursor, cursorStyle, p)
	p.X += rw.StringWidth(cursor)

	postCursor, trunc := self.scroll.PostCursor(self.Value)
	buf.SetString(postCursor, self.Style, p)
	p.X += rw.StringWidth(postCursor)

	if trunc {
		buf.SetString(ELLIPSIS, self.Style, p)
		p.X += rw.StringWidth(ELLIPSIS)
	} else if remaining := self.Max.X - p.X - rw.StringWidth(tail); remaining > 0 {
		buf.SetString(strings.Repeat(" ", remaining), self.Style, p)
		p.X += remaining
	}
	buf.SetString(tail, style, p)
}
