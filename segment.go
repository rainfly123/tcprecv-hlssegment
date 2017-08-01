package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
	BASE_DIR  = "/home/xiechc/git/"
)

func handleRequest(conn net.Conn) {
	buf := make([]byte, 128)
	reqLen, err := conn.Read(buf)
	if err != nil && reqLen == 0 {
		fmt.Println("Error reading:", err.Error())
		return
	}

	where := 0
	concat := false
	if buf[reqLen-1] != '\n' {
		for i := 0; i < reqLen; i++ {
			if buf[i] == '\n' {
				where = i
				concat = true
				break
			}
		}
	} else {
		where = reqLen - 1
	}

	path := string(buf[:where])
	workdir := BASE_DIR + path
	err = os.MkdirAll(workdir, 0777)
	l := os.Chdir(workdir)
	fmt.Println(workdir, err, l)

	var args = []string{"-i", "-", "-g", "60", "-f", "hls", "-hls_list_size",
		"3", "-hls_wrap", "10", "-hls_time", "10", "-hls_init_time", "10", "playlist.m3u8"}

	cmd := exec.Command("ffmpeg", args[0:]...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		defer conn.Close()
		//fl, err := os.Open("h264")
		if concat {
			stdin.Write(buf[where+1:])
		}
		w := bufio.NewWriter(stdin)
		n, er := w.ReadFrom(conn)
		fmt.Println(n, err, er)

	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

func main() {
	// Listen for incoming connections.
	port := flag.String("port", ":8080", "tcp receive data port")
	ipaddr := flag.String("ipaddr", "localhost", "tcp receive data port")
	flag.Parse()
	fmt.Println(*port, *ipaddr)
	l, err := net.Listen(CONN_TYPE, *ipaddr+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + *ipaddr + ":" + *port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
