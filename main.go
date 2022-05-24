package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var l int

func tree(dir string, firstSymbol string, level int) error {
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
		if i != len(res)-1 {
			symbol = " ╠═══"
			nextSymbol = " ║  "
		} else {
			symbol = " ╚═══"
			nextSymbol = "    "
		}
		if unit.IsDir() {
			fmt.Printf("%s\033[1;35m%s\033[0m\n", firstSymbol+symbol, unit.Name())
			tree(dir+"/"+unit.Name(), firstSymbol+nextSymbol, level+1)
		} else {
			fmt.Printf("%s\033[1;34m%s\033[0m\n", firstSymbol+symbol, unit.Name())
		}
	}
	return nil
}

func main() {
	var path string
	flag.StringVar(&path, "path", ".", "path for tree")
	flag.IntVar(&l, "level", -1, "deep level of tree")
	flag.Parse()
	err := tree(path, "", 0)
	if err != nil {
		log.Fatal(err)
	}
}
