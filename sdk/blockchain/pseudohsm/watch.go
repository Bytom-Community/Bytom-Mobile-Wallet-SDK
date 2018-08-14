// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// +build darwin,!ios freebsd linux,!arm64 netbsd solaris windows

package pseudohsm

//"fmt"
//"github.com/rjeczalik/notify"
//"time"

type watcher struct {
	kc       *keyCache
	starting bool
	running  bool
	//ev       chan notify.EventInfo
	quit chan struct{}
}

func newWatcher(kc *keyCache) *watcher {
	return &watcher{
		kc: kc,
		//ev:   make(chan notify.EventInfo, 10),
		quit: make(chan struct{}),
	}
}

// starts the watcher loop in the background.
// Start a watcher in the background if that's not already in progress.
// The caller must hold w.kc.mu.
func (w *watcher) start() {
	if w.starting || w.running {
		return
	}
	w.starting = true
	go w.loop()
}

func (w *watcher) close() {
	close(w.quit)
}

func (w *watcher) loop() {
	defer func() {
		w.kc.mu.Lock()
		w.running = false
		w.starting = false
		w.kc.mu.Unlock()
	}()
}
