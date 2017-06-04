CREATE TABLE users
(
  uid serial NOT NULL,
  name character varying(255) NOT NULL,
  email character varying(255) NOT NULL,
  age integer NOT NULL,
  city character varying(255) NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);