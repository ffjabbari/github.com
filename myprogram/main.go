package main

import (
    "github.com/samplepkg"
    "github.com/samplepkg/subpkg"
)

func main() {
    sample := samplepkg.New("Test Sample Package")
    sample.Print()

    sub := subpkg.New("Test Sub Package")
    sub.Print()
}
