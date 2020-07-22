package vocabulary

import (
	"math/rand"
	"os"

	"gopkg.in/yaml.v2"
)

type List []*Pair

func (list List) RandomPair() (*Pair, int) {
	index := rand.Intn(len(list))
	return list[index], index
}

// LoadList loads a vocabulary list from one or multiple paths
func LoadList(paths ...string) (List, error) {
	var master List

	for _, p := range paths {
		file, err := os.Open(p)
		if err != nil {
			return nil, err
		}
		var list List
		if err := yaml.NewDecoder(file).Decode(&list); err != nil {
			return nil, err
		}
		master = append(master, list...)
	}

	return master, nil
}
