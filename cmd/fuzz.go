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
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var ip string
var port string
var command string

// fuzzCmd represents the fuzz command
var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fuzz called")

		var d net.Dialer
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		conn, err := d.DialContext(ctx, "tcp", ip+":"+port)
		if err != nil {
			log.Fatalf("Faied to dial: %v", err)
		}
		defer conn.Close()

		for n := 100; ; n += 100 {
			garbage := strings.Repeat("A", n)
			if _, err := conn.Write([]byte(command + garbage)); err != nil {
				log.Printf("crashed at:%d bytes\n", n)
				log.Fatal(err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(fuzzCmd)
	fuzzCmd.PersistentFlags().StringVarP(&ip, "address", "a", "", "IP address of the target machine")
	fuzzCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "port on which the viln app is running on the target machine")
	fuzzCmd.PersistentFlags().StringVarP(&command, "command", "c", "", "command/function that needs to be fuzzed")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fuzzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fuzzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
