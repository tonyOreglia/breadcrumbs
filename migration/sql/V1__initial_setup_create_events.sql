CREATE TABLE IF NOT EXISTS numbers(
  number        	  TEXT NOT NULL,
  file_ref      	  UUID NOT NULL,
  country_ioc_code 	TEXT NOT NULL,
  PRIMARY KEY       (number, file_ref)
);

CREATE TABLE IF NOT EXISTS fixed_numbers (
  original_number   TEXT NOT NULL,
  changes           TEXT NOT NULL,
  fixed_number      TEXT NOT NULL,
  file_ref          UUID NOT NULL,
  PRIMARY KEY       (original_number, file_ref)
);

CREATE TABLE IF NOT EXISTS rejected_numbers (
  number        TEXT NOT NULL,
  file_ref      UUID NOT NULL,
  PRIMARY KEY   (number, file_ref)
);


GRANT ALL PRIVILEGES ON TABLE numbers TO breadcrumbs;
GRANT ALL PRIVILEGES ON TABLE fixed_numbers TO breadcrumbs;
GRANT ALL PRIVILEGES ON TABLE rejected_numbers TO breadcrumbs;