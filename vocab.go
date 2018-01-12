package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func write(f *os.File, term string) {
	b, err := f.WriteString(term)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%b bytes written\n", b)

	f.Sync()
}

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	dirName := user.HomeDir + "/.vocab"
	mainFileName := dirName + "/terms"

	addPtr := flag.Bool("a", true, "add a term")
	flag.Parse()
	args := flag.Args()

	if _, err = os.Stat(dirName); os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0777)

		if err != nil {
			panic(err)
		}
	}

	file, _ := os.OpenFile(mainFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	str := strings.Join(args, ",") + "\n"

	fmt.Printf("%T %v", str, str)

	if *addPtr {
		write(file, str)
	}
}
