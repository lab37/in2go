package data

type Contract struct {
	Id       int
	CcId     string
	CcData   string
	CcType   string
	CstmName string
	PrdtId   string
	Price    float64
	Quantity int
	Remark   string
}
type Income struct {
	Id       int
	CcId     string
	IcData   string
	PrdtId   string
	Quantity int
	Pnumber  string
	Remark   string
}
type Outgo struct {
	Id        int
	CcId      string
	OgData    string
	PrdtId    string
	Pnumber   string
	Quantity  int
	ExpName   string
	ExpNumber string
	Remark    string
}
type Invoice struct {
	Id        int
	IvId      string
	IvData    string
	CcId      string
	IvSum     string
	PostData  string
	ExpName   string
	ExpNumber string
	Remark    string
}

type Payment struct {
	Id     int
	CcId   string
	PmData string
	PmSum  float64
	Remark string
}

type Product struct {
	Id       int
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
	Id       int
	CstmId   string
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

func (r *Contract) Insert() (err error) {
	statement := "insert into contract (ccid,ccdata, cctype, cstmname, prdtid, price, quantity, remark) values (?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.CcId, r.CcData, r.CcType, r.CstmName, r.PrdtId, r.Price, r.Quantity, r.Remark)
	return
}

func (r *Contract) Select() (contracts []Contract, err error) {
	rows, err := Db.Query("SELECT * FROM contract WHERE ccid LIKE ? AND ccdata LIKE ? AND cctype LIKE ? AND cstmname LIKE ?", "%"+r.CcId+"%", "%"+r.CcData+"%", "%"+r.CcType+"%", "%"+r.CstmName+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		contract := Contract{}
		if err = rows.Scan(&contract.Id, &contract.CcId, &contract.CcData, &contract.CcType, &contract.CstmName, &contract.Price, &contract.Quantity, &contract.Remark); err != nil {
			return
		}
		contracts = append(contracts, contract)
	}
	rows.Close()
	return
}

func GetAllContract() (contracts []Contract, err error) {
	rows, err := Db.Query("SELECT * FROM contract")
	if err != nil {
		return
	}
	for rows.Next() {
		contract := Contract{}
		if err = rows.Scan(&contract.Id, &contract.CcId, &contract.CcData, &contract.CcType, &contract.CstmName, &contract.Price, &contract.Quantity, &contract.Remark); err != nil {
			return
		}
		contracts = append(contracts, contract)
	}
	rows.Close()
	return
}

func (r *Contract) Update() (err error) {
	_, err = Db.Exec("UPDATE contract set ccid=?,ccdata=?,cctype=?,cstmname=?,price=?,quantity=?,remark=? where id=?", r.CcId, r.CcData, r.CcType, r.CstmName, r.Price, r.Quantity, r.Remark, r.Id)
	return
}

func (r *Contract) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM contract WHERE id=?", r.Id)
	return
}

func (r *Income) Insert() (err error) {
	statement := "insert into income (ccid,icdata, prdtid, quantity, pnumber, remark) values (?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.CcId, r.IcData, r.PrdtId, r.Quantity, r.Pnumber, r.Remark)
	return
}

func (r *Income) Select() (incomes []Income, err error) {
	rows, err := Db.Query("SELECT * FROM income WHERE ccid LIKE ? AND icdata LIKE ? AND prdtid LIKE ?", "%"+r.CcId+"%", "%"+r.IcData+"%", "%"+r.PrdtId+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		income := Income{}
		if err = rows.Scan(&income.Id, &income.CcId, &income.IcData, &income.PrdtId, &income.Quantity, &income.Pnumber, &income.Remark); err != nil {
			return
		}
		incomes = append(incomes, income)
	}
	rows.Close()
	return
}

func GetAllIncome() (incomes []Income, err error) {
	rows, err := Db.Query("SELECT * FROM income")
	if err != nil {
		return
	}
	for rows.Next() {
		income := Income{}
		if err = rows.Scan(&income.Id, &income.CcId, &income.IcData, &income.PrdtId, &income.Quantity, &income.Pnumber, &income.Remark); err != nil {
			return
		}
		incomes = append(incomes, income)
	}
	rows.Close()
	return
}

func (r *Income) Update() (err error) {
	_, err = Db.Exec("UPDATE income set ccid=?,icdata=?,prdtid=?,quantity=?,pnumber=?,remark=? where id=?", r.CcId, r.IcData, r.PrdtId, r.Quantity, r.Pnumber, r.Remark, r.Id)
	return
}

func (r *Income) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM income WHERE id=?", r.Id)
	return
}

func (r *Outgo) Insert() (err error) {
	statement := "insert into outgo (ccid,ogdata, prdtid, pnumber, quantity, expname, expnumber, remark) values (?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.CcId, r.OgData, r.PrdtId, r.Pnumber, r.Quantity, r.ExpName, r.ExpName, r.Remark)
	return
}

func (r *Outgo) Select() (outgos []Outgo, err error) {
	rows, err := Db.Query("SELECT * FROM outgo WHERE ccid LIKE ? AND ogdata LIKE ? AND prdtid LIKE ?", "%"+r.CcId+"%", "%"+r.OgData+"%", "%"+r.PrdtId+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		outgo := Outgo{}
		if err = rows.Scan(&outgo.Id, &outgo.CcId, &outgo.OgData, &outgo.PrdtId, &outgo.Pnumber, &outgo.Quantity, &outgo.ExpName, &outgo.ExpNumber, &outgo.Remark); err != nil {
			return
		}
		outgos = append(outgos, outgo)
	}
	rows.Close()
	return
}

func GetAllOutgo() (outgos []Outgo, err error) {
	rows, err := Db.Query("SELECT * FROM outgo")
	if err != nil {
		return
	}
	for rows.Next() {
		outgo := Outgo{}
		if err = rows.Scan(&outgo.Id, &outgo.CcId, &outgo.OgData, &outgo.PrdtId, &outgo.Pnumber, &outgo.Quantity, &outgo.ExpName, &outgo.ExpNumber, &outgo.Remark); err != nil {
			return
		}
		outgos = append(outgos, outgo)
	}
	rows.Close()
	return
}

func (r *Outgo) Update() (err error) {
	_, err = Db.Exec("UPDATE outgo set ccid=?,ogdata=?,prdtid=?,pnumber=?,quantity=?,expname=?,expnumber=?,remark=? where id=?", r.CcId, r.OgData, r.PrdtId, r.Pnumber, r.Quantity, r.ExpName, r.ExpNumber, r.Remark, r.Id)
	return
}

func (r *Outgo) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM outgo WHERE id=?", r.Id)
	return
}

func (r *Invoice) Insert() (err error) {
	statement := "insert into invoice (ivid, ivdata, ccid, ivsum, postdata, expname, expnumber, remark) values (?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.IvId, r.IvData, r.CcId, r.IvSum, r.PostData, r.ExpName, r.ExpNumber, r.Remark)
	return
}

func (r *Invoice) Select() (invoices []Invoice, err error) {
	rows, err := Db.Query("SELECT * FROM invoice WHERE ivid LIKE ? AND ivdata LIKE ? AND ccid LIKE ?", "%"+r.CcId+"%", "%"+r.IvData+"%", "%"+r.CcId+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		invoice := Invoice{}
		if err = rows.Scan(&invoice.Id, &invoice.IvId, &invoice.IvData, &invoice.CcId, &invoice.IvSum, &invoice.PostData, &invoice.ExpName, &invoice.ExpNumber, &invoice.Remark); err != nil {
			return
		}
		invoices = append(invoices, invoice)
	}
	rows.Close()
	return
}

func GetAllInvoice() (invoices []Invoice, err error) {
	rows, err := Db.Query("SELECT * FROM invoice")
	if err != nil {
		return
	}
	for rows.Next() {
		invoice := Invoice{}
		if err = rows.Scan(&invoice.Id, &invoice.IvId, &invoice.IvData, &invoice.CcId, &invoice.IvSum, &invoice.PostData, &invoice.ExpName, &invoice.ExpNumber, &invoice.Remark); err != nil {
			return
		}
		invoices = append(invoices, invoice)
	}
	rows.Close()
	return
}

func (r *Invoice) Update() (err error) {
	_, err = Db.Exec("UPDATE invoice set ivid=?,ivdata=?,ccid=?,ivsum=?,postdata=?,expname=?,expnumber=?,remark=? where id=?", r.IvId, r.IvData, r.CcId, r.IvSum, r.PostData, r.ExpName, r.ExpNumber, r.Remark, r.Id)
	return
}

func (r *Invoice) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM invoice WHERE id=?", r.Id)
	return
}

func (r *Payment) Insert() (err error) {
	statement := "insert into payment (ccid, pmdata, pmsum, remark) values (?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.CcId, r.PmData, r.PmSum, r.Remark)
	return
}

func (r *Payment) Select() (payments []Payment, err error) {
	rows, err := Db.Query("SELECT * FROM payment WHERE ccid LIKE ? AND pmdata LIKE ?", "%"+r.CcId+"%", "%"+r.PmData+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		payment := Payment{}
		if err = rows.Scan(&payment.Id, &payment.CcId, &payment.PmData, &payment.PmSum, &payment.Remark); err != nil {
			return
		}
		payments = append(payments, payment)
	}
	rows.Close()
	return
}

func GetAllPayment() (payments []Payment, err error) {
	rows, err := Db.Query("SELECT * FROM payment")
	if err != nil {
		return
	}
	for rows.Next() {
		payment := Payment{}
		if err = rows.Scan(&payment.Id, &payment.CcId, &payment.PmData, &payment.PmSum, &payment.Remark); err != nil {
			return
		}
		payments = append(payments, payment)
	}
	rows.Close()
	return
}

func (r *Payment) Update() (err error) {
	_, err = Db.Exec("UPDATE payment set ccid=?,pmdata=?,pmsum=?,remark=? where id=?", r.CcId, r.PmData, r.PmSum, r.Remark, r.Id)
	return
}

func (r *Payment) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM payment WHERE id=?", r.Id)
	return
}

func (r *Product) Insert() (err error) {
	statement := "insert into product (prdtid, prdtname, specific, inventor, unit, boxnumb, inline, ivtype, remark) values (?,?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.PrdtId, r.PrdtName, r.Specific, r.Inventor, r.Unit, r.BoxNumb, r.Inline, r.IvType, r.Remark)
	return
}

func (r *Product) Select() (products []Product, err error) {
	rows, err := Db.Query("SELECT * FROM products WHERE prdtid LIKE ? AND prdtname LIKE ? AND specific LIKE ? AND inventor LIKE ? AND inline LIKE ? AND ivtype LIKE ?", "%"+r.PrdtId+"%", "%"+r.PrdtName+"%", "%"+r.Specific+"%", "%"+r.Inventor+"%", "%"+r.Inline+"%", "%"+r.IvType+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.Id, &product.PrdtId, &product.Specific, &product.Inventor, &product.Unit, &product.BoxNumb, &product.Inline, &product.IvType, &product.Remark); err != nil {
			return
		}
		products = append(products, product)
	}
	rows.Close()
	return
}

func GetAllProduct() (products []Product, err error) {
	rows, err := Db.Query("SELECT * FROM products")
	if err != nil {
		return
	}
	for rows.Next() {
		product := Product{}
		if err = rows.Scan(&product.Id, &product.PrdtId, &product.PrdtName, &product.Specific, &product.Inventor, &product.Unit, &product.BoxNumb, &product.Inline, &product.IvType, &product.Remark); err != nil {
			return
		}
		products = append(products, product)
	}
	rows.Close()
	return
}

func (r *Product) Update() (err error) {
	_, err = Db.Exec("UPDATE products set prdtid=?,prdtname=?,specific=?,inventor=?,unit=?,boxnumb=?,inline=?,ivtype=?,remark=? where id=?", r.PrdtId, r.PrdtName, r.Specific, r.Inventor, r.Unit, r.BoxNumb, r.Inline, r.IvType, r.Remark, r.Id)
	return
}

func (r *Product) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM products WHERE id=?", r.Id)
	return
}

func (r *Product) GetPrdtId() (prdtId string, err error) {
	err = Db.QueryRow("SELECT prdtid FROM products where prdtname LIKE ? AND specific LIKE ?", "%"+r.PrdtName+"%", "%"+r.Specific+"%").Scan(&prdtId)
	return
}

func (r *Customer) Insert() (err error) {
	statement := "insert into customers (cstmid, cstmname, cstmtype, gaddr, gname, gphone, ivaddr, ivname, ivphone, remark) values (?,?,?,?,?,?,?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.CstmId, r.CstmName, r.CstmType, r.Gaddr, r.Gname, r.Gphone, r.IvAddr, r.IvName, r.IvPhone, r.Remark)
	return
}

func (r *Customer) Select() (customers []Customer, err error) {
	rows, err := Db.Query("SELECT * FROM customers WHERE cstmname LIKE ? AND cstmtype LIKE ?", "%"+r.CstmName+"%", "%"+r.CstmType+"%")
	if err != nil {
		return
	}
	for rows.Next() {
		customer := Customer{}
		if err = rows.Scan(&customer.Id, &customer.CstmId, &customer.CstmName, &customer.CstmType, &customer.Gaddr, &customer.Gname, &customer.Gphone, &customer.IvAddr, &customer.IvName, &customer.IvPhone, &customer.Remark); err != nil {
			return
		}
		customers = append(customers, customer)
	}
	rows.Close()
	return
}

func GetAllCustomer() (customers []Customer, err error) {
	rows, err := Db.Query("SELECT * FROM customers")
	if err != nil {
		return
	}
	for rows.Next() {
		customer := Customer{}
		if err = rows.Scan(&customer.Id, &customer.CstmId, &customer.CstmName, &customer.CstmType, &customer.Gaddr, &customer.Gname, &customer.Gphone, &customer.IvAddr, &customer.IvName, &customer.IvPhone, &customer.Remark); err != nil {
			return
		}
		customers = append(customers, customer)
	}
	rows.Close()
	return
}

func (r *Customer) Update() (err error) {
	_, err = Db.Exec("UPDATE customers set cstmid=?,cstmname=?,cstmtype=?,gaddr=?,gname=?,gphone=?,ivaddr=?,ivname=?,ivphone=?,remark=? where id=?", r.CstmId, r.CstmName, r.CstmType, r.Gaddr, r.Gname, r.Gphone, r.IvAddr, r.IvName, r.IvPhone, r.Remark, r.Id)
	return
}

func (r *Customer) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM customers WHERE id=?", r.Id)
	return
}
