// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package startuphandler

import (
	"context"
	"sync"

	api "github.com/gravitational/teleport/lib/teleterm/api/protogen/golang/v1"

	"github.com/gravitational/trace"
)

// Handler holds values needed to orchestrate the procedure of setting up the tshd events service
// client on app startup.
type Handler struct {
	// mu is used to ensure we don't somehow end up with two concurrent calls to
	// ResolveTshdEventsServerAddress.
	mu sync.Mutex
	// closeContext gets canceled when apiserver.APIServer gets stopped.
	closeContext context.Context
	// waitForTshdEventsServerAddressC gets closed after the address becomes available.
	waitForTshdEventsServerAddressC chan struct{}
	// tshdEventsServerAddress becomes available after the Electron app makes a call to
	// ResolveTshdEventsServerAddress.
	tshdEventsServerAddress string
}

func New(closeContext context.Context) (*Handler, error) {
	return &Handler{
		closeContext:                    closeContext,
		waitForTshdEventsServerAddressC: make(chan struct{}),
	}, nil
}

// RPC handlers

// ResolveTshdEventsServerAddress is called by the Electron app after the tshd events server starts.
// It'll return an error if called more than once within the application lifetime – there's no need
// to do so, if it's called more than once then it's a sign the Electron app is buggy.
func (s *Handler) ResolveTshdEventsServerAddress(ctx context.Context, req *api.ResolveTshdEventsServerAddressRequest) (*api.ResolveTshdEventsServerAddressResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-s.waitForTshdEventsServerAddressC:
		// The channel is closed so the address must have been resolved already.
		return nil, trace.AlreadyExists("repeated call to resolve tshd events server address")
	default:
		s.tshdEventsServerAddress = req.Address
		close(s.waitForTshdEventsServerAddressC)
	}

	return &api.ResolveTshdEventsServerAddressResponse{}, nil
}

// Other methods

// TshdEventsServerAddress returns either after ctx gets canceled or after the
// ResolveTshdEventsServerAddress handler is called.
func (s *Handler) TshdEventsServerAddress(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", trace.Wrap(ctx.Err())
	case <-s.waitForTshdEventsServerAddressC:
	}

	return s.tshdEventsServerAddress, nil
}
