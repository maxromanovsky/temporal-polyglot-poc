package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func ScoreCalculation(ctx workflow.Context, input ScoreProfile) (string, error) {

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
	var profileOutput string

	profileErr := workflow.ExecuteActivity(ctx, RetrieveProfile, input).Get(ctx, &profileOutput)

	if profileErr != nil {
		return "", profileErr
	}

	// Calculate golang dimensions
	var golangDimensionOutput string
	// todo: pass profile as an input
	// todo: if dimension weight is zero, then skip calculation entirely
	golangDimensionErr := workflow.ExecuteActivity(ctx, CalculateDimensions, input).Get(ctx, &golangDimensionOutput)

	if golangDimensionErr != nil {
		return "", golangDimensionErr
	}

	//if golangDimensionErr != nil {
	//	// The deposit failed; put money back in original account.
	//
	//	var result string
	//
	//	refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &result)
	//
	//	if refundErr != nil {
	//		return "",
	//			fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
	//				input.TargetAccount, golangDimensionErr, input.SourceAccount, refundErr)
	//	}
	//
	//	return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
	//		input.TargetAccount, input.SourceAccount, golangDimensionErr)
	//}

	// Calculate golang dimensions
	var scoreOutput string
	// todo: pass profile as an input
	// todo: if dimension weight is zero, then skip calculation entirely
	scoreErr := workflow.ExecuteActivity(ctx, CalculateScore, input).Get(ctx, &scoreOutput)

	if scoreErr != nil {
		return "", scoreErr
	}

	result := fmt.Sprintf("Score calculation complete for profile %s, score %s, dimensions %s)", profileOutput, scoreOutput, golangDimensionOutput)
	return result, nil
}
