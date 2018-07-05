package colors

import "fmt"

const ANSI_ESC = "\x1b["

const ANSI_BOLD = "1;"

const ANSI_0 = ANSI_ESC + "30;1m"
const ANSI_R = ANSI_ESC + "31;1m"
const ANSI_G = ANSI_ESC + "32;1m"
const ANSI_Y = ANSI_ESC + "33;1m"
const ANSI_B = ANSI_ESC + "34;1m"
const ANSI_M = ANSI_ESC + "35;1m"
const ANSI_C = ANSI_ESC + "36;1m"
const ANSI_W = ANSI_ESC + "37;1m"
const ANSI_N = ANSI_ESC + "0m"

/*
_Black   = "\x1b[30;1m"
_Red     = "\x1b[31;1m"
_Green   = "\x1b[32;1m"
_Yellow  = "\x1b[33;1m"
_Blue    = "\x1b[34;1m"
_Magenta = "\x1b[35;1m"
_Cyan    = "\x1b[36;1m"
_White   = "\x1b[37;1m"
_?N???   = "\x1b[0m"
*/

func color(i int) string {
	if i < 0 || i > 7 {
		i = 7
	}

	i += 30

	return fmt.Sprintf("\x1b[%v;1m", i)
}

type AnsiColorString struct {
	color int
	text  string
}

func NewAnsiColorString(text string) *AnsiColorString {
	return &AnsiColorString{
		-1,
		text,
	}
}

func (w *AnsiColorString) String() string {
	return color(w.color) + text + ANSI_N
}

func (w *AnsiColorString) GoString() string {
	return w.String()
}

func (w *AnsiColorString) Black() *AnsiColorString {
	w.color = 0
	return w
}

func (w *AnsiColorString) Red() *AnsiColorString {
	w.color = 1
	return w
}

func (w *AnsiColorString) Green() *AnsiColorString {
	w.color = 2
	return w
}

func (w *AnsiColorString) Yellow() *AnsiColorString {
	w.color = 3
	return w
}

func (w *AnsiColorString) Blue() *AnsiColorString {
	w.color = 4
	return w
}

func (w *AnsiColorString) Magenta() *AnsiColorString {
	w.color = 5
	return w
}

func (w *AnsiColorString) Cyan() *AnsiColorString {
	w.color = 6
	return w
}

func (w *AnsiColorString) Gray() *AnsiColorString {
	w.color = 7
	return w
}
