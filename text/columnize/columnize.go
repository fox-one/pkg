package columnize

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

const tab = "\t"

func formatLine(line []string) string {
	if len(line) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, cell := range line {
		trimmed := strings.TrimSpace(cell)
		builder.WriteString(trimmed)
		builder.WriteString(tab)
	}

	return builder.String()
}

func Fprint(out io.Writer, lines [][]string) error {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)

	for idx := range lines {
		if _, err := fmt.Fprintln(w, formatLine(lines[idx])); err != nil {
			return err
		}
	}

	return w.Flush()
}

func Print(lines [][]string) error {
	return Fprint(os.Stdout, lines)
}

type Form struct {
	lines [][]string
}

func (f *Form) Append(line ...string) {
	f.lines = append(f.lines, line)
}

func (f *Form) Fprint(out io.Writer) error {
	return Fprint(out, f.lines)
}

func (f *Form) Print() error {
	return Print(f.lines)
}

func (f *Form) String() string {
	var buff bytes.Buffer
	_ = f.Fprint(&buff)
	return buff.String()
}

func (f *Form) WriteCSV(out io.Writer) error {
	w := csv.NewWriter(out)
	return w.WriteAll(f.lines)
}
