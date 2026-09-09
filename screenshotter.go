package main

import (
	"os"
	"os/exec"
)

var path string = os.Getenv("HOME") + "/Pictures/"

func main() {
	screen("image.jpg")
	move("image.jpg")
}

func screen(string) {
	cmd := exec.Command("spectacle", "--fullscreen", "--background", "--nonotify", "--output", "image.jpg")
	cmd.Run()
}

func move(filename string) {
	cmd := exec.Command("cp", filename, path)
	cmd.Run()
}
