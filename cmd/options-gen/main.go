package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	optionsgen "github.com/kazhuravlev/options-gen/options-gen"
)

func main() {
	var (
		inFilename        string
		outFilename       string
		optionsStructName string
		outPackageName    string
		defaultsFrom      string
		muteWarnings      bool
	)

	envGoFile := os.Getenv("GOFILE")
	envGoPackage := os.Getenv("GOPACKAGE")

	defaultOutFilename := strings.Replace(filepath.Base(envGoFile), ".go", "_generated.go", 1)

	flag.StringVar(&inFilename,
		"filename", envGoFile,
		"input filename")
	flag.StringVar(&outPackageName,
		"pkg", envGoPackage,
		"output package name")
	flag.StringVar(&outFilename,
		"out-filename", defaultOutFilename,
		"output filename")
	flag.StringVar(&optionsStructName,
		"from-struct", "",
		"struct that contains options")
	flag.StringVar(&defaultsFrom,
		"defaults-from", "tag=default",
		"where to get defaults for options. none, tag=TagName, func=FuncName, var=VarName")
	flag.BoolVar(&muteWarnings,
		"mute-warnings", false,
		"mute all warnings")
	flag.Parse()

	if isEmpty(inFilename, outFilename, outPackageName, optionsStructName, defaultsFrom) {
		flag.Usage()
		//nolint:forbidigo
		fmt.Println("missed required options")

		return
	}

	defaults, err := parseDefaults(defaultsFrom)
	if err != nil {
		//nolint:forbidigo
		fmt.Println("bad defaults spec", err.Error())

		return
	}

	errRun := optionsgen.Run(
		inFilename,
		outFilename,
		optionsStructName,
		outPackageName,
		*defaults,
		!muteWarnings,
	)
	if errRun != nil {
		//nolint:forbidigo
		fmt.Println("cannot run options gen", errRun.Error())

		return
	}
}

func parseDefaults(in string) (*optionsgen.Defaults, error) {
	parts := strings.Split(in, "=")

	from := optionsgen.DefaultsFrom(parts[0])

	switch from {
	case optionsgen.DefaultsFromNone:
		return &optionsgen.Defaults{
			From:  from,
			Param: "",
		}, nil
	case optionsgen.DefaultsFromTag:
		return &optionsgen.Defaults{
			From:  from,
			Param: get1(parts),
		}, nil
	case optionsgen.DefaultsFromVar:
		return &optionsgen.Defaults{
			From:  from,
			Param: get1(parts),
		}, nil
	case optionsgen.DefaultsFromFunc:
		return &optionsgen.Defaults{
			From:  from,
			Param: get1(parts),
		}, nil
	}

	return nil, errors.New("bad syntax")
}

func get1(parts []string) string {
	if len(parts) == 2 { //nolint:gomnd // expect exactly two part
		return parts[1]
	}

	return ""
}

func isEmpty(values ...string) bool {
	for i := range values {
		if values[i] == "" {
			return true
		}
	}

	return false
}
