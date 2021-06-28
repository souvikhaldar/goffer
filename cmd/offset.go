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
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/souvikhaldar/goffer/pkg/utils"
	"github.com/souvikhaldar/goffer/pkg/webfuzz"
	"github.com/souvikhaldar/gorand"
	"github.com/spf13/cobra"
)

// offsetCmd represents the offset command
var offsetCmd = &cobra.Command{
	Use:   "offset",
	Short: "Find address of EIP",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Finding the offset")
		fmt.Println("Generating garbage string of length: ", l)
		randStr := gorand.RandStr(l)
		payload := append([]byte(command), []byte(randStr)[:]...)
		if err := webfuzz.SendPayload(
			ip,
			port,
			payload,
			poolSize,
		); err != nil {
			fmt.Println("Unable to fuzz: ", err)
		}
		var eip string
		fmt.Println("Enter value of EIP:")
		reader := bufio.NewReader(os.Stdin)
		eip, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin: ", err)
		}

		eipAscii, err := hex.DecodeString(eip)
		if err != nil {
			// gettign error `invalid byte: U+000A` always
			//fmt.Println(err)
		}
		fmt.Println("EIP entered: ", eip)
		eipRevStr := utils.ReverseStr(string(eipAscii))
		//fmt.Println("EIP string: ", eipRevStr)
		idx := strings.Index(randStr, eipRevStr)
		fmt.Println("Starting address of EIP: ", idx)
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
