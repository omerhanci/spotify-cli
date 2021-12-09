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
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify/v2"
	"spotify-cli/internal/prompter"
	"spotify-cli/internal/templates"
	"spotify-cli/pkg"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for a song",
	Long: `query a song then select a track among them using arrow keys
e.g: spotify-cli search "comfortably numb"`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		searchText := args[0]
		searchResult := app.PlayerService.Search(searchText, spotify.SearchTypeTrack)

		trackList := searchResult.Tracks.Tracks

		selectedTrackIndex, err := prompter.PromptSelect("Select Track", trackList, templates.TracksTemplate, 20)

		if err != nil {
			return
		}

		queueOrPlay, err := prompter.PromptSelect("Queue or play right now?", []string{"play", "queue"}, nil, 2)

		if err != nil {
			return
		}

		if queueOrPlay == 0 {
			app.PlayerService.Play(trackList[selectedTrackIndex].URI, &app.PlayerState.Device.ID, nil)
		} else {
			app.PlayerService.Queue(trackList[selectedTrackIndex].ID, &app.PlayerState.Device.ID)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
