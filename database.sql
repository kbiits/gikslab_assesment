--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8
-- Dumped by pg_dump version 14.2

-- Started on 2023-07-29 21:25:24

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3367 (class 1262 OID 16385)
-- Name: gikslab; Type: DATABASE; Schema: -; Owner: root
--

CREATE DATABASE gikslab WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


\connect gikslab

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: root
--

CREATE SCHEMA IF NOT EXISTS public;


--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: root
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 214 (class 1259 OID 16426)
-- Name: activities; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.activities (
    id bigint NOT NULL,
    title character varying NOT NULL,
    description text NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    participants text NOT NULL,
    skill_id bigint NOT NULL
);


--
-- TOC entry 215 (class 1255 OID 16442)
-- Name: list_activities(bigint, character varying, character varying, integer, integer); Type: FUNCTION; Schema: public; Owner: root
--

CREATE FUNCTION public.list_activities(p_skill_id bigint, sort_by character varying, order_by character varying, limitval integer, offsetval integer) RETURNS TABLE("like" public.activities)
    LANGUAGE plpgsql
    AS $_$
BEGIN 
	RETURN QUERY EXECUTE format(
	'SELECT * FROM activities 
	WHERE skill_id = $1 
	ORDER BY %I %s
	LIMIT $2 OFFSET $3', sort_by, order_by)
	USING p_skill_id, limitVal, offsetVal;
END
$_$;


--
-- TOC entry 213 (class 1259 OID 16425)
-- Name: activities_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.activities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 213
-- Name: activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.activities_id_seq OWNED BY public.activities.id;


--
-- TOC entry 210 (class 1259 OID 16398)
-- Name: skills; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.skills (
    id bigint NOT NULL,
    name character varying NOT NULL
);


--
-- TOC entry 209 (class 1259 OID 16397)
-- Name: skills_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.skills_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 209
-- Name: skills_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.skills_id_seq OWNED BY public.skills.id;


--
-- TOC entry 212 (class 1259 OID 16408)
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text NOT NULL,
    email character varying NOT NULL,
    username character varying NOT NULL,
    password text NOT NULL,
    profile character varying NOT NULL,
    skill_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- TOC entry 211 (class 1259 OID 16407)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 211
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3202 (class 2604 OID 16429)
-- Name: activities id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.activities ALTER COLUMN id SET DEFAULT nextval('public.activities_id_seq'::regclass);


--
-- TOC entry 3198 (class 2604 OID 16401)
-- Name: skills id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.skills ALTER COLUMN id SET DEFAULT nextval('public.skills_id_seq'::regclass);


--
-- TOC entry 3199 (class 2604 OID 16411)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3361 (class 0 OID 16426)
-- Dependencies: 214
-- Data for Name: activities; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.activities VALUES (1, 'Activity 1', 'Desc of Activity 1', '2023-07-29 18:00:00+00', '2023-07-29 18:10:00+00', '2023-07-29 14:22:17.448737+00', '2023-07-29 14:22:17.448737+00', '[5]', 2);
INSERT INTO public.activities VALUES (2, 'Activity Coding', 'Desc of Activity Coding', '2023-07-29 18:00:00+00', '2023-07-29 20:10:00+00', '2023-07-29 14:23:26.087768+00', '2023-07-29 14:23:26.087768+00', '[2,6]', 1);


--
-- TOC entry 3357 (class 0 OID 16398)
-- Dependencies: 210
-- Data for Name: skills; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.skills VALUES (1, 'coding');
INSERT INTO public.skills VALUES (2, 'design');


--
-- TOC entry 3359 (class 0 OID 16408)
-- Dependencies: 212
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.users VALUES (1, 'Nabiel', 'nabiel@email.com', 'nabiel', '$2a$10$R3DsXi4iTi25gobSmsB0iutbY1LNOefpmy1H.7wZlRO6I./bKtAMu', 'board', NULL, '2023-07-29 14:16:59.150382+00', '2023-07-29 14:16:59.150382+00');
INSERT INTO public.users VALUES (2, 'Expert 1', 'expert_1@email.com', 'expert_1', '$2a$10$OTuDpHgbqjeJPrsOLcKvXOJcL3jKclsxENbG1IbM5iX5HGRZgm3ai', 'expert', 1, '2023-07-29 14:17:50.941995+00', '2023-07-29 14:17:50.941995+00');
INSERT INTO public.users VALUES (3, 'Expert 2', 'expert_2@email.com', 'expert_2', '$2a$10$WOxuBTrTxYwqW2LjrzFqW.S45XSaDh7reTQHS65RR5keVRAGUfJJG', 'expert', 2, '2023-07-29 14:17:58.516473+00', '2023-07-29 14:17:58.516473+00');
INSERT INTO public.users VALUES (5, 'Competitor 2', 'competitor_2@email.com', 'competitor_2', '$2a$10$zG9tI2mlL3RulMGx9aQiMed9rycq6KHkDXl0fXEJGG9/9RmwYbbjq', 'competitor', 2, '2023-07-29 14:18:29.948545+00', '2023-07-29 14:18:29.948545+00');
INSERT INTO public.users VALUES (4, 'Competitor 1', 'competitor_1@email.com', 'competitor_1', '$2a$10$AdC2jT2Lu8bzeLt4c5xt3.pj5UgANeGCDQIOWFeP5cmmUpupHOzN.', 'competitor', NULL, '2023-07-29 14:18:14.583536+00', '2023-07-29 14:18:14.583536+00');
INSERT INTO public.users VALUES (6, 'Trainer 1', 'trainer_1@email.com', 'trainer_1', '$2a$10$Ize5k4eA8OR/6HIMn2cvOuVRv6U4AQgfz10yMk4b8PGQbcLPnAVpC', 'trainer', 1, '2023-07-29 14:19:10.908895+00', '2023-07-29 14:19:10.908895+00');
INSERT INTO public.users VALUES (7, 'Trainer 2', 'trainer_2@email.com', 'trainer_2', '$2a$10$hzqWZliE69vBKk13PgZXp.XGhOZbo7slHWHpsK2YcPTen.w.JcGE6', 'trainer', 2, '2023-07-29 14:19:16.810458+00', '2023-07-29 14:19:16.810458+00');


--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 213
-- Name: activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.activities_id_seq', 2, true);


--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 209
-- Name: skills_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.skills_id_seq', 2, true);


--
-- TOC entry 3374 (class 0 OID 0)
-- Dependencies: 211
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.users_id_seq', 7, true);


--
-- TOC entry 3213 (class 2606 OID 16435)
-- Name: activities activities_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_pk PRIMARY KEY (id);


--
-- TOC entry 3207 (class 2606 OID 16405)
-- Name: skills skills_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_pk PRIMARY KEY (id);


--
-- TOC entry 3210 (class 2606 OID 16417)
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- TOC entry 3214 (class 1259 OID 16441)
-- Name: activities_skill_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX activities_skill_id_idx ON public.activities USING hash (skill_id);


--
-- TOC entry 3205 (class 1259 OID 16406)
-- Name: skills_name_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE UNIQUE INDEX skills_name_idx ON public.skills USING btree (name);


--
-- TOC entry 3208 (class 1259 OID 16423)
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- TOC entry 3211 (class 1259 OID 16424)
-- Name: users_username_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);


--
-- TOC entry 3216 (class 2606 OID 16436)
-- Name: activities activities_fk; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_fk FOREIGN KEY (skill_id) REFERENCES public.skills(id);


--
-- TOC entry 3215 (class 2606 OID 16418)
-- Name: users users_fk; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk FOREIGN KEY (skill_id) REFERENCES public.skills(id) ON DELETE SET NULL;


-- Completed on 2023-07-29 21:25:24

--
-- PostgreSQL database dump complete
--

