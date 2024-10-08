import asyncio
import logging

import temporalio.runtime
from temporalio.runtime import LoggingConfig, TelemetryFilter
from temporalio.client import Client
from temporalio.worker import Worker

from activities import DimensionActivities
from shared import PYTHON_DIMENSION_CALCULATION_TASK_QUEUE
# from workflows import MoneyTransfer


async def main() -> None:
    client: Client = await Client.connect("localhost:7233", namespace="default")
    # Run the worker
    activities = DimensionActivities()
    worker: Worker = Worker(
        client,
        task_queue=PYTHON_DIMENSION_CALCULATION_TASK_QUEUE,
        # workflows=[MoneyTransfer],
        activities=[activities.calculate_dimensions],
    )
    await worker.run()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    temporal_logger = logging.getLogger("temporal")

    runtime = temporalio.runtime.Runtime(
        telemetry=temporalio.runtime.TelemetryConfig(
            logging=temporalio.runtime.LoggingConfig(
                filter=temporalio.runtime.TelemetryFilter(core_level="DEBUG", other_level="INFO"),
                forwarding=temporalio.runtime.LogForwardingConfig(logger=temporal_logger),
            )
        )
    )

    temporalio.runtime.Runtime.set_default(runtime)
    # = LoggingConfig(
    # filter=TelemetryFilter(core_level="INFO", other_level="INFO")
    logging.info("starting worker...")
    asyncio.run(main())
