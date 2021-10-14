package cmd

import (
	"fmt"
	"github.com/kuzik/text-files-profile/file_profiler"
	"github.com/spf13/cobra"
	"os"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Build text files profile",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		if dir == "" {
			panic("missed required parameter")
		}
		profile(dir)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := profileCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	profileCmd.Flags().StringP("dir", "d", "", "Base profiling directory")
}

func profile(dir string) {

	profiler := file_profiler.NewProfiler(
		&file_profiler.Collector{},
		&file_profiler.Processor{},
	)
	profile, err := profiler.Profile(dir)
	if err != nil {
		panic("Error during profiling process")
	}

	profiler.PrintProfile(profile)
}
