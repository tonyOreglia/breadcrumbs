ALTER TABLE breadcrumbs DROP COLUMN date_created_unix;

ALTER TABLE breadcrumbs ADD COLUMN ts timestamp DEFAULT CURRENT_TIMESTAMP;
