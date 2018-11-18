create table contract (
  id          integer primary key,
  ccid        text not null,
  ccdata      text not null,
  cctype      text not null,
  cstmname    text not null,
  prdtid      text not null  references products(prdtid),
  price       real not null,
  quantity    integer not null,
  remark      text
);

create table income (
  id          integer primary key,
  ccid        text not null,
  icdata      text not null,
  prdtid      text not null references products(prdtid),
  quantity    integer not null,
  pnumber     text not null,
  remark      texts
);

create table outgo (
  id          integer primary key,
  ccid        text not null,
  ogdata      text not null,
  prdtid      text not null references products(prdtid),
  pnumber     text not null,
  quantity    integer not null,
  expname     text,
  expnumber   text,
  remark      text
);

create table invoice (
  id          integer primary key,
  ivid        text not null,
  ivdata      text not null,
  ccid        text not null ,
  ivsum       real not null,
  postdata    text,
  expname     text,
  expnumber   text,
  remark      text
);

create table payment (
  id          integer primary key,
  ccid        text ,
  pmdata      text not null,
  pmsum       real not null,
  remark      text
);

create table products (
  id          integer primary key,
  prdtid      text not null,
  prdtname    text not null,
  specific    text not null,
  inventor    text not null,
  unit        text not null,
  boxnumb     integer not null,
  inline      text,
  ivtype      text,
  remark      text
);

create table customers (
  id          integer primary key,
  cstmid      text not null,
  cstmname    text not null,
  cstmtype    text not null,
  gaddr       text,
  gname       text,
  gphone      text,
  ivaddr      text,
  ivname      text,
  ivphone     text,
  remark      text
);