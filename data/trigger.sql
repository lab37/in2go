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
    from contracts0
    where contracts0.ccid=new.ccid
  limit 1);
END;

create TRIGGER add_stocks4outgos after
delete
ON
outgos
BEGIN
  update stocks set quantity = quantity+old.quantity where prdtid=old.prdtid and cstmid=(select cstmid
    from contracts0
    where contracts0.ccid=old.ccid
  limit 1);
END;

create TRIGGER upt_stocks4outgos after
update
ON
outgos
BEGIN
  update stocks set quantity = quantity+old.quantity where prdtid=old.prdtid and cstmid=(select cstmid
    from contracts0
    where contracts0.ccid=old.ccid
  limit 1);
  update stocks set quantity = quantity-new.quantity where prdtid=new.prdtid and cstmid=(select cstmid
    from contracts0
    where contracts0.ccid=new.ccid
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
  values(new.ccid, new.prdtid, new.quantity, 0);
END;


create TRIGGER update_onway_products after
update
ON
contracts
BEGIN
  insert into onway_products
    (ccid,prdtid,quantity,cksum)
  values(new.ccid, new.prdtid, new.quantity, 0);
END;

create TRIGGER add_onway_products4incomes after
insert
ON
incomes
BEGIN
  update onway_products set cksum=cksum+new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;

create TRIGGER del_onway_products4incomes after
delete
ON
incomes
BEGIN
  update onway_products set cksum=cksum-old.quantity where ccid=old.ccid and prdtid=old.prdtid;
END;

create TRIGGER upt_onway_products4incomes after
update
ON
incomes
BEGIN
  update onway_products set cksum=cksum-old.quantity where ccid=old.ccid and prdtid=old.prdtid;
  update onway_products set cksum=cksum+new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;




create TRIGGER add_onway_products4outgos after
insert
ON
outgos
BEGIN
  update onway_products set cksum=cksum+new.quantity where ccid=new.ccid and prdtid=new.prdtid;
END;

create TRIGGER del_onway_products4outgos after
delete
ON
outgos
BEGIN
  update onway_products set cksum=cksum-old.quantity where ccid=old.ccid and prdtid=old.prdtid;
END;

create TRIGGER upt_onway_products4outgos after
update
ON
outgos
BEGIN
  update onway_products set cksum=cksum-old.quantity where ccid=old.ccid and prdtid=old.prdtid;
  update onway_products set cksum=cksum+new.quantity where ccid=new.ccid and prdtid=new.prdtid;
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