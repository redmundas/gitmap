package cmd

import (
	"gitmap/sum"
	"log"
	"sort"
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
	Short: "gitmap",
	Long:  "git heat map",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := gitLog(limit)

		if err != nil {
			log.Fatal(err)
		}

		content := string(data)
		lines := strings.Split(content, "\n")
		changes := sum.HeatMap(lines, threshold, ignoreList)

		sort.Slice(changes, func(i, j int) bool {
			if reverse {
				return changes[i].Count > changes[j].Count
			}
			return changes[i].Count < changes[j].Count
		})

		if json {
			printJson(changes)
		} else {
			printTable(changes)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&json, "json", false, "json output format")
	rootCmd.PersistentFlags().BoolVar(&reverse, "reverse", false, "reverse the order of the output")
	rootCmd.PersistentFlags().IntVar(&limit, "limit", 0, "limit the number of log entries")
	rootCmd.PersistentFlags().IntVar(&threshold, "threshold", 1, "threshold to include entries in the output")
	rootCmd.PersistentFlags().StringSliceVar(&ignoreList, "ignore", []string{}, "list of file patterns to ignore")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
