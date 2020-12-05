CREATE TABLE IF NOT EXISTS users (
 id SERIAL NOT NULL PRIMARY KEY,
 CreateDate timestamp,
 username varchar(160) NOT NULL UNIQUE,
 password text NOT NULL,
 email varchar(160) NOT NULL,
 data jsonb)