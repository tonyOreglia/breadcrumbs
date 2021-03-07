CREATE TABLE IF NOT EXISTS bc_user(
  id            SERIAL PRIMARY KEY,
  full_name     TEXT UNIQUE NOT NULL,
  ts timestamp  DEFAULT CURRENT_TIMESTAMP
);

GRANT ALL PRIVILEGES ON TABLE bc_user TO breadcrumbs_user;

ALTER TABLE breadcrumbs ADD COLUMN user_id INTEGER REFERENCES bc_user (id) DEFAULT NULL;
