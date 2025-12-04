DROP INDEX IF EXISTS idx_transactions_user;
DROP INDEX IF EXISTS idx_positions_user_market;
DROP INDEX IF EXISTS idx_orders_market_status;
DROP INDEX IF EXISTS idx_markets_category;
DROP INDEX IF EXISTS idx_accounts_user_id;

DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS market_resolutions;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS positions;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS markets;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "pgcrypto";
