package service

import (
	"ekgo/app/model"
	"ekgo/lib/cache"
	"ekgo/lib/conv"
	"ekgo/lib/jwt"
	"ekgo/lib/orm"
	"ekgo/lib/response"
	"ekgo/lib/secret"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type AdminInterface interface {
	Login() *response.Write  //登录
	Index() *response.Write  //分页
	Store() *response.Write  //添加
	Update() *response.Write //修改
	Delete() *response.Write //删除
	GetCache() model.Admin   //获取缓存
	Show() *response.Write   //查询
}

type Admin struct {
	PageParam response.PageParam
	Model     model.Admin
	Db        *gorm.DB
}

//登录
func (this *Admin) Login() *response.Write {
	//用户名查询
	row := model.Admin{}
	err := this.Db.Where("username=?", this.Model.Username).First(&row).Error

	if err != nil {
		log.Println(err.Error())
	} else {
		this.Model.Password = secret.Decode(this.Model.Password)
		password := secret.Sha256(row.Salt + this.Model.Password)

		if password == row.Password {
			//创建JWT登录token
			token, _ := jwt.Create(map[string]interface{}{
				"id":  row.Id,
				"exp": time.Now().Add(time.Hour * 168).Unix(),
			})

			return &response.Write{
				Code: 200,
				Msg:  "登录成功",
				Data: map[string]string{
					"Authorization": token,
					"username":      row.Username,
					"name":          row.Name,
				},
			}
		}
	}

	return &response.Write{
		Code: 403,
		Msg:  "用户名或者密码不正确",
	}
}

//分页
func (this *Admin) Index() *response.Write {
	var list = []model.Admin{}

	err := this.Db.Select("id,name,username,status,super,role_id,created_at").Scopes(
		orm.WhereQueryBuild(this.PageParam.Filter),
		orm.Order(this.PageParam.Order),
		orm.Paginate(this.PageParam.PageSize, this.PageParam.CurrentPage),
	).Find(&list).Offset(0).Count(&this.PageParam.Total).Error

	if err != nil {
		return response.Fail("获取失败")
	}
	return &response.Write{
		Code: 200,
		Data: response.Page{
			Total: this.PageParam.Total,
			List:  list,
		},
	}
}

//添加
func (this *Admin) Store() *response.Write {
	var countAdmin = model.CountAdminByUsername(this.Model.Username)
	if countAdmin > 0 {
		return response.Success("用户已名已存在")
	} else {
		uuid := secret.GetUUID()
		this.Model.Password = secret.Sha256(uuid + this.Model.Password)
		this.Model.Salt = uuid
		err := this.Db.Create(&this.Model).Error

		if err == nil {
			return response.Success("保存成功")
		}

		return response.Fail("保存失败", err.Error())
	}
}

//修改
func (this *Admin) Update() *response.Write {
	countAdmin := model.CountAdminByUsername(this.Model.Name)

	if countAdmin > 1 {
		return response.Fail("用户已名已存在")
	} else {
		if this.Model.Password != "" {
			uuid := secret.GetUUID()
			this.Model.Password = secret.Sha256(uuid + this.Model.Password)
			this.Model.Salt = uuid
		}
		err := this.Db.Model(&this.Model).Update(this.Model).Error
		if err == nil {
			return response.Success("修改成功")
		}
		return response.Fail("修改失败", err.Error())
	}
}

//删除
func (this *Admin) Delete() *response.Write {
	err := this.Db.First(&this.Model, this.Model.Id).Error

	if err == nil {
		this.Db.Where("id=?", this.Model.Id).Delete(&this.Model)
		return response.Success("删除成功")
	}
	return response.Fail("删除失败", err.Error())
}

//查询单个
func (this *Admin) Show() *response.Write {
	err := this.Db.First(&this.Model, this.Model.Id).Error

	if err == nil {
		return response.Success("获取成功", this.Model)
	}
	return response.Fail("获取失败")
}

//查询缓存数据
func (this *Admin) GetCache() model.Admin {
	cache := cache.New{Key: conv.String(this.Model.Id), Group: "admin"}
	getCache := cache.Get()
	if getCache == nil {
		this.Db.First(&this.Model, this.Model.Id)
		cache.Set(this.Model)
	}

	conv.BytesToStruct(getCache, &this.Model)
	return this.Model
}
