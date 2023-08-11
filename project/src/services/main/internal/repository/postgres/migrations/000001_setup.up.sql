create table if not exists public.users(
  id serial primary key,
  version smallint default 1,
  uuid uuid not null,
  name varchar(50) not null,
  email varchar(100) not null,
  password varchar(255) not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  deleted_at timestamp
);
