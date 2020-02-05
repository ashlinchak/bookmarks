package cmd

import (
	"fmt"

	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/spf13/cobra"
)

func updateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update url",
		Short: "Update bookmark",
		Args:  cobra.ExactArgs(1),
		Run:   updateCmdHandler,
		Example: "	bookmarks update https://google.com -u https://google.com/imghp -i \"search by images\" -t search,images",
	}

	cmd.Flags().StringP("url", "u", "", "New URL")
	cmd.Flags().StringP("title", "i", "", "Title for the bookmark")
	cmd.Flags().StringSliceP("tags", "t", []string{}, "Comma-separated tags for the bookmark. E.g. \"tag, second tag\"")

	return cmd
}

func init() {
	rootCmd.AddCommand(updateCmd())
}

func updateCmdHandler(cmd *cobra.Command, args []string) {
	db := database.GetDatabase()

	url := args[0]
	newURL, _ := cmd.Flags().GetString("url")
	title, _ := cmd.Flags().GetString("title")
	tags, _ := cmd.Flags().GetStringSlice("tags")
	notes, _ := cmd.Flags().GetString("notes")

	bookmark, err := db.BookmarkRepository.Update(url, newURL, title, tags, notes)

	if err != nil {
		if len(bookmark.Errors) > 0 {
			for _, validationMessage := range bookmark.Errors {
				fmt.Println(validationMessage)
			}
		} else {
			fmt.Println(err)
		}

		defer db.Conn.Close()
		return
	}

	defer db.Conn.Close()

	fmt.Println("1 bookmark updated.")
}
