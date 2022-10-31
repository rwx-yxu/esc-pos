package escpos

import (
	"log"
	"os"

	"github.com/rwx-yxu/esc-pos/seq"
)

//Run will create a data.txt containing escpos sequences to be ran on
//a escpos printer. The file can then be printed off using cups lp
//command or manually printed from gui.
func Run() {
	fileName := "data.txt"

	f, err := os.Create(fileName)

	if err != nil {

		log.Fatal(err)
	}

	defer f.Close()
	words := []string{seq.Center, "Hello world", seq.Cut}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
}
