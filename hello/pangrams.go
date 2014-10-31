package main

import (
	"os"
	"fmt"
	"bufio"
	"github.com/miry/string"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')

	fmt.Println(string.IsPangram(line))
}
