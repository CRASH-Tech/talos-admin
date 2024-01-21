--CREATE USER talos-admin;
--CREATE DATABASE talos-admin;
--GRANT ALL PRIVILEGES ON DATABASE talos-admin TO talos-admin;

/*
CREATE DATABASE "talos-admin"
    WITH
    OWNER = "talos-admin"
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
*/

DROP TABLE IF EXISTS clusters;
CREATE TABLE IF NOT EXISTS clusters (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS nodes;
CREATE TABLE IF NOT EXISTS nodes (
	id serial PRIMARY KEY,
	cluster_id integer,
	template_id integer,
	machine_config text,
	name VARCHAR ( 50 ) NOT NULL,
	created_on TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS variables;
CREATE TABLE IF NOT EXISTS variables (
	id serial PRIMARY KEY,
	cluster_id integer,
	key VARCHAR ( 50 ),
	value text,
	created_on TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS templates;
CREATE TABLE IF NOT EXISTS templates (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ),
	data text,
	created_on TIMESTAMP NOT NULL
);

INSERT INTO clusters (name, created_on) VALUES ('k-test-a', NOW());
INSERT INTO clusters (name, created_on) VALUES ('k-test-b', NOW());

INSERT INTO nodes (cluster_id, template_id, machine_config, name, created_on) VALUES (1, 1, 'k-test-a-m1 mchine config sample','k-test-a-m1', NOW());
INSERT INTO nodes (cluster_id, template_id, machine_config, name, created_on) VALUES (1, 1, 'k-test-a-m2 mchine config sample','k-test-a-m2', NOW());
INSERT INTO nodes (cluster_id, template_id, machine_config, name, created_on) VALUES (1, 1, 'k-test-a-m3 mchine config sample','k-test-a-m3', NOW());

INSERT INTO nodes (cluster_id, template_id, machine_config, name, created_on) VALUES (2, 2, 'k-test-b-m1 mchine config sample','k-test-b-m1', NOW());
INSERT INTO nodes (cluster_id, template_id, machine_config, name, created_on) VALUES (2, 2, 'k-test-b-m2 mchine config sample','k-test-b-m2', NOW());
INSERT INTO nodes (cluster_id, template_id, machine_config, name, created_on) VALUES (2, 2, 'k-test-b-m3 mchine config sample','k-test-b-m3', NOW());

INSERT INTO templates (name, data, created_on) VALUES ('k-test-a-master', 'k-test-a master template', NOW());
INSERT INTO templates (name, data, created_on) VALUES ('k-test-b-master', 'k-test-b master template', NOW());

INSERT INTO variables (cluster_id, key, value, created_on) VALUES (1, 'VIP', '10.0.1.1', NOW());
INSERT INTO variables (cluster_id, key, value, created_on) VALUES (2, 'VIP', '10.0.2.1', NOW());