CREATE TABLE IF NOT EXISTS tokens
(
    id SERIAL NOT NULL PRIMARY KEY,
    creationdate timestamp NOT NULL,
    expirationdate timestamp,
    token text,
    userid INTEGER NOT NULL,
    CONSTRAINT fk_alert
      FOREIGN KEY(userid) 
	  REFERENCES users(id)
	  ON DELETE SET NULL)
