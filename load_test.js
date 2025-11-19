import http from 'k6/http';
import { sleep } from 'k6';

// Количество параллельных "процессов"
export let options = {
    // можно сделать stages для ступенчатой нагрузки
  scenarios: {
        constant_request_rate: {
            executor: 'constant-arrival-rate',
            rate: 3000,
            timeUnit: '1s', // 1000 итераций в секунду, т.е.1000 запросов секунду
            duration: '1m',
	    preAllocatedVUs: 500,
        }
    }
};

const URL = __ENV.TARGET_URL || "http://localhost:8000/events";

// Список типов событий
const eventTypes = [
    "UserCreated",
    "UserUpdated",
    "OrderCreated",
];

// Генерация случайного события
function randomEvent() {
    const type = eventTypes[Math.floor(Math.random() * eventTypes.length)];

    switch (type) {
        case "UserCreated":
            return {
                event_type: "user_created",
                data: {
                    id: Math.floor(Math.random() * 1_000_000),
                    name: "Test User",
                }
            };
        case "UserUpdated":
            return {
                event_type: "user_updated",
                data: {
                    id: Math.floor(Math.random() * 1_000_000),
                    name: "Test User",
                }
            };

        case "OrderCreated":
            return {
                event_type: "order_created",
                data: {
                    id: Math.floor(Math.random() * 1000000),
                    user_id: Math.floor(Math.random() * 1000000),
                    product_id: Math.floor(Math.random() * 1000000),
                    amount: Math.floor(Math.random() * 1000),
                }
            };
    }
}

// Основная логика нагрузки
export default function () {
    const evt = randomEvent();

    const res = http.post(URL, JSON.stringify(evt), {
        headers: { "Content-Type": "application/json" },
    });

    // Можно замедлять, чтобы стабилизировать RPS
    // sleep(0.1);
}
