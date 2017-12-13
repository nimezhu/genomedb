package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./genomeVersion.txt")
	c, _ := ioutil.ReadAll(f)
	lines := strings.Split(string(c), "\n")
	for _, l := range lines {
		l := strings.Trim(l, "\n")
		if len(l) > 0 {
			fmt.Println(strings.Replace(os.Args[1], "%s", l, -1))
		}
	}
}
