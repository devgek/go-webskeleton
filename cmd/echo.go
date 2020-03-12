package cmd

import (
	"log"

	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/webecho"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "start web server and serve html with echo http server",
	Long:  `webskeleton serve; a typical go web app`,
	Run: func(cmd *cobra.Command, args []string) {
		runEcho(cmd)
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	echoCmd.Flags().String("port", "8080", "The port this app listens")
}

func runEcho(cmd *cobra.Command) {
	// start the web server
	port, _ := cmd.Flags().GetString("port")
	log.Println("Starting webskeleton with echo on port ", port)

	env := config.InitEnv()

	e := webecho.InitWeb(env)

	log.Fatal(e.Start(":" + port))
}