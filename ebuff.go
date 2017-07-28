package main

import "os/exec"

import "io/ioutil"
import "bytes"
import "fmt"
import "os"

func main() {
	l := os.Chdir("/home/xiechc/")
	fmt.Println(l)
	var args = []string{"-i", "-", "-g", "80", "-f", "hls", "-hls_list_size",
		"3", "-hls_wrap", "10", "-hls_time", "10", "-hls_init_time", "10", "playlist.m3u8"}

	cmd := exec.Command("ffmpeg", args[0:]...)
	var in bytes.Buffer
	data, er := ioutil.ReadFile("b.ts")
	fmt.Println(er)
	len, er := in.Write(data)
	fmt.Println(len, er)
	cmd.Stdin = &in
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	fmt.Println(err)

	fmt.Println(out.String())
}
