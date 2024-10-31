package somedemo

import (
	"sync"
	"sync/atomic"
)

type ReentrantLock struct {
	sync.Mutex
	owner     int64 // goroutine id
	recursion int32 // lock times
}

func (l *ReentrantLock) Lock() {
	gid := int64(0) // get goroutine from runtime.Stack
	if atomic.LoadInt64(&l.owner) == gid {
		l.recursion++
		return
	}

	l.Lock()
	// get lock and record goroutine id
	atomic.StoreInt64(&l.owner, gid)
	l.recursion++
}

func (l *ReentrantLock) Unlock() {
	gid := int64(0)
	// panic when other goroutine try to Unlock
	if atomic.LoadInt64(&l.owner) != gid {
		panic("current goroutine not the lock owner")
	}

	l.recursion--
	if l.recursion != 0 {
		return // keep lock
	}

	atomic.StoreInt64(&l.owner, -1)
	l.Unlock()
}
