CREATE TABLE pessoas (
	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
	apelido VARCHAR ( 32 ) UNIQUE NOT NULL,
	nome VARCHAR ( 100 ) NOT NULL,
	nascimento DATE NOT null,
	stack  VARCHAR []
);
