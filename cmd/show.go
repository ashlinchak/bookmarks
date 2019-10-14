package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/ashlinchak/bookmarks/lib/models"
	"github.com/spf13/cobra"
)

var (
	err error
)

func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show bookmarks",
		Run:   showCmdHandler,
		Example: `
	bookmarks show -t tag_1,tag_2
	bookmarks show -t "tag, another tag"`,
	}

	cmd.Flags().StringSliceP("tags", "t", []string{}, "Comma-separated tags for search bookmarks. E.g. 'tag, second tag")

	return cmd
}

func init() {
	rootCmd.AddCommand(showCmd())
}

func showCmdHandler(cmd *cobra.Command, args []string) {
	db := database.GetDatabase()
	tags, _ := cmd.Flags().GetStringSlice("tags")

	bookmarks, err := db.BookmarkRepository.List(tags)
	defer db.Conn.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch total := len(bookmarks); {
	case total == 0:
		fmt.Println("0 bookmarks found")
		return
	case total == 1:
		fmt.Println("1 bookmark found")
	default:
		fmt.Printf("%d bookmarks found:\n", total)
	}

	fmt.Println()
	for index, bookmark := range bookmarks {
		print(index+1, &bookmark)
		if index+1 < len(bookmarks) {
			fmt.Println()
		}
	}
}

func print(n int, b *models.Bookmark) {
	// Title
	fmt.Printf("%d. %v\n", n, b.Title)
	// URL
	fmt.Printf("   %v\n", b.URL)

	// Tags
	var tagNames []string
	for _, tag := range b.Tags {
		tagNames = append(tagNames, tag.Name)
	}
	fmt.Printf("   # %v\n", strings.Join(tagNames, ", "))
}
