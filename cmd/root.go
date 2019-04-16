package cmd

import (
	"github.com/radoondas/owmbeat/beater"

	"github.com/elastic/beats/libbeat/cmd"
	"github.com/elastic/beats/libbeat/cmd/instance"
)

// Name of this beat
var Name = "owmbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name})
