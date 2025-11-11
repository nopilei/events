import asyncio
import os
import random
import concurrent.futures

import httpx
from polyfactory.factories.pydantic_factory import ModelFactory

from src.event_schemas.base import BaseEventData
from src.event_schemas.registry import events_registry

URL = f'http://{os.getenv("API_HOST")}:{os.getenv("API_PORT")}/events'


def run_in_process(t):
    asyncio.run(run_tasks())


async def run_tasks():
    tasks = [_task() for _ in range(10)]
    await asyncio.gather(*tasks)


async def _task():
    while True:
        event_type, event_class = random.choice(list(events_registry.items()))
        event_json_data = ModelFactory.create_factory(event_class).build().model_dump()
        event = BaseEventData(event_type=event_type, data=event_json_data).model_dump()
        async with httpx.AsyncClient() as client:
            response = await client.post(URL, json=event)
            response.raise_for_status()
            print(f"Status Code: {response.status_code} {os.getpid()}")
            print("Response JSON:")
            print(response.json())


if __name__ == '__main__':
    with concurrent.futures.ProcessPoolExecutor() as executor:
        executor.map(run_in_process, range(10))

    asyncio.run(run_tasks())
