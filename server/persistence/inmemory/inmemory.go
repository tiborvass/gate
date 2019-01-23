// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package inmemory implements an in-process AccessTracker.
package inmemory

import (
	"context"
	"time"

	cache "github.com/patrickmn/go-cache"
	"github.com/tsavola/gate/server"
	"github.com/tsavola/gate/server/state"
)

const DefaultCleanupInterval = 15 * time.Second // Value is subject to change.

type AccessTracker struct {
	state.AccessTrackerBase
	cache *cache.Cache
}

// New AccessTracker with custom cleanup interval.
func New(cleanupInterval time.Duration) *AccessTracker {
	return &AccessTracker{
		cache: cache.New(0, cleanupInterval),
	}
}

// NewDefault AccessTracker with DefaultCleanupInterval.
func NewDefault() *AccessTracker {
	return New(DefaultCleanupInterval)
}

func (at *AccessTracker) TrackNonce(ctx context.Context, pri *server.PrincipalKey, nonce string, expires time.Time) error {
	d := time.Until(expires)
	if d < 1 {
		d = 1
	}
	return at.cache.Add(pri.PrincipalID+" "+nonce, true, d)
}