package session

type Command struct {
	New newCmd `cmd:"" help:"start a new session" default:"1"`
}
