package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	v   bool
	l   int
	ext string
	reg *regexp.Regexp
)

func tree(dir string, firstSymbol string, level int) error {
	if level > l && l != -1 {
		return nil
	}
	res, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for i, unit := range res {
		if !v && unit.Name()[0:1] == "." {
			continue
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
			if err := tree(dir+"/"+unit.Name(), firstSymbol+nextSymbol, level+1); err != nil {
				return err
			}
		} else {
			if match := reg.MatchString(unit.Name()); match {
				fmt.Printf("%s\033[1;34m%s\033[0m\n", firstSymbol+symbol, unit.Name())
			}
		}
	}
	return nil
}

func main() {
	var path string
	var err error
	flag.StringVar(&path, "p", ".", "path for tree")
	flag.IntVar(&l, "l", -1, "deep level of tree")
	flag.StringVar(&ext, "e", "", "extension of files")
	flag.BoolVar(&v, "v", false, "show hidden elements")
	flag.Parse()
	if ext != "" {
		ext = "." + ext
	}
	reg, err = regexp.Compile(fmt.Sprintf("\\w%s$", ext))
	if err != nil {
		log.Fatal(err)
	}
	err = tree(path, "", 0)
	if err != nil {
		log.Fatal(err)
	}
}
