CREATE TABLE ptrs (
    address VARCHAR(15) PRIMARY KEY,
    record VARCHAR,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE EXTENSION pg_trgm;

CREATE INDEX trgm_idx_ptrs_record ON ptrs USING gin (record gin_trgm_ops);