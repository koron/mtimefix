package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func getMtime(name string) (time.Time, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return time.Time{}, err
	}
	return fi.ModTime(), nil
}

func fixMtime(name string, d time.Duration) error {
	mt, err := getMtime(name)
	if err != nil {
		return err
	}
	mt = mt.Add(d)
	return os.Chtimes(name, mt, mt)
}

func main() {
	d := flag.Duration("d", 0, "diff for mtime")
	flag.Parse()
	if *d == 0 {
		fmt.Println(`"-d" should not be zero`)
		os.Exit(1)
	}
	n := flag.NArg()
	if n == 0 {
		fmt.Println("need some files")
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		f := flag.Arg(i)
		err := fixMtime(f, *d)
		if err != nil {
			fmt.Printf("failed to fix MTIME of %s: %s\n", f, err)
		}
	}
}
