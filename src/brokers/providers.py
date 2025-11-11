from typing import Annotated

from fastapi import Depends

from src.brokers.base import AsyncProducer, AsyncProducerManager
from src.brokers.kafka.manager import KafkaManager

kafka_manager = KafkaManager()


def get_manager() -> AsyncProducerManager:
    return kafka_manager


def get_producer() -> AsyncProducer:
    return get_manager().get_producer()


ProducerDep = Annotated[AsyncProducer, Depends(get_producer)]
