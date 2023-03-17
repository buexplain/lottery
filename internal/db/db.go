package db

import (
	"strconv"
	"strings"
	"sync"
)

type User struct {
	UniqId  string `json:"uniqId"`
	Id      uint32 `json:"id"`
	IsAdmin bool   `json:"isAdmin"`
}

type collect struct {
	data   map[string]*User
	number map[uint32]struct{}
	mux    *sync.RWMutex
}

func (r *collect) Del(uniqId string) *User {
	r.mux.Lock()
	defer r.mux.Unlock()
	if u, ok := r.data[uniqId]; ok {
		r.number[u.Id] = struct{}{}
		delete(r.data, uniqId)
		return u
	}
	return nil
}

func (r *collect) Add(uniqId string, isAdmin bool) User {
	r.mux.Lock()
	defer r.mux.Unlock()
	var id uint32
	for id = range r.number {
		break
	}
	u := User{
		Id:      id,
		UniqId:  uniqId,
		IsAdmin: isAdmin,
	}
	r.data[uniqId] = &u
	return u
}

func (r *collect) GetALl() []User {
	r.mux.RLock()
	defer r.mux.RUnlock()
	u := make([]User, 0, len(r.data))
	for _, v := range r.data {
		u = append(u, *v)
	}
	return u
}
func (r *collect) Count() int {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return len(r.data)
}

func (r *collect) GetRandomList(limit int, uniqId string) []User {
	r.mux.RLock()
	defer r.mux.RUnlock()
	u := make([]User, 0, limit)
	//先把自己放进去
	if v, ok := r.data[uniqId]; ok {
		u = append(u, *v)
	}
	for _, v := range r.data {
		if len(u) == limit {
			break
		}
		//跳过自己
		if v.UniqId == uniqId {
			continue
		}
		u = append(u, *v)
	}
	return u
}

var Collect *collect

func init() {
	Collect = &collect{
		data:   map[string]*User{},
		number: make(map[uint32]struct{}, 728),
		mux:    &sync.RWMutex{},
	}
	for i := 1; i < 1000; i++ {
		if !strings.Contains(strconv.Itoa(i), "4") {
			Collect.number[uint32(i)] = struct{}{}
		}
	}
}
