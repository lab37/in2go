create table contracts0
(
  ccid text primary key not null,
  cstmid integer not null references customers(cstmid),
  vector INTEGER not null check (vector<2 and vector >=0),
  remark text DEFAULT '_'
);
create table contracts
(
  id integer primary key not null,
  ccid text not null,
  ccdate text DEFAULT current_date,
  cctype text DEFAULT '-',
  cstmid integer not null references customers(cstmid),
  prdtid text not null references products(prdtid),
  price real not null,
  quantity integer not null,
  remark text DEFAULT '-',
  unique(ccid,prdtid)
);


create table incomes
(
  id integer primary key not null,
  ccid text not null references contract0(ccid),
  icdate text DEFAULT current_date,
  prdtid text not null references products(prdtid),
  quantity integer not null,
  pnumber text DEFAULT '-',
  remark text DEFAULT '-'
);

create table outgos
(
  id integer primary key not null,
  ccid text not null references contract0(ccid),
  ogdate text DEFAULT current_date,
  prdtid text not null references products(prdtid),
  pnumber text DEFAULT '-',
  quantity integer not null,
  expname text DEFAULT '-',
  expnumber text DEFAULT '-',
  remark text DEFAULT '-'
);

create table invoices
(
  id integer primary key not null,
  ivid text not null,
  ivdate text DEFAULT current_date,
  ccid text not null references contracts0(ccid),
  ivsum real not null,
  postdate text DEFAULT '-',
  expname text DEFAULT '-',
  expnumber text DEFAULT '-',
  remark text DEFAULT '-',
  unique(ccid,ivid)
);

create table payments
(
  id integer primary key not null,
  ccid text not null references contracts0(ccid),
  pmdate text DEFAULT current_date,
  pmsum real not null,
  remark text DEFAULT '-'
);

create table products
(
  prdtid text primary key not null,
  prdtname text not null,
  specific text not null,
  inventor text not null,
  unit text not null,
  boxnumb integer not null,
  inline text DEFAULT '-',
  ivtype text DEFAULT '-',
  remark text DEFAULT '-',
  unique(prdtname,specific)
);

create table customers
(
  cstmid integer primary key not null,
  cstmname text not null UNIQUE,
  cstmtype text DEFAULT '-',
  gaddr text DEFAULT '-',
  gname text DEFAULT '-',
  gphone text DEFAULT '-',
  ivaddr text DEFAULT '-',
  ivname text DEFAULT '-',
  ivphone text DEFAULT '-',
  remark text DEFAULT '-'
);

create table stocks
(
  id integer PRIMARY key not null,
  prdtid text not null references products(prdtid),
  cstmid integer not null references customers(cstmid),
  pnumber text DEFAULT '-',
  quantity integer default 0 check(quantity>=0),
  remark text DEFAULT '-',
  unique(cstmid,prdtid)
);

create table debts
(
  id integer PRIMARY key not null,
  srcid integer not null references customers(cstmid),
  cstmid integer not null references customers(cstmid),
  dbtsum real default 0 check(dbtsum>=0),
  remark text DEFAULT '-',
  UNIQUE(srcid,cstmid)
);

create table onway_products
(
  id integer PRIMARY key not null,
  ccid text not null REFERENCES contracts0(ccid),
  prdtid text not null REFERENCES products(prdtid),
  quantity INTEGER default 0,
  cksum INTEGER default 0 check(cksum<=quantity),
  remark text DEFAULT '-',
  unique(ccid,prdtid)
);

create table onway_invoices
(
  id integer PRIMARY key not null,
  ccid text not null unique REFERENCES contracts0(ccid),
  csum real default 0,
  cksum real default 0 check(cksum<=csum),
  remark text DEFAULT '-'
);
