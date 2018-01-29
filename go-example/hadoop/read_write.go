package main

import (
	"fmt"
	"github.com/colinmarc/hdfs"
)

func main() {
	client, err := hdfs.New("127.0.0.1:9000")
	{
		fmt.Println(111, err)
		file, err := client.Open("/test.txt")
		fmt.Println(222, err)

		buf := make([]byte, 10)
		n, err := file.Read(buf)

		fmt.Println(333, n, err, string(buf))
	}

	{
		writer, err := client.Create("/_test/create/1.txt")
		fmt.Println(444, err)
		n, err := writer.Write([]byte("foo"))
		fmt.Println(555, n, err)
		writer.Close()
	}
}
