package data

type Contract0 struct {
	CcId     string
	CstmId   int
	CstmName string
	Vector   int
	Remark   string
}

type Contract struct {
	Id       int
	CcId     string
	Vector   int
	SrcName  string
	CcDate   string
	CcType   string
	CstmId   int
	CstmName string
	PrdtId   string
	PrdtName string
	Specific string
	Price    float64
	Quantity int
	Remark   string
}
type Income struct {
	Id       int
	CcId     string
	SrcName  string
	CstmName string
	IcDate   string
	PrdtId   string
	PrdtName string
	Specific string
	Price    float64
	Quantity int
	Pnumber  string
	Remark   string
}
type Outgo struct {
	Id        int
	CcId      string
	CstmName  string
	SrcName   string
	OgDate    string
	PrdtId    string
	PrdtName  string
	Specific  string
	Price     float64
	Pnumber   string
	Quantity  int
	ExpName   string
	ExpNumber string
	Remark    string
}
type Invoice struct {
	Id        int
	IvId      string
	IvDate    string
	CcId      string
	SrcName   string
	CstmName  string
	IvSum     float64
	PostDate  string
	ExpName   string
	ExpNumber string
	Remark    string
}

type Payment struct {
	Id       int
	CcId     string
	SrcName  string
	CstmName string
	PmDate   string
	PmSum    float64
	Remark   string
}

type Product struct {
	PrdtId   string
	PrdtName string
	Specific string
	Inventor string
	Unit     string
	BoxNumb  int
	Inline   string
	IvType   string
	Remark   string
}
type Customer struct {
	CstmId   int
	CstmName string
	CstmType string
	Gaddr    string
	Gname    string
	Gphone   string
	IvAddr   string
	IvName   string
	IvPhone  string
	Remark   string
}

type Stock struct {
	Id       int
	CstmId   int
	CstmName string
	PrdtId   string
	PrdtName string
	Specific string
	Pnumber  string
	Quantity int
	Remark   string
}

type Debt struct {
	Id       int
	SrcId    int
	SrcName  string
	CstmId   int
	CstmName string
	DbtSum   float64
	Remark   string
}

type OnwayProduct struct {
	Id       int
	CcId     string
	SrcName  string
	CstmName string
	Quantity int
	CkSum    int
	PrdtId   int
	PrdtName string
	Specific string
	Vector   int
	Remark   string
}

type OnwayInvoice struct {
	Id       int
	CcId     string
	SrcName  string
	CstmName string
	Csum     float64
	CkSum    float64
	Remark   string
	Vector   int
}

func (r *Contract) Insert() (err error) {
	statement := "insert into contracts(ccid,ccdate, cctype, cstmid, prdtid, price, quantity, remark) values (?,?,?,(select cstmid from customers where customers.cstmname = ? limit 1),(select prdtid from products where products.prdtname like ? and products.specific like ? limit 1),?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CcId, r.CcDate, r.CcType, r.CstmName, "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Price, r.Quantity, r.Remark)
	return
}

func (r *Contract) Select() (contracts []Contract, err error) {
	rows, err := Db.Query("SELECT * ,(select vector from contracts0 where contracts0.ccid=contracts.ccid limit 1) AS vector,(SELECT cstmname FROM customers WHERE customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=contracts.ccid limit 1)  limit 1) AS srcname, (SELECT cstmname FROM customers WHERE contracts.cstmid = customers.cstmid limit 1) AS cstmname,(SELECT prdtname FROM products WHERE contracts.prdtid = products.prdtid limit 1) AS prdtname, (SELECT specific FROM products WHERE contracts.prdtid = products.prdtid limit 1) AS specific FROM contracts where prdtid in (select prdtid from products where products.prdtname like ? and products.specific like ?) and ccid like ? and ccdate like ? and cctype like ? and cstmid in (select cstmid from customers where customers.cstmname like ?) and remark like ? and ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?)) and ccid in (select ccid from contracts0 where contracts0.vector <> ?) order by ccdate desc", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.CcId+"%", "%"+r.CcDate+"%", "%"+r.CcType+"%", "%"+r.CstmName+"%", "%"+r.Remark+"%", "%"+r.SrcName+"%", r.Vector)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		contract := Contract{}
		if err = rows.Scan(&contract.Id, &contract.CcId, &contract.CcDate, &contract.CcType, &contract.CstmId, &contract.PrdtId, &contract.Price, &contract.Quantity, &contract.Remark, &contract.Vector, &contract.SrcName, &contract.CstmName, &contract.PrdtName, &contract.Specific); err != nil {

			return
		}
		contracts = append(contracts, contract)
	}

	return
}

func GetAllContract() (contracts []Contract, err error) {

	rows, err := Db.Query("SELECT * ,(select vector from contracts0 where contracts0.ccid=contracts.ccid limit 1) AS vector,(SELECT cstmname FROM customers WHERE customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=contracts.ccid limit 1) limit 1) AS srcname, (SELECT cstmname FROM customers WHERE contracts.cstmid = customers.cstmid limit 1) AS cstmname,(SELECT prdtname FROM products WHERE contracts.prdtid = products.prdtid limit 1) AS prdtname, (SELECT specific FROM products WHERE contracts.prdtid = products.prdtid limit 1) AS specific FROM contracts  order by ccdate desc")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		contract := Contract{}
		if err = rows.Scan(&contract.Id, &contract.CcId, &contract.CcDate, &contract.CcType, &contract.CstmId, &contract.PrdtId, &contract.Price, &contract.Quantity, &contract.Remark, &contract.Vector, &contract.SrcName, &contract.CstmName, &contract.PrdtName, &contract.Specific); err != nil {

			return
		}
		contracts = append(contracts, contract)
	}

	return
}

func (r *Contract) Update() (err error) {
	_, err = Db.Exec("UPDATE contracts set ccid=?, ccdate=?,cctype=?,cstmid=(select cstmid from customers where customers.cstmname like ? limit 1),prdtid=(select prdtid from products where products.prdtname like ? and products.specific like ? limit 1),price=?,quantity=?,remark=? where id=?", r.CcId, r.CcDate, r.CcType, "%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Price, r.Quantity, r.Remark, r.Id)
	return
}

func (r *Contract) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM contracts WHERE id=?", r.Id)
	return
}

func (r *Income) Insert() (err error) {
	statement := "insert into incomes(ccid,icdate, prdtid, quantity, pnumber, remark) values (?,?,(select prdtid from products where products.prdtname like ? limit 1),?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CcId, r.IcDate, "%"+r.PrdtName+"%", r.Quantity, r.Pnumber, r.Remark)
	return
}

func (r *Income) Select() (incomes []Income, err error) {
	rows, err := Db.Query("SELECT *, (SELECT prdtname FROM products WHERE incomes.prdtid = products.prdtid) AS prdtname, (SELECT specific FROM products WHERE incomes.prdtid = products.prdtid) AS specific, (SELECT cstmname FROM customers WHERE customers.cstmid=(select cstmid from contracts0 where contracts0.ccid=incomes.ccid)) AS cstmname, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=incomes.ccid limit 1) limit 1) as srcname,(select price from contracts where contracts.ccid=incomes.ccid limit 1) as price FROM incomes WHERE ccid LIKE ? AND icdate LIKE ? AND prdtid in (select prdtid from products where prdtname LIKE ? AND specific LIKE ?) AND pnumber LIKE ? AND remark LIKE ? and ccid in (select ccid from contracts where contracts.cstmid in (select cstmid from customers where customers.cstmname like ?)) and ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?)) order by icdate desc", "%"+r.CcId+"%", "%"+r.IcDate+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.Pnumber+"%", "%"+r.Remark+"%", "%"+r.CstmName+"%", "%"+r.SrcName+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		income := Income{}
		if err = rows.Scan(&income.Id, &income.CcId, &income.IcDate, &income.PrdtId, &income.Quantity, &income.Pnumber, &income.Remark, &income.PrdtName, &income.Specific, &income.CstmName, &income.SrcName, &income.Price); err != nil {
			return
		}
		incomes = append(incomes, income)
	}

	return
}

func GetAllIncome() (incomes []Income, err error) {
	rows, err := Db.Query("SELECT *, (SELECT prdtname FROM products WHERE incomes.prdtid = products.prdtid) AS prdtname, (SELECT specific FROM products WHERE incomes.prdtid = products.prdtid) AS specific, (SELECT cstmname FROM customers WHERE customers.cstmid=(select cstmid from contracts0 where contracts0.ccid=incomes.ccid)) AS cstmname, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=incomes.ccid limit 1) limit 1) as srcname,(select price from contracts where contracts.ccid=incomes.ccid limit 1) as price FROM incomes  order by icdate desc")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		income := Income{}
		if err = rows.Scan(&income.Id, &income.CcId, &income.IcDate, &income.PrdtId, &income.Quantity, &income.Pnumber, &income.Remark, &income.PrdtName, &income.Specific, &income.CstmName, &income.SrcName, &income.Price); err != nil {
			return
		}
		incomes = append(incomes, income)
	}

	return
}

func (r *Income) Update() (err error) {
	_, err = Db.Exec("UPDATE incomes set ccid=?,icdate=?,prdtid=(select prdtid from products where products.prdtname like ? limit 1),quantity=?,pnumber=?,remark=? where id=?", r.CcId, r.IcDate, "%"+r.PrdtName+"%", r.Quantity, r.Pnumber, r.Remark, r.Id)
	return
}

func (r *Income) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM incomes WHERE id=?", r.Id)
	return
}

func (r *Outgo) Insert() (err error) {
	statement := "insert into outgos (ccid,ogdate, prdtid, pnumber, quantity, expname, expnumber, remark) values (?,?,(select prdtid from products where products.prdtname like ? and products.specific like ?),?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CcId, r.OgDate, "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Pnumber, r.Quantity, r.ExpName, r.ExpName, r.Remark)
	return
}

func (r *Outgo) Select() (outgos []Outgo, err error) {
	rows, err := Db.Query("SELECT *, (SELECT prdtname FROM products WHERE products.prdtid = outgos.prdtid limit 1) AS prdtname, (SELECT specific FROM products WHERE  products.prdtid = outgos.prdtid limit 1) AS specific, (SELECT cstmname FROM customers WHERE customers.cstmid=(select cstmid from contracts where contracts.ccid = outgos.ccid limit 1) limit 1) AS cstmname, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid = outgos.ccid limit 1) limit 1) as srcname,(select price form contracts where contracts.ccid = outgos.ccid limit 1) as price FROM outgos WHERE ccid LIKE ? AND ogdate LIKE ? AND prdtid in (select prdtid from products where prdtname LIKE ? AND specific LIKE ?)  AND pnumber LIKE ? AND remark LIKE ? and ccid in (select ccid from contracts where contracts.cstmid in (select cstmid from customers where customers.cstmname like ?)) and ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?))  order by ogdate desc", "%"+r.CcId+"%", "%"+r.OgDate+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.Pnumber+"%", "%"+r.Remark+"%", "%"+r.CstmName+"%", "%"+r.SrcName+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		outgo := Outgo{}
		if err = rows.Scan(&outgo.Id, &outgo.CcId, &outgo.OgDate, &outgo.PrdtId, &outgo.Pnumber, &outgo.Quantity, &outgo.ExpName, &outgo.ExpNumber, &outgo.Remark, &outgo.PrdtName, &outgo.Specific, &outgo.CstmName, &outgo.SrcName, &outgo.Price); err != nil {
			return
		}
		outgos = append(outgos, outgo)
	}

	return
}

func GetAllOutgo() (outgos []Outgo, err error) {
	rows, err := Db.Query("SELECT *, (SELECT prdtname FROM products WHERE products.prdtid = outgos.prdtid limit 1) AS prdtname, (SELECT specific FROM products WHERE  products.prdtid = outgos.prdtid limit 1) AS specific, (SELECT cstmname FROM customers WHERE customers.cstmid=(select cstmid from contracts where contracts.ccid = outgos.ccid limit 1) limit 1) AS cstmname, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid = outgos.ccid limit 1) limit 1) as srcname,(select price form contracts where contracts.ccid = outgos.ccid limit 1) as price from outgos order by ogdate desc")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		outgo := Outgo{}
		if err = rows.Scan(&outgo.Id, &outgo.CcId, &outgo.OgDate, &outgo.PrdtId, &outgo.Pnumber, &outgo.Quantity, &outgo.ExpName, &outgo.ExpNumber, &outgo.Remark, &outgo.PrdtName, &outgo.Specific, &outgo.CstmName, &outgo.SrcName, &outgo.Price); err != nil {
			return
		}
		outgos = append(outgos, outgo)
	}

	return
}

func (r *Outgo) Update() (err error) {
	_, err = Db.Exec("UPDATE outgos set ccid=?,ogdate=?,prdtid=(select prdtid from products where products.prdtname like ? and products.specific like ? limit 1),pnumber=?,quantity=?,expname=?,expnumber=?,remark=? where id=?", r.CcId, r.OgDate, "%"+r.PrdtName+"%", "%"+r.Specific+"%", r.Pnumber, r.Quantity, r.ExpName, r.ExpNumber, r.Remark, r.Id)
	return
}

func (r *Outgo) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM outgos WHERE id=?", r.Id)
	return
}

func (r *Invoice) Insert() (err error) {
	statement := "insert into invoices(ivid, ivdate, ccid, ivsum, postdate, expname, expnumber, remark) values(?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.IvId, r.IvDate, r.CcId, r.IvSum, r.PostDate, r.ExpName, r.ExpNumber, r.Remark)
	return
}

func (r *Invoice) Select() (invoices []Invoice, err error) {
	rows, err := Db.Query("SELECT * , (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=invoices.ccid limit 1) limit 1) as srcname,(SELECT cstmname FROM customers where customers.cstmid = (select cstmid from contracts WHERE contracts.ccid = invoices.ccid limit 1) limit 1) AS cstmname FROM invoices WHERE ivid LIKE ? AND ivdate LIKE ? AND ccid LIKE ? AND postdate LIKE ? AND remark LIKE ? and ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?)) and ccid in (select ccid from contracts where contracts.cstmid in (select cstmid from customers where customers.cstmname like ?)) order by ivdate desc", "%"+r.CcId+"%", "%"+r.IvDate+"%", "%"+r.CcId+"%", "%"+r.PostDate+"%", "%"+r.Remark+"%", "%"+r.SrcName+"%", "%"+r.CstmName+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		invoice := Invoice{}
		if err = rows.Scan(&invoice.Id, &invoice.IvId, &invoice.IvDate, &invoice.CcId, &invoice.IvSum, &invoice.PostDate, &invoice.ExpName, &invoice.ExpNumber, &invoice.Remark, &invoice.SrcName, &invoice.CstmName); err != nil {
			return
		}
		invoices = append(invoices, invoice)
	}

	return
}

func GetAllInvoice() (invoices []Invoice, err error) {
	rows, err := Db.Query("SELECT * , (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=invoices.ccid limit 1) limit 1) as srcname,(SELECT cstmname FROM customers where customers.cstmid = (select cstmid from contracts WHERE contracts.ccid = invoices.ccid limit 1) limit 1) AS cstmname FROM invoices  order by ivdate desc")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		invoice := Invoice{}
		if err = rows.Scan(&invoice.Id, &invoice.IvId, &invoice.IvDate, &invoice.CcId, &invoice.IvSum, &invoice.PostDate, &invoice.ExpName, &invoice.ExpNumber, &invoice.Remark, &invoice.SrcName, &invoice.CstmName); err != nil {
			return
		}
		invoices = append(invoices, invoice)
	}

	return
}

func (r *Invoice) Update() (err error) {
	_, err = Db.Exec("UPDATE invoices set ivid=?,ivdate=?,ccid=?,ivsum=?,postdate=?,expname=?,expnumber=?,remark=? where id=?", r.IvId, r.IvDate, r.CcId, r.IvSum, r.PostDate, r.ExpName, r.ExpNumber, r.Remark, r.Id)
	return
}

func (r *Invoice) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM invoices WHERE id=?", r.Id)
	return
}

func (r *Payment) Insert() (err error) {
	statement := "insert into payments (ccid, pmdate, pmsum, remark) values (?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CcId, r.PmDate, r.PmSum, r.Remark)
	return
}

func (r *Payment) Select() (payments []Payment, err error) {
	rows, err := Db.Query("SELECT * ,(SELECT cstmname FROM customers where customers.cstmid = (select cstmid from contracts WHERE contracts.ccid = payments.ccid limit 1) limit 1) AS cstmname, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=payments.ccid limit 1) limit 1) as srcname FROM payments WHERE pmdate LIKE ?  AND  remark LIKE ? AND ccid like ? and ccid IN (SELECT ccid FROM contracts WHERE contracts.cstmid in (select cstmid from customers where customers.cstmname like ?)) and ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?)) order by pmdate desc", "%"+r.PmDate+"%", "%"+r.Remark+"%", "%"+r.CcId+"%", "%"+r.CstmName+"%", "%"+r.SrcName+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		payment := Payment{}
		if err = rows.Scan(&payment.Id, &payment.CcId, &payment.PmDate, &payment.PmSum, &payment.Remark, &payment.CstmName, &payment.SrcName); err != nil {
			return
		}
		payments = append(payments, payment)
	}

	return
}

func GetAllPayment() (payments []Payment, err error) {
	rows, err := Db.Query("SELECT * ,(SELECT cstmname FROM customers where customers.cstmid = (select cstmid from contracts WHERE contracts.ccid = payments.ccid limit 1) limit 1) AS cstmname, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid=payments.ccid limit 1) limit 1) as srcname FROM payments order by pmdate desc")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		payment := Payment{}
		if err = rows.Scan(&payment.Id, &payment.CcId, &payment.PmDate, &payment.PmSum, &payment.Remark, &payment.CstmName, &payment.SrcName); err != nil {
			return
		}
		payments = append(payments, payment)
	}

	return
}

func (r *Payment) Update() (err error) {
	_, err = Db.Exec("UPDATE payments set ccid=?,pmdate=?,pmsum=?,remark=? where id=?", r.CcId, r.PmDate, r.PmSum, r.Remark, r.Id)
	return
}

func (r *Payment) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM payments WHERE id=?", r.Id)
	return
}

func (r *Product) Insert() (err error) {
	statement := "insert into products (prdtid, prdtname, specific, inventor, unit, boxnumb, inline, ivtype, remark) values (?,?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.PrdtId, r.PrdtName, r.Specific, r.Inventor, r.Unit, r.BoxNumb, r.Inline, r.IvType, r.Remark)
	return
}

func (r *Product) Select() (products []Product, err error) {
	rows, err := Db.Query("SELECT * FROM products WHERE prdtid LIKE ? AND prdtname LIKE ? AND specific LIKE ? AND inventor LIKE ? AND inline LIKE ?", "%"+r.PrdtId+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.Inventor+"%", "%"+r.Inline+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.PrdtId, &product.PrdtName, &product.Specific, &product.Inventor, &product.Unit, &product.BoxNumb, &product.Inline, &product.IvType, &product.Remark); err != nil {
			return
		}
		products = append(products, product)
	}

	return
}

func GetAllProduct() (products []Product, err error) {
	rows, err := Db.Query("SELECT * FROM products")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.PrdtId, &product.PrdtName, &product.Specific, &product.Inventor, &product.Unit, &product.BoxNumb, &product.Inline, &product.IvType, &product.Remark); err != nil {

			return
		}
		products = append(products, product)
	}

	return
}

func (r *Product) Update() (err error) {
	_, err = Db.Exec("UPDATE products set prdtid=?,prdtname=?,specific=?,inventor=?,unit=?,boxnumb=?,inline=?,ivtype=?,remark=? where prdtid=?", r.PrdtId, r.PrdtName, r.Specific, r.Inventor, r.Unit, r.BoxNumb, r.Inline, r.IvType, r.Remark, r.PrdtId)
	return
}

func (r *Product) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM products WHERE prdtid=?", r.PrdtId)
	return
}

func GetProductNS() (products []Product, err error) {

	rows, err := Db.Query("SELECT prdtname, specific FROM products")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.PrdtName, &product.Specific); err != nil {
			return
		}
		products = append(products, product)
	}

	return
}

func (r *Customer) Insert() (err error) {
	statement := "insert into customers (cstmname, cstmtype, gaddr, gname, gphone, ivaddr, ivname, ivphone, remark) values (?,?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CstmName, r.CstmType, r.Gaddr, r.Gname, r.Gphone, r.IvAddr, r.IvName, r.IvPhone, r.Remark)
	return
}

func (r *Customer) Select() (customers []Customer, err error) {
	rows, err := Db.Query("SELECT * FROM customers WHERE cstmname LIKE ? AND cstmtype LIKE ? AND gname LIKE ? AND ivname LIKE ? AND remark LIKE ?", "%"+r.CstmName+"%", "%"+r.CstmType+"%", "%"+r.Gname+"%", "%"+r.IvName+"%", "%"+r.Remark+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		customer := Customer{}
		if err = rows.Scan(&customer.CstmId, &customer.CstmName, &customer.CstmType, &customer.Gaddr, &customer.Gname, &customer.Gphone, &customer.IvAddr, &customer.IvName, &customer.IvPhone, &customer.Remark); err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return
}

func GetCustomerName() (customers []Customer, err error) {

	rows, err := Db.Query("SELECT cstmname FROM customers")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		customer := Customer{}
		if err = rows.Scan(&customer.CstmName); err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return
}

func GetAllCustomer() (customers []Customer, err error) {

	rows, err := Db.Query("SELECT * FROM customers")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		customer := Customer{}
		if err = rows.Scan(&customer.CstmId, &customer.CstmName, &customer.CstmType, &customer.Gaddr, &customer.Gname, &customer.Gphone, &customer.IvAddr, &customer.IvName, &customer.IvPhone, &customer.Remark); err != nil {
			return
		}
		customers = append(customers, customer)
	}

	return
}

func (r *Customer) Update() (err error) {
	_, err = Db.Exec("UPDATE customers set cstmname=?,cstmtype=?,gaddr=?,gname=?,gphone=?,ivaddr=?,ivname=?,ivphone=?,remark=? where cstmid=?", r.CstmName, r.CstmType, r.Gaddr, r.Gname, r.Gphone, r.IvAddr, r.IvName, r.IvPhone, r.Remark, r.CstmId)
	return
}

func (r *Customer) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM customers WHERE cstmid=?", r.CstmId)
	return
}

func (r *Contract0) Insert() (err error) {
	statement := "insert into contracts0(ccid, cstmid, vector, remark) values (?,(select cstmid from customers where customers.cstmname like ? limit 1),?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.CcId, "%"+r.CstmName+"%", r.Vector, r.Remark)
	return
}

func (r *Contract0) Select() (contracts0 []Contract0, err error) {
	rows, err := Db.Query("SELECT *,(select cstmname from customers where customers.cstmid=contracts0.cstmid limit 1) as cstmname FROM contracts0 WHERE ccid LIKE ?  AND cstmid in (select cstmid from customers where customers.cstmname like ?) AND  remark LIKE ? and vector <> ?", "%"+r.CcId+"%", "%"+r.CstmName+"%", "%"+r.Remark+"%", r.Vector)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		contract0 := Contract0{}
		if err = rows.Scan(&contract0.CcId, &contract0.CstmId, &contract0.Vector, &contract0.Remark, &contract0.CstmName); err != nil {
			return
		}
		contracts0 = append(contracts0, contract0)
	}

	return
}

func GetAllContract0() (contracts0 []Contract0, err error) {
	rows, err := Db.Query("SELECT *,(select cstmname from customers where customers.cstmid=contracts0.cstmid limit 1) as cstmname FROM contracts0")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		contract0 := Contract0{}
		if err = rows.Scan(&contract0.CcId, &contract0.CstmId, &contract0.Vector, &contract0.Remark, &contract0.CstmName); err != nil {
			return
		}
		contracts0 = append(contracts0, contract0)
	}

	return
}

func (r *Contract0) Update() (err error) {
	_, err = Db.Exec("UPDATE contracts0 set cstmid=(select cstmid from customers where customers.cstmname like ? limit 1),vector=?,remark=? where ccid=?", "%"+r.CstmName+"%", r.Vector, r.Remark, r.CcId)
	return
}

func (r *Contract0) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM contracts0 WHERE ccid=?", r.CcId)
	return
}

func (r *Stock) Select() (stocks []Stock, err error) {
	rows, err := Db.Query("select id, prdtid, cstmid, pnumber, quantity, remark, (select prdtname from products where products.prdtid = stocks.prdtid limit 1) as prdtname, (select specific from products where products.prdtid = stocks.prdtid limit 1) as specific, (select cstmname from customers where customers.cstmid=stocks.cstmid limit 1) as cstmname from stocks where cstmid in (select cstmid from customers where customers.cstmname like ?) and prdtid in (select prdtid from products where products.prdtname like ? and products.specific like ?) and remark like ? ", "%"+r.CstmName+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.Remark+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		stock := Stock{}
		if err = rows.Scan(&stock.Id, &stock.PrdtId, &stock.CstmId, &stock.Pnumber, &stock.Quantity, &stock.Remark, &stock.PrdtName, &stock.Specific, &stock.CstmName); err != nil {
			return
		}
		stocks = append(stocks, stock)
	}

	return
}

func GetAllStocks() (stocks []Stock, err error) {
	rows, err := Db.Query("select id, prdtid, cstmid, pnumber, quantity, remark, (select prdtname from products where products.prdtid = stocks.prdtid limit 1) as prdtname, (select specific from products where products.prdtid = stocks.prdtid limit 1) as specific, (select cstmname from customers where customers.cstmid=stocks.cstmid limit 1) as cstmname from stocks")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		stock := Stock{}
		if err = rows.Scan(&stock.Id, &stock.PrdtId, &stock.CstmId, &stock.Pnumber, &stock.Quantity, &stock.Remark, &stock.PrdtName, &stock.Specific, &stock.CstmName); err != nil {
			return
		}
		stocks = append(stocks, stock)
	}

	return
}

func (r *Stock) Insert() (err error) {
	statement := "insert into stocks(prdtid, cstmid, pnumber, quantity,remark) values ((select prdtid from products where products.prdtname like ? and products.specific like ? limit 1),(select cstmid from customers where customers.cstmname like ? limit 1),?,?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec("%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.CstmName+"%", r.Pnumber, r.Quantity, r.Remark)
	return
}

func (r *Stock) Update() (err error) {
	_, err = Db.Exec("UPDATE stocks set prdtid=(select prdtid from products where products.prdtname like ? and products.specific like ? limit 1),cstmid=(select cstmid from customers where customers.cstmname like ? limit 1),pnumber=?, quantity=?, remark=? where id=?", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.CstmName+"%", r.Pnumber, r.Quantity, r.Remark, r.Id)
	return
}

func (r *Stock) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM stocks WHERE id=?", r.Id)
	return
}

func (r *Debt) Insert() (err error) {
	statement := "insert into debts(srcid, cstmid, dbtsum, remark) values((select cstmid from customers where customers.cstmname like ? limit 1),(select cstmid from customers where customers.cstmname like ?),?,?)"
	stmt, err := Db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	_, err = stmt.Exec(r.SrcName, r.CstmName, r.DbtSum, r.Remark)
	return
}

func (r *Debt) Select() (debts []Debt, err error) {
	rows, err := Db.Query("select id, srcid, cstmid, dbtsum, remark, (select cstmname from customers where customers.cstmid=debts.srcid limit 1) as srcname, (select cstmname from customers where customers.cstmid=debts.cstmid limit 1) as cstmname from debts WHERE srcid in (select cstmid from customers where customers.cstmname like ?) and cstmid in (select cstmid from customers where customers.cstmname like ?) and remark like ?", "%"+r.SrcName+"%", "%"+r.CstmName+"%", "%"+r.Remark+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		debt := Debt{}
		if err = rows.Scan(&debt.Id, &debt.SrcId, &debt.CstmId, &debt.DbtSum, &debt.Remark, &debt.SrcName, &debt.CstmName); err != nil {
			return
		}
		debts = append(debts, debt)
	}

	return
}

func GetAllDebts() (debts []Debt, err error) {
	rows, err := Db.Query("select id, srcid, cstmid, dbtsum, remark, (select cstmname from customers where customers.cstmid=debts.srcid limit 1) as srcname, (select cstmname from customers where customers.cstmid=debts.cstmid limit 1) as cstmname from debts")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		debt := Debt{}
		if err = rows.Scan(&debt.Id, &debt.SrcId, &debt.CstmId, &debt.DbtSum, &debt.Remark, &debt.SrcName, &debt.CstmName); err != nil {
			return
		}
		debts = append(debts, debt)
	}

	return
}

func (r *Debt) Update() (err error) {
	_, err = Db.Exec("UPDATE debts set srcid=(select cstmid from customers where customers.cstmname like ? limit 1),cstmid=(select cstmid from customers where customers.cstmname like ? limit 1),dbtsum=?, remark=? where id=?", "%"+r.SrcName+"%", "%"+r.CstmName+"%", r.DbtSum, r.Remark, r.Id)
	return
}

func (r *Debt) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM debts WHERE id=?", r.Id)
	return
}

func (r *OnwayProduct) Select() (onwayProducts []OnwayProduct, err error) {
	rows, err := Db.Query("select id, ccid, prdtid, quantity, isum,remark, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid = inway_products.ccid limit 1) limit 1) as srcname, (select cstmname from customers where customers.cstmid = inwayproducts.ccid limit 1) as cstmname,(select prdtname from products where products.prdtid=inwayproducts.prdtid) as prdtname,(select specific from products where products.prdtid=inwayproducts.prdtid) as specific FROM inwayproducts where (quantity-isum) <> 0 and ccid like ? and prdtid in (select prdtid from products where products.prdtname like ? and products.specific like ?) and ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?) and contracts0.Vector=?) and ccid in (select ccid from contracts where contracts.cstmid in (select cstmid from customers where customers.cstmname like ?))", "%"+r.CcId+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.SrcName+"%", r.Vector, "%"+r.CstmName+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		onwayProduct := OnwayProduct{}
		if err = rows.Scan(&onwayProduct.Id, &onwayProduct.CcId, &onwayProduct.PrdtId, &onwayProduct.Quantity, &onwayProduct.CkSum, &onwayProduct.Remark, &onwayProduct.SrcName, &onwayProduct.CstmName, &onwayProduct.PrdtName, &onwayProduct.Specific); err != nil {
			return
		}
		onwayProducts = append(onwayProducts, onwayProduct)
	}

	return
}

func GetAllOnwayProducts(vector int) (onwayProducts []OnwayProduct, err error) {
	rows, err := Db.Query("select id,ccid, prdtid, quantity, isum,remark, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid = inway_products.ccid limit 1) limit 1) as srcname, (select cstmname from customers where customers.cstmid = inwayproducts.ccid limit 1) as cstmname,(select prdtname from products where products.prdtid=inwayproducts.prdtid) as prdtname,(select specific from products where products.prdtid=inwayproducts.prdtid) as specific FROM inwayproducts where (quantity-isum) <> 0 and ccid in (select ccid from contracts0 where contracts0.Vector=?)", vector)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		onwayProduct := OnwayProduct{}
		if err = rows.Scan(&onwayProduct.Id, &onwayProduct.CcId, &onwayProduct.PrdtId, &onwayProduct.Quantity, &onwayProduct.CkSum, &onwayProduct.Remark, &onwayProduct.SrcName, &onwayProduct.CstmName, &onwayProduct.PrdtName, &onwayProduct.Specific); err != nil {
			return
		}
		onwayProducts = append(onwayProducts, onwayProduct)
	}

	return
}

func (r *OnwayInvoice) Select() (onwayInvoices []OnwayInvoice, err error) {
	rows, err := Db.Query("select id, ccid, csum,osum,remark, (select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid = inway_invoices.ccid limit 1) limit 1) as srcname, (select cstmname from customers where customers.cstmid = inway_invoices.ccid limit 1) as cstmname FROM inway_invoices where (csum-osum) <> 0 and ccid like ? and  ccid in (select ccid from contracts0 where contracts0.cstmid in (select cstmid from customers where customers.cstmname like ?) and contracts0.vector=?) and ccid in (select ccid from contracts where contracts.cstmid in (select cstmid from customers where customers.cstmname like ?))", "%"+r.CcId+"%", "%"+r.SrcName+"%", r.Vector, "%"+r.CstmName+"%")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		onwayInvoice := OnwayInvoice{}
		if err = rows.Scan(&onwayInvoice.Id, &onwayInvoice.CcId, &onwayInvoice.Csum, &onwayInvoice.CkSum, &onwayInvoice.Remark, &onwayInvoice.SrcName, &onwayInvoice.CstmName); err != nil {
			return
		}
		onwayInvoices = append(onwayInvoices, onwayInvoice)
	}

	return
}

func GetAllOnwayInvoices(vector int) (onwayInvoices []OnwayInvoice, err error) {
	rows, err := Db.Query("select id,ccid, csum,osum, remark,(select cstmname from customers where customers.cstmid = (select cstmid from contracts0 where contracts0.ccid = inway_invoices.ccid limit 1) limit 1) as srcname, (select cstmname from customers where customers.cstmid = inway_invoices.ccid limit 1) as cstmname FROM inway_invoices where (csum-osum) <> 0 and  ccid in (select ccid from contracts0 where contracts0.vector=?)", vector)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		onwayInvoice := OnwayInvoice{}
		if err = rows.Scan(&onwayInvoice.Id, &onwayInvoice.CcId, &onwayInvoice.Csum, &onwayInvoice.CkSum, &onwayInvoice.Remark, &onwayInvoice.SrcName, &onwayInvoice.CstmName); err != nil {
			return
		}
		onwayInvoices = append(onwayInvoices, onwayInvoice)
	}

	return
}
