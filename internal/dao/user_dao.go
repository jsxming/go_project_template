package dao

import "go_project_template/internal/model"

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (d *Dao) CreateUser(param *User) (*model.User, error) {
	u := &model.User{
		Name: "tm",
		Age:  21,
	}
	return u, nil

}
func (d *Dao) GetUser(name string) *model.User {
	u := model.User{Name: "tm"}
	return &u

}
