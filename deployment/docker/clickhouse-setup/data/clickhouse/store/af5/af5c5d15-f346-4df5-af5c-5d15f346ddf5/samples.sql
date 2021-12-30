ATTACH TABLE _ UUID 'cc6192fc-ba17-4b61-8c61-92fcba17eb61'
(
    `fingerprint` UInt64 CODEC(DoubleDelta, LZ4),
    `timestamp_ms` Int64 CODEC(DoubleDelta, LZ4),
    `value` Float64 CODEC(Gorilla, LZ4)
)
ENGINE = MergeTree
PARTITION BY toDate(timestamp_ms / 1000)
ORDER BY (fingerprint, timestamp_ms)
SETTINGS index_granularity = 8192
