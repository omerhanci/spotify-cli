package pkg

import (
	"fmt"
	"github.com/zmb3/spotify/v2"
	"log"
	"spotify-cli/internal/auth"
	"spotify-cli/internal/cli"
	"spotify-cli/internal/prompter"
	"spotify-cli/internal/templates"
)

var App Application

type Application struct {
	PlayerService *cli.PlayerService
	PlayerState   *spotify.PlayerState
	AuthClient    *auth.AuthClient
}

func Init() *Application {
	authClient := auth.NewAuthClient()
	App = Application{
		AuthClient:    authClient,
		PlayerService: cli.NewPlayerService(authClient.Client),
	}

	App.PlayerState = App.PlayerService.GetPlayerState()
	if !App.PlayerState.Device.Active {
		devices := App.PlayerService.GetDevicesOfUser()
		fmt.Println("No active device found, please select the device you wish to play on")
		selectedDeviceIndex, err := prompter.PromptSelect("Select Device", devices, templates.DeviceTemplate, len(devices))
		if err != nil {
			log.Fatal(err)
		}
		App.PlayerState.Device.ID = devices[selectedDeviceIndex].ID
	}
	return &App
}
