package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// used for flags
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "hrbac",
		Short: "The command-line tool for hrbac system",
	}

	// queryCmd = &cobra.Command{
	// 	Use:   "query [path] [tag or expression]",
	// 	Short: "query by the tag recursive the given path and subpath",
	// 	Long: ` query by the tag
	// 			Complete documentation is available at https://github.com/QingYunTasha/TagPyrenees`,
	// 	Args: cobra.ExactArgs(2),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		startTime := time.Now()

	// 		var err error
	// 		if cmd.Flags().Lookup("expression").Value.String() == "true" {
	// 			err = usecase.QueryByExpression(args[0], args[1])
	// 		} else {
	// 			err = usecase.QueryByTag(args[0], args[1])
	// 		}

	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}

	// 		fmt.Println("execute time: " + time.Since(startTime).String())
	// 	},
	// }

)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)

	// queryCmd.PersistentFlags().BoolP("expression", "e", false, "use expression to search for tags")

	// rootCmd.AddCommand(queryCmd)
	// rootCmd.AddCommand(listTagsCmd)
}
