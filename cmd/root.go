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
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// logic
var ip string
var port string
var command string
var poolSize int
var l int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goffer",
	Short: "Perform buffer overflow exploit",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goffer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "IP address of the target machine")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "port on which the viln app is running on the target machine")
	rootCmd.PersistentFlags().StringVarP(&command, "command", "c", "", "command/function that needs to be fuzzed")
	rootCmd.PersistentFlags().IntVarP(&poolSize, "pool", "s", 1, "Size for the goroutine pool for concurrent execution")
	rootCmd.PersistentFlags().IntVarP(&l, "length", "l", 0, "Length of the garbage string")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".goffer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goffer")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
