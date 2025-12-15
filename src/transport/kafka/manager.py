from aiokafka import AIOKafkaProducer
from pydantic.v1 import BaseSettings

from src.brokers.base import AsyncProducerManager
from src.brokers.kafka.producer import KafkaProducer


class KafkaSettings(BaseSettings):
    KAFKA_BOOTSTRAP_SERVERS: str


class KafkaManager(AsyncProducerManager):
    def __init__(self):
        self.kafka_producer = None
        self.kafka_settings = KafkaSettings()

    def get_producer(self) -> KafkaProducer:
        if not self.kafka_producer:
            aioproducer = AIOKafkaProducer(
                bootstrap_servers=self.kafka_settings.KAFKA_BOOTSTRAP_SERVERS,
            )
            self.kafka_producer = KafkaProducer(aioproducer)
        return self.kafka_producer

    async def start_producer(self) -> None:
        await self.get_producer().producer.start()

    async def stop_producer(self) -> None:
        await self.get_producer().producer.stop()
