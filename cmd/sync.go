/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/koksmat-com/koksmat/model"
	"github.com/spf13/cobra"
)

var hubSiteID string = ""

// importCmd represents the import command
var syncCmd = &cobra.Command{
	Use:   "sync [collection name]",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Importing")

		switch fmt.Sprint(args[0]) {
		case "domains":

			err := model.SyncDomains()
			if err != nil {
				log.Fatalln(err)
			}
		case "rooms":

			err := model.ImportRooms()
			if err != nil {
				log.Fatalln(err)
			}
		case "sites":
			if hubSiteID == "" {
				log.Fatalln("Need to specify hubSiteId")
			}
			err := model.SyncHubSitePages(hubSiteID)
			if err != nil {
				log.Fatalln(err)
			}
			err = model.SyncSitesNavigation()
			if err != nil {
				log.Fatalln(err)
			}
		default:

			log.Fatalln("Cannot use that collection name", subject)
			return
		}

		log.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVarP(&hubSiteID, "hubSiteId", "", "", "Hub Site ID")

}
