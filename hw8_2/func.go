package hw82

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var rounds = make(chan Round)
var results = make(chan bool)

func generator(ctx context.Context, rounds chan<- Round) {
	var questions = []string{"2+2=", "2+3=", "2+4=", "2+5"}
	var answers = []string{"4", "5", "6", "7"}

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
			idx := rand.Intn(len(questions))
			rounds <- Round{question: questions[idx], answer: answers[idx]}
		}
	}
}

func player(ctx context.Context, playerID int, gameRounds <-chan Round, results chan<- bool) {
	for {
		select {
		case <-ctx.Done():
			return
		case round := <-gameRounds:
			fmt.Printf("player %v -> %v ", playerID, round.question)
			time.Sleep(600 * time.Millisecond)
			var answer = strconv.Itoa(rand.Intn(9))
			correct := round.answer == answer
			fmt.Printf("%v\n", answer)
			results <- correct
		}
	}
}
func counter(ctx context.Context, players int) {
	correctAnswers := 0
	for i := 0; i < players; i++ {
		select {
		case <-ctx.Done():
			return
		case isCorrect := <-results:
			if isCorrect {
				correctAnswers++
			}
		}
	}
	ctx.Done()
}

func Run() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go generator(ctx, rounds)

	players := 5
	for i := 1; i <= players; i++ {
		go player(ctx, i, rounds, results)
	}

	go counter(ctx, players)
	<-done
	cancel()
	<-ctx.Done()
}
