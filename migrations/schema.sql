--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.9
-- Dumped by pg_dump version 10.4 (Debian 10.4-2.pgdg90+1)

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: current_parking_occupancy; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.current_parking_occupancy (
    floor_label character(1) NOT NULL,
    space_label text NOT NULL,
    status text NOT NULL,
    started_at timestamp without time zone,
    expires_at timestamp without time zone,
    ended_at timestamp without time zone,
    occupant_id uuid NOT NULL
);


ALTER TABLE public.current_parking_occupancy OWNER TO postgres;

--
-- Name: delivered_sms_messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivered_sms_messages (
    id uuid NOT NULL,
    parking_occupant_id uuid NOT NULL,
    content text NOT NULL,
    floor_label character(1) NOT NULL,
    space_label text NOT NULL,
    sent_at timestamp without time zone NOT NULL,
    sent_to text[] NOT NULL
);


ALTER TABLE public.delivered_sms_messages OWNER TO postgres;

--
-- Name: parking_occupant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.parking_occupant (
    id uuid NOT NULL,
    resident_name text NOT NULL,
    resident_phone text NOT NULL,
    other_phones text[] NOT NULL,
    resident_room_number text NOT NULL
);


ALTER TABLE public.parking_occupant OWNER TO postgres;

--
-- Name: resident_information; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.resident_information (
    id uuid NOT NULL,
    first_name text,
    last_name text,
    phone_number text,
    room_number text
);


ALTER TABLE public.resident_information OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: current_parking_occupancy current_parking_occupancy_floor_label_space_label_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.current_parking_occupancy
    ADD CONSTRAINT current_parking_occupancy_floor_label_space_label_key UNIQUE (floor_label, space_label);


--
-- Name: delivered_sms_messages delivered_sms_messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivered_sms_messages
    ADD CONSTRAINT delivered_sms_messages_pkey PRIMARY KEY (id);


--
-- Name: parking_occupant parking_occupant_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.parking_occupant
    ADD CONSTRAINT parking_occupant_pkey PRIMARY KEY (id);


--
-- Name: resident_information resident_information_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.resident_information
    ADD CONSTRAINT resident_information_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

