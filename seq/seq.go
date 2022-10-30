package seq

import (
	"fmt"
	"log"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	Esc     = "\x1b"
	GS      = "\x1d"
	LF      = "\x1A"
	HT      = "\x09"
	UL      = "\x1B\x2D\x32"
	ULOff   = Esc + "\x2D\x00"
	Cut     = GS + "VA\x03"
	Left    = Esc + "\x61\x00"
	Center  = Esc + "\x61\x01"
	Right   = Esc + "\x61\x02"
	Default = Esc + "\x40"
	FS      = Esc + "\x1C"
)

func CharSize(n int) string {
	return fmt.Sprintf(GS+"\x21%v", n)
}

func ChnEncode(str String) String {
	chn, err := simplifiedchinese.GB18030.NewEncoder().String(str)
	if err != nil {
		log.Println(err)
	}
	return chn
}
