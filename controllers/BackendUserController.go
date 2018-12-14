package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/json"

	"myproject/enums"
	"myproject/models"
	"myproject/utils"

	"github.com/astaxie/beego/orm"

	)


type  BackendUserController struct {
	BaseController
}

func (c *BackendUserController) Prepare() {
	//先执行
	c.BaseController.Prepare()
    //如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()

}

func (c *BackendUserController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("BackendUserController","Edit")
	c.Data["canDelete"] = c.checkActionAuthor("BackendUserController","Delete")

}


func (c *BackendUserController) DataGrid(){
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.BackendUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody,&params)

	data,total := models.BackendUserPageList(&params)

	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *BackendUserController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}

	Id,_ := c.GetInt(":id",0)
	m := &models.BackendUser{}
	var err error
	if Id > 0 {
		m,err = models.BackendUserOne(Id)
		if err != nil {
			c.pageError("数据无效,请刷新后重试")
		}
		o := orm.NewOrm()
		o.LoadRelated(m,"RoleBackendUserRel")
	} else {
		m.Status = enums.Enabled
	}

	c.Data["m"] = m

	var roleIds []string
	for _,item := range m.RoleBackendUserRel {
		roleIds = append(roleIds,strconv.Itoa(item.Role.Id))
	}
	c.Data["roles"] = strings.Join(roleIds,",")
	c.setTpl("backenduser/edit.html","shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "backenduser/edit_footerjs.html"
}

func (c *BackendUserController) Save() {
	m := models.BackendUser{}
	o := orm.NewOrm()
	var err error
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed,"删除历史关系失败","")
	}
	if m.Id == 0 {
		m.UserPwd = utils.String2md5(m.UserPwd)
		if _,err := o.Insert(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed,"添加失败",m.Id)
		}
	}else {
		if oM,err := models.BackendUserOne(m.Id); err != nil {
			c.jsonResult(enums.JRCodeFailed,"数据无效,请刷新后重试",m.Id)
		}else {
			m.UserPwd = strings.TrimSpace(m.UserPwd)
			if len(m.UserPwd) == 0 {
				m.UserPwd = oM.UserPwd
			}else{
				m.UserPwd = utils.String2md5(m.UserPwd)
			}
			m.Avatar = oM.Avatar
		}
		if _,err := o.Update(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed,"编辑失败",m.Id)
		}
	}
	var relations []models.RoleBackendUserRel
	for _,roleId := range m.RoleIds {
		r := models.Role{Id:roleId}
		relation := models.RoleBackendUserRel{BackendUser:&m,Role:&r}
		relations = append(relations,relation)
	}
	if len(relations) > 0 {
		if _,err := o.InsertMulti(len(relations),relations); err == nil {
			c.jsonResult(enums.JRCodeSucc,"保持成功",m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed,"保持失败",m.Id)
		}
	} else {
		c.jsonResult(enums.JRCodeSucc,"保持成功",m.Id)
	}
}

func (c *BackendUserController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int,0,len(strs))
	for _,str := range strings.Split(strs,","){
		if id,err := strconv.Atoi(str); err == nil {
			ids = append(ids,id)
		}
	}
	query := orm.NewOrm().QueryTable(models.BackendUserTBName())
	if num,err := query.Filter("id_in",ids).Delete(); err == nil {
		c.jsonResult(enums.JRCodeSucc,fmt.Sprintf("成功删除 %d 项",num),0)
	} else {
		c.jsonResult(enums.JRCodeFailed,"删除失败",0)
	}
}