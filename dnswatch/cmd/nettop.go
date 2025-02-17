/*
Copyright (c) Facebook, Inc. and its affiliates.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/facebookincubator/dns/dnswatch/snoop"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(netTopCmd)
}

var netTopCmd = &cobra.Command{
	Use:   "nettop",
	Short: "Interactive nettop-like display",
	Long: `Interactive nettop like display for DNS traffic

Control refresh time using <period> flag

Usage example:
  dnswatch nettop --period 3s
`,

	Run: func(cmd *cobra.Command, args []string) {
		ConfigureVerbosity()

		cfg.NetTop = true
		if err := snoop.Run(&cfg); err != nil {
			log.Fatalf("unable to run snoop: %v", err)
		}
	},
}
