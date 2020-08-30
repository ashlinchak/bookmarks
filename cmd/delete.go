package cmd

import (
	"fmt"

	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete url",
		Short:   "Delete bookmark",
		Args:    cobra.ExactArgs(1),
		Run:     deleteCmdHandler,
		Example: "  bookmarks delete https://google.com",
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(deleteCmd())
}

func deleteCmdHandler(cmd *cobra.Command, args []string) {
	db := database.GetDatabase()

	url := args[0]
	if err := db.BookmarkRepository.DeleteByURL(url); err != nil {
		fmt.Println(err)

		return
	}

	if err := db.TagRepository.DeleteNotActive(); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("1 bookmark deleted.")
}
