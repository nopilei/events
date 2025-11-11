import abc

from src.event_schemas.base import BaseEventData


class AsyncProducer(abc.ABC):

    @abc.abstractmethod
    async def send(self, event: BaseEventData) -> None:
        """Отправить сообщение в брокер."""


class AsyncProducerManager(abc.ABC):

    @abc.abstractmethod
    def get_producer(self) -> AsyncProducer:
        """Вернуть продюсера."""

    @abc.abstractmethod
    async def start_producer(self) -> None:
        """Запустить продюсер."""

    @abc.abstractmethod
    async def stop_producer(self) -> None:
        """Остановить продюсер."""
