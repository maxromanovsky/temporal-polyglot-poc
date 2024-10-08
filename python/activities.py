import random

from temporalio import activity

from shared import CalculationConfig, ScoreProfile, Dimension


class DimensionActivities:

    @activity.defn
    async def calculate_dimensions(self, config: CalculationConfig, profile: ScoreProfile) -> list[Dimension]:
        print(f"Calculating dimensions for id {config.ProfileID}, profile {profile.Name}")

        should_fail = random.randint(0, 1)
        if should_fail:
            raise Exception("oops. Python failed!")

        return [Dimension("DimPython1", 0.91, "Because it's Python, baby!")]
        # todo: exception
