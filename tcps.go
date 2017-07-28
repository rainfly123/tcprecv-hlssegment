package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
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
	path := string(buf[:reqLen-1])
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
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
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
