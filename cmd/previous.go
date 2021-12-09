/*
Copyright © 2021 Ömer Hancı <hanciomer@gmail.com>

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
	"spotify-cli/pkg"

	"github.com/spf13/cobra"
)

// previousCmd represents the previous command
var previousCmd = &cobra.Command{
	Use:   "previous",
	Short: "Play the previous song of the album/playlist",
	Long:  `Skips to the previous song. Usage: spotify-cli previous`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		app.PlayerService.Previous()
	},
}

func init() {
	rootCmd.AddCommand(previousCmd)
}
