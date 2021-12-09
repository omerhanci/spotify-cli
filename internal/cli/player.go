package cli

import (
	"context"
	"github.com/zmb3/spotify/v2"
	"log"
)

type PlayerService struct {
	Client *spotify.Client
}

func NewPlayerService(client *spotify.Client) *PlayerService {
	return &PlayerService{
		Client: client,
	}
}

func (p PlayerService) Play(songUri spotify.URI, deviceId *spotify.ID, contextUri *spotify.URI) {
	ctx := context.Background()
	opts := &spotify.PlayOptions{
		DeviceID:   deviceId,
	}

	if contextUri != nil {
		opts.PlaybackContext = contextUri
		opts.PlaybackOffset = &spotify.PlaybackOffset{
			URI: songUri,
		}
	} else {
		songUris := make([]spotify.URI, 0)
		songUris = append(songUris, songUri)
		opts.URIs = songUris
	}

	err := p.Client.PlayOpt(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

}

func (p PlayerService) Queue(trackId spotify.ID, deviceId *spotify.ID) error {
	ctx := context.Background()
	songUris := make([]spotify.URI, 0)
	opts := &spotify.PlayOptions{
		DeviceID:        deviceId,
		PlaybackContext: nil,
		URIs:            songUris,
		PlaybackOffset:  nil,
		PositionMs:      0,
	}
	err := p.Client.QueueSongOpt(ctx, trackId, opts)

	if err != nil {
		return err
	}

	return nil
}

func (p PlayerService) PlayActiveSong(deviceId *spotify.ID) {
	ctx := context.Background()
	opts := &spotify.PlayOptions{
		DeviceID: deviceId,
	}
	err := p.Client.PlayOpt(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
}

func (p PlayerService) Next() {
	ctx := context.Background()

	err := p.Client.Next(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (p PlayerService) Previous() {
	ctx := context.Background()

	err := p.Client.Previous(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (p PlayerService) Pause() *spotify.SearchResult {
	ctx := context.Background()
	err := p.Client.Pause(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (p PlayerService) Volume(volumePercentage int) {
	ctx := context.Background()
	err := p.Client.Volume(ctx, volumePercentage)
	if err != nil {
		log.Fatal(err)
	}
}

func (p PlayerService) Search(query string, searchType spotify.SearchType) *spotify.SearchResult {
	ctx := context.Background()

	results, err := p.Client.Search(ctx, query, searchType)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func (p PlayerService) GetTracksOfPlaylist(id spotify.ID) *spotify.PlaylistTrackPage {
	ctx := context.Background()

	results, err := p.Client.GetPlaylistTracks(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func (p PlayerService) GetTracksOfAlbum(id spotify.ID) *spotify.SimpleTrackPage {
	ctx := context.Background()

	results, err := p.Client.GetAlbumTracks(ctx, id)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func (p PlayerService) GetDevicesOfUser() []spotify.PlayerDevice {
	ctx := context.Background()

	results, err := p.Client.PlayerDevices(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

func (p PlayerService) GetPlayerState() *spotify.PlayerState {
	ctx := context.Background()

	results, err := p.Client.PlayerState(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

func (p PlayerService) GetCurrentUserPlaylists() *spotify.SimplePlaylistPage {
	ctx := context.Background()

	results, err := p.Client.CurrentUsersPlaylists(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return results
}
