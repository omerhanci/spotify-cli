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
	"spotify-cli/internal/prompter"
	"spotify-cli/internal/templates"
	"spotify-cli/pkg"

	"github.com/spf13/cobra"
)

// playlistsCmd represents the playlists command
var playlistsCmd = &cobra.Command{
	Use:   "playlists",
	Short: "Command to return playlists of the active users",
	Long:  `Command to return playlists of the active users. Usage spotify-cli playlists`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		userPlaylists := app.PlayerService.GetCurrentUserPlaylists()

		playlistList := userPlaylists.Playlists
		selectedPlaylistIndex, err := prompter.PromptSelect("Select Playlist", playlistList, templates.PlaylistTemplate, 20)

		if err != nil {
			return
		}

		trackResponse := app.PlayerService.GetTracksOfPlaylist(playlistList[selectedPlaylistIndex].ID)
		trackList := trackResponse.Tracks

		selectedTrackIndex, err2 := prompter.PromptSelect("Select Track", trackList, templates.PlaylistTracksTemplate, 5)

		if err2 != nil {
			return
		}

		app.PlayerService.Play(trackList[selectedTrackIndex].Track.URI, &app.PlayerState.Device.ID, &playlistList[selectedPlaylistIndex].URI)
	},
}

func init() {
	rootCmd.AddCommand(playlistsCmd)
}
