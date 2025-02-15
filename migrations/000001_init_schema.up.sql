CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id               BIGSERIAL PRIMARY KEY,
    username         VARCHAR(255) NOT NULL UNIQUE,
    hashed_password  TEXT NOT NULL,
    balance          INTEGER NOT NULL DEFAULT 0,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS products (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    price       INTEGER NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS purchases (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity   INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_purchases_user
      FOREIGN KEY (user_id)
      REFERENCES users (id)
      ON DELETE CASCADE,

    CONSTRAINT fk_purchases_product
      FOREIGN KEY (product_id)
      REFERENCES products (id)
      ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS transfers (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_user_id  BIGINT NOT NULL,
    to_user_id    BIGINT NOT NULL,
    amount        INTEGER NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_transfers_from_user
      FOREIGN KEY (from_user_id)
      REFERENCES users (id)
      ON DELETE CASCADE,

    CONSTRAINT fk_transfers_to_user
      FOREIGN KEY (to_user_id)
      REFERENCES users (id)
      ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_users_username
  ON users (username);

CREATE INDEX IF NOT EXISTS idx_transfers_from_user
  ON transfers (from_user_id);

CREATE INDEX IF NOT EXISTS idx_transfers_to_user
  ON transfers (to_user_id);

CREATE INDEX IF NOT EXISTS idx_transfers_created_at
  ON transfers (created_at);
