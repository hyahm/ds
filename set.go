package ds

import (
	"errors"
	"sync"
)

type Arrint []int

type Set struct {
	Length int
	S      Arrint
	M      map[int]struct{}
	Mu     *sync.RWMutex
}

func (ai Arrint) add(value int) Arrint {
	if len(ai) == 0 {
		ai = append(ai, value)
		return ai
	}

	var temp Arrint
	if cap(ai) > len(ai)+1 {
		temp = make([]int, 0, cap(ai))
	} else {
		temp = make([]int, 0, cap(ai)*2)
	}
	for i, v := range ai {
		// 1, 2, 4, 5
		// 如果值大于最大的值， 直接添加到末尾
		if i == len(ai)-1 && value > v {
			temp = append(ai, value)
			return temp
		}
		if value < v {
			if i == 0 {
				temp = append(temp, value)
				temp = append(temp, ai...)
				return temp
			} else {
				temp = append(temp, ai[:i]...)
				temp = append(temp, value)
				temp = append(temp, ai[i:]...)
				return temp
			}
		}
	}
	return ai
}

func NewSet() *Set {
	return &Set{
		S:  make([]int, 0),
		M:  make(map[int]struct{}),
		Mu: &sync.RWMutex{},
	}
}

func (s *Set) Add(value int) *Set {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if _, ok := s.M[value]; !ok {
		s.Length++
		s.M[value] = struct{}{}
		s.S = s.S.add(value)
	}
	return s
}

func (s *Set) Get() []int {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.S
}

func (s *Set) GetIndex(i int) (int, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	if s.Length >= i || i < 0 {
		return 0, errors.New("not found this index")
	}
	return s.S[i], nil
}
