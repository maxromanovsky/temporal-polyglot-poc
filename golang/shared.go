package app

const ScoreCalculationTaskQueueName = "SCORE_CALCULATION_TASK_QUEUE"

type CalculationConfig struct {
	ProfileID   string
	ReferenceID string
}

type ScoreProfile struct {
	SpaceID          string
	Name             string
	NodeType         string
	DimensionWeights map[string]float64
}

type Dimension struct {
	Name        string
	Value       float64
	Explanation string
}

type Score struct {
	Score      float64
	Dimensions []Dimension
}
