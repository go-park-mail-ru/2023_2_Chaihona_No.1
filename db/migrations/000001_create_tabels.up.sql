BEGIN;

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE public.user
(
    PRIMARY KEY (id),
    id              serial       NOT NULL,
    nickname        varchar(30)  NOT NULL,
    email           varchar(300) NOT NULL UNIQUE,
    password        varchar(50)  NOT NULL,
    is_author       boolean      NOT NULL,
    status          varchar(100),
    avatar_path     text DEFAULT 'static/default_avatar.png',
    background_path text DEFAULT 'static/default_background.png',
    description     text,
    created_at      timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notification
(
    PRIMARY KEY (id),
    id            serial   NOT NULL,
    event_type    smallint NOT NULL,
    user_id       serial   NOT NULL,
                  CONSTRAINT FK_user_id 
                  FOREIGN KEY (user_id) 
                  REFERENCES public.user (id) 
                  ON DELETE CASCADE,
    created_at    timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscription_level
(
  PRIMARY KEY (id),
  id              serial      NOT NULL,
  level           smallint    NOT NULL,
                  CONSTRAINT level_range CHECK(level >= 0),
  name            varchar(50) NOT NULL,
  description     text        NOT NULL,
  cost_integer    bigint      NOT NULL,
                  CONSTRAINT cost_integer_range CHECK(cost_integer >= 0),
  cost_fractional bigint      NOT NULL,
                  CONSTRAINT cost_fractional_range CHECK(cost_fractional >= 0),
  currency        char(3)     NOT NULL,
  creator_id      serial      NOT NULL,
                  CONSTRAINT FK_creator_id FOREIGN KEY (creator_id) REFERENCES public.user (id) ON DELETE CASCADE,
  created_at   timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at     timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post
(
  PRIMARY KEY (id),
  id                        serial       NOT NULL,
  header                    varchar(200) NOT NULL,
  body                      text         NOT NULL,
  creator_id                serial       NOT NULL,
                            CONSTRAINT FK_creator_id FOREIGN KEY (creator_id) REFERENCES public.user (id) ON DELETE CASCADE,
  min_subscription_level_id serial,
                            CONSTRAINT FK_min_subscription_level_id FOREIGN KEY (min_subscription_level_id) REFERENCES subscription_level (id) ON DELETE RESTRICT,
  created_at                timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at                timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post_attach
(
  PRIMARY KEY (id),
  id              serial       NOT NULL,
  file_path       text         NOT NULL,
  name            text         NOT NULL,
  post_id         serial       NOT NULL,
                  CONSTRAINT FK_post_id FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
  is_media        boolean      NOT NULL,
  created_at      timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment
(
  PRIMARY KEY (id),
  id                 serial   NOT NULL,
  uuid               text NOT NULL UNIQUE,
  payment_integer    bigint   NOT NULL,
                     CONSTRAINT payment_integer_range CHECK(payment_integer >= 0),
  payment_fractional bigint   NOT NULL,
                     CONSTRAINT payment_fractional_range CHECK(payment_fractional >= 0),
  status             smallint NOT NULL,
  donater_id         serial   NOT NULL,
                     CONSTRAINT FK_donater_id FOREIGN KEY (donater_id) REFERENCES public.user (id) ON DELETE RESTRICT,
  creator_id         serial   NOT NULL,
                     CONSTRAINT FK_creator_id FOREIGN KEY (creator_id) REFERENCES public.user (id) ON DELETE RESTRICT,
  payment_type       smallint NOT NULL,
  currency           text NOT NULL,
  payment_method_id  text,
  created_at         timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at         timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post_comment
(
  PRIMARY KEY (id),
  id            serial NOT NULL,
  text          text   NOT NULL,
  user_id       serial NOT NULL,
                CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES public.user (id) ON DELETE SET NULL,
  post_id       serial NOT NULL,
                CONSTRAINT FK_post_id FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
  created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at   timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post_like
(
  PRIMARY KEY (id),
  id            serial NOT NULL,
  user_id       serial NOT NULL,
                CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES public.user (id) ON DELETE CASCADE,
  post_id       serial NOT NULL,
                CONSTRAINT FK_post_id FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
  created_at    timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subscription
(
 PRIMARY KEY (id),
 id                    serial    NOT NULL,
 subscriber_id         serial    NOT NULL,
                       CONSTRAINT FK_subscribe_id FOREIGN KEY (subscriber_id) REFERENCES public.user (id) ON DELETE CASCADE,
 creator_id            serial    NOT NULL,
                       CONSTRAINT FK_creator_id FOREIGN KEY (creator_id) REFERENCES public.user (id) ON DELETE CASCADE,
 subscription_level_id serial NOT NULL,
                       CONSTRAINT FK_subscription_level_id FOREIGN KEY (subscription_level_id) REFERENCES subscription_level (id) ON DELETE RESTRICT,
 created_at            timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 updated_at            timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE question
(
  PRIMARY KEY (id),
  id              serial    NOT NULL,
  question_type   smallint  NOT NULL,
  question        text      NOT NULL
);

CREATE TABLE answer
(
  PRIMARY KEY (id),
  id          serial    NOT NULL,
  question_id serial    NOT NULL,
              CONSTRAINT FK_question_id FOREIGN KEY (question_id) REFERENCES question (id) ON DELETE CASCADE,
  user_id     serial    NOT NULL,
              CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES public.user (id) ON DELETE SET NULL,
  rating      smallint  NOT NULL
);

CREATE TABLE analytics
(
  PRIMARY KEY (id),
  id serial NOT NULL,
  user_id serial NOT NULL,
          CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES public.user (id) ON DELETE CASCADE,
  total_posts int NOT NULL,
  total_likes int NOT NULL,
  total_comments int NOT NULL,
  total_donations int NOT NULL,
  total_donations_earned_integer int NOT NULL,
  total_donations_earned_fractional int NOT NULL,
  total_earned_integer int NOT NULL,
  total_earned_fractional int NOT NULL,
  total_subscribers int NOT NULL,
  difference_posts int NOT NULL,
  difference_likes int NOT NULL,
  difference_comments int NOT NULL,
  difference_donations int NOT NULL,
  difference_donations_earned_integer int NOT NULL,
  difference_donations_earned_fractional int NOT NULL,
  difference_earned_integer int NOT NULL,
  difference_earned_fractional int NOT NULL,
  difference_subscribers int NOT NULL,
  created_at            timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_device
(
  PRIMARY KEY (id),
  id            serial NOT NULL,
  user_id       serial NOT NULL,
                CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES public.user (id) ON DELETE CASCADE,
  device_id     text NOT NULL
);


COMMIT;