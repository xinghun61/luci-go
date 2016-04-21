// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tumble

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/luci/luci-go/appengine/gaemiddleware"
	"github.com/luci/luci-go/common/logging"
	"golang.org/x/net/context"
)

// Service is an instance of a Tumble service. It installs its handlers into an
// HTTP router and services Tumble request tasks.
type Service struct {
	// Middleware is an optional function which allows your application to add
	// application-specific resources to the context used by ProcessShardHandler.
	//
	// Context will already be setup with BaseProd.
	Middleware func(context.Context) context.Context
}

// InstallHandlers installs http handlers.
func (s *Service) InstallHandlers(r *httprouter.Router) {
	// GET so that this can be invoked from cron
	r.GET(fireAllTasksURL,
		gaemiddleware.BaseProd(gaemiddleware.RequireCron(s.FireAllTasksHandler)))

	r.POST(processShardPattern,
		gaemiddleware.BaseProd(gaemiddleware.RequireTaskQueue(baseName, s.ProcessShardHandler)))
}

// FireAllTasksHandler is a http handler suitable for installation into
// a httprouter. It expects `logging` and `luci/gae` services to be installed
// into the context.
//
// FireAllTasksHandler verifies that it was called within an Appengine Cron
// request, and then invokes the FireAllTasks function.
func (s *Service) FireAllTasksHandler(c context.Context, rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := s.FireAllTasks(c); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "fire_all_tasks failed: %s", err)
	} else {
		rw.Write([]byte("ok"))
	}
}

// FireAllTasks fires off 1 task per shard to ensure that no tumble work
// languishes forever. This may not be needed in a constantly-loaded system with
// good tumble key distribution.
func (s *Service) FireAllTasks(c context.Context) error {
	cfg := getConfig(c)
	shards := make(map[taskShard]struct{}, cfg.NumShards)
	for i := uint64(0); i < cfg.NumShards; i++ {
		shards[taskShard{i, minTS}] = struct{}{}
	}

	err := error(nil)
	if !fireTasks(c, cfg, shards) {
		err = errors.New("unable to fire all tasks")
	}

	return err
}

// ProcessShardHandler is a http handler suitable for installation into
// a httprouter. It expects `logging` and `luci/gae` services to be installed
// into the context.
//
// ProcessShardHandler verifies that its being run as a taskqueue task and that
// the following parameters exist and are well-formed:
//   * timestamp: decimal-encoded UNIX/UTC timestamp in seconds.
//   * shard_id: decimal-encoded shard identifier.
//
// ProcessShardHandler then invokes ProcessShard with the parsed parameters.
func (s *Service) ProcessShardHandler(c context.Context, rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if s.Middleware != nil {
		c = s.Middleware(c)
	}

	tstampStr := p.ByName("timestamp")
	sidStr := p.ByName("shard_id")

	tstamp, err := strconv.ParseInt(tstampStr, 10, 64)
	if err != nil {
		logging.Errorf(c, "bad timestamp %q", tstampStr)
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "bad timestamp")
		return
	}

	sid, err := strconv.ParseUint(sidStr, 10, 64)
	if err != nil {
		logging.Errorf(c, "bad shardID %q", tstampStr)
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "bad shardID")
		return
	}

	cfg := getConfig(c)
	err = processShard(c, cfg, time.Unix(tstamp, 0).UTC(), sid)
	if err != nil {
		logging.Errorf(c, "failure! %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "error: %s", err)
	} else {
		rw.Write([]byte("ok"))
	}
}
