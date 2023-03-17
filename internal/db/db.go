package db

import "sync"

type User struct {
	UniqId   string
	Nickname string
	IsAdmin  bool
}

type collect struct {
	data map[string]*User
	mux  *sync.RWMutex
}

func (r *collect) Del(uniqId string) *User {
	r.mux.Lock()
	defer r.mux.Unlock()
	if u, ok := r.data[uniqId]; ok {
		delete(r.data, uniqId)
		return u
	}
	return nil
}

func (r *collect) Add(uniqId string, nickname string, isAdmin bool) User {
	r.mux.Lock()
	defer r.mux.Unlock()
	u := User{
		UniqId:   uniqId,
		Nickname: nickname,
		IsAdmin:  isAdmin,
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

var Collect *collect

func init() {
	Collect = &collect{
		data: map[string]*User{},
		mux:  &sync.RWMutex{},
	}
}
