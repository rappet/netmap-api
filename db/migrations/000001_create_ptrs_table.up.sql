CREATE TABLE ptrs (
    address VARCHAR(15) PRIMARY KEY,
    record VARCHAR,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    domain VARCHAR GENERATED ALWAYS AS (lower(substring(record from '[\w-]+\.[\w-]+$'))) STORED
);

CREATE EXTENSION pg_trgm;

CREATE INDEX trgm_idx_ptrs_record ON ptrs USING gin (record gin_trgm_ops);
CREATE INDEX idx_ptrs_domain ON ptrs(domain);