ALTER TABLE users
ADD COLUMN IF NOT EXISTS isAdmin boolean DEFAULT FALSE;

ALTER TABLE portfolio
ADD COLUMN IF NOT EXISTS name varchar DEFAULT 'Моё портфолио';

ALTER TABLE symbol
ADD UNIQUE(abbr);

ALTER TABLE portfolio
ALTER COLUMN account TYPE DECIMAL;

ALTER TABLE portfolio
ADD CHECK (account >= 0);

ALTER TABLE active_share
ADD CHECK (number >= 0);

ALTER TABLE  deal
ALTER COLUMN date TYPE timestamp;