-- +migrate Up
ALTER TABLE state.monitored_txs
    ADD COLUMN gas_offset DECIMAL(78, 0) NOT NULL DEFAULT 0;
ALTER TABLE state.monitored_txs ALTER COLUMN gas_offset DROP DEFAULT;
CREATE TABLE IF NOT EXISTS state.attestation_id
(
    attestation_id      BIGINT NOT NULL PRIMARY KEY,
    block_num           BIGINT NOT NULL REFERENCES state.block (block_num) ON DELETE CASCADE
);

-- +migrate Down
ALTER TABLE state.monitored_txs
    DROP COLUMN gas_offset;
DROP TABLE IF EXISTS state.attestation_id;