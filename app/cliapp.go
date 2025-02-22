package app

import "github.com/dknathalage/pkg/log"

type CliApp struct {
	Log *log.Logger
}

func NewCliApp() *CliApp {
	return &CliApp{
		Log: log.NewLogger(),
	}
}
