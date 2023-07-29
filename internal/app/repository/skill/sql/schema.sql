-- public.skills definition

-- Drop table

-- DROP TABLE public.skills;

CREATE TABLE public.skills (
	id bigserial NOT NULL,
	"name" varchar NOT NULL,
	CONSTRAINT skills_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX skills_name_idx ON public.skills USING btree (name);