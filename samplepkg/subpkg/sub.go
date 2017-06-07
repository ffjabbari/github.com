package subpkg

import (
    "fmt"
)

type Sub struct {
    Name string
}

func New(name string) (sub * Sub) {
    return &Sub{
        Name: name,
    }
}

func (sub * Sub) Print() {
    fmt.Println("Sub Name:", sub.Name)
}
