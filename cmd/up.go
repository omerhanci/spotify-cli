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
	"fmt"
	"log"
	"spotify-cli/pkg"
	"strconv"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := pkg.Init()
		var levelUp int
		var err error

		if len(args) == 0 {
			levelUp = 5
		} else {
			levelUp, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}

		curLevel := app.PlayerState.Device.Volume
		levelToBe := curLevel + levelUp

		if levelToBe > 100 {
			fmt.Println("Level can't be higher than 100")
			levelToBe = 100
		}

		app.PlayerService.Volume(levelToBe)
		fmt.Println(fmt.Sprintf("Volume is set to %d", levelToBe))
	},
}

func init() {
	volumeCmd.AddCommand(upCmd)
}
