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
	"log"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var ip string
var port string
var command string
var poolSize int

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
		log.Println("Fuzzing the machine: ", ip)
		conn, err := net.DialTimeout("tcp", ip+":"+port, 1000000*time.Microsecond)
		if err != nil {
			log.Fatalf("Faied to dial: %v", err)
		}
		fmt.Println("Command: ", command)
		defer conn.Close()
		crashed := make(chan int)
		pool := make(chan int, poolSize)
		fmt.Println("Pool size: ", poolSize)

		done := make(chan bool)
		go func() {
			cd := <-crashed
			log.Println("***********")
			log.Printf("crashed at:%d bytes\n", cd)
			log.Println("***********")
			done <- true
			return

		}()

		notCrashed := true
		for n := 100; notCrashed; n += 100 {
			pool <- 1
			go func(n int, pool chan int) {
				defer func() {
					<-pool
				}()

				log.Println("Fuzzing: ", n)
				garbage := strings.Repeat("A", n)

				if err := conn.SetDeadline(time.Now().Add(1000000 * time.Microsecond)); err != nil {
					log.Println("Can't set write deadline: ", err)
					return
				}

				if _, err := conn.Write([]byte(command + garbage)); err != nil && notCrashed {
					//log.Println("Can't write more because: ", err)
					notCrashed = false
					crashed <- n
					return
				}
				rcv := make([]byte, 2048)
				_, err = conn.Read(rcv)
				if err != nil && notCrashed {
					//log.Println("Can't read: ", err)
					notCrashed = false
					crashed <- n
					return
				}

			}(n, pool)

		}

		<-done
		return
	},
}

func init() {
	rootCmd.AddCommand(fuzzCmd)
	fuzzCmd.PersistentFlags().StringVarP(&ip, "address", "a", "", "IP address of the target machine")
	fuzzCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "port on which the viln app is running on the target machine")
	fuzzCmd.PersistentFlags().StringVarP(&command, "command", "c", "", "command/function that needs to be fuzzed")
	fuzzCmd.PersistentFlags().IntVarP(&poolSize, "pool", "s", 1, "Size for the goroutine pool for concurrent execution")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fuzzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fuzzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
