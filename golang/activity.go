package app

import (
	"context"
	"log"
)

func RetrieveProfile(ctx context.Context, data ScoreProfile) (string, error) {
	log.Printf("Retrieving profile %s.\n\n",
		data.ProfileID,
	)

	//referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	//confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return "profile data that should not be string", nil
}

func CalculateDimensions(ctx context.Context, data ScoreProfile) (string, error) {
	log.Printf("Calculating dimensions for %s.\n\n",
		data.ProfileID,
	)

	//referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	//confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return "dimension data that should not be string", nil
}

func CalculateScore(ctx context.Context, data ScoreProfile) (string, error) {
	log.Printf("Calculating score for %s.\n\n",
		data.ProfileID,
	)

	//referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	//confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return "score data that should not be string", nil
}
