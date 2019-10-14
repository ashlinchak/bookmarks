package cmd

import (
	"fmt"

	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add url",
		Short:   "Create a bookmark for the URL",
		Args:    cobra.ExactArgs(1),
		Run:     addCmdHandler,
		Example: "  bookmarks add google.com -i \"search engine\" -t \"search\"",
	}

	cmd.Flags().StringP("title", "i", "", "Title for the bookmark")
	cmd.Flags().StringSliceP("tags", "t", []string{}, "Comma-separated tags for the bookmark. E.g. \"tag, second tag\"")

	return cmd
}

func init() {
	rootCmd.AddCommand(addCmd())
}

func addCmdHandler(cmd *cobra.Command, args []string) {
	db := database.GetDatabase()

	url := args[0]
	title, _ := cmd.Flags().GetString("title")
	tags, _ := cmd.Flags().GetStringSlice("tags")

	bookmark, err := db.BookmarkRepository.Add(url, title, tags)
	defer db.Conn.Close()

	if err != nil {
		if len(bookmark.Errors) > 0 {
			for _, validationMessage := range bookmark.Errors {
				fmt.Println(validationMessage)
			}
		} else {
			fmt.Println(err)
		}

		return
	}

	fmt.Println("1 bookmark added.")
}
