package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"gorm.io/gorm"
)

type AdminRepository struct {
	Model model.Admin
	List  []model.Admin
	Db    *gorm.DB
	Base  Base
}

//Factory 继承增删改查公共工厂
func (r *AdminRepository) Factory(model *model.Admin, Db ...*gorm.DB) *Base {
	if len(Db) > 0 {
		r.Base.Db = Db[0]
	} else {
		r.Base.Db = db.Master
	}
	r.Base.New(model, r.Base.Db)
	return &r.Base
}

//GetModel
func (r *AdminRepository) GetModel() model.Admin {
	return *r.Base.Model.(*model.Admin)
}

//New
func (r *AdminRepository) New(model model.Admin, Db ...*gorm.DB) *AdminRepository {
	if len(Db) > 0 {
		r.Db = Db[0]
	} else {
		r.Db = db.Master
	}
	r.Model = model
	return r
}

//FindByUserName 用户名查询
func (r *AdminRepository) FindByUserName() (model.Admin, error) {
	err := r.Db.Where("status='true' AND username = ?", r.Model.Username).First(&r.Model).Error
	return r.Model, err
}
