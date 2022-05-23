package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	path string
	l    int
)

func read(dir string, firstSymbol string, level int) error {
	res, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for i, unit := range res {
		if level > l && l != -1 {
			return nil
		}
		var symbol string
		var nextSymbol string
		if level > 0 {
			if i != len(res)-1 {
				symbol = " ╠═"
				nextSymbol = " ║"
			} else {
				symbol = " ╚═"
				nextSymbol = "  "
			}
		}

		if unit.IsDir() {
			fmt.Printf("%s\033[1;35m%s\033[0m", firstSymbol+symbol, unit.Name()+"\n")
			read(dir+"/"+unit.Name(), firstSymbol+nextSymbol, level+1)
		} else {
			fmt.Printf("%s\033[1;34m%s\033[0m", firstSymbol+symbol, unit.Name()+"\n")
		}
	}
	return nil
}

func main() {
	flag.StringVar(&path, "path", ".", "path for tree")
	flag.IntVar(&l, "level", -1, "deep of tree")
	flag.Parse()
	read(path, "", 0)
}
