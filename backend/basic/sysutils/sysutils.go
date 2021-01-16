package sysutils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
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
		return "", err
	}

	defer os.Remove(tempfile.Name())

	tempfile.Write(content)

	cmd, err := exec.Command("/bin/bash", tempfile.Name()).Output()

	if err != nil {
		return "", err
	}

	return string(cmd), nil
}

func GetSysInfo() (model.SysInfo, error) {
	cmd := exec.Command("cat", "/proc/meminfo")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return model.SysInfo{}, err
	}

	meminfo := out.String()

	re := regexp.MustCompile(`MemTotal:\s+(\d+)\skB`)
	submatches := re.FindStringSubmatch(meminfo)

	var memTotal string
	if len(submatches) >= 2 {
		memTotal = submatches[1]
	} else {
		return model.SysInfo{}, errors.New("total memory not found")
	}

	re = regexp.MustCompile(`MemAvailable:\s+(\d+)\skB`)
	submatches = re.FindStringSubmatch(meminfo)

	var memAvailable string
	if len(submatches) >= 2 {
		memAvailable = submatches[1]
	} else {
		return model.SysInfo{}, errors.New("available memory not found")
	}

	out.Reset()

	cmd = exec.Command("lscpu")

	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return model.SysInfo{}, err
	}

	cpuinfo := out.String()

	re = regexp.MustCompile(`Model name:\s+([^\n]+)\n`)
	submatches = re.FindStringSubmatch(cpuinfo)

	var cpuModelName string
	if len(submatches) >= 2 {
		cpuModelName = submatches[1]
	} else {
		re = regexp.MustCompile(`Имя модели:\s+([^\n]+)\n`)
		submatches = re.FindStringSubmatch(cpuinfo)

		if len(submatches) >= 2 {
			cpuModelName = submatches[1]
		} else {
			return model.SysInfo{}, errors.New("cpu model name not found")
		}
	}

	re = regexp.MustCompile(`CPU\(s\):\s+(\d+)`)
	submatches = re.FindStringSubmatch(cpuinfo)

	var cpus int
	if len(submatches) >= 2 {
		cpus, _ = strconv.Atoi(submatches[1])
	} else {
		return model.SysInfo{}, errors.New("available memory not found")
	}

	cmd = exec.Command("uptime")

	out.Reset()

	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return model.SysInfo{}, err
	}

	uptime := out.String()

	uptime = strings.ReplaceAll(uptime, ",", "")

	words := strings.FieldsFunc(uptime, unicode.IsSpace)

	var currentTime string
	var runningTime string
	var countUser string
	var loadLastOneMin float64
	var loadLastFiveMin float64
	var loadLastFifteenMin float64
	if len(words) == 10 {
		currentTime = words[0]
		runningTime = words[2]
		countUser = words[3]

		loadLastOneMin, _ = strconv.ParseFloat(words[7], 64)
		loadLastOneMin = loadLastOneMin / float64(cpus)

		loadLastFiveMin, _ = strconv.ParseFloat(words[8], 64)
		loadLastFiveMin = loadLastFiveMin / float64(cpus)

		loadLastFifteenMin, _ = strconv.ParseFloat(words[9], 64)
		loadLastFifteenMin = loadLastFifteenMin / float64(cpus)
	} else {
		return model.SysInfo{}, errors.New("uptime doesn't available")
	}

	return model.SysInfo{
		MemTotal:           memTotal,
		MemAvailable:       memAvailable,
		CpuModelName:       cpuModelName,
		CPUs:               cpus,
		CurrentTime:        currentTime,
		RunningTime:        runningTime,
		CountUser:          countUser,
		LoadLastOneMin:     loadLastOneMin,
		LoadLastFiveMin:    loadLastFiveMin,
		LoadLastFifteenMin: loadLastFifteenMin,
	}, nil
}
