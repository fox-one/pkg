package qrcode

import (
	"io"
	"os"

	"github.com/mdp/qrterminal"
)

func Fprint(out io.Writer, content string) {
	qrterminal.GenerateHalfBlock(content, qrterminal.H, out)
}

func Print(content string) {
	Fprint(os.Stdout, content)
}
