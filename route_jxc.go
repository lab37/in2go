package main

import (
	"001/data"
	"fmt"
	"net/http"
	"strconv"
)

func SelectItem(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		topic := request.PostFormValue("topic")
		tt := request.PostFormValue("tt")
		switch topic {
		case "contract":
			if tt == "all" {
				resp, err := data.GetAllContract()
				if err != nil {
					danger(err, "Cannot get all contract")
				}
				fmt.Println(resp)
			} else {
				contract := data.Contract{}
				contract.CcId = request.PostFormValue("ccid")
				contract.CcData = request.PostFormValue("ccdata")
				contract.CcType = request.PostFormValue("cctype")
				contract.CstmName = request.PostFormValue("cstmname")
				contract.PrdtId = request.PostFormValue("prdtid")
				contract.Remark = request.PostFormValue("remark")
				resp, err := contract.Select()
				if err != nil {
					danger(err, "Cannot get some contract")
				}
				fmt.Println(resp)
			}
		case "income":
			if tt == "all" {
				resp, err := data.GetAllIncome()
				if err != nil {
					danger(err, "Cannot get all income")
				}
				fmt.Println(resp)
			} else {
				income := data.Income{}
				income.CcId = request.PostFormValue("ccid")
				income.IcData = request.PostFormValue("icdata")
				income.PrdtId = request.PostFormValue("prdtid")
				income.Remark = request.PostFormValue("remark")
				resp, err := income.Select()
				if err != nil {
					danger(err, "Cannot get some income")
				}
				fmt.Println(resp)
			}
		case "outgo":
			if tt == "all" {
				resp, err := data.GetAllOutgo()
				if err != nil {
					danger(err, "Cannot get all outgo")
				}
				fmt.Println(resp)
			} else {
				outgo := data.Outgo{}
				outgo.CcId = request.PostFormValue("ccid")
				outgo.OgData = request.PostFormValue("ogdata")
				outgo.PrdtId = request.PostFormValue("prdtid")
				outgo.Pnumber = request.PostFormValue("pnumber")
				outgo.ExpName = request.PostFormValue("expname")
				outgo.ExpNumber = request.PostFormValue("expnumber")
				outgo.Remark = request.PostFormValue("remark")
				resp, err := outgo.Select()
				if err != nil {
					danger(err, "Cantnot get some outgo")
				}
				fmt.Println(resp)
			}
		case "invoice":
			if tt == "all" {
				resp, err := data.GetAllInvoice()
				if err != nil {
					danger(err, "Cannot get all invoices")
				}
				fmt.Println(resp)
			} else {
				invoice := data.Invoice{}
				invoice.IvId = request.PostFormValue("ivid")
				invoice.IvData = request.PostFormValue("ivdata")
				invoice.CcId = request.PostFormValue("ccid")
				invoice.IvSum = request.PostFormValue("ivsum")
				invoice.PostData = request.PostFormValue("postdata")
				invoice.ExpName = request.PostFormValue("expname")
				invoice.ExpNumber = request.PostFormValue("expnumber")
				invoice.Remark = request.PostFormValue("remark")
				resp, err := invoice.Select()
				if err != nil {
					danger(err, "Cannot get some invoices")
				}
				fmt.Println(resp)
			}
		case "payment":
			if tt == "all" {
				resp, err := data.GetAllPayment()
				if err != nil {
					danger(err, "Cannot get all payments")
				}
				fmt.Println(resp)
			} else {
				payment := data.Payment{}
				payment.CcId = request.PostFormValue("ccid")
				payment.PmData = request.PostFormValue("pmdata")
				payment.Remark = request.PostFormValue("remark")
				resp, err := payment.Select()
				if err != nil {
					danger(err, "Cannot get some payments")
				}
				fmt.Println(resp)
			}
		case "products":
			if tt == "all" {
				resp, err := data.GetAllProduct()
				if err != nil {
					danger(err, "Cannot get all products")
				}
				fmt.Println(resp)
			} else {
				product := data.Product{}
				product.PrdtId = request.PostFormValue("prdtid")
				product.PrdtName = request.PostFormValue("prdtname")
				product.Specific = request.PostFormValue("specific")
				product.Inventor = request.PostFormValue("inventor")
				product.Inline = request.PostFormValue("inline")
				product.IvType = request.PostFormValue("ivtype")
				product.Remark = request.PostFormValue("remark")
				resp, err := product.Select()
				if err != nil {
					danger(err, "Cannot get some products")
				}
				fmt.Println(resp)
			}
		case "customers":
			if tt == "all" {
				resp, err := data.GetAllCustomer()
				if err != nil {
					danger(err, "Cannot get all customers")
				}
				fmt.Println(resp)
			} else {
				customer := data.Customer{}
				customer.CstmName = request.PostFormValue("cstmname")
				customer.CstmType = request.PostFormValue("cstmtype")
				resp, err := customer.Select()
				if err != nil {
					danger(err, "Cannt get some customers")
				}
				fmt.Println(resp)
			}
		default:
		}
		http.Redirect(writer, request, "/", 302)
	}

}

func updateItem(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}

		topic := request.PostFormValue("topic")
		switch topic {
		case "contract":
			contract := data.Contract{}
			contract.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			contract.CcId = request.PostFormValue("ccid")
			contract.CcData = request.PostFormValue("ccdata")
			contract.CcType = request.PostFormValue("cctype")
			contract.CstmName = request.PostFormValue("cstmname")
			contract.PrdtId = request.PostFormValue("prdtid")
			contract.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
			contract.Quantity, _ = strconv.Atoi(request.PostFormValue("quantity"))
			contract.Remark = request.PostFormValue("remark")
			err := contract.Update()
			if err != nil {
				danger(err, "Cannot update contract")
			}
		case "income":
			income := data.Income{}
			income.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			income.CcId = request.PostFormValue("ccid")
			income.IcData = request.PostFormValue("icdata")
			income.PrdtId = request.PostFormValue("prdtid")
			income.Quantity, _ = strconv.Atoi(request.PostFormValue("quantity"))
			income.Pnumber = request.PostFormValue("pnumber")
			income.Remark = request.PostFormValue("remark")
			err := income.Update()
			if err != nil {
				danger(err, "Cannot update contract")
			}
		case "outgo":
			outgo := data.Outgo{}
			outgo.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			outgo.CcId = request.PostFormValue("ccid")
			outgo.OgData = request.PostFormValue("ogdata")
			outgo.PrdtId = request.PostFormValue("prdtid")
			outgo.Pnumber = request.PostFormValue("pnumber")
			outgo.Quantity, _ = strconv.Atoi(request.PostFormValue("quantity"))
			outgo.ExpName = request.PostFormValue("expname")
			outgo.ExpNumber = request.PostFormValue("expnumber")
			outgo.Remark = request.PostFormValue("remark")
			err := outgo.Update()
			if err != nil {
				danger(err, "Cannot update outgo")
			}
		case "invoice":
			invoice := data.Invoice{}
			invoice.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			invoice.IvId = request.PostFormValue("ivid")
			invoice.IvData = request.PostFormValue("ivdata")
			invoice.CcId = request.PostFormValue("ccid")
			invoice.IvSum = request.PostFormValue("ivsum")
			invoice.PostData = request.PostFormValue("postdata")
			invoice.ExpName = request.PostFormValue("expname")
			invoice.ExpNumber = request.PostFormValue("expnumber")
			invoice.Remark = request.PostFormValue("remark")
			err := invoice.Update()
			if err != nil {
				danger(err, "Cannot update invoice")
			}
		case "payment":
			payment := data.Payment{}
			payment.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			payment.CcId = request.PostFormValue("ccid")
			payment.PmData = request.PostFormValue("pmdata")
			payment.PmSum, _ = strconv.ParseFloat(request.PostFormValue("pmsum"), 64)
			payment.Remark = request.PostFormValue("remark")
			err := payment.Update()
			if err != nil {
				danger(err, "Cannot update payment")
			}
		case "product":
			product := data.Product{}
			product.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			product.PrdtId = request.PostFormValue("prdtid")
			product.PrdtName = request.PostFormValue("prdtname")
			product.Specific = request.PostFormValue("specific")
			product.Inventor = request.PostFormValue("inventor")
			product.Unit = request.PostFormValue("unit")
			product.BoxNumb, _ = strconv.Atoi(request.PostFormValue("boxnumb"))
			product.Inline = request.PostFormValue("inline")
			product.IvType = request.PostFormValue("ivtype")
			product.Remark = request.PostFormValue("remark")
			err := product.Update()
			if err != nil {
				danger(err, "Cant update products")
			}
		case "customer":
			customer := data.Customer{}
			customer.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			customer.CstmId = request.PostFormValue("cstmid")
			customer.CstmName = request.PostFormValue("cstmname")
			customer.Gaddr = request.PostFormValue("gaddr")
			customer.Gname = request.PostFormValue("gname")
			customer.Gphone = request.PostFormValue("gphone")
			customer.IvAddr = request.PostFormValue("ivaddr")
			customer.IvName = request.PostFormValue("ivname")
			customer.IvPhone = request.PostFormValue("ivphone")
			customer.Remark = request.PostFormValue("remark")
			err := customer.Update()
			if err != nil {
				danger(err, "Cant update customer")
			}
		default:
		}
	}

}

func insertItem(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}

		topic := request.PostFormValue("topic")
		fmt.Println(topic)

		switch topic {
		case "contract":
			contract := data.Contract{}
			product := data.Product{}
			product.PrdtName = request.PostFormValue("prdtname")
			product.Specific = request.PostFormValue("specific")
			contract.PrdtId, err = product.GetPrdtId()
			if err != nil {
				danger(err, "Cant get prdtid")
			}
			contract.CcId = request.PostFormValue("ccid")
			contract.CcData = request.PostFormValue("ccdata")
			contract.CcType = request.PostFormValue("cctype")
			contract.CstmName = request.PostFormValue("cstmname")

			contract.Price, _ = strconv.ParseFloat(request.PostFormValue("price"), 64)
			contract.Quantity, _ = strconv.Atoi(request.PostFormValue("quantity"))
			contract.Remark = request.PostFormValue("remark")
			fmt.Println(contract)
			err := contract.Insert()
			if err != nil {
				danger(err, "Cannot insert contract")
			}
		case "income":
			income := data.Income{}

			income.CcId = request.PostFormValue("ccid")
			income.IcData = request.PostFormValue("icdata")
			income.PrdtId = request.PostFormValue("prdtid")
			income.Quantity, _ = strconv.Atoi(request.PostFormValue("quantity"))
			income.Pnumber = request.PostFormValue("pnumber")
			income.Remark = request.PostFormValue("remark")
			err := income.Insert()
			if err != nil {
				danger(err, "Cannot insert contract")
			}
		case "outgo":
			outgo := data.Outgo{}

			outgo.CcId = request.PostFormValue("ccid")
			outgo.OgData = request.PostFormValue("ogdata")
			outgo.PrdtId = request.PostFormValue("prdtid")
			outgo.Pnumber = request.PostFormValue("pnumber")
			outgo.Quantity, _ = strconv.Atoi(request.PostFormValue("quantity"))
			outgo.ExpName = request.PostFormValue("expname")
			outgo.ExpNumber = request.PostFormValue("expnumber")
			outgo.Remark = request.PostFormValue("remark")
			err := outgo.Insert()
			if err != nil {
				danger(err, "Cannot insert outgo")
			}
		case "invoice":
			invoice := data.Invoice{}

			invoice.IvId = request.PostFormValue("ivid")
			invoice.IvData = request.PostFormValue("ivdata")
			invoice.CcId = request.PostFormValue("ccid")
			invoice.IvSum = request.PostFormValue("ivsum")
			invoice.PostData = request.PostFormValue("postdata")
			invoice.ExpName = request.PostFormValue("expname")
			invoice.ExpNumber = request.PostFormValue("expnumber")
			invoice.Remark = request.PostFormValue("remark")
			err := invoice.Insert()
			if err != nil {
				danger(err, "Cannot insert invoice")
			}
		case "payment":
			payment := data.Payment{}

			payment.CcId = request.PostFormValue("ccid")
			payment.PmData = request.PostFormValue("pmdata")
			payment.PmSum, _ = strconv.ParseFloat(request.PostFormValue("pmsum"), 64)
			payment.Remark = request.PostFormValue("remark")
			err := payment.Insert()
			if err != nil {
				danger(err, "Cannot insert payment")
			}
		case "product":
			product := data.Product{}

			product.PrdtId = request.PostFormValue("prdtid")
			product.PrdtName = request.PostFormValue("prdtname")
			product.Specific = request.PostFormValue("specific")
			product.Inventor = request.PostFormValue("inventor")
			product.Unit = request.PostFormValue("unit")
			product.BoxNumb, _ = strconv.Atoi(request.PostFormValue("boxnumb"))
			product.Inline = request.PostFormValue("inline")
			product.IvType = request.PostFormValue("ivtype")
			product.Remark = request.PostFormValue("remark")
			err := product.Insert()
			if err != nil {
				danger(err, "Cant insert products")
			}
		case "customer":
			customer := data.Customer{}

			customer.CstmId = request.PostFormValue("cstmid")
			customer.CstmName = request.PostFormValue("cstmname")
			customer.Gaddr = request.PostFormValue("gaddr")
			customer.Gname = request.PostFormValue("gname")
			customer.Gphone = request.PostFormValue("gphone")
			customer.IvAddr = request.PostFormValue("ivaddr")
			customer.IvName = request.PostFormValue("ivname")
			customer.IvPhone = request.PostFormValue("ivphone")
			customer.Remark = request.PostFormValue("remark")
			err := customer.Insert()
			if err != nil {
				danger(err, "Cant insert customer")
			}
		default:
		}
	}

}

func DeleteItem(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}

		topic := request.PostFormValue("topic")
		switch topic {
		case "contract":
			contract := data.Contract{}
			contract.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := contract.Delete()
			if err != nil {
				danger(err, "Cannot delete contract")
			}
		case "income":
			income := data.Income{}
			income.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := income.Delete()
			if err != nil {
				danger(err, "Cannot delete contract")
			}
		case "outgo":
			outgo := data.Outgo{}
			outgo.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := outgo.Delete()
			if err != nil {
				danger(err, "Cannot delete outgo")
			}
		case "invoice":
			invoice := data.Invoice{}
			invoice.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := invoice.Delete()
			if err != nil {
				danger(err, "Cannot delete invoice")
			}
		case "payment":
			payment := data.Payment{}
			payment.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := payment.Delete()
			if err != nil {
				danger(err, "Cannot delete payment")
			}
		case "product":
			product := data.Product{}
			product.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := product.Delete()
			if err != nil {
				danger(err, "Cant delete products")
			}
		case "customer":
			customer := data.Customer{}
			customer.Id, _ = strconv.Atoi(request.PostFormValue("id"))
			err := customer.Delete()
			if err != nil {
				danger(err, "Cant delete customer")
			}
		default:
		}
	}

}
