// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"
	"time"

	"my-golang-playground/ch9/exercise9.3"
	"my-golang-playground/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	cancel := make(chan struct{})
	done := make(chan struct{})
	m := memo.New(httpGetBody)
	defer m.Close()
	timer := time.NewTimer(time.Second * 50)
	defer timer.Stop()
	go func() {
		memotest.Sequential(t, m, cancel)
		if timer.Stop() {
			done <- struct{}{}
		}
	}()
	select {
	case <-timer.C:
		close(cancel)
	case <-done:
	}
}

func TestConcurrent(t *testing.T) {
	cancel := make(chan struct{})
	done := make(chan struct{})
	m := memo.New(httpGetBody)
	defer m.Close()
	timer := time.NewTimer(time.Second * 20)
	defer timer.Stop()
	go func() {
		memotest.Concurrent(t, m, cancel)
		if timer.Stop() {
			done <- struct{}{}
		}
	}()
	select {
	case <-timer.C:
		close(cancel)
	case <-done:
	}
}
