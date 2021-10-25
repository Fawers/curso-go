package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleIndexError(conn net.Conn, cmd string) {
	err := recover()

	if err != nil {
		if cmd == "ADD" || cmd == "UPDATE" {
			fmt.Fprintln(conn, "ADD requires two arguments")
		} else {
			fmt.Fprintf(conn, "%s requires one argument\n", cmd)
		}
	}
}

func processConn(conn net.Conn) {
	r := bufio.NewReader(conn)

	data, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Reading data failed")
		fmt.Println(err)
	}

	fmt.Printf("read: %+q\n", data)

	line := strings.TrimSpace(data)
	pieces := strings.Split(line, " ")
	if len(pieces) < 1 {
		fmt.Println("nothing to read")
		return
	}

	cmd := pieces[0]

	defer handleIndexError(conn, cmd)
	switch cmd {
	case "UPDATE":
		fallthrough
	case "ADD":
		key := pieces[1]
		val := pieces[2]
		fmt.Printf("%s %+q %q\n", cmd, key, val)
		if cmd == "ADD" {
			err := gCacheD.Add(key, val)
			if err != nil {
				fmt.Fprintf(conn, "Error adding key: %s\n", err)
				return
			}
		} else {
			gCacheD.Update(key, val)
		}
		fmt.Fprintln(conn, "OK")

	case "DEL":
		key := pieces[1]
		fmt.Printf("DEL %+q\n", key)
		gCacheD.Del(key)
		fmt.Fprintln(conn, "OK")

	case "GET":
		key := pieces[1]
		fmt.Printf("GET %+q\n", key)
		v, err := gCacheD.Get(key)

		if err != nil {
			fmt.Fprintf(conn, "GETting failed: %s\n", err)
		} else {
			fmt.Fprintln(conn, v)
		}

	case "GETALL":
		fmt.Println("GETALL")
		items := gCacheD.GetAll()

		for _, item := range items {
			fmt.Fprintf(conn, "%s\t%s\n", item[0], item[1])
		}
	}
}

func main() {
	host := ":8080"

	if len(os.Args) >= 2 {
		host = os.Args[1]
	}

	server, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not initialize server:", err)
		return
	}
	defer server.Close()

	fmt.Printf("Listening on %s...\n", server.Addr())

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Println("received:", conn)

		go func() {
			processConn(conn)
			conn.Close()
		}()
	}
}
