package cmd

import (
    "myapp/server"

    "github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Starts the web server",
    Run: func(cmd *cobra.Command, args []string) {
        srv := server.NewServer()
        srv.Start()
    },
}
