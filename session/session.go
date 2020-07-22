package session

import "gitlab.com/jhackenberg/vtrain/vocabulary"

type Session struct {
	Vocabulary vocabulary.List
}

func NewSession(vocabulary vocabulary.List) *Session {
	return &Session{Vocabulary: vocabulary}
}
