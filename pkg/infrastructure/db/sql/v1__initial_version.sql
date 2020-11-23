
CREATE SEQUENCE dummy_id START 1;

CREATE TABLE dummy (
    id int NOT NULL DEFAULT nextval('dummy_id'),
    name varchar(100) not null
);

CREATE INDEX dummy_value ON dummy(name);

