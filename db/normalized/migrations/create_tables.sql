CREATE TABLE "USER"
(
 PRIMARY KEY ("id"),
 "id"              serial NOT NULL,
 nickname        varchar(20) NOT NULL,
 email           varchar(300) NOT NULL UNIQUE,
 password        varchar(50) NOT NULL,
 is_author       boolean NOT NULL,
 status          varchar(100),
 avatar_path     varchar(100) DEFAULT 'static/default_avatar.png',
 background_path varchar(100) DEFAULT 'static/default_background.png',
 creation_date   timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 last_update     timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 description     text,
);

CREATE TABLE NOTIFICATION
(
 PRIMARY KEY ("id"),
 "id"            serial NOT NULL,
 event_type    smallint NOT NULL,
 user_id       serial NOT NULL,
 creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 last_update   timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 CONSTRAINT FK_user_id FOREIGN KEY ( user_id ) REFERENCES USER ( "id" ) ON DELETE CASCADE
);

CREATE TABLE SUBSCRIPTION_LEVEL
(
 PRIMARY KEY ("id"),
 "id"              serial NOT NULL,
 level           smallint NOT NULL,
 name            varchar(30) NOT NULL,
 description     text NOT NULL,
 cost_integer    bigint NOT NULL,
 cost_fractional bigint NOT NULL,
 currency        char(3) NOT NULL,
 creation_date   timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 last_update     timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 creator_id      serial NOT NULL,
 CONSTRAINT FK_creator_id FOREIGN KEY ( creator_id ) REFERENCES USER ( "id" ) ON DELETE CASCADE,
 CONSTRAINT level_range CHECK(level >= 0),
 CONSTRAINT cost_integer_range CHECK(cost_integer >= 0),
 CONSTRAINT cost_fractional_range CHECK(cost_fractional >= 0)
);

CREATE TABLE POST
(
 PRIMARY KEY ("id"),
 "id"                     serial NOT NULL,
 header                 varchar(200) NOT NULL,
 body                   text NOT NULL,
 creation_date          timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 last_update            timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 creator_id             serial NOT NULL,
 min_subscription_level_id serial,
 CONSTRAINT FK_min_subscription_level_id FOREIGN KEY ( min_subscription_level_id ) REFERENCES SUBSCRIPTION_LEVEL ( "id" ) ON DELETE RESTRICT,
 CONSTRAINT FK_creator_id FOREIGN KEY ( creator_id ) REFERENCES USER ( "id" ) ON DELETE CASCADE
);

CREATE TABLE POST_ATTACH
(
 PRIMARY KEY ("id"),
 "id"            serial NOT NULL,
 file_path     varchar(100) NOT NULL,
 creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 post_id       serial NOT NULL,
 CONSTRAINT FK_post_id FOREIGN KEY ( post_id ) REFERENCES POST ( "id" ) ON DELETE CASCADE
);

CREATE TABLE PAYMENT
(
  PRIMARY KEY (id),
  id                 serial NOT NULL,
  payment_integer    bigint NOT NULL,
  payment_fractional bigint NOT NULL,
  status             smallint NOT NULL,
  creation_date      timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  last_update        timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  donater_id         serial NOT NULL,
  creator_id         serial NOT NULL,
  CONSTRAINT FK_creator_id FOREIGN KEY ( creator_id ) REFERENCES USER ( "id" ) ON DELETE RESTRICT,
  CONSTRAINT FK_donater_id FOREIGN KEY ( donater_id ) REFERENCES USER ( "id" ) ON DELETE RESTRICT,
  CONSTRAINT payment_integer_range CHECK(payment_integer >= 0),
  CONSTRAINT payment_fractional_range CHECK(payment_fractional >= 0)
);

CREATE TABLE POST_COMMENT
(
  PRIMARY KEY ("id"),
 "id"            serial NOT NULL,
 creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 last_update   timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 text          text NOT NULL,
 user_id       serial NOT NULL,
 post_id       serial NOT NULL,
 CONSTRAINT FK_post_id FOREIGN KEY ( post_id ) REFERENCES POST ( "id" ) ON DELETE CASCADE,
 CONSTRAINT FK_user_id FOREIGN KEY ( user_id ) REFERENCES USER ( "id" ) ON DELETE SET NULL
);

CREATE TABLE POST_LIKE
(
 PRIMARY KEY ("id"),
 "id"            serial NOT NULL,
 creation_date timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 user_id       serial NOT NULL,
 post_id       serial NOT NULL,
 CONSTRAINT FK_post_id FOREIGN KEY ( post_id ) REFERENCES POST ( "id" ) ON DELETE CASCADE,
 CONSTRAINT FK_user_id FOREIGN KEY ( user_id ) REFERENCES "USER" ( "id" ) ON DELETE CASCADE
);

CREATE TABLE SUBSCRIPTION
(
 PRIMARY KEY ("id"),
 "id"                 serial NOT NULL,
 creation_date      timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 last_update        timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 subscriber_id      serial NOT NULL,
 creator_id         serial NOT NULL,
 subscription_level_id serial NOT NULL,
 CONSTRAINT FK_subscription_level_id FOREIGN KEY ( subscription_level_id ) REFERENCES SUBSCRIPTION_LEVEL ( "id" ) ON DELETE RESTRICT,
 CONSTRAINT FK_creator_id FOREIGN KEY ( creator_id ) REFERENCES "USER" ( "id" ) ON DELETE CASCADE,
 CONSTRAINT FK_subscribe_id FOREIGN KEY ( subscriber_id ) REFERENCES "USER" ( "id" ) ON DELETE CASCADE
);
