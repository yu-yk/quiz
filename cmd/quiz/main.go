package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yu-yk/quiz/pkg/trivia"
)

var (
	amount int
	qtype  string
)

func init() {
	flag.Usage = usage
	flag.IntVar(&amount, "n", 1, "numbers of question.")
	flag.StringVar(&qtype, "type", "multiple", `type of question, "multiple" or "boolean".`)
}

func usage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
}

func main() {
	// Parse the input
	flag.Parse()

	if !(qtype == "multiple" || qtype == "boolean") {
		fmt.Println(`Please input "multiple" or "boolean".`)
		os.Exit(1)
	}

	// Create a context listening to singal termination.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opt := trivia.Options{Amount: amount, Type: qtype}
	api := trivia.NewAPI()
	api.Opt = opt

	qq, err := api.GetQuestions(ctx)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Println(trivia.PrettyPrint(qq))
}
