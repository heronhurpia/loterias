CREATE TABLE if not exists roles (
  id serial PRIMARY KEY,
  alias VARCHAR(32) NOT NULL,
  description VARCHAR(32) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);