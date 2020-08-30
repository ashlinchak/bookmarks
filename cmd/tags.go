package cmd

import (
	"fmt"

	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/spf13/cobra"
)

// list tags
func tagsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tags",
		Short: "List all tags",
		Run:   tagsCmdHandler,
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(tagsCmd())
}

func tagsCmdHandler(cmd *cobra.Command, args []string) {
	db := database.GetDatabase()
	tags := db.TagRepository.List(true)

	for _, tag := range tags {
		fmt.Println(tag.Name)
	}
}
