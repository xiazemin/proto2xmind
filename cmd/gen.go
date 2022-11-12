/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	gen "github.com/xiazemin/proto2xmind/proto"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen called")
		fmt.Println("ready to gen:", dstName, srcNames)
		gen.GenFiles(dstName, srcNames...)
	},
}

var (
	srcNames []string
)

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")
	genCmd.PersistentFlags().StringSliceVarP(&srcNames, "src", "s", srcNames, "源proto文件，可以输入多个,eg: proto2xmind gen -s ./example/sub.proto -s ./example/request.proto")
	//genCmd.PersistentFlags().StringSliceVarP()
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
