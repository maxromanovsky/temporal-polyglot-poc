package app

import (
	"log"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func ScoreCalculation(ctx workflow.Context, input CalculationConfig) (Score, error) {

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    500, // 0 is unlimited retries
		// todo: Error types from Python?
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	// Retrieve profile.
	var profileOutput ScoreProfile

	profileErr := workflow.ExecuteActivity(ctx, RetrieveProfile, input).Get(ctx, &profileOutput)

	if profileErr != nil {
		return Score{}, profileErr
	}

	var dimensions []Dimension

	if profileOutput.DimensionWeights["DimGo1"] > 0 || profileOutput.DimensionWeights["DimGo2"] > 0 {
		// Calculate golang dimensions if weight > 0
		var golangDimensionOutput []Dimension
		golangDimensionErr := workflow.ExecuteActivity(ctx, CalculateDimensions, input, profileOutput).Get(ctx, &golangDimensionOutput)

		if golangDimensionErr != nil {
			return Score{}, golangDimensionErr
		}

		dimensions = append(dimensions, golangDimensionOutput...)
	}
	// Calculate score dimensions
	var scoreOutput Score
	scoreErr := workflow.ExecuteActivity(ctx, CalculateScore, input, profileOutput, dimensions).Get(ctx, &scoreOutput)

	if scoreErr != nil {
		return Score{}, scoreErr
	}

	log.Printf("Score calculation complete for profile %s, score %+v)", profileOutput.Name, scoreOutput)
	return scoreOutput, nil
}
