create view stocks_auto
AS
  select cstmid, prdtid, isum, (select osum
    from (select (select cstmid
        from contracts0
        where contracts0.ccid=outgos.ccid) as cstmid, prdtid, sum(quantity) as osum
      from outgos
      group by cstmid,prdtid) o
    where o.cstmid=i.cstmid and o.prdtid = i.prdtid )
  from ((select cstmid, prdtid, sum(quantity) as isum
    from incomes
    group by cstmid,prdtid) i);

create view debts_auto
as
  select ccid, ivsum, (select pmtsum
    from (select ccid, sum(pmsum) as pmtsum
      from payments
      group by ccid) o
    where o.ccid=c.ccid) as osum
  from (select ccid, sum(ivsum) as ivsum
    from invoices
    group by ccid) c;


create view inway_products_auto
as
  select ccid, prdtid, quantity, (select qsum
    from (select ccid, prdtid, sum(quantity) as qsum
      from incomes
      group by ccid,prdtid) tmp
    where tmp.ccid=contracts.ccid and tmp.prdtid=contracts.prdtid ) as isum
  from contracts
  where ccid in (select ccid
  from contracts0
  where contracts0.vector = 0);

create view outway_products_auto
as
  select ccid, prdtid, quantity, (select qsum
    from (select ccid, prdtid, sum(quantity) as qsum
      from outgos
      group by ccid,prdtid) tmp
    where tmp.ccid=contracts.ccid and tmp.prdtid=contracts.prdtid ) as osum
  from contracts
  where ccid in (select ccid
  from contracts0
  where contracts0.vector = 1);

create view inway_invoices_auto
AS
  select ccid, csum, (select ivtsum
    from (select ccid, sum(ivsum) as ivtsum
      from invoices
      group by ccid) t2
    where t2.ccid=t1.ccid ) as isum
  from (select ccid, sum(price*quantity) as csum
    from contracts
    where ccid in (select ccid
    from contracts0
    where contracts0.vector=0)
    group by ccid) t1;


create view outway_invoices_auto
AS
  select ccid, csum, (select ovtsum
    from (select ccid, sum(ivsum) as ovtsum
      from invoices
      group by ccid) t2
    where t2.ccid=t1.ccid ) as osum
  from (select ccid, sum(price*quantity) as csum
    from contracts
    where ccid in (select ccid
    from contracts0
    where contracts0.vector=1)
    group by ccid) t1;
