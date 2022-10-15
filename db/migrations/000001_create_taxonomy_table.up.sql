CREATE TABLE IF NOT EXISTS taxonomy (
  id serial PRIMARY KEY,
  language_key VARCHAR(10) NOT NULL,
  key VARCHAR(255) NOT NULL,
  value VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT lang_key UNIQUE (language_key, key)
);
