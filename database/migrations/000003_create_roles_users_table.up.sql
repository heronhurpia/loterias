CREATE TABLE roles_users (
  user_id integer NOT NULL REFERENCES users(id),
  role_id integer NOT NULL REFERENCES roles(id),
  PRIMARY KEY (user_id, role_id)
);
