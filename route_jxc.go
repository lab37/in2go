package main

import (
	"001/data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func selectItem(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {

		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		topic := strings.TrimFunc(request.PostFormValue("topic"), unicode.IsSpace)
		tt := strings.TrimFunc(request.PostFormValue("tt"), unicode.IsSpace)
		writer.Header().Set("Content-Type", "application/json;charset:utf-8;")
		switch topic {
		case "contract":
			if tt == "all" {
				rsts, err := data.GetAllContract()
				if err != nil {
					danger(err, "Cannot get all contract")
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				contract := data.Contract{}
				product := data.Product{}
				product.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				product.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)

				contract.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				contract.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				contract.Vector, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
				contract.CcDate = strings.TrimFunc(request.PostFormValue("ccdate"), unicode.IsSpace)
				contract.CcType = strings.TrimFunc(request.PostFormValue("cctype"), unicode.IsSpace)
				contract.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				contract.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := contract.Select()
				if err != nil {
					danger("When select the contract get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "income":
			if tt == "all" {
				rsts, err := data.GetAllIncome()
				if err != nil {
					danger("When select all income get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				income := data.Income{}
				income.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				income.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				income.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)

				income.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				income.IcDate = strings.TrimFunc(request.PostFormValue("icdate"), unicode.IsSpace)
				income.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
				income.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := income.Select()
				if err != nil {
					danger("When select incomes get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "outgo":
			if tt == "all" {
				rsts, err := data.GetAllOutgo()
				if err != nil {
					danger("When select all outgos get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				outgo := data.Outgo{}
				outgo.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				outgo.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				outgo.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				outgo.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)

				outgo.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				outgo.OgDate = strings.TrimFunc(request.PostFormValue("ogdate"), unicode.IsSpace)
				outgo.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
				outgo.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := outgo.Select()
				if err != nil {
					danger("When select some outgos get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "invoice":
			if tt == "all" {
				rsts, err := data.GetAllInvoice()
				if err != nil {
					danger("When select all invoices get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				invoice := data.Invoice{}
				invoice.IvId = strings.TrimFunc(request.PostFormValue("ivid"), unicode.IsSpace)
				invoice.IvDate = strings.TrimFunc(request.PostFormValue("ivdate"), unicode.IsSpace)
				invoice.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				invoice.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				invoice.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				invoice.PostDate = strings.TrimFunc(request.PostFormValue("PostDate"), unicode.IsSpace)
				invoice.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := invoice.Select()
				if err != nil {
					danger("When select some invoice get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "payment":
			if tt == "all" {
				rsts, err := data.GetAllPayment()
				if err != nil {
					danger("When select all payments get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				payment := data.Payment{}
				payment.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				payment.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				payment.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				payment.PmDate = strings.TrimFunc(request.PostFormValue("pmdate"), unicode.IsSpace)
				payment.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := payment.Select()
				if err != nil {
					danger("When select some payments get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "product":
			if tt == "all" {
				rsts, err := data.GetAllProduct()
				if err != nil {
					danger("When select all products get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				product := data.Product{}
				product.PrdtId = strings.TrimFunc(request.PostFormValue("prdtid"), unicode.IsSpace)
				product.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				product.Inventor = strings.TrimFunc(request.PostFormValue("inventor"), unicode.IsSpace)
				product.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := product.Select()
				if err != nil {
					danger("When select some products get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "customer":
			if tt == "all" {
				rsts, err := data.GetAllCustomer()
				if err != nil {
					danger("When select all customers get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				customer := data.Customer{}
				customer.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				customer.CstmType = strings.TrimFunc(request.PostFormValue("cstmtype"), unicode.IsSpace)
				customer.Gname = strings.TrimFunc(request.PostFormValue("gname"), unicode.IsSpace)
				customer.IvName = strings.TrimFunc(request.PostFormValue("ivname"), unicode.IsSpace)
				customer.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := customer.Select()
				if err != nil {
					danger("When select some customers get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "contract0":
			if tt == "all" {
				rsts, err := data.GetAllContract0()
				if err != nil {
					danger("When select all contracts0 get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				contract0 := data.Contract0{}
				contract0.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				contract0.Vector, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
				contract0.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				contract0.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := contract0.Select()
				if err != nil {
					danger("When select some contract0 get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "stock":
			if tt == "all" {
				rsts, err := data.GetAllStocks()
				if err != nil {
					danger("When select all stocks get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				stock := data.Stock{}
				stock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				stock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				stock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				stock.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
				stock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := stock.Select()
				if err != nil {
					danger("When select some stocks get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "debt":
			if tt == "all" {
				rsts, err := data.GetAllDebts()
				if err != nil {
					danger("When select all stocks get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				debt := data.Debt{}
				debt.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				debt.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				debt.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
				rsts, err := debt.Select()
				if err != nil {
					danger("When select some stocks get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "onway_product":
			if tt == "all" {
				vector, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
				rsts, err := data.GetAllOnwayProducts(vector)
				if err != nil {
					danger("When select all inway products get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				onwayProduct := data.OnwayProduct{}
				onwayProduct.Vector, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
				onwayProduct.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				onwayProduct.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				onwayProduct.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				onwayProduct.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
				onwayProduct.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
				rsts, err := onwayProduct.Select()
				if err != nil {
					danger("When select some inway products get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		case "onway_invoice":
			if tt == "all" {
				vector, _ := strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
				rsts, err := data.GetAllOnwayInvoices(vector)
				if err != nil {
					danger("When select all onway invoices get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			} else {
				onwayInvoice := data.OnwayInvoice{}
				onwayInvoice.Vector, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
				onwayInvoice.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
				onwayInvoice.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
				onwayInvoice.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
				rsts, err := onwayInvoice.Select()
				if err != nil {
					danger("When select some onway invoice get an error", err)
				}
				resp, _ := json.Marshal(rsts)
				fmt.Fprintf(writer, string(resp))
			}
		default:
		}
		//http.Redirect(writer, request, "/", 302)
	}

}

func updateItem(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger("Cannot parse form in updateitem", err)
		}

		topic := request.PostFormValue("topic")
		switch topic {
		case "contract":
			contract := data.Contract{}
			contract.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			contract.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			contract.CcDate = strings.TrimFunc(request.PostFormValue("ccdate"), unicode.IsSpace)
			contract.CcType = strings.TrimFunc(request.PostFormValue("cctype"), unicode.IsSpace)
			contract.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			contract.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			contract.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			contract.Price, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("price"), unicode.IsSpace), 64)
			contract.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			contract.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := contract.Update()
			if err != nil {
				danger("When update the contract get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "income":
			income := data.Income{}
			income.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			income.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			income.IcDate = strings.TrimFunc(request.PostFormValue("icdate"), unicode.IsSpace)
			income.PrdtId = strings.TrimFunc(request.PostFormValue("prdtid"), unicode.IsSpace)
			income.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			income.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
			income.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := income.Update()
			if err != nil {
				danger("When update the income get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "outgo":
			outgo := data.Outgo{}
			outgo.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			outgo.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			outgo.OgDate = strings.TrimFunc(request.PostFormValue("ogdate"), unicode.IsSpace)
			outgo.PrdtId = strings.TrimFunc(request.PostFormValue("prdtid"), unicode.IsSpace)
			outgo.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
			outgo.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			outgo.ExpName = strings.TrimFunc(request.PostFormValue("expname"), unicode.IsSpace)
			outgo.ExpNumber = strings.TrimFunc(request.PostFormValue("expnumber"), unicode.IsSpace)
			outgo.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := outgo.Update()
			if err != nil {
				danger("When update the outgo get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "invoice":
			invoice := data.Invoice{}
			invoice.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			invoice.IvId = strings.TrimFunc(request.PostFormValue("ivid"), unicode.IsSpace)
			invoice.IvDate = strings.TrimFunc(request.PostFormValue("ivdate"), unicode.IsSpace)
			invoice.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			invoice.IvSum, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("ivsum"), unicode.IsSpace), 64)
			invoice.PostDate = strings.TrimFunc(request.PostFormValue("PostDate"), unicode.IsSpace)
			invoice.ExpName = strings.TrimFunc(request.PostFormValue("expname"), unicode.IsSpace)
			invoice.ExpNumber = strings.TrimFunc(request.PostFormValue("expnumber"), unicode.IsSpace)
			invoice.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := invoice.Update()
			if err != nil {
				danger("When update the invoice get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "payment":
			payment := data.Payment{}
			payment.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			payment.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			payment.PmDate = strings.TrimFunc(request.PostFormValue("pmdate"), unicode.IsSpace)
			payment.PmSum, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("pmsum"), unicode.IsSpace), 64)
			payment.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := payment.Update()
			if err != nil {
				danger("When update the payment get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "product":
			product := data.Product{}
			product.PrdtId = strings.TrimFunc(request.PostFormValue("prdtid"), unicode.IsSpace)
			product.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			product.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			product.Inventor = strings.TrimFunc(request.PostFormValue("inventor"), unicode.IsSpace)
			product.Unit = strings.TrimFunc(request.PostFormValue("unit"), unicode.IsSpace)
			product.BoxNumb, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("boxnumb"), unicode.IsSpace))
			product.Inline = strings.TrimFunc(request.PostFormValue("inline"), unicode.IsSpace)
			product.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			product.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := product.Update()
			if err != nil {
				danger("When update the product get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "customer":
			customer := data.Customer{}
			customer.CstmId, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("cstmid"), unicode.IsSpace))
			customer.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			customer.CstmType = strings.TrimFunc(request.PostFormValue("cstmtype"), unicode.IsSpace)
			customer.Gaddr = strings.TrimFunc(request.PostFormValue("gaddr"), unicode.IsSpace)
			customer.Gname = strings.TrimFunc(request.PostFormValue("gname"), unicode.IsSpace)
			customer.Gphone = strings.TrimFunc(request.PostFormValue("gphone"), unicode.IsSpace)
			customer.IvAddr = strings.TrimFunc(request.PostFormValue("ivaddr"), unicode.IsSpace)
			customer.IvName = strings.TrimFunc(request.PostFormValue("ivname"), unicode.IsSpace)
			customer.IvPhone = strings.TrimFunc(request.PostFormValue("ivphone"), unicode.IsSpace)
			customer.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := customer.Update()
			if err != nil {
				danger("When update the cutomer get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "contract0":
			contract0 := data.Contract0{}
			contract0.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			contract0.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			contract0.Vector, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
			contract0.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := contract0.Update()
			if err != nil {
				danger("When update the contract0 get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "stock":
			stock := data.Stock{}
			stock.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			stock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			stock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			stock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			stock.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
			stock.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			stock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := stock.Update()
			if err != nil {
				danger("When update the stocks get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}
		case "debt":
			debt := data.Debt{}
			debt.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			debt.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
			debt.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			debt.DbtSum, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("dbtsum"), unicode.IsSpace), 64)
			debt.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := debt.Insert()
			if err != nil {
				danger("When update the debts get an error", err)
				fmt.Fprintf(writer, "更新失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "更新成功")
			}

		default:
		}
	}

}

func insertItem(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger("Cannot parse form in function insertItem", err)
		}

		topic := request.PostFormValue("topic")

		switch topic {
		case "contract":
			contract := data.Contract{}
			contract.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			contract.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)

			contract.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			contract.CcDate = strings.TrimFunc(request.PostFormValue("ccdate"), unicode.IsSpace)
			contract.CcType = strings.TrimFunc(request.PostFormValue("cctype"), unicode.IsSpace)
			contract.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)

			contract.Price, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("price"), unicode.IsSpace), 64)
			contract.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			contract.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := contract.Insert()
			if err != nil {
				danger("When insert the contract get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}

		case "income":
			income := data.Income{}
			income.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			income.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)

			income.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			income.IcDate = strings.TrimFunc(request.PostFormValue("icdate"), unicode.IsSpace)
			income.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			income.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
			income.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := income.Insert()
			if err != nil {
				danger("When insert the income get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "outgo":
			outgo := data.Outgo{}
			outgo.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			outgo.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)

			outgo.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			outgo.OgDate = strings.TrimFunc(request.PostFormValue("ogdate"), unicode.IsSpace)
			outgo.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
			outgo.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			outgo.ExpName = strings.TrimFunc(request.PostFormValue("expname"), unicode.IsSpace)
			outgo.ExpNumber = strings.TrimFunc(request.PostFormValue("expnumber"), unicode.IsSpace)
			outgo.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := outgo.Insert()
			if err != nil {
				danger("When insert the outgo get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "invoice":
			invoice := data.Invoice{}

			invoice.IvId = strings.TrimFunc(request.PostFormValue("ivid"), unicode.IsSpace)
			invoice.IvDate = strings.TrimFunc(request.PostFormValue("ivdate"), unicode.IsSpace)
			invoice.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			invoice.IvSum, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("ivsum"), unicode.IsSpace), 64)
			invoice.PostDate = strings.TrimFunc(request.PostFormValue("PostDate"), unicode.IsSpace)
			invoice.ExpName = strings.TrimFunc(request.PostFormValue("expname"), unicode.IsSpace)
			invoice.ExpNumber = strings.TrimFunc(request.PostFormValue("expnumber"), unicode.IsSpace)
			invoice.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := invoice.Insert()
			if err != nil {
				danger("When insert the invoice get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "payment":
			payment := data.Payment{}

			payment.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			payment.PmDate = strings.TrimFunc(request.PostFormValue("pmdate"), unicode.IsSpace)
			payment.PmSum, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("pmsum"), unicode.IsSpace), 64)
			payment.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := payment.Insert()
			if err != nil {
				danger("When insert the payment get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "product":
			product := data.Product{}

			product.PrdtId = strings.TrimFunc(request.PostFormValue("prdtid"), unicode.IsSpace)
			product.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			product.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			product.Inventor = strings.TrimFunc(request.PostFormValue("inventor"), unicode.IsSpace)
			product.Unit = strings.TrimFunc(request.PostFormValue("unit"), unicode.IsSpace)
			product.BoxNumb, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("boxnumb"), unicode.IsSpace))
			product.Inline = strings.TrimFunc(request.PostFormValue("inline"), unicode.IsSpace)
			product.IvType = strings.TrimFunc(request.PostFormValue("ivtype"), unicode.IsSpace)
			product.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := product.Insert()
			if err != nil {
				danger("When insert the product get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "customer":
			customer := data.Customer{}

			customer.CstmId, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("cstmid"), unicode.IsSpace))
			customer.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			customer.CstmType = strings.TrimFunc(request.PostFormValue("cstmtype"), unicode.IsSpace)
			customer.Gaddr = strings.TrimFunc(request.PostFormValue("gaddr"), unicode.IsSpace)
			customer.Gname = strings.TrimFunc(request.PostFormValue("gname"), unicode.IsSpace)
			customer.Gphone = strings.TrimFunc(request.PostFormValue("gphone"), unicode.IsSpace)
			customer.IvAddr = strings.TrimFunc(request.PostFormValue("ivaddr"), unicode.IsSpace)
			customer.IvName = strings.TrimFunc(request.PostFormValue("ivname"), unicode.IsSpace)
			customer.IvPhone = strings.TrimFunc(request.PostFormValue("ivphone"), unicode.IsSpace)
			customer.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := customer.Insert()
			if err != nil {
				danger("When insert the customer get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "contract0":
			contract0 := data.Contract0{}

			contract0.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			contract0.Vector, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("vector"), unicode.IsSpace))
			contract0.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			contract0.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := contract0.Insert()
			if err != nil {
				danger("When insert the contract0 get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "stock":
			stock := data.Stock{}

			stock.PrdtName = strings.TrimFunc(request.PostFormValue("prdtname"), unicode.IsSpace)
			stock.Specific = strings.TrimFunc(request.PostFormValue("specific"), unicode.IsSpace)
			stock.Quantity, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("quantity"), unicode.IsSpace))
			stock.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			stock.Pnumber = strings.TrimFunc(request.PostFormValue("pnumber"), unicode.IsSpace)
			stock.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := stock.Insert()
			if err != nil {
				danger("When insert the stocks get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		case "debt":
			debt := data.Debt{}

			debt.SrcName = strings.TrimFunc(request.PostFormValue("srcname"), unicode.IsSpace)
			debt.CstmName = strings.TrimFunc(request.PostFormValue("cstmname"), unicode.IsSpace)
			debt.DbtSum, _ = strconv.ParseFloat(strings.TrimFunc(request.PostFormValue("dbtsum"), unicode.IsSpace), 64)
			debt.Remark = strings.TrimFunc(request.PostFormValue("remark"), unicode.IsSpace)
			err := debt.Insert()
			if err != nil {
				danger("When insert the debts get an error", err)
				fmt.Fprintf(writer, "<p class='results'>添加失败！服务器错误</p>")
			} else {
				fmt.Fprintf(writer, "<p class='results'>添加成功！</p>")
			}
		default:
		}
	}

}

func deleteItem(writer http.ResponseWriter, request *http.Request) {
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
			contract.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := contract.Delete()
			if err != nil {
				danger("When delete the contract get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "income":
			income := data.Income{}
			income.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := income.Delete()
			if err != nil {
				danger("When delete the income get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "outgo":
			outgo := data.Outgo{}
			outgo.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := outgo.Delete()
			if err != nil {
				danger("When delete the outgo get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "invoice":
			invoice := data.Invoice{}
			invoice.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := invoice.Delete()
			if err != nil {
				danger("When delete the invoice get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "payment":
			payment := data.Payment{}
			payment.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := payment.Delete()
			if err != nil {
				danger("When delete the payment get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "product":
			product := data.Product{}
			product.PrdtId = strings.TrimFunc(request.PostFormValue("prdtid"), unicode.IsSpace)
			err := product.Delete()
			if err != nil {
				danger("When delete the product get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "customer":
			customer := data.Customer{}
			customer.CstmId, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("cstmid"), unicode.IsSpace))
			err := customer.Delete()
			if err != nil {
				danger("When delete the customer get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "contract0":
			contract0 := data.Contract0{}
			contract0.CcId = strings.TrimFunc(request.PostFormValue("ccid"), unicode.IsSpace)
			err := contract0.Delete()
			if err != nil {
				danger("When delete the contract0 get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "stock":
			stock := data.Stock{}
			stock.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := stock.Delete()
			if err != nil {
				danger("When delete the stocks get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		case "debt":
			debt := data.Debt{}
			debt.Id, _ = strconv.Atoi(strings.TrimFunc(request.PostFormValue("id"), unicode.IsSpace))
			err := debt.Delete()
			if err != nil {
				danger("When delete the debts get an error", err)
				fmt.Fprintf(writer, "删除失败！服务器错误")
			} else {
				fmt.Fprintf(writer, "删除成功！")
			}
		default:
		}
	}

}

func getCustomerName(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		rsts, err := data.GetCustomerName()
		if err != nil {
			danger(err, "Cannot get all customername")
		}
		resp, _ := json.Marshal(rsts)
		fmt.Fprintf(writer, string(resp))
	}
}

func getProductNS(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		rsts, err := data.GetProductNS()
		if err != nil {
			danger(err, "Cannot get all prdtns")
		}
		resp, _ := json.Marshal(rsts)
		fmt.Fprintf(writer, string(resp))
	}
}

func getAllCcId(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		rsts, err := data.GetAllCcId()
		if err != nil {
			danger(err, "Cannot get all contractccid")
		}
		resp, _ := json.Marshal(rsts)
		fmt.Fprintf(writer, string(resp))
	}
}
