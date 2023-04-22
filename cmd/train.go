package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Train the phishing detection model",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Training model...")
		trainModel()
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
}

func trainModel() {
	cmd := exec.Command("python3", "training/python_scripts/train.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running train.py:", err)
		fmt.Println("Output from train.py:", string(output)) // Add this line to print the output
		return
	}
	fmt.Println("Training completed:", strings.TrimSpace(string(output)))
}
