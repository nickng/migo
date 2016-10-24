# migo [![Build Status](https://travis-ci.org/nickng/migo.svg?branch=master)](https://travis-ci.org/nickng/migo) [![GoDoc](https://godoc.org/github.com/nickng/migo?status.svg)](http://godoc.org/github.com/nickng/migo)

## `nickng/migo` is a MiGo Types library in Go.

MiGo (mini-go) calculus is a introduced in [this
paper](http://mrg.doc.ic.ac.uk/publications/fencing-off-go-liveness-and-safety-for-channel-based-programming/)
to capture core concurrency features of Go.

This library was designed to work with MiGo types, i.e. the types of
communication primitives in the MiGo calculus, where the values to be
sent/received are abstracted away, for static analysis and verification.

## Build

The package can be installed using `go get`:

    go get github.com/nickng/migo

Some tests uses `golang/mock`, to (re)generate mock files, install `mockgen` and
run `go generate`:

    go get github.com/golang/mock/mockgen
    go generate
    go test
