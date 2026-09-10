package main

import (
	"os"
	"os/exec"
)

var output string = "/tmp/temp/data/prosses_info.txt"

func getinfo() {
	var result string

	result += "TOP 10 PROCESSES BY CPU\n\n"
	cpu, _ := exec.Command("ps", "-eo", "pid,comm,pcpu,pmem", "--sort=-pcpu").Output()
	result += string(cpu)

	result += "\nTOP 10 PROCESSES BY MEMORY\n\n"
	mem, _ := exec.Command("ps", "-eo", "pid,comm,pcpu,pmem", "--sort=-pmem").Output()
	result += string(mem)

	os.WriteFile(output, []byte(result), 0644)
}

func procinfo() {
	getinfo()
}
