package cmd

import (
	. "../story"
	"errors"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "story",
		Short: "Kanban as code in projects.",
		Long:  `story`,
		Args: func(cmd *cobra.Command, args []string) error {
			userName := viper.GetString("user")
			if userName == "" {
				return errors.New("require a user name")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			userName := viper.GetString("user")

			pick := cmd.Flag("pick").Value.String()
			if pick != "" {
				PickStory(pick, userName)
			}

			status := cmd.Flag("status").Value.String()
			if pick != "" && status != "" {
				ChangeStoryStatus(pick, status)
			}
		},
	}

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "create show",
		Run: func(cmd *cobra.Command, args []string) {
			_ = os.MkdirAll("stories", os.ModePerm)

			create := cmd.Flag("create").Value.String()
			if create != "" {
				CreateStory(create)
				return
			}
		},
	}

	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "sync the stories",
		Run: func(cmd *cobra.Command, args []string) {
			SyncStory()
		},
	}

	showCmd = &cobra.Command{
		Use:   "list",
		Short: "list the stories",
		Run: func(cmd *cobra.Command, args []string) {
			stories := ListStory()

			table := tablewriter.NewWriter(os.Stdout)
			table.SetRowLine(true)
			table.SetHeader([]string{"Id", "Title", "Date", "Status", "Author"})

			for _, v := range stories {
				str := []string{v.Id, v.Title, v.StartDate, v.Status, v.Author}
				table.Append(str)
			}
			table.Render()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(createCmd)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("user", "u", "", "set author")

	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
