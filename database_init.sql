CREATE TABLE tasks (
	id integer GENERATED ALWAYS AS IDENTITY NOT NULL,
	"name" varchar NOT NULL,
	is_done boolean DEFAULT false NOT NULL,
	CONSTRAINT newtable_pk PRIMARY KEY (id)
);