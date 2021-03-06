package cmd

import (
	"fmt"
	"os"

	"github.com/hlindberg/tiler/internal/check"
	"github.com/hlindberg/tiler/internal/logging"
	"github.com/puppetlabs/scarp/scarp"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tiler",
	Short: "A utility tiling",
	Long: `A utility for tiling
  `,
	// this is what is run if no subcommand or arguments have been given
	Run: func(cmd *cobra.Command, args []string) {
		if !scarp.RunAsTask() {
			fmt.Println("No input was given. See help below:")
			cmd.HelpFunc()(cmd, args)
		}
	}, PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logging.SetLevelFromName(Loglevel)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if !check.Loglevel(Loglevel) {
			return fmt.Errorf("Given loglevel '%s' is not a recognized loglevel", Loglevel)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Loglevel set on command line (defaults to "warning")
var Loglevel string

func init() {
	cobra.OnInitialize(initConfig)

	// support persisted flags - global for this application and allow giving config file as option
	flags := RootCmd.PersistentFlags()
	flags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scarp.yaml)")
	flags.StringVar(&Loglevel, "loglevel", "warning", "sets log filtering to this level and above")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".scarp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tiler")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
