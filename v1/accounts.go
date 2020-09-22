package v1

import (
	"github.com/labstack/echo"
	"github.com/nats-io/nsc/cmd"
	"log"
)

type Account struct {}


func AccountRouter(g *echo.Group) {
	account := Account{}
	g.POST("", account.Save)
}


func (a Account) Save(context echo.Context) error {
	_,_,err:= ExecuteCmd(cmd.CreateAddAccountCmd(),"-n","abc")
	if err!=nil{
		log.Println(err.Error())
		return GenerateErrorResponse(context,nil,err.Error())
	}
	return GenerateSuccessResponse(context, nil, "Operation successful!")
}
