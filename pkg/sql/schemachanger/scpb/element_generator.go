// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

//go:build generator
// +build generator

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"text/template"

	"github.com/cockroachdb/cockroach/pkg/cli/exit"
)

var (
	in  = flag.String("in", "", "input proto file")
	out = flag.String("out", "", "output file for generated go code")
)

func main() {
	flag.Parse()
	if err := run(*in, *out); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		exit.WithCode(exit.FatalError())
	}
}

func run(in, out string) error {
	if out == "" {
		return fmt.Errorf("output required")
	}
	elementNames, err := getElementNames(in)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := template.Must(template.New("templ").Parse(`{{- /**/ -}}
// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by element_generator.go. DO NOT EDIT.

package scpb

type ElementStatusIterator interface {
	ForEachElementStatus(fn func(status, targetStatus Status, elem Element))
}
{{ range . }}

func (e {{ . }}) element() {}

// ForEach{{ . }} iterates over nodes of type {{ . }}.
func ForEach{{ . }}(
	b ElementStatusIterator, elementFunc func(status, targetStatus Status, element *{{ . }}),
) {
	b.ForEachElementStatus(func(status, targetStatus Status, elem Element) {
		if e, ok := elem.(*{{ . }}); ok {
			elementFunc(status, targetStatus, e)
		}
	})
}
{{- end -}}
`)).Execute(&buf, elementNames); err != nil {
		return err
	}
	return ioutil.WriteFile(out, buf.Bytes(), 0777)
}

// getElementNames parses the ElementsProto struct definition and extracts
// the names of the types of its members.
func getElementNames(inProtoFile string) (names []string, _ error) {
	var (
		elementProtoBufMeta = `(\s+\[\([A-z\.]+\)\s+=\s+\"[A-z\:\",\s]+\])?`
		elementFieldPat     = `\s*(?P<type>\w+)\s+(?P<name>\w+)\s+=\s+\d+` +
			elementProtoBufMeta + `;`
		elementProtoRegexp = regexp.MustCompile(`(?s)message ElementProto {
  option \(gogoproto.onlyone\) = true;
(?P<fields>(` + elementFieldPat + "\n)+)" +
			"}",
		)
		elementFieldRegexp  = regexp.MustCompile(elementFieldPat)
		elementFieldTypeIdx = elementFieldRegexp.SubexpIndex("type")
		elementFieldsIdx    = elementProtoRegexp.SubexpIndex("fields")
	)

	got, err := ioutil.ReadFile(inProtoFile)
	if err != nil {
		return nil, err
	}
	submatch := elementProtoRegexp.FindSubmatch(got)
	if submatch == nil {
		return nil, fmt.Errorf("failed to find ElementProto in %s: %s",
			inProtoFile, elementProtoRegexp)
	}
	fieldMatches := elementFieldRegexp.FindAllSubmatch(submatch[elementFieldsIdx], -1)
	for _, m := range fieldMatches {
		names = append(names, string(m[elementFieldTypeIdx]))
	}
	return names, nil
}
