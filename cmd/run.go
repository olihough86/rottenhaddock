package cmd

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

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
				predictDomain(domain)
			}

		case err := <-errStream:
			log.Error(err)
		}
	}
}

func predictDomain(domain string) {
	preprocessedDomain := preprocessDomain(domain) // THIS IS WHY IT DOESN'T WORK AS THIS DOES NOT EXIST YET

	// Remove the 'td' field from the preprocessedDomain
	preprocessedDomain = append(preprocessedDomain[:0], preprocessedDomain[1:]...)

	cmd := exec.Command("python3", "training/python_scripts/predict.py")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Errorf("Error creating stdin pipe for predict.py: %v", err)
		return
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, strings.Join(preprocessedDomain, ",")+"\n")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error running predict.py: %v", err)
		return
	}

	fmt.Println("Prediction for", domain+":", strings.TrimSpace(string(out)))
}
