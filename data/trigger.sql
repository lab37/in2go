

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

create TRIGGER add_stocks4incoms after
insert
ON
incomes
BEGIN
  insert or
  ignore into stocks (cstmid,prdtid)
  values
    ((select cstmid
      from contracts
      where contracts.ccid=new.ccid
  limit 1),new.prdtid);
update stocks set quantity = quantity+new.quantity where prdtid=new.prdtid and cstmid=(select cstmid
  from contracts
  where contracts.ccid=new.ccid
limit 1);
END;

create TRIGGER del_stocks4incoms after
delete
ON
incomes
BEGIN
update stocks set quantity = quantity-old.quantity where prdtid=old.prdtid and cstmid=(select cstmid
  from contracts
  where contracts.ccid=old.ccid
limit 1);
END;

create TRIGGER upt_stocks4incoms after
update
ON
incomes
BEGIN
update stocks set quantity = quantity-old.quantity where prdtid=old.prdtid and cstmid=(select cstmid
  from contracts
  where contracts.ccid=old.ccid
limit 1);
update stocks set quantity = quantity+new.quantity where prdtid=new.prdtid and cstmid=(select cstmid
  from contracts
  where contracts.ccid=new.ccid
limit 1);
END;



create TRIGGER del_stocks4outgos after
insert
ON
outgos
BEGIN
  update stocks set quantity = quantity-new.quantity where prdtid=new.prdtid and cstmid=(select cstmid
    from contracts
    where contracts.ccid=new.ccid
  limit 1);
END;

create TRIGGER add_stocks4outgos after
delete
ON
outgos
BEGIN
  update stocks set quantity = quantity+old.quantity where prdtid=old.prdtid and cstmid=(select cstmid
    from contracts
    where contracts.ccid=old.ccid
  limit 1);
END;

create TRIGGER upt_stocks4outgos after
update
ON
outgos
BEGIN
  update stocks set quantity = quantity+old.quantity where prdtid=old.prdtid and cstmid=(select cstmid
    from contracts
    where contracts.ccid=old.ccid
  limit 1);
  update stocks set quantity = quantity-new.quantity where prdtid=new.prdtid and cstmid=(select cstmid
    from contracts
    where contracts.ccid=new.ccid
  limit 1);
END;





create TRIGGER del_debts4payments after
insert 
ON
payments
BEGIN
  insert or
  ignore into debts (srcid,cstmid)
  values
    ((select cstmid
      from contracts0
      where contracts0.ccid=new.ccid
  limit 1),
  (select cstmid
  from contracts
  where contracts.ccid=new.ccid
  limit 1));
  update debts set dbtsum = dbtsum-new.pmsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=new.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=new.ccid);
END;

create TRIGGER add_debts4payments after
delete 
ON
payments
BEGIN
  
  update debts set dbtsum = dbtsum+old.pmsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=old.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=old.ccid);
END;

create TRIGGER upt_debts4payments after
update 
ON
payments
BEGIN
  
  update debts set dbtsum = dbtsum+old.pmsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=old.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=old.ccid);
  update debts set dbtsum = dbtsum-new.pmsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=new.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=new.ccid);
END;


create TRIGGER add_debts4invoices after
insert
ON
invoices
BEGIN
  insert or
  ignore into debts (srcid,cstmid)
  values
    ((select cstmid
      from contracts0
      where contracts0.ccid=new.ccid
  limit 1),
  (select cstmid
  from contracts
  where contracts.ccid=new.ccid
  limit 1));

  update debts set dbtsum = dbtsum+new.ivsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=new.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=new.ccid
  limit 1);
END;

create TRIGGER del_debts4invoices after
delete
ON
invoices
BEGIN
  
  update debts set dbtsum = dbtsum-old.ivsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=old.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=old.ccid
  limit 1);
END;

create TRIGGER upt_debts4invoices after
update
ON
invoices
BEGIN
  
  update debts set dbtsum = dbtsum-old.ivsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=old.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=old.ccid
  limit 1);
  update debts set dbtsum = dbtsum+new.ivsum where srcid=(select cstmid
  from contracts0
  where contracts0.ccid=new.ccid
  limit 1) and cstmid=
  (select cstmid
  from contracts
  where contracts.ccid=new.ccid
  limit 1);
END;

create TRIGGER insert_onway_products after
insert
ON
contracts
BEGIN
  insert into onway_products
    (ccid,prdtid,quantity,cksum)
  values(new.ccid, new.prdtid, new.quantity, new.quantity);
END;
create TRIGGER update_onway_products after
update
ON
contracts
BEGIN
  insert into onway_products
    (ccid,prdtid,quantity,cksum)
  values(new.ccid, new.prdtid, new.quantity, new.quantity);
END;

create TRIGGER add_onway_products4incomes after
insert
ON
incomes
BEGIN
  update onway_products set cksum=cksum-new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;

create TRIGGER del_onway_products4incomes after
delete
ON
incomes
BEGIN
  update onway_products set cksum=cksum+old.quantity where ccid=old.ccid and prdtid=old.prdtid;
END;

create TRIGGER upt_onway_products4incomes after
update
ON
incomes
BEGIN
  update onway_products set cksum=cksum+old.quantity where ccid=old.ccid and prdtid=old.prdtid;
  update onway_products set cksum=cksum-new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;




create TRIGGER del_onway_products4outgos after
insert
ON
outgos
BEGIN
  update onway_products set cksum=cksum-new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;

create TRIGGER add_onway_products4outgos after
delete
ON
outgos
BEGIN
  update onway_products set cksum=cksum+old.quantity where ccid=old.ccid and prdtid=old.prdtid;
END;

create TRIGGER del_onway_products4outgos after
update
ON
outgos
BEGIN
  update onway_products set cksum=cksum+old.quantity where ccid=old.ccid and prdtid=old.prdtid;
  update onway_products set cksum=cksum-new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;



create TRIGGER insert_onway_invoice4contracts after
insert
ON
contracts
BEGIN
  insert or
  ignore into onway_invoices (ccid)
  values
    (new.ccid);
  update onway_invoices set csum=(csum+new.quantity * new.price) where ccid=new.ccid;
END;

create TRIGGER del_onway_invoice4contracts after
delete
ON
contracts
BEGIN
  update onway_invoices set csum=csum-(old.quantity * old.price) where ccid=old.ccid;
END;

create TRIGGER upt_onway_invoice4contracts after
update
ON
contracts
BEGIN
  update onway_invoices set csum=csum-(old.quantity * old.price) where ccid=old.ccid;
  update onway_invoices set csum=csum+(new.quantity * new.price) where ccid=new.ccid;
END;




create TRIGGER add_onway_invoices4invoices after
insert
ON
invoices
BEGIN
  update onway_invoices set cksum=cksum+new.ivsum where ccid=new.ccid;
END;

create TRIGGER del_onway_invoices4invoices after
delete
ON
invoices
BEGIN
  update onway_invoices set cksum=cksum-old.ivsum where ccid=old.ccid;
END;

create TRIGGER upt_onway_invoices4invoices after
update
ON
invoices
BEGIN
  update onway_invoices set cksum=cksum-old.ivsum where ccid=old.ccid;
  update onway_invoices set cksum=cksum+new.ivsum where ccid=new.ccid;
END;