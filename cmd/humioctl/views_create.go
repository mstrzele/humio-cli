// Copyright © 2018 Humio Ltd.
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
	"fmt"
	"github.com/spf13/cobra"
)

func newViewsCreateCmd() *cobra.Command {
	connections := make(map[string] string)
	description := ""

	c := &cobra.Command{
		Use:   "create <view-name>",
		Short: "Create a view.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			viewName := args[0]

			client := NewApiClient(cmd)

			apiErr := client.Views().Create(viewName, description, connections)
			exitOnError(cmd, apiErr, "Error creating view")
			fmt.Printf("Successfully created view %s\n", viewName)

			view, apiErr := client.Views().Get(viewName)
			exitOnError(cmd, apiErr, "error fetching view")

			printViewTable(view)

			fmt.Println()
		},
	}

	c.Flags().StringToStringVar(&connections, "connection", connections, "Sets a repository connection with the chosen filter.")
	c.Flags().StringVar(&description, "description", description, "Sets an optional description")

	return c
}