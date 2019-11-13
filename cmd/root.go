package cmd

import (
	"github.com/liu-xiao-guo/readjson/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
	"github.com/elastic/beats/libbeat/cmd/instance"
)

// Name of this beat
var Name = "readjson"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name})
