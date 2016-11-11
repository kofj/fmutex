// +build linux darwin freebsd netbsd openbsd

package fmutex

import "syscall"

const (
	perm = 0750
)

type FMutex struct {
	fd     int
	locked bool
}

func New(filename string) (*FMutex, error) {
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, perm)
	return &FMutex{fd: fd}, err
}

func (f *FMutex) Lock() bool {
	err := syscall.Flock(f.fd, syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return false
	}
	f.locked = true
	return true
}

func (f *FMutex) Unlock() (err error) {
	err = syscall.Flock(f.fd, syscall.LOCK_UN)
	if err == nil {
		f.locked = false
	}
	return
}

func (f *FMutex) IsLocked() bool {
	return f.locked
}
