create extension postgis;
create extension "uuid-ossp";

create table partner(
  id     uuid             not null primary key default gen_random_uuid(),
  name   text             not null,
  geo    geography(point) not null,
  radius float            not null,
  rating int              not null
);
create index on partner using gist (geo);

create table partner_skill(
  partner_id uuid not null references partner(id) on update restrict on delete restrict,
  code       text not null,
  unique(partner_id, code)
);

create index on partner_skill(code);
