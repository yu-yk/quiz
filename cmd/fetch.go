package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/yu-yk/quiz/pkg/trivia"
)

var (
	// Used for flags.
	number int
	qtype  string

	// fetchCmd represents the fetch command
	fetchCmd = &cobra.Command{
		Use:   "fetch [flags]",
		Short: "Fetch returns the question data",
		Long: `Fetch returns the question data in JSON format,
by default it will return 1 random type question.`,
		Example: "  quiz fetch -n 10 -t multiple",
		Run:     fetch,
	}
)

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().IntVarP(&number, "number", "n", 1, "Number of questions")
	fetchCmd.Flags().StringVarP(&qtype, "type", "t", "", `Type of question, "multiple" or "boolean" (default is random)`)
}

func fetch(cmd *cobra.Command, args []string) {
	// Create a context listening to singal termination.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opt := trivia.Options{Amount: number, Type: qtype}
	api := trivia.NewAPI()
	api.Opt = opt

	qq, err := api.GetQuestions(ctx)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			os.Exit(1)
		} else {
			log.Fatalf("%+v\n", err)
		}
	}

	fmt.Println(trivia.PrettyPrint(qq))
}
