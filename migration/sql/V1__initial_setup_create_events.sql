CREATE TABLE IF NOT EXISTS breadcrumbs(
  id          SERIAL PRIMARY KEY,           
  data_type   CHAR(6) NOT NULL,
  data_id	    INTEGER,
  geog        GEOGRAPHY(POINT, 4326)
);

GRANT ALL PRIVILEGES ON TABLE breadcrumbs TO breadcrumbs;
