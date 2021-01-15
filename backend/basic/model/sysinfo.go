package model

type SysInfo struct {
	MemTotal           string
	MemAvailable       string
	CpuModelName       string
	CPUs               int
	CurrentTime        string
	RunningTime        string
	CountUser          string
	LoadLastOneMin     float64
	LoadLastFiveMin    float64
	LoadLastFifteenMin float64
}
