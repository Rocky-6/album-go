CREATE TABLE IF NOT EXISTS album (
  id SERIAL PRIMARY KEY NOT NULL ,
  title TEXT NOT NULL,
  artist TEXT NOT NULL,
  price NUMERIC(5, 2) NOT NULL
);