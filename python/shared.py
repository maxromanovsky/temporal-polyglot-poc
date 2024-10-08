# @@@SNIPSTART python-money-transfer-project-template-shared
from dataclasses import dataclass

PYTHON_DIMENSION_CALCULATION_TASK_QUEUE = "PYTHON_DIMENSION_CALCULATION_TASK_QUEUE"

@dataclass
class CalculationConfig:
    ProfileID: str
    ReferenceID: str


@dataclass
class ScoreProfile:
    SpaceID: str
    Name: str
    NodeType: str
    DimensionWeights: dict[str, float]


@dataclass
class Dimension:
    Name: str
    Value: float
    Explanation: str
