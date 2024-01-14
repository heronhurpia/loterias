CREATE TABLE if not exists roles (
  id serial PRIMARY KEY,
  alias VARCHAR(31) NOT NULL,
  description VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);