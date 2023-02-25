
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id		    SERIAL PRIMARY KEY ,
    email       TEXT UNIQUE NOT NULL,
    password    TEXT NOT NULL,
    role        TEXT NOT NULL CHECK (role = 'ADMIN' OR role = 'MEMBER' OR role = 'TECHNICIAN'),
    firstName   TEXT NOT NULL,
    lastName    TEXT NOT NULL,
    company     TEXT,
    designation TEXT
);
