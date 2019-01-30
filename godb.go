package main

import (
	"C"
	"fmt"
	"log"
	"net"
	"os"

	"./core"
)

func main() {
	port := "8080"
	if len(os.Args) < 2 {
		log.Fatal("ERROR: Usage godb <dir_databases>")
	}

	if len(os.Args) >= 3 {
		port = os.Args[2]
	}

	if port == "8080" {
		log.Println("Use Default Port(8080)")
	} else {
		log.Println("Use Custom Port(" + port + ")")
	}

	ls, err := net.Listen("tcp6", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	core.DirDB = os.Args[1]

	hdir, herr := os.Open(core.DirDB)
	if herr != nil {
		log.Fatal(herr)
	}
	log.Println("Directory for databases is " + core.DirDB)
	core.HdirDB = hdir
	defer core.HdirDB.Close()

	fis, ferr := hdir.Readdir(-1)
	core.FdirDB = fis
	if ferr != nil {
		log.Fatal(ferr)
	}
	fmt.Println(fis[1].Name())
	ls.Accept()
}
