package sysutils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"unicode"
)


func GetProcessTable() ([]map[string]string, error) {
	cmd := exec.Command("ps", "-eF")

	var out bytes.Buffer
    cmd.Stdout = &out

	err := cmd.Run()

    if err != nil {
		println(err)
		log.Fatal(err)
		return nil, err
	}
	
	lines := strings.Split(out.String(), "\n")

	var process []map[string]string

	for _, line := range lines[1:len(lines) - 1] {
		ps_map := make(map[string]string)

		process = append(process, ps_map)
		words := strings.FieldsFunc(line, unicode.IsSpace)

		ps_map["UID"] = words[0]
		ps_map["PID"] = words[1]
		ps_map["PPID"] = words[2]
		ps_map["C"] = words[3]
		ps_map["SZ"] = words[4]
		ps_map["RSS"] = words[5]
		ps_map["PSR"] = words[6]
		process = append(process, ps_map)
		ps_map["STIME"] = words[7]
		ps_map["TTY"] = words[8]
		ps_map["TIME"] = words[9]
		ps_map["CMD"] = words[10]

		process = append(process, ps_map)
	}

	return process, err
}