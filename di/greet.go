package di

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) (int, error) {
	return fmt.Fprintf(writer, "Hello, %s", name)
}
