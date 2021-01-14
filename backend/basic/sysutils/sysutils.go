package sysutils

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"

	"example.com/basic/model"
)

func GetProcessTable() ([]model.Process, error) {
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

	var processes []model.Process

	for _, line := range lines[1 : len(lines)-1] {
		words := strings.FieldsFunc(line, unicode.IsSpace)

		process := model.Process{
			UID:   words[0],
			PID:   words[1],
			PPID:  words[2],
			C:     words[3],
			SZ:    words[4],
			RSS:   words[5],
			PSR:   words[6],
			STIME: words[7],
			TTY:   words[8],
			TIME:  words[9],
			CMD:   words[10],
		}

		processes = append(processes, process)
	}

	return processes, nil
}

func KillProcess(PID int) error {
	cmd := exec.Command("kill", "-TERM", strconv.Itoa(PID))

	err := cmd.Run()

	return err
}

func ExecuteScript(script string) (string, error) {
	content := []byte(script)

	tempfile, err := ioutil.TempFile("", "prefix")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	defer os.Remove(tempfile.Name())

	tempfile.Write(content)

	cmd, err := exec.Command("/bin/bash", tempfile.Name()).Output()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(cmd), nil
}
