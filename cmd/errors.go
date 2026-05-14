package main

import "errors"

var (
	ErrNoSrcOrDst = errors.New("gomagine: must provide flags 'src' and 'dst'")
	ErrNoOptions  = errors.New("gomagine: no editing options provided")
)
