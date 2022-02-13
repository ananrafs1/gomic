package model

type Process struct {
	Start func()
	Finish func()
}