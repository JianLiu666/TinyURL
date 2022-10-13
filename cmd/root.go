package cmd

import (
	"fmt"
	"os"
	"tinyurl/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "",
		Long:  ``,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./conf.d/env.yaml")

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(integrationCmd)
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
