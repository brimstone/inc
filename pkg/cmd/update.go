// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/blang/semver"
	"github.com/brimstone/inc/pkg/version"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
func UpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Updates the local binary to a newer version",
		Long:  `Update the local binary to a newer version`,
		RunE:  update,
	}
}

func update(cmd *cobra.Command, args []string) error {
	v := semver.MustParse(version.Version)
	pubkey, err := version.PublicKey()
	if err != nil {
		return err
	}
	up, err := selfupdate.NewUpdater(selfupdate.Config{
		Validator: &selfupdate.ECDSAValidator{
			PublicKey: pubkey,
		},
		Filters: []string{
			version.Binary,
		},
	})
	latest, err := up.UpdateSelf(v, "brimstone/inc")
	if err != nil {
		return err
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", version.Version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}
