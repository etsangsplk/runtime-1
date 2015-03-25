// Copyright 2015 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2015 Petar Maymounkov <p@gocircuit.org>

package use

import (
	"github.com/gocircuit/core/sys"
	"github.com/gocircuit/core/sys/backdial"
	"github.com/gocircuit/core/sys/codec"
	"github.com/gocircuit/core/sys/pool"
	"github.com/gocircuit/core/sys/tcp"
)

func NewClearTCP(addr string) (sys.Peer, error) {
	a, err := tcp.ResolveAddr(addr)
	if err != nil {
		return nil, err
	}
	carrier, err := tcp.New(a)
	if err != nil {
		return nil, err
	}
	return pool.New(backdial.New(codec.NewGob(carrier))), nil
}
