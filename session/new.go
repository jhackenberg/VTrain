package session

import (
	"fmt"

	"gitlab.com/jhackenberg/vtrain/vocabulary"
)

type newCmd struct {
	ListPaths []string `arg:"" required:"" name:"list" type:"existingpath"`
}

func (cmd *newCmd) Run() error {
	list, err := vocabulary.LoadList(cmd.ListPaths...)
	if err != nil {
		return err
	}

	fmt.Println(list)

	return nil
}
