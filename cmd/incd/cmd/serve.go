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
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brimstone/logger"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
func ServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start an Inc server",
		Long:  `Starts an Inc server on a TCP port.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.New()

			ctx, allDone := context.WithCancel(context.Background())

			s := &http.Server{
				Addr:           ":8080", // TODO make this a flag
				ReadTimeout:    60 * time.Second,
				WriteTimeout:   60 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}
			go func() {
				log.Info("Starting http server")
				err := s.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					log.Error("http server closed",
						log.Field("err", err.Error()),
					)
				}
			}()

			go func() {
				// Set up channel on which to send signal notifications.
				// We must use a buffered channel or risk missing the signal
				// if we're not ready to receive when the signal is sent.
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)

				// Block until a signal is received.
				s := <-c
				log.Info("Got signal",
					log.Field("signal", s),
				)
				allDone()
			}()

			me, _ := os.Executable()
			fi, _ := os.Stat(me)
			origTime := fi.ModTime()
			needReload := false
			log.Info("Waiting for executable to change")
			for ctx.Err() == nil {
				fi, err := os.Stat(me)
				if err != nil {
					continue
				}
				if origTime != fi.ModTime() {
					needReload = true
					allDone()
					log.Warn("Detected change in executable, relaunching")
					break
				}
				select {
				case <-time.After(time.Second):
				case <-ctx.Done():
					break
				}
			}
			s.Shutdown(context.Background())

			if needReload {
				if err := syscall.Exec(me, os.Args, os.Environ()); err != nil {
					log.Fatal(err)
				}
			}
			return nil
		},
	}
}
