package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "GitRapid - The missing link b/w Git CLI & SCM Providers",
	Long: `Tool to navigate to important features of SCM efficiently right from the command line.
No more searching for clone url, simply use the browser url to clone repos.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitr.yaml)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "verbose logging")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gitr")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
	}
}