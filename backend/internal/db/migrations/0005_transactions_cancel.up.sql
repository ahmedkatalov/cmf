-- 0005_transactions_cancel.up.sql

ALTER TABLE transactions
ADD COLUMN IF NOT EXISTS is_cancelled BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN IF NOT EXISTS cancelled_at TIMESTAMP NULL,
ADD COLUMN IF NOT EXISTS cancelled_by UUID NULL REFERENCES users(id) ON DELETE SET NULL,
ADD COLUMN IF NOT EXISTS cancel_reason TEXT NULL;

CREATE INDEX IF NOT EXISTS idx_tx_cancelled ON transactions(is_cancelled);
