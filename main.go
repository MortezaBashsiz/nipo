package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	database := CreateDB()
    reader := bufio.NewReader(os.Stdin)
    for {
            fmt.Print("nipo > ")
            cmd, err := reader.ReadString('\n')
            if err != nil {
                    fmt.Fprintln(os.Stderr, err)
            }
            cmd = strings.TrimSuffix(cmd, "\n")
            database.cmdCheck(cmd)
 	}
}