-- public.activities definition

-- Drop table

-- DROP TABLE public.activities;

CREATE TABLE public.activities (
	id bigserial NOT NULL,
	title varchar NOT NULL,
	description text NOT NULL,
	start_date timestamptz NOT NULL,
	end_date timestamptz NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	participants text NOT NULL,
	skill_id int8 NOT NULL,
	CONSTRAINT activities_pk PRIMARY KEY (id),
	CONSTRAINT activities_fk FOREIGN KEY (skill_id) REFERENCES public.skills(id)
);
CREATE INDEX activities_skill_id_idx ON public.activities USING hash (skill_id);

create or replace function list_activities(p_skill_id int8, sort_by character varying, order_by character varying, limitVal int, offsetVal int)
returns TABLE (LIKE activities) AS 
$$
BEGIN 
	RETURN QUERY EXECUTE format(
	'SELECT * FROM activities 
	WHERE skill_id = $1 
	ORDER BY %I %s
	LIMIT $2 OFFSET $3', sort_by, order_by)
	USING p_skill_id, limitVal, offsetVal;
END
$$ language plpgsql;