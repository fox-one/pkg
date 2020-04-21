package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

var payload struct {
	Tag        string
	ConfigFile string
	GenFile    string
	Package    string
	Data       string
}

func main() {
	flag.StringVar(&payload.Tag, "tag", "", "build tag")
	flag.StringVar(&payload.ConfigFile, "config", "", "config file path")
	flag.StringVar(&payload.GenFile, "out", "config_gen.go", "gen config file path")
	flag.StringVar(&payload.Package, "package", "main", "package name")
	flag.Parse()

	f, err := os.Open(payload.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := &bytes.Buffer{}
	r := io.TeeReader(f, buf)

	if err := validateYaml(r); err != nil {
		log.Fatal(err)
	}

	payload.Data = base64.StdEncoding.EncodeToString(buf.Bytes())

	out, err := os.Create(payload.GenFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	if err := template.Must(
		template.New("_").Parse(templ),
	).Execute(out, payload); err != nil {
		log.Fatal(err)
	}
}

func validateYaml(r io.Reader) error {
	m := make(map[interface{}]interface{})
	return yaml.NewDecoder(r).Decode(&m)
}

const templ = `
{{- if .Tag -}}
// +build {{.Tag}}
{{- end}}

package {{.Package}}

import "github.com/fox-one/pkg/config"

func init() {
	config.DATA = "{{.Data}}"
}
`
