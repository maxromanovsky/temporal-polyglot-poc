package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"score-polyglot-go/app"
)

// @@@SNIPSTART money-transfer-project-template-go-worker
func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, app.ScoreCalculationTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(app.ScoreCalculation)
	w.RegisterActivity(app.RetrieveProfile)
	w.RegisterActivity(app.CalculateDimensions)
	w.RegisterActivity(app.CalculateScore)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}

// @@@SNIPEND
