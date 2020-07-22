package main

import "gitlab.com/jhackenberg/vtrain/session"

var cli struct {
	Session session.Command `cmd:"" name:"session"`
}
