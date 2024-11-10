--
-- PostgreSQL database dump
--

-- Dumped from database version 14.13
-- Dumped by pg_dump version 16.4

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
-- Data for Name: partner; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.partner (id, name, geo, radius, rating) VALUES ('d6d7bff5-2b61-4adc-8740-82e1f1312ae9', 'alice', '0101000020E6100000BF2B82FFADC42A4057957D5704434A40', 10000, 9);
INSERT INTO public.partner (id, name, geo, radius, rating) VALUES ('a2d0304a-02a4-4752-8ff5-aff2e1c045c5', 'bob', '0101000020E610000076711B0DE01D294032207BBDFB334A40', 200000, 8);
INSERT INTO public.partner (id, name, geo, radius, rating) VALUES ('a17e0cd1-3377-4010-92df-aa72dfe0105e', 'carlos', '0101000020E6100000840D4FAF9445284041481630810B4B40', 100000, 8);
INSERT INTO public.partner (id, name, geo, radius, rating) VALUES ('e1802b2f-d2f2-4902-af1d-0df98e0767b9', 'chad', '0101000020E610000049D74CBED9C628401EFE9AAC51AB4940', 80000, 2);


--
-- Data for Name: partner_skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.partner_skill (partner_id, code) VALUES ('d6d7bff5-2b61-4adc-8740-82e1f1312ae9', 'wood');
INSERT INTO public.partner_skill (partner_id, code) VALUES ('a2d0304a-02a4-4752-8ff5-aff2e1c045c5', 'wood');
INSERT INTO public.partner_skill (partner_id, code) VALUES ('a2d0304a-02a4-4752-8ff5-aff2e1c045c5', 'carpet');
INSERT INTO public.partner_skill (partner_id, code) VALUES ('a17e0cd1-3377-4010-92df-aa72dfe0105e', 'carpet');
INSERT INTO public.partner_skill (partner_id, code) VALUES ('e1802b2f-d2f2-4902-af1d-0df98e0767b9', 'wood');
INSERT INTO public.partner_skill (partner_id, code) VALUES ('e1802b2f-d2f2-4902-af1d-0df98e0767b9', 'tiles');


--
-- PostgreSQL database dump complete
--

