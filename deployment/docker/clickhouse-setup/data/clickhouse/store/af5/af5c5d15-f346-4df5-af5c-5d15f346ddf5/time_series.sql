ATTACH TABLE _ UUID '55dd3a19-1531-4a39-95dd-3a1915314a39'
(
    `date` Date CODEC(DoubleDelta, LZ4),
    `fingerprint` UInt64 CODEC(DoubleDelta, LZ4),
    `labels` String CODEC(ZSTD(5))
)
ENGINE = ReplacingMergeTree
PARTITION BY date
ORDER BY fingerprint
SETTINGS index_granularity = 8192
