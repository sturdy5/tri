/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tri/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears out done items",
	Long:  `Clears out all done items. If you pass in --all, then it will clear out all todos`,
	Run:   clearItems,
}

var (
	clearAllOpt bool
)

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clearCmd.Flags().BoolVar(&clearAllOpt, "all", false, "Clear all Todos")
}

func clearItems(cmd *cobra.Command, args []string) {
	if clearAllOpt {
		todo.SaveItems(viper.GetString("datafile"), []todo.Item{})
		fmt.Println("All Todos cleared")
	} else {
		items, err := todo.ReadItems(viper.GetString("datafile"))
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		remaining := []todo.Item{}
		for _, i := range items {
			if !i.Done {
				remaining = append(remaining, i)
			}
		}
		todo.SaveItems(viper.GetString("datafile"), remaining)
		fmt.Println("Completed Todos cleared")
	}
}
