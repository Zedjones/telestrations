--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 11.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: difficulty_t; Type: TYPE; Schema: public; Owner: zedjones
--

CREATE TYPE public.difficulty_t AS ENUM (
    'Easy',
    'Medium',
    'Hard',
    'Hardcore'
);


ALTER TYPE public.difficulty_t OWNER TO zedjones;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: guess; Type: TABLE; Schema: public; Owner: zedjones
--

CREATE TABLE public.guess (
    user_id integer NOT NULL,
    word character varying(255) NOT NULL,
    round integer NOT NULL
);


ALTER TABLE public.guess OWNER TO zedjones;

--
-- Name: picture; Type: TABLE; Schema: public; Owner: zedjones
--

CREATE TABLE public.picture (
    user_id integer NOT NULL,
    svg text NOT NULL,
    round integer NOT NULL
);


ALTER TABLE public.picture OWNER TO zedjones;

--
-- Name: settings; Type: TABLE; Schema: public; Owner: zedjones
--

CREATE TABLE public.settings (
    id integer NOT NULL,
    time_limit integer DEFAULT 0,
    word_difficulty public.difficulty_t DEFAULT 'Easy'::public.difficulty_t
);


ALTER TABLE public.settings OWNER TO zedjones;

--
-- Name: user; Type: TABLE; Schema: public; Owner: zedjones
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    name character varying(24)
);


ALTER TABLE public."user" OWNER TO zedjones;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: zedjones
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO zedjones;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zedjones
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: guess; Type: TABLE DATA; Schema: public; Owner: zedjones
--

COPY public.guess (user_id, word, round) FROM stdin;
\.


--
-- Data for Name: picture; Type: TABLE DATA; Schema: public; Owner: zedjones
--

COPY public.picture (user_id, svg, round) FROM stdin;
\.


--
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: zedjones
--

COPY public.settings (id, time_limit, word_difficulty) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: zedjones
--

COPY public."user" (id, name) FROM stdin;
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zedjones
--

SELECT pg_catalog.setval('public.user_id_seq', 1, false);


--
-- Name: guess guess_pk; Type: CONSTRAINT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public.guess
    ADD CONSTRAINT guess_pk PRIMARY KEY (user_id);


--
-- Name: picture picture_pk; Type: CONSTRAINT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public.picture
    ADD CONSTRAINT picture_pk PRIMARY KEY (user_id);


--
-- Name: settings settings_pk; Type: CONSTRAINT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pk PRIMARY KEY (id);


--
-- Name: user user_pk; Type: CONSTRAINT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: settings_id_uindex; Type: INDEX; Schema: public; Owner: zedjones
--

CREATE UNIQUE INDEX settings_id_uindex ON public.settings USING btree (id);


--
-- Name: user_id_uindex; Type: INDEX; Schema: public; Owner: zedjones
--

CREATE UNIQUE INDEX user_id_uindex ON public."user" USING btree (id);


--
-- Name: guess guess_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public.guess
    ADD CONSTRAINT guess_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id);


--
-- Name: picture picture_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: zedjones
--

ALTER TABLE ONLY public.picture
    ADD CONSTRAINT picture_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id);


--
-- PostgreSQL database dump complete
--

