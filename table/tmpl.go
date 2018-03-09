package table

var (
	StructModelTPL = `package models
{{importTimePkg}}
{{modelStruct}}
`

	ModelTPL = `package models

import (
	{{timePkg}}
	"zcm_crm_app/utils"
	"github.com/astaxie/beego/orm"
)

{{modelStruct}}

func init() {
	orm.RegisterModel(new({{modelName}}))
}

// Add {{modelName}}
func Add{{modelName}}(m *{{modelName}}) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err!=nil{
		utils.SendErrEmail("Add{{modelName}}",err)
	}
	return
}

// Get {{modelName}} by {{modelPkName}}
func Get{{modelName}}By{{modelBy}}(key {{modelPkType}}) (v *{{modelName}}, err error) {
	o := orm.NewOrm()
	v = &{{modelName}}{}
	err = o.QueryTable(new({{modelName}})).Filter("{{modelPkName}}", key).One(v)
	if err!=nil && err!= orm.ErrNoRows{
		utils.SendErrEmail("Get{{modelName}}By{{modelPkName}}",err)
	}
	return v,err
}

// Get {{modelName}} list by {{modelPkName}}
func Get{{modelName}}List(key string)(list []*{{modelName}},err error){
	o := orm.NewOrm()
	_,err = o.QueryTable(new({{modelName}})).Filter("{{modelPkName}}", key).All(&list)
	if err!=nil {
		utils.SendErrEmail("Get{{modelName}}List",err)
	}
	return list,err
}

// Update {{modelName}}
func Update{{modelName}}(m *{{modelName}}) (err error) {
	o := orm.NewOrm()
	_,err = o.Update(m)
	if err!=nil{
		utils.SendErrEmail("Update{{modelName}}",err)
	}
	return
}

// Delete {{modelName}}
func Delete{{modelName}}(pk {{modelPkType}}) (err error) {
	o := orm.NewOrm()
	v := {{modelName}}{{{modelBy}}: pk}
	// ascertain id exists in the database
	_,err = o.Delete(&v)
	if err!=nil{
		utils.SendErrEmail("Delete{{modelName}}",err)
	}
	return
}
`
)
