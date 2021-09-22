package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tryCmd)
}

var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		//if err := someFunc(); err != nil {
		//	return err
		//}
		fmt.Println("this is a try cmd showcase....")
		return nil
	},
}
