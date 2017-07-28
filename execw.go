package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	var args = []string{"-i", "-", "-g", "80", "-f", "hls", "-hls_list_size",
		"3", "-hls_wrap", "10", "-hls_time", "10", "-hls_init_time", "10", "playlist.m3u8"}

	cmd := exec.Command("ffmpeg", args[0:]...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		fl, err := os.Open("h264")
		w := bufio.NewWriter(stdin)
		n, er := w.ReadFrom(fl)
		fmt.Println(n, err, er)

	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
