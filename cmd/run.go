package cmd

import (
	"fmt"

	"github.com/CaliDog/certstream-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start monitoring certificate transparency logs for potential phishing domains",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoring started...")
		monitorCertStream()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func monitorCertStream() {
	stream, errStream := certstream.CertStreamEventStream(true)

	for {
		select {
		case jq := <-stream:
			domains, err := jq.ArrayOfStrings("data", "leaf_cert", "all_domains")
			if err != nil {
				log.Error("Error decoding jq string")
				continue
			}

			for _, domain := range domains {
				log.Info("Domain: ", domain)
				// Process the domain here
			}

		case err := <-errStream:
			log.Error(err)
		}
	}
}
