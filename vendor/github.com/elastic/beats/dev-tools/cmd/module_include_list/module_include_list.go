// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"

	devtools "github.com/elastic/beats/dev-tools/mage"
	"github.com/elastic/beats/licenses"
)

var usageText = `
Usage: module_include_list [flags]
  module_include_list generates a list.go file containing import statements for
  the specified imports and module directories. An import is a directory or
  directory glob containing .go files. A moduleDir is a directory to search
  for modules and datasets.

  Packages without .go files or without an init() method are omitted from the
  generated file. The output file is written to the include/list.go in the
  Beat's root directory by default.
Options:
`[1:]

var (
	license    string
	pkg        string
	outFile    string
	moduleDirs stringSliceFlag
	importDirs stringSliceFlag
)

func init() {
	flag.StringVar(&license, "license", "ASL2", "License header for generated file (ASL2 or Elastic).")
	flag.StringVar(&pkg, "pkg", "include", "Package name.")
	flag.StringVar(&outFile, "out", "include/list.go", "Output file.")
	flag.Var(&moduleDirs, "moduleDir", "Directory to search for modules to include")
	flag.Var(&importDirs, "import", "Directory to include")
	flag.Usage = usageFlag
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	license, err := licenses.Find(license)
	if err != nil {
		log.Fatalf("Invalid license specifier: %v", err)
	}

	if len(moduleDirs) == 0 && len(importDirs) == 0 {
		log.Fatal("At least one -import or -moduleDir must be specified.")
	}

	dirs, err := findModuleAndDatasets()
	if err != nil {
		log.Fatal(err)
	}

	if imports, err := findImports(); err != nil {
		log.Fatal(err)
	} else {
		dirs = append(dirs, imports...)
	}

	// Get the current directories Go import path.
	repo, err := devtools.GetProjectRepoInfo()
	if err != nil {
		log.Fatalf("Failed to determine import path: %v", err)
	}

	// Build import paths.
	var imports []string
	for _, dir := range dirs {
		// Skip packages without an init() function because that cannot register
		// anything as a side-effect of being imported (e.g. filebeat/input/file).
		var foundInitMethod bool
		goFiles, err := filepath.Glob(filepath.Join(dir, "*.go"))
		if err != nil {
			log.Fatalf("Failed checking for .go files in package dir: %v", err)
		}
		for _, f := range goFiles {
			// Skip test files
			if strings.HasSuffix(f, "_test.go") {
				continue
			}
			if hasInitMethod(f) {
				foundInitMethod = true
				break
			}
		}
		if !foundInitMethod {
			continue
		}

		importDir := dir
		if filepath.IsAbs(dir) {
			// Make it relative to the current package if it's absolute.
			importDir, err = filepath.Rel(devtools.CWD(), dir)
			if err != nil {
				log.Fatalf("Failure creating import for dir=%v: %v", dir, err)
			}
		}

		imports = append(imports, filepath.ToSlash(
			filepath.Join(repo.ImportPath, importDir)))
	}

	sort.Strings(imports)

	// Populate the template.
	var buf bytes.Buffer
	err = Template.Execute(&buf, Data{
		License: license,
		Package: pkg,
		Imports: imports,
	})
	if err != nil {
		log.Fatal("Failed executing template: %v", err)
	}

	// Create the output directory.
	if err = os.MkdirAll(filepath.Dir(outFile), 0755); err != nil {
		log.Fatal("Failed to create output directory: %v", err)
	}

	// Write the output file.
	if err = ioutil.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
		log.Fatal("Failed writing output file: %v", err)
	}
}

func usageFlag() {
	fmt.Fprintf(os.Stderr, usageText)
	flag.PrintDefaults()
}

var Template = template.Must(template.New("normalizations").Funcs(map[string]interface{}{
	"trim": strings.TrimSpace,
}).Parse(`
{{ .License | trim }}

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

package {{ .Package }}

import (
	// Import packages that need to register themselves.
{{- range $import := .Imports }}
	_ "{{ $import }}"
{{- end }}
)
`[1:]))

type Data struct {
	License string
	Package string
	Imports []string
}

//stringSliceFlag is a flag type that allows more than one value to be specified.
type stringSliceFlag []string

func (f *stringSliceFlag) String() string { return strings.Join(*f, ", ") }

func (f *stringSliceFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

// findModuleAndDatasets searches the specified moduleDirs for packages that
// should be imported. They are designated by the presence of a _meta dir.
func findModuleAndDatasets() ([]string, error) {
	var dirs []string
	for _, moduleDir := range moduleDirs {
		// Find modules and datasets as indicated by the _meta dir.
		metaDirs, err := devtools.FindFiles(
			filepath.Join(moduleDir, "*/_meta"),
			filepath.Join(moduleDir, "*/*/_meta"),
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed finding modules and datasets")
		}

		for _, metaDir := range metaDirs {
			// Strip off _meta.
			dirs = append(dirs, filepath.Dir(metaDir))
		}
	}
	return dirs, nil
}

// findImports expands the given import values in case they contain globs.
func findImports() ([]string, error) {
	return devtools.FindFiles(importDirs...)
}

// hasInitMethod returns true if the file contains 'func init()'.
func hasInitMethod(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("failed to read from %v: %v", file, err)
	}
	defer f.Close()

	var initSignature = []byte("func init()")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), initSignature) {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("failed scanning %v: %v", file, err)
	}
	return false
}
