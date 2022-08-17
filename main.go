package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jlaffaye/ftp"
)

func getStringEntryType(t ftp.EntryType) string {
	switch t {
	case ftp.EntryTypeFile:
		return "(file)"
	case ftp.EntryTypeFolder:
		return "(folder)"
	case ftp.EntryTypeLink:
		return "(link)"
	default:
		return ""
	}
}

func main() {
	fmt.Println("init...")
	client, err := ftp.Dial("192.168.1.XX:21", ftp.DialWithDisabledEPSV(true))
	if err != nil {
		panic(err)
	}

	if err := client.Login("user_name", "secret_password"); err != nil {
		panic(err)
	}

	r, err := client.Retr("test.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	println(string(buf))

	fmt.Println("======= PATH ./")

	entries, err1 := client.List(".")
	if err1 != nil {
		fmt.Println("ERROR")
		fmt.Println(entries)
	}
	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}

	fmt.Println("end...")
}
