--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: notification; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.notification (
    id integer NOT NULL,
    event_type smallint NOT NULL,
    user_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.notification OWNER TO kopilka;

--
-- Name: notification_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.notification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notification_id_seq OWNER TO kopilka;

--
-- Name: notification_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.notification_id_seq OWNED BY public.notification.id;


--
-- Name: notification_user_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.notification_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notification_user_id_seq OWNER TO kopilka;

--
-- Name: notification_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.notification_user_id_seq OWNED BY public.notification.user_id;


--
-- Name: payment; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.payment (
    id integer NOT NULL,
    payment_integer bigint NOT NULL,
    payment_fractional bigint NOT NULL,
    status smallint NOT NULL,
    donater_id integer NOT NULL,
    creator_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT payment_fractional_range CHECK ((payment_fractional >= 0)),
    CONSTRAINT payment_integer_range CHECK ((payment_integer >= 0))
);


ALTER TABLE public.payment OWNER TO kopilka;

--
-- Name: payment_creator_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.payment_creator_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_creator_id_seq OWNER TO kopilka;

--
-- Name: payment_creator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.payment_creator_id_seq OWNED BY public.payment.creator_id;


--
-- Name: payment_donater_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.payment_donater_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_donater_id_seq OWNER TO kopilka;

--
-- Name: payment_donater_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.payment_donater_id_seq OWNED BY public.payment.donater_id;


--
-- Name: payment_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.payment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_id_seq OWNER TO kopilka;

--
-- Name: payment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.payment_id_seq OWNED BY public.payment.id;


--
-- Name: post; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.post (
    id integer NOT NULL,
    header character varying(200) NOT NULL,
    body text NOT NULL,
    creator_id integer NOT NULL,
    min_subscription_level_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.post OWNER TO kopilka;

--
-- Name: post_attach; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.post_attach (
    id integer NOT NULL,
    file_path character varying(100) NOT NULL,
    post_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.post_attach OWNER TO kopilka;

--
-- Name: post_attach_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_attach_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_attach_id_seq OWNER TO kopilka;

--
-- Name: post_attach_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_attach_id_seq OWNED BY public.post_attach.id;


--
-- Name: post_attach_post_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_attach_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_attach_post_id_seq OWNER TO kopilka;

--
-- Name: post_attach_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_attach_post_id_seq OWNED BY public.post_attach.post_id;


--
-- Name: post_comment; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.post_comment (
    id integer NOT NULL,
    text text NOT NULL,
    user_id integer NOT NULL,
    post_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.post_comment OWNER TO kopilka;

--
-- Name: post_comment_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_comment_id_seq OWNER TO kopilka;

--
-- Name: post_comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_comment_id_seq OWNED BY public.post_comment.id;


--
-- Name: post_comment_post_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_comment_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_comment_post_id_seq OWNER TO kopilka;

--
-- Name: post_comment_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_comment_post_id_seq OWNED BY public.post_comment.post_id;


--
-- Name: post_comment_user_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_comment_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_comment_user_id_seq OWNER TO kopilka;

--
-- Name: post_comment_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_comment_user_id_seq OWNED BY public.post_comment.user_id;


--
-- Name: post_creator_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_creator_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_creator_id_seq OWNER TO kopilka;

--
-- Name: post_creator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_creator_id_seq OWNED BY public.post.creator_id;


--
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO kopilka;

--
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_id_seq OWNED BY public.post.id;


--
-- Name: post_like; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.post_like (
    id integer NOT NULL,
    user_id integer NOT NULL,
    post_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.post_like OWNER TO kopilka;

--
-- Name: post_like_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_like_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_like_id_seq OWNER TO kopilka;

--
-- Name: post_like_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_like_id_seq OWNED BY public.post_like.id;


--
-- Name: post_like_post_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_like_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_like_post_id_seq OWNER TO kopilka;

--
-- Name: post_like_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_like_post_id_seq OWNED BY public.post_like.post_id;


--
-- Name: post_like_user_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_like_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_like_user_id_seq OWNER TO kopilka;

--
-- Name: post_like_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_like_user_id_seq OWNED BY public.post_like.user_id;


--
-- Name: post_min_subscription_level_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.post_min_subscription_level_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_min_subscription_level_id_seq OWNER TO kopilka;

--
-- Name: post_min_subscription_level_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.post_min_subscription_level_id_seq OWNED BY public.post.min_subscription_level_id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO kopilka;

--
-- Name: subscription; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.subscription (
    id integer NOT NULL,
    subscriber_id integer NOT NULL,
    creator_id integer NOT NULL,
    subscription_level_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.subscription OWNER TO kopilka;

--
-- Name: subscription_creator_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.subscription_creator_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscription_creator_id_seq OWNER TO kopilka;

--
-- Name: subscription_creator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.subscription_creator_id_seq OWNED BY public.subscription.creator_id;


--
-- Name: subscription_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.subscription_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscription_id_seq OWNER TO kopilka;

--
-- Name: subscription_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.subscription_id_seq OWNED BY public.subscription.id;


--
-- Name: subscription_level; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public.subscription_level (
    id integer NOT NULL,
    level smallint NOT NULL,
    name character varying(30) NOT NULL,
    description text NOT NULL,
    cost_integer bigint NOT NULL,
    cost_fractional bigint NOT NULL,
    currency character(3) NOT NULL,
    creator_id integer NOT NULL,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT cost_fractional_range CHECK ((cost_fractional >= 0)),
    CONSTRAINT cost_integer_range CHECK ((cost_integer >= 0)),
    CONSTRAINT level_range CHECK ((level >= 0))
);


ALTER TABLE public.subscription_level OWNER TO kopilka;

--
-- Name: subscription_level_creator_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.subscription_level_creator_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscription_level_creator_id_seq OWNER TO kopilka;

--
-- Name: subscription_level_creator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.subscription_level_creator_id_seq OWNED BY public.subscription_level.creator_id;


--
-- Name: subscription_level_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.subscription_level_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscription_level_id_seq OWNER TO kopilka;

--
-- Name: subscription_level_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.subscription_level_id_seq OWNED BY public.subscription_level.id;


--
-- Name: subscription_subscriber_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.subscription_subscriber_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscription_subscriber_id_seq OWNER TO kopilka;

--
-- Name: subscription_subscriber_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.subscription_subscriber_id_seq OWNED BY public.subscription.subscriber_id;


--
-- Name: subscription_subscription_level_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.subscription_subscription_level_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subscription_subscription_level_id_seq OWNER TO kopilka;

--
-- Name: subscription_subscription_level_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.subscription_subscription_level_id_seq OWNED BY public.subscription.subscription_level_id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: kopilka
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    nickname character varying(20) NOT NULL,
    email character varying(300) NOT NULL,
    password character varying(50) NOT NULL,
    is_author boolean NOT NULL,
    status character varying(100),
    avatar_path character varying(100) DEFAULT 'static/default_avatar.png'::character varying,
    background_path character varying(100) DEFAULT 'static/default_background.png'::character varying,
    description text,
    creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public."user" OWNER TO kopilka;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: kopilka
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO kopilka;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kopilka
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: notification id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.notification ALTER COLUMN id SET DEFAULT nextval('public.notification_id_seq'::regclass);


--
-- Name: notification user_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.notification ALTER COLUMN user_id SET DEFAULT nextval('public.notification_user_id_seq'::regclass);


--
-- Name: payment id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.payment ALTER COLUMN id SET DEFAULT nextval('public.payment_id_seq'::regclass);


--
-- Name: payment donater_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.payment ALTER COLUMN donater_id SET DEFAULT nextval('public.payment_donater_id_seq'::regclass);


--
-- Name: payment creator_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.payment ALTER COLUMN creator_id SET DEFAULT nextval('public.payment_creator_id_seq'::regclass);


--
-- Name: post id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post ALTER COLUMN id SET DEFAULT nextval('public.post_id_seq'::regclass);


--
-- Name: post creator_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post ALTER COLUMN creator_id SET DEFAULT nextval('public.post_creator_id_seq'::regclass);


--
-- Name: post min_subscription_level_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post ALTER COLUMN min_subscription_level_id SET DEFAULT nextval('public.post_min_subscription_level_id_seq'::regclass);


--
-- Name: post_attach id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_attach ALTER COLUMN id SET DEFAULT nextval('public.post_attach_id_seq'::regclass);


--
-- Name: post_attach post_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_attach ALTER COLUMN post_id SET DEFAULT nextval('public.post_attach_post_id_seq'::regclass);


--
-- Name: post_comment id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_comment ALTER COLUMN id SET DEFAULT nextval('public.post_comment_id_seq'::regclass);


--
-- Name: post_comment user_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_comment ALTER COLUMN user_id SET DEFAULT nextval('public.post_comment_user_id_seq'::regclass);


--
-- Name: post_comment post_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_comment ALTER COLUMN post_id SET DEFAULT nextval('public.post_comment_post_id_seq'::regclass);


--
-- Name: post_like id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_like ALTER COLUMN id SET DEFAULT nextval('public.post_like_id_seq'::regclass);


--
-- Name: post_like user_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_like ALTER COLUMN user_id SET DEFAULT nextval('public.post_like_user_id_seq'::regclass);


--
-- Name: post_like post_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_like ALTER COLUMN post_id SET DEFAULT nextval('public.post_like_post_id_seq'::regclass);


--
-- Name: subscription id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription ALTER COLUMN id SET DEFAULT nextval('public.subscription_id_seq'::regclass);


--
-- Name: subscription subscriber_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription ALTER COLUMN subscriber_id SET DEFAULT nextval('public.subscription_subscriber_id_seq'::regclass);


--
-- Name: subscription creator_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription ALTER COLUMN creator_id SET DEFAULT nextval('public.subscription_creator_id_seq'::regclass);


--
-- Name: subscription subscription_level_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription ALTER COLUMN subscription_level_id SET DEFAULT nextval('public.subscription_subscription_level_id_seq'::regclass);


--
-- Name: subscription_level id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription_level ALTER COLUMN id SET DEFAULT nextval('public.subscription_level_id_seq'::regclass);


--
-- Name: subscription_level creator_id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription_level ALTER COLUMN creator_id SET DEFAULT nextval('public.subscription_level_creator_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: notification; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.notification (id, event_type, user_id, creation_date, last_update) FROM stdin;
\.


--
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.payment (id, payment_integer, payment_fractional, status, donater_id, creator_id, creation_date, last_update) FROM stdin;
\.


--
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.post (id, header, body, creator_id, min_subscription_level_id, creation_date, last_update) FROM stdin;
1	haha	heheh	50	10	2023-10-27 15:59:43.386084+03	2023-10-27 15:59:43.386084+03
2	haha	heheh	50	10	2023-10-27 15:59:47.673613+03	2023-10-27 15:59:47.673613+03
3	haha	heheh	50	10	2023-10-27 15:59:48.571455+03	2023-10-27 15:59:48.571455+03
4	haha	heheh	50	10	2023-10-27 15:59:48.806651+03	2023-10-27 15:59:48.806651+03
\.


--
-- Data for Name: post_attach; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.post_attach (id, file_path, post_id, creation_date) FROM stdin;
2	static/img1.png	1	2023-10-29 22:42:56.895943+03
3	static/img2.png	1	2023-10-29 22:43:13.676466+03
4	static/img2.png	2	2023-10-29 22:43:43.739598+03
5	static/img2.png	3	2023-10-29 22:44:48.794579+03
6	static/img2.png	4	2023-10-29 22:44:52.069118+03
\.


--
-- Data for Name: post_comment; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.post_comment (id, text, user_id, post_id, creation_date, last_update) FROM stdin;
\.


--
-- Data for Name: post_like; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.post_like (id, user_id, post_id, creation_date) FROM stdin;
1	5	1	2023-10-30 16:32:35.645509+03
2	6	1	2023-10-30 16:32:42.563862+03
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: subscription; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.subscription (id, subscriber_id, creator_id, subscription_level_id, creation_date, last_update) FROM stdin;
2	1	50	10	2023-10-26 23:18:14.391086+03	2023-10-26 23:18:14.391086+03
3	3	50	10	2023-10-26 23:53:02.168703+03	2023-10-26 23:53:02.168703+03
4	5	50	2	2023-10-28 19:41:29.552855+03	2023-10-28 19:41:29.552855+03
5	6	50	3	2023-10-28 19:41:36.803413+03	2023-10-28 19:41:36.803413+03
\.


--
-- Data for Name: subscription_level; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public.subscription_level (id, level, name, description, cost_integer, cost_fractional, currency, creator_id, creation_date, last_update) FROM stdin;
10	1	hehe	hehe	55	55	rub	50	2023-10-26 23:17:43.186825+03	2023-10-26 23:17:43.186825+03
1	0	zero_level	zero_level	0	0	rub	1	2023-10-28 19:18:34.628177+03	2023-10-28 19:18:34.628177+03
2	0	zero_level	zero_level	0	0	rub	50	2023-10-28 19:19:57.926271+03	2023-10-28 19:19:57.926271+03
3	2	advanced	advanced	500	500	rub	50	2023-10-28 19:20:20.154555+03	2023-10-28 19:20:20.154555+03
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: kopilka
--

COPY public."user" (id, nickname, email, password, is_author, status, avatar_path, background_path, description, creation_date, last_update) FROM stdin;
1		Abcd1234	Abcd1234**	f					2023-10-26 17:40:04.192589+03	2023-10-26 17:40:04.192589+03
3	chel	Abcd12345	Abcd12345**	f		static/default_avatar.png	static/default_background.png		2023-10-26 19:29:06.144815+03	2023-10-26 19:29:06.144815+03
50	chel	Abcd123456	Abcd123456**	t		static/default_avatar.png	static/default_background.png		2023-10-26 23:16:14.348326+03	2023-10-26 23:16:14.348326+03
4		Abcd123467	Abcd123467**	t					2023-10-27 00:38:14.665112+03	2023-10-27 00:38:14.665112+03
6	kek	kek	Abcd213**	f		static/default_avatar.png	static/default_background.png		2023-10-28 19:21:53.563832+03	2023-10-28 19:21:53.563832+03
5	lol	Abcd4321	Abcd4321**	f		static/default_avatar.png	static/default_background.png		2023-10-28 19:21:44.571832+03	2023-10-28 19:21:44.571832+03
\.


--
-- Name: notification_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.notification_id_seq', 1, false);


--
-- Name: notification_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.notification_user_id_seq', 1, false);


--
-- Name: payment_creator_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.payment_creator_id_seq', 1, false);


--
-- Name: payment_donater_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.payment_donater_id_seq', 1, false);


--
-- Name: payment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.payment_id_seq', 1, false);


--
-- Name: post_attach_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_attach_id_seq', 6, true);


--
-- Name: post_attach_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_attach_post_id_seq', 1, false);


--
-- Name: post_comment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_comment_id_seq', 1, false);


--
-- Name: post_comment_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_comment_post_id_seq', 1, false);


--
-- Name: post_comment_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_comment_user_id_seq', 1, false);


--
-- Name: post_creator_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_creator_id_seq', 1, false);


--
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_id_seq', 4, true);


--
-- Name: post_like_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_like_id_seq', 2, true);


--
-- Name: post_like_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_like_post_id_seq', 1, false);


--
-- Name: post_like_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_like_user_id_seq', 1, false);


--
-- Name: post_min_subscription_level_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.post_min_subscription_level_id_seq', 1, false);


--
-- Name: subscription_creator_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.subscription_creator_id_seq', 1, false);


--
-- Name: subscription_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.subscription_id_seq', 5, true);


--
-- Name: subscription_level_creator_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.subscription_level_creator_id_seq', 1, true);


--
-- Name: subscription_level_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.subscription_level_id_seq', 3, true);


--
-- Name: subscription_subscriber_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.subscription_subscriber_id_seq', 1, false);


--
-- Name: subscription_subscription_level_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.subscription_subscription_level_id_seq', 1, false);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: kopilka
--

SELECT pg_catalog.setval('public.user_id_seq', 6, true);


--
-- Name: notification notification_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT notification_pkey PRIMARY KEY (id);


--
-- Name: payment payment_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (id);


--
-- Name: post_attach post_attach_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_attach
    ADD CONSTRAINT post_attach_pkey PRIMARY KEY (id);


--
-- Name: post_comment post_comment_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_comment
    ADD CONSTRAINT post_comment_pkey PRIMARY KEY (id);


--
-- Name: post_like post_like_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_like
    ADD CONSTRAINT post_like_pkey PRIMARY KEY (id);


--
-- Name: post post_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: subscription_level subscription_level_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription_level
    ADD CONSTRAINT subscription_level_pkey PRIMARY KEY (id);


--
-- Name: subscription subscription_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription
    ADD CONSTRAINT subscription_pkey PRIMARY KEY (id);


--
-- Name: user user_email_key; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: subscription_level fk_creator_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription_level
    ADD CONSTRAINT fk_creator_id FOREIGN KEY (creator_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: post fk_creator_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT fk_creator_id FOREIGN KEY (creator_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: payment fk_creator_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT fk_creator_id FOREIGN KEY (creator_id) REFERENCES public."user"(id) ON DELETE RESTRICT;


--
-- Name: subscription fk_creator_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription
    ADD CONSTRAINT fk_creator_id FOREIGN KEY (creator_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: payment fk_donater_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT fk_donater_id FOREIGN KEY (donater_id) REFERENCES public."user"(id) ON DELETE RESTRICT;


--
-- Name: post fk_min_subscription_level_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT fk_min_subscription_level_id FOREIGN KEY (min_subscription_level_id) REFERENCES public.subscription_level(id) ON DELETE RESTRICT;


--
-- Name: post_attach fk_post_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_attach
    ADD CONSTRAINT fk_post_id FOREIGN KEY (post_id) REFERENCES public.post(id) ON DELETE CASCADE;


--
-- Name: post_comment fk_post_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_comment
    ADD CONSTRAINT fk_post_id FOREIGN KEY (post_id) REFERENCES public.post(id) ON DELETE CASCADE;


--
-- Name: post_like fk_post_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_like
    ADD CONSTRAINT fk_post_id FOREIGN KEY (post_id) REFERENCES public.post(id) ON DELETE CASCADE;


--
-- Name: subscription fk_subscribe_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription
    ADD CONSTRAINT fk_subscribe_id FOREIGN KEY (subscriber_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: subscription fk_subscription_level_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.subscription
    ADD CONSTRAINT fk_subscription_level_id FOREIGN KEY (subscription_level_id) REFERENCES public.subscription_level(id) ON DELETE RESTRICT;


--
-- Name: notification fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.notification
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- Name: post_comment fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_comment
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE SET NULL;


--
-- Name: post_like fk_user_id; Type: FK CONSTRAINT; Schema: public; Owner: kopilka
--

ALTER TABLE ONLY public.post_like
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

