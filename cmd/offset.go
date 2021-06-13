/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/souvikhaldar/goffer/pkg/webfuzz"
	"github.com/souvikhaldar/gorand"
	"github.com/spf13/cobra"
)

// offsetCmd represents the offset command
var offsetCmd = &cobra.Command{
	Use:   "offset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("offset called")
		fmt.Println("Generating garbage string of length: ", l)
		randStr := gorand.RandStr(l)
		if err := webfuzz.FuzzContent(ip, port, command, randStr, poolSize); err != nil {
			fmt.Println("Unable to fuzz: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(offsetCmd)
	//offsetCmd.LocalFlags().IntVarP(&l, "length", "l", 0, "Length of the garbage string")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offsetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
