ALTER SYSTEM SET max_connections = 300;

CREATE EXTENSION pg_trgm;

ALTER DATABASE rinha SET synchronous_commit=OFF;
-- using 25% of memory as suggested in the docs:
--    https://www.postgresql.org/docs/9.1/runtime-config-resource.html
ALTER SYSTEM SET shared_buffers TO "425MB";

CREATE TABLE pessoas (
	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
	apelido VARCHAR ( 32 ) UNIQUE NOT NULL,
	nome VARCHAR ( 100 ) NOT NULL,
	nascimento DATE NOT null,
	stack  VARCHAR []
);

CREATE INDEX pessoas_search_index_idx ON pessoas (apelido, nome, stack);