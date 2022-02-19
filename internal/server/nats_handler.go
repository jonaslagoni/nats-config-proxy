// Copyright 2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"fmt"

	"github.com/nats-io/nats-config-proxy/internal/models"
	"github.com/nats-io/nats-config-proxy/internal/natsclient"
)

type Empty struct {
}
type ErrorMessage struct {
	description string
}
type NatsHandler struct {
	store *Store
	nc    *natsclient.Client
}

// HandlePerm handles a request to create/update permissions.
func (nsh *NatsHandler) HandleNats() {
	nsh.nc.SubscribeToAccountsV1AuthIdentsDelete(func(m *models.User) {
		err := nsh.store.deleteAllUsers()
		if err != nil {
			fmt.Errorf("Could not delete all users %q", err)
		} else {
			fmt.Printf("deleted all users, OK.")
		}
	})
}
