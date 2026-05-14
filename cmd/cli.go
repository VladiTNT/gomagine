package main

import (
	"flag"
	"strings"
)

type Cli struct {
	Src  string
	Dst  string
	Opts []string
}

func NewCli() (*Cli, error) {
	src := flag.String("src", "", "The source file")
	dst := flag.String("dst", "", "The destination file")
	opts := flag.String("opts", "", "Editing options")
	flag.Parse()

	// If there is no src or dst flags
	if *src == "" || *dst == "" {
		return nil, ErrNoSrcOrDst
	}

	// If there are no options
	if *opts == "" {
		return nil, ErrNoOptions
	}

	// Parse editing options
	fields := strings.Fields(*opts)

	return &Cli{
		Src:  *src,
		Dst:  *dst,
		Opts: fields,
	}, nil
}
