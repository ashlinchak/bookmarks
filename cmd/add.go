package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ashlinchak/bookmarks/lib/database"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add url",
		Short:   "Create a bookmark for the URL",
		Args:    cobra.ExactArgs(1),
		Run:     addCmdHandler,
		Example: "  bookmarks add https://google.com -t \"search,google\"",
	}

	cmd.Flags().StringP("title", "i", "", "Title for the bookmark. If you don't specify the title it will be set from the HTML page title tag")
	cmd.Flags().StringP("notes", "n", "", "Notes for the bookmark")
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
	notes, _ := cmd.Flags().GetString("notes")
	tags, _ := cmd.Flags().GetStringSlice("tags")

	if len(title) == 0 {
		pageTitle := getPageTitle(url)
		if len(*pageTitle) > 0 {
			title = *pageTitle
		}
	}

	bookmark, err := db.BookmarkRepository.Add(url, title, tags, notes)
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

func getPageTitle(url string) *string {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	dataInBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	pageContent := string(dataInBytes)

	// Find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		return nil
	}
	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += 7

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		return nil
	}

	pageTitle := pageContent[titleStartIndex:titleEndIndex]

	return &pageTitle
}
