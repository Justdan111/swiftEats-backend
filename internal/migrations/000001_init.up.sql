-- enable uuid extension (Postgres)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- users
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  name TEXT,
  phone TEXT,
  created_at timestamptz DEFAULT now()
);

-- migrations table (migrate creates this itself; kept here in case)
