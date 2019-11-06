package lruset

import (
	"github.com/hashicorp/golang-lru/simplelru"
)

type Set struct {
	cache *simplelru.LRU
}

func New(size int) *Set {
	cache, err := simplelru.NewLRU(size, nil)
	if err != nil {
		panic(err)
	}

	return &Set{cache: cache}
}

func (s *Set) Add(value interface{}) {
	s.cache.Add(value, struct{}{})
}

func (s *Set) Contains(value interface{}) bool {
	_, ok := s.cache.Get(value)
	return ok
}

func (s *Set) Remove(value interface{}) {
	s.cache.Remove(value)
}
