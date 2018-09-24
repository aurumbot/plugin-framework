package main

import (
	f "github.com/aurumbot/lib/foundation"
	dsg "github.com/bwmarrin/discordgo"
)

var Commands = make(map[string]*f.Command)

func init() {
	Commands["commandname"] = &f.Command{
		Name:    "Somename",
		Help:    "Help page, make sure to define flags and other noteworthy things",
		Perms:   -1,
		Version: "v1.0",
		Action:  command,
	}
}

// Must have these arguments with these names.
func command(session *dsg.Session, message *dsg.Message) {}
