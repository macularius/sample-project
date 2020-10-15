--
-- PostgreSQL database dump
--

-- Dumped from database version 11.8 (Ubuntu 11.8-1.pgdg18.04+1)
-- Dumped by pg_dump version 12.0

-- Started on 2020-10-15 18:17:04

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
-- TOC entry 3023 (class 1262 OID 1200202)
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
-- TOC entry 3024 (class 0 OID 0)
-- Dependencies: 204
-- Name: TABLE ref_event_types; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.ref_event_types IS 'справочник типов событий';


--
-- TOC entry 3025 (class 0 OID 0)
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
-- TOC entry 3026 (class 0 OID 0)
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
-- TOC entry 3027 (class 0 OID 0)
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
-- TOC entry 3028 (class 0 OID 0)
-- Dependencies: 200
-- Name: TABLE ref_statuses; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.ref_statuses IS 'справочник статусов книг';


--
-- TOC entry 3029 (class 0 OID 0)
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
-- TOC entry 3030 (class 0 OID 0)
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
-- TOC entry 3031 (class 0 OID 0)
-- Dependencies: 198
-- Name: TABLE t_books; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.t_books IS 'таблица книг';


--
-- TOC entry 3032 (class 0 OID 0)
-- Dependencies: 198
-- Name: COLUMN t_books.c_isbn; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON COLUMN library.t_books.c_isbn IS 'уникальный id';


--
-- TOC entry 3033 (class 0 OID 0)
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
-- TOC entry 3034 (class 0 OID 0)
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
-- TOC entry 3035 (class 0 OID 0)
-- Dependencies: 206
-- Name: TABLE t_employees; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.t_employees IS 'справочник должностей';


--
-- TOC entry 3036 (class 0 OID 0)
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
-- TOC entry 3037 (class 0 OID 0)
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
-- TOC entry 3038 (class 0 OID 0)
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
-- TOC entry 3039 (class 0 OID 0)
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
-- TOC entry 3040 (class 0 OID 0)
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
-- TOC entry 3041 (class 0 OID 0)
-- Dependencies: 207
-- Name: t_users_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.t_users_pk_id_seq OWNED BY library.t_users.pk_id;


--
-- TOC entry 212 (class 1259 OID 1200400)
-- Name: toc_books_employees; Type: TABLE; Schema: library; Owner: -
--

CREATE TABLE library.toc_books_employees (
    pk_id integer NOT NULL,
    fk_book integer NOT NULL,
    fk_employee integer NOT NULL
);


--
-- TOC entry 3042 (class 0 OID 0)
-- Dependencies: 212
-- Name: TABLE toc_books_employees; Type: COMMENT; Schema: library; Owner: -
--

COMMENT ON TABLE library.toc_books_employees IS 'таблица книг сотрудника в текущий момент';


--
-- TOC entry 211 (class 1259 OID 1200398)
-- Name: toc_books_employees_pk_id_seq; Type: SEQUENCE; Schema: library; Owner: -
--

CREATE SEQUENCE library.toc_books_employees_pk_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3043 (class 0 OID 0)
-- Dependencies: 211
-- Name: toc_books_employees_pk_id_seq; Type: SEQUENCE OWNED BY; Schema: library; Owner: -
--

ALTER SEQUENCE library.toc_books_employees_pk_id_seq OWNED BY library.toc_books_employees.pk_id;


--
-- TOC entry 2850 (class 2604 OID 1200316)
-- Name: ref_event_types pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_event_types ALTER COLUMN pk_id SET DEFAULT nextval('library.ref_event_types_pk_id_seq'::regclass);


--
-- TOC entry 2853 (class 2604 OID 1200348)
-- Name: ref_positions pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_positions ALTER COLUMN pk_id SET DEFAULT nextval('library.ref_positions_pk_id_seq'::regclass);


--
-- TOC entry 2848 (class 2604 OID 1200300)
-- Name: ref_statuses pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_statuses ALTER COLUMN pk_id SET DEFAULT nextval('library.ref_statuses_pk_id_seq'::regclass);


--
-- TOC entry 2846 (class 2604 OID 1200291)
-- Name: t_books pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books ALTER COLUMN pk_id SET DEFAULT nextval('library.t_books_pk_id_seq'::regclass);


--
-- TOC entry 2851 (class 2604 OID 1200332)
-- Name: t_employees pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_employees ALTER COLUMN pk_id SET DEFAULT nextval('library.t_employees_pk_id_seq'::regclass);


--
-- TOC entry 2849 (class 2604 OID 1200308)
-- Name: t_event pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event ALTER COLUMN pk_id SET DEFAULT nextval('library.t_event_pk_id_seq'::regclass);


--
-- TOC entry 2852 (class 2604 OID 1200340)
-- Name: t_users pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users ALTER COLUMN pk_id SET DEFAULT nextval('library.t_users_pk_id_seq'::regclass);


--
-- TOC entry 2854 (class 2604 OID 1200403)
-- Name: toc_books_employees pk_id; Type: DEFAULT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.toc_books_employees ALTER COLUMN pk_id SET DEFAULT nextval('library.toc_books_employees_pk_id_seq'::regclass);


--
-- TOC entry 3009 (class 0 OID 1200313)
-- Dependencies: 204
-- Data for Name: ref_event_types; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.ref_event_types (pk_id, c_value) VALUES (1, 'выдача');
INSERT INTO library.ref_event_types (pk_id, c_value) VALUES (2, 'сдача');


--
-- TOC entry 3015 (class 0 OID 1200345)
-- Dependencies: 210
-- Data for Name: ref_positions; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.ref_positions (pk_id, c_name) VALUES (1, 'Инженер-программист');


--
-- TOC entry 3005 (class 0 OID 1200297)
-- Dependencies: 200
-- Data for Name: ref_statuses; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.ref_statuses (pk_id, c_value) VALUES (1, 'доступна');
INSERT INTO library.ref_statuses (pk_id, c_value) VALUES (2, 'не доступна');


--
-- TOC entry 3003 (class 0 OID 1200288)
-- Dependencies: 198
-- Data for Name: t_books; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.t_books (pk_id, fk_status, c_name, c_isbn, c_author, c_publisher, c_year, c_is_archive) VALUES (6, 1, 'Секреты javascript ниндзя', '0000000002', 'Бибо Бэер', 'САНКТ ПЕТЕРБУРГ', '2016-01-01', 0);
INSERT INTO library.t_books (pk_id, fk_status, c_name, c_isbn, c_author, c_publisher, c_year, c_is_archive) VALUES (3, 2, 'Программирование на GO', '0000000001', 'Алан Донован Керниган', 'МОСКВА', '2012-01-01', 0);


--
-- TOC entry 3011 (class 0 OID 1200329)
-- Dependencies: 206
-- Data for Name: t_employees; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.t_employees (pk_id, fk_position, c_firstname, c_lastname, c_middlename, c_phone_number, c_email, c_is_archive) VALUES (1, 1, 'Игорь', 'Коваценко', 'Николаевич', '89873036298', 'почта', 0);


--
-- TOC entry 3007 (class 0 OID 1200305)
-- Dependencies: 202
-- Data for Name: t_event; Type: TABLE DATA; Schema: library; Owner: -
--

INSERT INTO library.t_event (pk_id, fk_book, fk_employee, fk_event_type, c_date) VALUES (10, 6, 1, 1, '2020-10-15');
INSERT INTO library.t_event (pk_id, fk_book, fk_employee, fk_event_type, c_date) VALUES (11, 3, 1, 1, '2020-10-15');
INSERT INTO library.t_event (pk_id, fk_book, fk_employee, fk_event_type, c_date) VALUES (12, 3, 1, 1, '2020-10-15');
INSERT INTO library.t_event (pk_id, fk_book, fk_employee, fk_event_type, c_date) VALUES (13, 6, 1, 1, '2020-10-15');
INSERT INTO library.t_event (pk_id, fk_book, fk_employee, fk_event_type, c_date) VALUES (14, 3, 1, 1, '2020-10-15');


--
-- TOC entry 3013 (class 0 OID 1200337)
-- Dependencies: 208
-- Data for Name: t_users; Type: TABLE DATA; Schema: library; Owner: -
--



--
-- TOC entry 3017 (class 0 OID 1200400)
-- Dependencies: 212
-- Data for Name: toc_books_employees; Type: TABLE DATA; Schema: library; Owner: -
--



--
-- TOC entry 3044 (class 0 OID 0)
-- Dependencies: 203
-- Name: ref_event_types_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.ref_event_types_pk_id_seq', 2, true);


--
-- TOC entry 3045 (class 0 OID 0)
-- Dependencies: 209
-- Name: ref_positions_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.ref_positions_pk_id_seq', 1, true);


--
-- TOC entry 3046 (class 0 OID 0)
-- Dependencies: 199
-- Name: ref_statuses_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.ref_statuses_pk_id_seq', 2, true);


--
-- TOC entry 3047 (class 0 OID 0)
-- Dependencies: 197
-- Name: t_books_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_books_pk_id_seq', 6, true);


--
-- TOC entry 3048 (class 0 OID 0)
-- Dependencies: 205
-- Name: t_employees_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_employees_pk_id_seq', 1, true);


--
-- TOC entry 3049 (class 0 OID 0)
-- Dependencies: 201
-- Name: t_event_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_event_pk_id_seq', 15, true);


--
-- TOC entry 3050 (class 0 OID 0)
-- Dependencies: 207
-- Name: t_users_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.t_users_pk_id_seq', 1, false);


--
-- TOC entry 3051 (class 0 OID 0)
-- Dependencies: 211
-- Name: toc_books_employees_pk_id_seq; Type: SEQUENCE SET; Schema: library; Owner: -
--

SELECT pg_catalog.setval('library.toc_books_employees_pk_id_seq', 1, false);


--
-- TOC entry 2864 (class 2606 OID 1200318)
-- Name: ref_event_types ref_event_types_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_event_types
    ADD CONSTRAINT ref_event_types_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2860 (class 2606 OID 1200302)
-- Name: ref_statuses ref_statuses_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_statuses
    ADD CONSTRAINT ref_statuses_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2856 (class 2606 OID 1200294)
-- Name: t_books t_books_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books
    ADD CONSTRAINT t_books_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2858 (class 2606 OID 1200390)
-- Name: t_books t_books_un_c_isbn; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books
    ADD CONSTRAINT t_books_un_c_isbn UNIQUE (c_isbn);


--
-- TOC entry 2866 (class 2606 OID 1200334)
-- Name: t_employees t_employees_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_employees
    ADD CONSTRAINT t_employees_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2862 (class 2606 OID 1200310)
-- Name: t_event t_event_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2872 (class 2606 OID 1200350)
-- Name: ref_positions t_positions_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.ref_positions
    ADD CONSTRAINT t_positions_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2868 (class 2606 OID 1200342)
-- Name: t_users t_users_pk; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users
    ADD CONSTRAINT t_users_pk PRIMARY KEY (pk_id);


--
-- TOC entry 2870 (class 2606 OID 1200392)
-- Name: t_users t_users_un_c_login; Type: CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users
    ADD CONSTRAINT t_users_un_c_login UNIQUE (c_login);


--
-- TOC entry 2873 (class 2606 OID 1200351)
-- Name: t_books t_books_fk; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_books
    ADD CONSTRAINT t_books_fk FOREIGN KEY (fk_status) REFERENCES library.ref_statuses(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2877 (class 2606 OID 1200376)
-- Name: t_employees t_employees_fk_position; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_employees
    ADD CONSTRAINT t_employees_fk_position FOREIGN KEY (fk_position) REFERENCES library.ref_positions(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2875 (class 2606 OID 1200361)
-- Name: t_event t_event_fk_book; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_fk_book FOREIGN KEY (fk_book) REFERENCES library.t_books(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2876 (class 2606 OID 1200366)
-- Name: t_event t_event_fk_employee; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_fk_employee FOREIGN KEY (fk_employee) REFERENCES library.t_employees(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2874 (class 2606 OID 1200356)
-- Name: t_event t_event_fk_event_types; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_event
    ADD CONSTRAINT t_event_fk_event_types FOREIGN KEY (fk_event_type) REFERENCES library.ref_event_types(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2878 (class 2606 OID 1200393)
-- Name: t_users t_users_fk_employee; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.t_users
    ADD CONSTRAINT t_users_fk_employee FOREIGN KEY (fk_employee) REFERENCES library.t_employees(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2879 (class 2606 OID 1200404)
-- Name: toc_books_employees toc_books_employees_fk_book; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.toc_books_employees
    ADD CONSTRAINT toc_books_employees_fk_book FOREIGN KEY (fk_book) REFERENCES library.t_books(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- TOC entry 2880 (class 2606 OID 1200409)
-- Name: toc_books_employees toc_books_employees_fk_employee; Type: FK CONSTRAINT; Schema: library; Owner: -
--

ALTER TABLE ONLY library.toc_books_employees
    ADD CONSTRAINT toc_books_employees_fk_employee FOREIGN KEY (fk_employee) REFERENCES library.t_employees(pk_id) ON UPDATE CASCADE ON DELETE RESTRICT;


-- Completed on 2020-10-15 18:17:04

--
-- PostgreSQL database dump complete
--

