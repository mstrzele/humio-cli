// Copyright © 2020 Humio Ltd.
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

package main

import (
	"github.com/humio/cli/cmd/internal/format"
	"github.com/spf13/cobra"
)

func newActionsShowCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "show <repo-or-view> <name>",
		Short: "Show details about an action in a repository or view.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			repoOrViewName := args[0]
			name := args[1]
			client := NewApiClient(cmd)

			action, err := client.Actions().Get(repoOrViewName, name)
			exitOnError(cmd, err, "Error fetching action")

			details := [][]format.Value{
				{format.String("Name"), format.String(action.Name)},
				{format.String("Type"), format.String(action.Type)},
			}

			printDetailsTable(cmd, details)
		},
	}

	return &cmd
}
