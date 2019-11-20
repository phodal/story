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
			_ = os.MkdirAll("stories", os.ModePerm)
			_ = os.MkdirAll("stories/db", os.ModePerm)
			InitStory()

			create := cmd.Flag("create").Value.String()

			if create != "" {
				CreateStory(create)
				return
			}

			list := cmd.Flag("list").Value.String()
			if list != "" {
				stories := ListStory()

				table := tablewriter.NewWriter(os.Stdout)
				table.SetRowLine(true)
				table.SetHeader([]string{"Id", "Title", "Date", "Status", "Author"})

				for _, v := range stories {
					str := []string{v.Id, v.Title, v.StartDate.String(), v.Status, v.Author}
					table.Append(str)
				}
				table.Render()
			}

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
)


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("create", "c", "", "create a story")
	rootCmd.PersistentFlags().StringP("list", "l", "", "list a story")
	rootCmd.PersistentFlags().StringP("pick", "p", "", "pick a story")
	rootCmd.PersistentFlags().StringP("status", "s", "", "change status of story")
	rootCmd.PersistentFlags().StringP("journal", "j", "", "show user journal")
	rootCmd.PersistentFlags().StringP("user", "u", "", "list author")

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
