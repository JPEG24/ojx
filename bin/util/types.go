package util

type Contest struct {
	Platform  string
	ContestID string
	Tasks     []Task
}

type Task struct {
	TaskID      string
	URL         string
	TimeLimit   float64
	MemoryLimit int
}
