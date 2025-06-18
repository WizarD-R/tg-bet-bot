CREATE TABLE IF NOT EXISTS current_event(
   id serial PRIMARY KEY,
   name VARCHAR (300) NOT NULL,
   team_1 VARCHAR (100) NOT NULL,
   team_2 VARCHAR (100) NOT NULL,
   date TIMESTAMPTZ NOT NULL,
   result VARCHAR (5)
);

CREATE TABLE IF NOT EXISTS current_predictions(
   username VARCHAR (100) NOT NULL,
   id_event INT REFERENCES current_event(id),
   prediction VARCHAR (5) NOT NULL
);