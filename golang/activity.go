package app

import (
	"context"
	"log"
)

func RetrieveProfile(ctx context.Context, data CalculationConfig) (ScoreProfile, error) {
	log.Printf("Retrieving profile %s.\n\n",
		data.ProfileID,
	)

	profile := ScoreProfile{
		SpaceID:  "fooid:4242",
		Name:     "Awesome profile",
		NodeType: "Person",
		DimensionWeights: map[string]float64{
			"DimGo1":     0.42,
			"DimGo2":     0.02,
			"DimPython1": 0.95,
		},
	}

	return profile, nil
}

func CalculateDimensions(ctx context.Context, config CalculationConfig, profile ScoreProfile) ([]Dimension, error) {
	log.Printf("Calculating dimensions for profile %s, node type %s.\n\n",
		config.ProfileID,
		profile.NodeType,
	)

	result := []Dimension{
		{
			Name:        "DimGo1",
			Value:       0.01,
			Explanation: "I don't trust it",
		},
		{
			Name:        "DimGo2",
			Value:       0.9,
			Explanation: "This looks fine",
		},
	}

	return result, nil
}

func CalculateScore(ctx context.Context, config CalculationConfig, profile ScoreProfile, dimensions []Dimension) (Score, error) {
	log.Printf("Calculating score for %s, node type %s.\n\n",
		config.ProfileID,
		profile.NodeType,
	)

	score := 0.0
	for _, dim := range dimensions {
		score += dim.Value * profile.DimensionWeights[dim.Name]
	}
	score /= float64(len(dimensions))

	result := Score{
		Score:      score,
		Dimensions: dimensions,
	}

	return result, nil
}
