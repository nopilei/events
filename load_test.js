import http from "k6/http";

export const options = {
  scenarios: {
    high_rps: {
      executor: "constant-arrival-rate",
      rate: 1400,
      timeUnit: "1s",
      duration: "30s",
      preAllocatedVUs: 500,
      maxVUs: 2000,
    },
  },
};

const URL = __ENV.TARGET_URL || "http://109.196.103.36:8000/events";
const payload = JSON.stringify({
  event_type: "order_created",
  data: {
    id: 1,
    user_id: 1,
    product_id: 1,
    amount: 1,
  },
});

export default function () {
  http.post(URL, payload, {
    headers: { "Content-Type": "application/json" },
    timeout: '200ms',
  });
}
