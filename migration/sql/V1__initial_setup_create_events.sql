CREATE TABLE IF NOT EXISTS notes(
  id          SERIAL PRIMARY KEY,
  note        TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS breadcrumbs(
  id          SERIAL PRIMARY KEY,
  data_type   CHAR(6) NOT NULL,
  data_id     INTEGER REFERENCES notes (id),
  geog        GEOGRAPHY(POINTZ, 4326)
);

GRANT ALL PRIVILEGES ON TABLE breadcrumbs TO breadcrumbs_user;
GRANT ALL PRIVILEGES ON TABLE notes TO breadcrumbs_user;
