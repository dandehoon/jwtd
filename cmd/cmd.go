package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "jwtd [JWT_TOKEN]",
	Short:   "A simple command-line JSON Web Tokens decoder tool",
	Long:    `jwtd decodes JWT tokens and prints their contents in a human-readable format.`,
	Version: VERSION,
	Example: `  jwtd $JWT_TOKEN
	echo $JWT_TOKEN | jwtd
  pbpaste | jwtd
  `,
	Args: cobra.MaximumNArgs(1),
	RunE: runDecode,
}

func init() {
	log.SetFlags(0)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

// runDecode is the main command handler
func runDecode(cmd *cobra.Command, args []string) error {
	token, err := getToken(args)
	if err != nil {
		return err
	}
	if err := decodeToken(token); err != nil {
		return err
	}
	return nil
}
