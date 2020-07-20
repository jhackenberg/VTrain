package vocabulary

import "math/rand"

type List []*Pair

func (list List) RandomPair() (*Pair, int) {
	index := rand.Intn(len(list))
	return list[index], index
}
