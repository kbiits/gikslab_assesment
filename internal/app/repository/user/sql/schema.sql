-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id bigserial NOT NULL,
	"name" text NOT NULL,
	email varchar NOT NULL,
	username varchar NOT NULL,
	"password" text NOT NULL,
	profile varchar NOT NULL,
	skill_id int8 NULL,
	created_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_fk FOREIGN KEY (skill_id) REFERENCES public.skills(id) ON DELETE SET NULL
);
CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);
CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);