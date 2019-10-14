package cmd

import (
	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/spf13/cobra"
)

// setup database
func setupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup Bookmark database",
		Run:   setupCmdHandler,
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(setupCmd())
}

func setupCmdHandler(cmd *cobra.Command, args []string) {
	db := database.GetDatabase()
	db.Setup()

	defer db.Conn.Close()
}
