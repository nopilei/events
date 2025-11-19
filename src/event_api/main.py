import logging
from typing import Annotated

from fastapi import FastAPI, HTTPException, Depends

from src.brokers.base import AsyncProducer
from src.brokers.exceptions import BrokerError
from src.brokers.providers import get_producer, get_manager
from src.event_schemas.base import BaseEventData
from src.event_schemas.registry import events_registry

from prometheus_fastapi_instrumentator import Instrumentator

logger = logging.getLogger(__name__)

app = FastAPI(title="Event API")

instrumentator = Instrumentator().instrument(app)


@app.on_event("startup")
async def startup_event():
    instrumentator.expose(app)
    await get_manager().start_producer()
    logger.info("producer started")


@app.on_event("shutdown")
async def shutdown_event():
    await get_manager().stop_producer()
    logger.info("Kafka producer stopped")


@app.get("/health")
async def health_check():
    return {"status": "healthy"}


@app.post("/events")
async def send_event(event: BaseEventData, producer: Annotated[AsyncProducer, Depends(get_producer)]):
    if event.event_type not in events_registry:
        raise HTTPException(status_code=400, detail="Invalid event_type.")

    event_class = events_registry[event.event_type]
    try:
        event_class(**event.data.model_dump())
    except ValueError:
        raise HTTPException(status_code=400, detail="Invalid data for given type.")

    try:
        await producer.send(event)
        return {"status": "ok"}
    except BrokerError as e:
        logger.error(f"Kafka error: {e}")
        raise HTTPException(status_code=500, detail="Failed to send event")
