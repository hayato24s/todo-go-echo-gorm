BEGIN;

CREATE TABLE users (
  id uuid primary key,
  name varchar(255) not null,
  password varchar(255) not null,
  unique (name)
);

CREATE TABLE tasks (
  id uuid primary key,
  user_id uuid not null,
  title varchar(255) not null,
  completed boolean not null,
  created_at timestamp without time zone not null,
  foreign key (user_id) references users (id)
);

COMMIT;
