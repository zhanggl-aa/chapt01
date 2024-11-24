package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://datacenter-web.eastmoney.com/api/data/v1/get?callback=jQuery112303992917399653777_1704108937232&sortColumns=PUBLIC_START_DATE%2CSECURITY_CODE&sortTypes=-1%2C-1&pageSize=99999&pageNumber=1&reportName=RPT_BOND_CB_LIST&columns=ALL&quoteColumns=f2~01~CONVERT_STOCK_CODE~CONVERT_STOCK_PRICE%2Cf235~10~SECURITY_CODE~TRANSFER_PRICE%2Cf236~10~SECURITY_CODE~TRANSFER_VALUE%2Cf2~10~SECURITY_CODE~CURRENT_BOND_PRICE%2Cf237~10~SECURITY_CODE~TRANSFER_PREMIUM_RATIO%2Cf239~10~SECURITY_CODE~RESALE_TRIG_PRICE%2Cf240~10~SECURITY_CODE~REDEEM_TRIG_PRICE%2Cf23~01~CONVERT_STOCK_CODE~PBV_RATIO&quoteType=0&source=WEB&client=WEB"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(body))
	strBody := string(body)
	codeparts := strings.Split(strBody, "Y_CODE\":\"")
	// nameparts := strings.Split(strBody, "Y_NAME_ABBR\":\"")
	// stcodeparts:= strings.Split(strBody, "T_STOCK_CODE\":\"")
	// priceparts := strings.Split(strBody, "T_BOND_PRICE\":")
	// tprparts:=strings.Split(strBody, "R_PREMIUM_RATIO\":") //转股溢价
	// yearparts:=strings.Split(strBody, "RE_DATE\":\"")
	// curr_iss_amtparts:=strings.Split(strBody, "ACTUAL_ISSUE_SCALE\":")
	// RATINGparts:=strings.Split(strBody, "RATING\":\"")
	// count := len(nameparts)

	// fmt.Printf("codeparts=%s\tstcodeparts=%s\tpriceparts=%s\ttprparts=%syearparts=%scurr_iss_amtparts=%sRATINGparts=%scount=%s",
	// codeparts,&stcodeparts,&priceparts,&tprparts,&yearparts,&curr_iss_amtparts,&RATINGparts,&count)

	fmt.Println(codeparts)
}
