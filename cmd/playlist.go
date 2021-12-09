/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/zmb3/spotify/v2"
	"spotify-cli/internal/prompter"
	"spotify-cli/internal/templates"
	"spotify-cli/pkg"

	"github.com/spf13/cobra"
)

// playlistCmd represents the playlist command
var playlistCmd = &cobra.Command{
	Use:   "playlist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		searchText := args[0]
		searchResult := app.PlayerService.Search(searchText, spotify.SearchTypePlaylist)

		playlistList := searchResult.Playlists.Playlists
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
	searchCmd.AddCommand(playlistCmd)
}
