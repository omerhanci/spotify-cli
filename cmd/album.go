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

// AlbumCmd represents the album command
var AlbumCmd = &cobra.Command{
	Use:   "album",
	Short: "Search albums",
	Long: `query an album then select an album among them using arrow keys, after that the tracks in that album will be listed, you can play them by choosing one of them
e.g: spotify-cli search album "dark side of the moon"`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		searchText := args[0]
		searchResult := app.PlayerService.Search(searchText, spotify.SearchTypeAlbum)

		albumList := searchResult.Albums.Albums
		selectedAlbumIndex, err := prompter.PromptSelect("Select Album", albumList, templates.AlbumTemplate, 10)

		if err != nil {
			return
		}

		trackResponse := app.PlayerService.GetTracksOfAlbum(albumList[selectedAlbumIndex].ID)
		trackList := trackResponse.Tracks

		selectedTrackIndex, err := prompter.PromptSelect("Select Track", trackList, templates.AlbumTracksTemplate, len(trackList))

		if err != nil {
			return
		}

		app.PlayerService.Play(trackList[selectedTrackIndex].URI,  &app.PlayerState.Device.ID, &albumList[selectedAlbumIndex].URI)

	},
}

func init() {
	searchCmd.AddCommand(AlbumCmd)
}
