--
-- PostgreSQL database dump
--

-- Dumped from database version 11.8 (Ubuntu 11.8-1.pgdg18.04+1)
-- Dumped by pg_dump version 12.0

-- Started on 2020-10-15 11:16:01

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
-- TOC entry 3012 (class 1262 OID 1200202)
-- Name: library; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE library WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'ru_RU.UTF-8' LC_CTYPE = 'ru_RU.UTF-8';


\connect library

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
-- TOC entry 7 (class 2615 OID 1200203)
-- Name: library; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA library;


--
-- TOC entry 204 (class 1259 OID 1200313)
-- Name: ref_event_types; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.ref_event_types (
    pk_id integer NOT NULL,
    c_value character varying(20) NOT NULL
);


--
-- TOC entry 3013 (class 0 OID 0)
-- Dependencies: 204
-- Name: TABLE ref_event_types; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.ref_event_types IS 'справочник типов событий';


--
-- TOC entry 3014 (class 0 OID 0)
-- Dependencies: 204
-- Name: COLUMN ref_event_types.c_value; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON COLUMN library.ref_event_types.c_value IS 'строковое представление типа события';


--
-- TOC entry 203 (class 1259 OID 1200311)
-- Name: ref_event_types_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.ref_event_types_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3015 (class 0 OID 0)
-- Dependencies: 203
-- Name: ref_event_types_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.ref_event_types_pk_id_seq OWNED BY library.ref_event_types.pk_id;


--
-- TOC entry 210 (class 1259 OID 1200345)
-- Name: ref_positions; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.ref_positions (
    pk_id integer NOT NULL,
    c_name character varying(100) NOT NULL
);


--
-- TOC entry 209 (class 1259 OID 1200343)
-- Name: ref_positions_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.ref_positions_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3016 (class 0 OID 0)
-- Dependencies: 209
-- Name: ref_positions_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.ref_positions_pk_id_seq OWNED BY library.ref_positions.pk_id;


--
-- TOC entry 200 (class 1259 OID 1200297)
-- Name: ref_statuses; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.ref_statuses (
    pk_id integer NOT NULL,
    c_value character varying(20) NOT NULL
);


--
-- TOC entry 3017 (class 0 OID 0)
-- Dependencies: 200
-- Name: TABLE ref_statuses; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.ref_statuses IS 'справочник статусов книг';


--
-- TOC entry 3018 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN ref_statuses.c_value; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON COLUMN library.ref_statuses.c_value IS 'строковое представление статуса';


--
-- TOC entry 199 (class 1259 OID 1200295)
-- Name: ref_statuses_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.ref_statuses_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3019 (class 0 OID 0)
-- Dependencies: 199
-- Name: ref_statuses_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.ref_statuses_pk_id_seq OWNED BY library.ref_statuses.pk_id;


--
-- TOC entry 198 (class 1259 OID 1200288)
-- Name: t_books; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.t_books (
    pk_id integer NOT NULL,
    fk_status integer NOT NULL,
    c_name character varying(100) NOT NULL,
    c_isbn character varying(13) NOT NULL,
    c_author character varying(100),
    c_publisher character varying(50),
    c_year date,
    c_is_archive integer DEFAULT 0 NOT NULL
);


--
-- TOC entry 3020 (class 0 OID 0)
-- Dependencies: 198
-- Name: TABLE t_books; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.t_books IS 'таблица книг';


--
-- TOC entry 3021 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_books.c_isbn; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON COLUMN library.t_books.c_isbn IS 'уникальный id';


--
-- TOC entry 3022 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_books.c_is_archive; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON COLUMN library.t_books.c_is_archive IS 'признак архивности';


--
-- TOC entry 197 (class 1259 OID 1200286)
-- Name: t_books_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.t_books_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3023 (class 0 OID 0)
-- Dependencies: 197
-- Name: t_books_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.t_books_pk_id_seq OWNED BY library.t_books.pk_id;


--
-- TOC entry 206 (class 1259 OID 1200329)
-- Name: t_employees; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.t_employees (
    pk_id integer NOT NULL,
    fk_position integer NOT NULL,
    c_firstname character varying(50) NOT NULL,
    c_lastname character varying(50) NOT NULL,
    c_middlename character varying(50),
    c_phone_number character varying(11) NOT NULL,
    c_email character varying(50) NOT NULL,
    c_is_archive integer NOT NULL
);


--
-- TOC entry 3024 (class 0 OID 0)
-- Dependencies: 206
-- Name: TABLE t_employees; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.t_employees IS 'справочник должностей';


--
-- TOC entry 3025 (class 0 OID 0)
-- Dependencies: 206
-- Name: COLUMN t_employees.c_is_archive; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON COLUMN library.t_employees.c_is_archive IS 'признак архивности';


--
-- TOC entry 205 (class 1259 OID 1200327)
-- Name: t_employees_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.t_employees_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3026 (class 0 OID 0)
-- Dependencies: 205
-- Name: t_employees_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.t_employees_pk_id_seq OWNED BY library.t_employees.pk_id;


--
-- TOC entry 202 (class 1259 OID 1200305)
-- Name: t_event; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.t_event (
    pk_id integer NOT NULL,
    fk_book integer NOT NULL,
    fk_employee integer NOT NULL,
    fk_event_type integer NOT NULL,
    c_date date
);


--
-- TOC entry 3027 (class 0 OID 0)
-- Dependencies: 202
-- Name: TABLE t_event; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.t_event IS 'таблица событий выдачи/сдачи книг';


--
-- TOC entry 201 (class 1259 OID 1200303)
-- Name: t_event_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.t_event_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3028 (class 0 OID 0)
-- Dependencies: 201
-- Name: t_event_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.t_event_pk_id_seq OWNED BY library.t_event.pk_id;


--
-- TOC entry 208 (class 1259 OID 1200337)
-- Name: t_users; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.t_users (
    pk_id integer NOT NULL,
    c_login character varying(50) NOT NULL,
    c_password character varying(50) NOT NULL,
    fk_employee integer NOT NULL
);


--
-- TOC entry 3029 (class 0 OID 0)
-- Dependencies: 208
-- Name: TABLE t_users; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.t_users IS 'таблица пользователей';


--
-- TOC entry 207 (class 1259 OID 1200335)
-- Name: t_users_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.t_users_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3030 (class 0 OID 0)
-- Dependencies: 207
-- Name: t_users_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.t_users_pk_id_seq OWNED BY library.t_users.pk_id;


--
-- TOC entry 2844 (class 2604 OID 1200316)
-- Name: ref_event_types pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_event_types ALTER COLUMN pk_id SET DEFAULT nextval('library.ref_event_types_pk_id_seq'::regclass);


--
-- TOC entry 2847 (class 2604 OID 1200348)
-- Name: ref_positions pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_positions ALTER COLUMN pk_id SET DEFAULT nextval('library.ref_positions_pk_id_seq'::regclass);


--
-- TOC entry 2842 (class 2604 OID 1200300)
-- Name: ref_statuses pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_statuses ALTER COLUMN pk_id SET DEFAULT nextval('library.ref_statuses_pk_id_seq'::regclass);


--
-- TOC entry 2840 (class 2604 OID 1200291)
-- Name: t_books pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books ALTER COLUMN pk_id SET DEFAULT nextval('library.t_books_pk_id_seq'::regclass);


--
-- TOC entry 2845 (class 2604 OID 1200332)
-- Name: t_employees pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_employees ALTER COLUMN pk_id SET DEFAULT nextval('library.t_employees_pk_id_seq'::regclass);


--
-- TOC entry 2843 (class 2604 OID 1200308)
-- Name: t_event pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event ALTER COLUMN pk_id SET DEFAULT nextval('library.t_event_pk_id_seq'::regclass);


--
-- TOC entry 2846 (class 2604 OID 1200340)
-- Name: t_users pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users ALTER COLUMN pk_id SET DEFAULT nextval('library.t_users_pk_id_seq'::regclass);


--
-- TOC entry 3000 (class 0 OID 1200313)
-- Dependencies: 204
-- Data for Name: ref_event_types; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.ref_event_types (pk_id, c_value) VALUES (1, 'выдача');
INSERT INTO library.ref_event_types (pk_id, c_value) VALUES (2, 'сдача');


--
-- TOC entry 3006 (class 0 OID 1200345)
-- Dependencies: 210
-- Data for Name: ref_positions; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.ref_positions (pk_id, c_name) VALUES (1, 'Инженер-программист');


--
-- TOC entry 2996 (class 0 OID 1200297)
-- Dependencies: 200
-- Data for Name: ref_statuses; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.ref_statuses (pk_id, c_value) VALUES (1, 'доступна');
INSERT INTO library.ref_statuses (pk_id, c_value) VALUES (2, 'не доступна');


--
-- TOC entry 2994 (class 0 OID 1200288)
-- Dependencies: 198
-- Data for Name: t_books; Type: TABLE DATA; Schema: library; Owner: -
--



--
-- TOC entry 3002 (class 0 OID 1200329)
-- Dependencies: 206
-- Data for Name: t_employees; Type: TABLE DATA; Schema: library; Owner: -
--



--
-- TOC entry 2998 (class 0 OID 1200305)
-- Dependencies: 202
-- Data for Name: t_event; Type: TABLE DATA; Schema: library; Owner: -
--



--
-- TOC entry 3004 (class 0 OID 1200337)
-- Dependencies: 208
-- Data for Name: t_users; Type: TABLE DATA; Schema: library; Owner: -
--



--
-- TOC entry 3031 (class 0 OID 0)
-- Dependencies: 203
-- Name: ref_event_types_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.ref_event_types_pk_id_seq', 2, true);


--
-- TOC entry 3032 (class 0 OID 0)
-- Dependencies: 209
-- Name: ref_positions_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.ref_positions_pk_id_seq', 1, true);


--
-- TOC entry 3033 (class 0 OID 0)
-- Dependencies: 199
-- Name: ref_statuses_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.ref_statuses_pk_id_seq', 2, true);


--
-- TOC entry 3034 (class 0 OID 0)
-- Dependencies: 197
-- Name: t_books_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_books_pk_id_seq', 2, true);


--
-- TOC entry 3035 (class 0 OID 0)
-- Dependencies: 205
-- Name: t_employees_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_employees_pk_id_seq', 1, false);


--
-- TOC entry 3036 (class 0 OID 0)
-- Dependencies: 201
-- Name: t_event_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_event_pk_id_seq', 1, false);


--
-- TOC entry 3037 (class 0 OID 0)
-- Dependencies: 207
-- Name: t_users_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_users_pk_id_seq', 1, false);


--
-- TOC entry 2857 (class 2606 OID 1200318)
-- Name: ref_event_types ref_event_types_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_event_types
    ADD CONSTRAINT ref_event_types_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2853 (class 2606 OID 1200302)
-- Name: ref_statuses ref_statuses_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_statuses
    ADD CONSTRAINT ref_statuses_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2849 (class 2606 OID 1200294)
-- Name: t_books t_books_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books
    ADD CONSTRAINT t_books_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2851 (class 2606 OID 1200390)
-- Name: t_books t_books_un_c_isbn; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books
    ADD CONSTRAINT t_books_un_c_isbn UNIQUE (c_isbn);


--
-- TOC entry 2859 (class 2606 OID 1200334)
-- Name: t_employees t_employees_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_employees
    ADD CONSTRAINT t_employees_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2855 (class 2606 OID 1200310)
-- Name: t_event t_event_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2865 (class 2606 OID 1200350)
-- Name: ref_positions t_positions_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_positions
    ADD CONSTRAINT t_positions_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2861 (class 2606 OID 1200342)
-- Name: t_users t_users_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users
    ADD CONSTRAINT t_users_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2863 (class 2606 OID 1200392)
-- Name: t_users t_users_un_c_login; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users
    ADD CONSTRAINT t_users_un_c_login UNIQUE (c_login);


--
-- TOC entry 2866 (class 2606 OID 1200351)
-- Name: t_books t_books_fk; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books
    ADD CONSTRAINT t_books_fk FOREIGN KEY (fk_status) REFERENCES library.ref_statuses(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2870 (class 2606 OID 1200376)
-- Name: t_employees t_employees_fk_position; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_employees
    ADD CONSTRAINT t_employees_fk_position FOREIGN KEY (fk_position) REFERENCES library.ref_positions(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2868 (class 2606 OID 1200361)
-- Name: t_event t_event_fk_book; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_fk_book FOREIGN KEY (fk_book) REFERENCES library.t_books(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2869 (class 2606 OID 1200366)
-- Name: t_event t_event_fk_employee; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_fk_employee FOREIGN KEY (fk_employee) REFERENCES library.t_employees(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2867 (class 2606 OID 1200356)
-- Name: t_event t_event_fk_event_types; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_fk_event_types FOREIGN KEY (fk_event_type) REFERENCES library.ref_event_types(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2871 (class 2606 OID 1200393)
-- Name: t_users t_users_fk_employee; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users
    ADD CONSTRAINT t_users_fk_employee FOREIGN KEY (fk_employee) REFERENCES library.t_employees(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


-- Completed on 2020-10-15 11:16:03

--
-- PostgreSQL database dump complete
--

