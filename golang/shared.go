package app

const ScoreCalculationTaskQueueName = "SCORE_CALCULATION_TASK_QUEUE"

type ScoreProfile struct {
	//SpaceID          string
	//NodeType         string
	//DimensionWeights map[string]float64
	ProfileID   string
	ReferenceID string
}
