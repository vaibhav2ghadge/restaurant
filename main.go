package main

import (
	"fmt"
	"os"
	"strings"
	dbrepo "./dbrepository"
	mongoutils "./utils"
	domain "./domain"

)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	//Run sample commands


//second assign
var id domain.ID;
id = "5c45653b79492b6216397592"
x,y:=repoaccess.Get(id)
fmt.Println("get",x,y)
xx,yy:=repoaccess.GetAll()
fmt.Println("getAll",xx,yy)
for _,obj:=range xx {
			fmt.Println(obj,yy)
	}
//third assignment
xc,mc := repoaccess.FindByName("pizza")
fmt.Println(xc,mc,id)
dbrepo.PrintRestaurant(xc)
cmdArgument := os.Args[1:]
if len(cmdArgument)>0 {
	if strings.Contains(cmdArgument[0],"find") {
			cmdArgument = strings.SplitAfter(cmdArgument[1],"=")
			if strings.Contains(cmdArgument[0],"--type_of_food") {
				fmt.Println("Types Of Food")
				rest,_ := repoaccess.FindByTypeOfFood(cmdArgument[1])
				dbrepo.PrintRestaurant(rest) //print the restarant
			} else if strings.Contains(cmdArgument[0],"--postcode") && len(cmdArgument)==2{
				fmt.Println("argument",cmdArgument,cmdArgument[1])
				rest,_ := repoaccess.FindByTypeOfPostCode(cmdArgument[1])
				dbrepo.PrintRestaurant(rest) //print the restarant
			} else {
				fmt.Println("invalid argument with --type_of_food")
			}//count number of restuarnt in give food type
	} else if strings.Contains(cmdArgument[0],"count") {
			cmdArgument = strings.SplitAfter(cmdArgument[1],"=")
			if strings.Contains(cmdArgument[0],"--type_of_food") {
				count,err := repoaccess.CountByTypeOfFood(cmdArgument[1])
				fmt.Println("Total Restarant : ",count,err)
	}
}
}}
