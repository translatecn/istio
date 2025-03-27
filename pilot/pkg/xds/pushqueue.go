// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xds

import (
	"sync"

	"istio.io/istio/pilot/pkg/model"
)

type PushQueue struct {
	cond *sync.Cond

	// pending stores all connections in the queue. If the same connection is enqueued again,
	// the PushRequest will be merged.
	pending map[*ConnectionServer]*model.PushRequest

	// queue maintains ordering of the queue
	queue []*ConnectionServer

	// processing stores all connections that have been Dequeue(), but not MarkDone().
	// The value stored will be initially be nil, but may be populated if the connection is Enqueue().
	// If model.PushRequest is not nil, it will be Enqueued again once MarkDone has been called.
	processing map[*ConnectionServer]*model.PushRequest

	shuttingDown bool
}

// ShutDown will cause queue to ignore all new items added to it. As soon as the
// worker goroutines have drained the existing items in the queue, they will be
// instructed to exit.
func (p *PushQueue) ShutDown() {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	p.shuttingDown = true
	p.cond.Broadcast()
}

func NewPushQueue() *PushQueue {
	return &PushQueue{
		pending:    make(map[*ConnectionServer]*model.PushRequest),
		processing: make(map[*ConnectionServer]*model.PushRequest),
		cond:       sync.NewCond(&sync.Mutex{}),
	}
}

// Enqueue 将代理标记为等待推送。如果它已经挂起，pushInfo将被合并。
// ServiceEntry更新将一起添加，如果其中一个已满，则设置为full
func (p *PushQueue) Enqueue(con *ConnectionServer, pushRequest *model.PushRequest) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	if p.shuttingDown {
		return
	}

	// If its already in progress, merge the info and return
	if request, f := p.processing[con]; f {
		p.processing[con] = request.CopyMerge(pushRequest)
		return
	}

	if request, f := p.pending[con]; f {
		p.pending[con] = request.CopyMerge(pushRequest)
		return
	}

	p.pending[con] = pushRequest
	p.queue = append(p.queue, con)
	// Signal waiters on Dequeue that a new item is available
	p.cond.Signal()
}

// Dequeue Remove a proxy from the queue. If there are no proxies ready to be removed, this will block
func (p *PushQueue) Dequeue() (con *ConnectionServer, request *model.PushRequest, shutdown bool) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	// Block until there is one to remove. Enqueue will signal when one is added.
	for len(p.queue) == 0 && !p.shuttingDown {
		p.cond.Wait()
	}

	if len(p.queue) == 0 {
		// We must be shutting down.
		return nil, nil, true
	}

	con = p.queue[0]
	// The underlying array will still exist, despite the slice changing, so the object may not GC without this
	// See https://github.com/grpc/grpc-go/issues/4758
	p.queue[0] = nil
	p.queue = p.queue[1:]

	request = p.pending[con]
	delete(p.pending, con)

	// Mark the connection as in progress
	p.processing[con] = nil

	return con, request, false
}

func (p *PushQueue) MarkDone(con *ConnectionServer) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	request := p.processing[con]
	delete(p.processing, con)

	// If the info is present, that means Enqueue was called while connection was not yet marked done.
	// This means we need to add it back to the queue.
	if request != nil {
		p.pending[con] = request
		p.queue = append(p.queue, con) // ✅
		p.cond.Signal()
	}
}

// Pending Get number of pending proxies
func (p *PushQueue) Pending() int {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	return len(p.queue)
}
