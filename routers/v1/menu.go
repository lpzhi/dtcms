package v1

import (
	"dtcms/const"
	"dtcms/models"
	"dtcms/routers/contract"
	"github.com/devfeel/dotweb"
)


func QueryMenuLists(ctx dotweb.Context) error {

	response := contract.NewResonseInfo()
	//query data
	menu := new(models.AuthMenu)
	menuLists := menu.MenuList()
	response.RetCode = _const.ApiRetCode_Ok
	response.RetMsg = _const.ApiRetMsg_Ok
	response.Message = menuLists

	return ctx.WriteJson(response)
}