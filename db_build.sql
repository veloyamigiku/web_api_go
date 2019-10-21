# user(gwp) create

# db(gwp) create

create table posts(
    id SERIAL NOT NULL,
    content text,
    author varchar(255));

create table users(
    id SERIAL NOT NULL,
    name text NOT NULL,
    password text NOT NULL);
