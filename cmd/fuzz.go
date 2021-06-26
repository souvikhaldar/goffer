/*
Copyright © 2021 Souvik Haldar <souvikhaldar32@gmail.com>

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
	"log"

	"github.com/souvikhaldar/goffer/pkg/webfuzz"
	"github.com/spf13/cobra"
)

// fuzzCmd represents the fuzz command
var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "",
	Long:  `Fuzzing!`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Fuzzing the machine: ", ip)
		err := webfuzz.Fuzz(ip, port, command, poolSize)
		if err != nil {
			fmt.Println("Error fuzzing: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fuzzCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fuzzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fuzzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
