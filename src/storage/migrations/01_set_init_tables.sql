-- +goose Up
CREATE TABLE IF NOT EXISTS events_raw (
    event_type String,
    event_timestamp DateTime,
    payload JSON
)
ENGINE = MergeTree
ORDER BY (event_timestamp, event_type)
PARTITION BY event_type;

-- +goose ENVSUB ON
CREATE TABLE IF NOT EXISTS queue (
    event_type String,
    data JSON
) ENGINE = Kafka('${KAFKA_BOOTSTRAP_SERVERS}', 'orders,users', 'clickhouse_group', 'JSONEachRow');
-- +goose ENVSUB OFF

CREATE MATERIALIZED VIEW IF NOT EXISTS consumer TO events_raw
AS
    SELECT toDateTime(_timestamp) AS event_timestamp,
           event_type,
           data AS payload FROM queue
;


-- +goose Down
DROP TABLE queue;
DROP VIEW consumer;
DROP TABLE events_raw;