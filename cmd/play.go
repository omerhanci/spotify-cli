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
	"fmt"
	"spotify-cli/pkg"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Plays the active song",
	Long:  `Plays the current song of the active spotify session. Usage: spotify-cli play`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		if app.PlayerState.Item != nil {
			fmt.Println(fmt.Sprintf("Playing %s", app.PlayerState.Item.Name))
		}
		app.PlayerService.PlayActiveSong(&app.PlayerState.Device.ID)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
