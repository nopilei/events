from aiokafka import AIOKafkaProducer

from src.brokers.base import AsyncProducer
from src.brokers.exceptions import BrokerError
from src.brokers.kafka.topics import topics
from src.event_schemas.base import BaseEventData
from src.event_schemas.registry import events_registry


class KafkaProducer(AsyncProducer):
    def __init__(self, producer: AIOKafkaProducer):
        self.producer = producer

    async def send(self, event: BaseEventData) -> None:
        topic = self.get_topic_name(event)
        data = event.model_dump_json().encode()
        try:
            await self.producer.send(topic=topic, value=data)
        except Exception as exc:
            raise BrokerError from exc

    def get_topic_name(self, event: BaseEventData) -> str:
        event_type = events_registry[event.event_type]
        return topics.get(event_type, event.event_type)
