package cmd

import (
	"fmt"
	"gitmap/git"
	"gitmap/out"
	"gitmap/sum"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var json bool
var limit int
var reverse bool
var threshold int
var ignoreList []string = []string{}

var rootCmd = &cobra.Command{
	Use:   "gitmap",
	Short: "git heat map",
	Long:  "heat map of git changes per file",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := git.Log(limit)

		if err != nil {
			log.Fatal(err)
		}

		content := string(data)
		lines := strings.Split(content, "\n")
		changes := sum.HeatMap(lines, threshold, reverse, ignoreList)

		var output string
		if json {
			output = out.Json(changes)
		} else {
			output = out.Table(changes)
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&json, "json", false, "output in a json format")
	rootCmd.PersistentFlags().BoolVar(&reverse, "reverse", false, "reverse the order of the output")
	rootCmd.PersistentFlags().IntVar(&limit, "limit", 0, "limit the number of git log entries")
	rootCmd.PersistentFlags().IntVar(&threshold, "threshold", 1, "threshold to include entries in the output")
	rootCmd.PersistentFlags().StringSliceVar(&ignoreList, "ignore", []string{}, "list of file patterns to ignore")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
