package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"
)

var (
	filePath = flag.String("file", "", "path to local JSON file containing disposable domains")
	urlPath  = flag.String("url", "", "path to remote JSON file containing disposable domains")
	outPath  = flag.String("out", "", "output file name where result is written")
	outPkg   = flag.String("pkg", "disposable", "package name of resulting Go file")
	varName  = flag.String("var", "disposable", "name of generated array variable")
)

const outputTemplate = `// Code generated; DO NOT EDIT.
// {{ .Timestamp }}

package {{ .PackageName }}

var {{ .Name }} = []string{
  {{- range $i, $item := .List }}
  "{{ $item }}",
  {{- end }}
}
`

var t = template.Must(template.New("output").Parse(outputTemplate))

func main() {
	flag.Parse()
	if *filePath != "" && *urlPath != "" {
		printerr(errors.New("can not specify both file path and remote url"))
	}
	if *filePath != "" {
		parseFile(*filePath)
	} else if *urlPath != "" {
		parseURL(*urlPath)
	} else {
		printerr(errors.New("neither file path or url specified"))
	}
}

func printerr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func parseFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		printerr(err)
	}
	var disposable []string
	if err = json.Unmarshal(data, &disposable); err != nil {
		printerr(err)
	}
	output(disposable)
}

func parseURL(url string) {
	resp, err := http.Get(*urlPath)
	if err != nil {
		printerr(err)
	}
	defer resp.Body.Close()
	var disposable []string
	err = json.NewDecoder(resp.Body).Decode(&disposable)
	if err != nil {
		printerr(err)
	}
	output(disposable)
}

func output(disposable []string) {
	out := os.Stdout
	if *outPath != "" {
		f, err := os.Create(*outPath)
		if err != nil {
			printerr(err)
		}
		out = f
	}
	err := t.Execute(out, struct {
		Timestamp   time.Time
		PackageName string
		Name        string
		List        []string
	}{time.Now().UTC(), *outPkg, *varName, disposable})
	if err != nil {
		printerr(err)
	}
}
