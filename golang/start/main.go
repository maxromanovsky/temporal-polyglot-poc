package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"log"
	"math/rand"
	"strconv"

	"score-polyglot-go/app"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	referenceID := rand.Int()

	input := app.ScoreProfile{
		//SpaceID:  "fooid:4242",
		//NodeType: "Person",
		//DimensionWeights: map[string]float64{
		//	"DimGo1":     0.42,
		//	"DimGo2":     0.2,
		//	"DimPython1": 0.95,
		//},
		ProfileID:   uuid.NewString(),
		ReferenceID: strconv.Itoa(referenceID),
	}

	options := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("score-calculation-%d", referenceID),
		TaskQueue: app.ScoreCalculationTaskQueueName,
	}

	log.Printf("Starting calculation for %s", input.ProfileID)

	we, err := c.ExecuteWorkflow(context.Background(), options, app.ScoreCalculation, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string

	err = we.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable to get Workflow result:", err)
	}

	log.Println(result)
}
