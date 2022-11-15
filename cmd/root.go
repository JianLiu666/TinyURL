package cmd

import (
	"fmt"
	"os"
	"tinyurl/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "",
	Long:  ``,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file (default is ./conf.d/env.yaml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("./conf.d/env.yaml")
	}

	viper.AutomaticEnv()
	config.LoadFromViper()
}
