package models

// import (
// 	"errors"
// 	"reflect"
// 	"strings"
// 	"time"

// 	"github.com/beego/beego/v2/client/orm"

// 	// "github.com/beego/beego/v2/core/utils"
// 	_ "github.com/go-sql-driver/mysql"
// )

// type AuthUser struct {
// 	Id          int    `json:"id, omitempty" orm:"column(id);pk;unique;auto_increment"`
// 	Name        string `json:"name, omitempty" orm:"column(name);size(20)"`
// 	Username    string `json:"username, omitempty" orm:"column(username);size(32)"`
// 	Password    string `json:"password, omitempty" orm:"column(password);size(128)"`
// 	Salt        string `json:"salt" orm:"column(salt);size(128)"`
// 	Active      bool   `json:"active, omitempty" orm:"column(active)"`
// 	Email       string `json:"email, omitempty" orm:"column(email);size(50)"`
// 	Address     string `json:"address, omitempty" orm:"column(address);size(50)"`
// 	CreatedAt   int64  `json:"created_at, omitempty" orm:"column(created_at);size(11)"`
// 	UpdatedAt   int64  `json:"updated_at, omitempty" orm:"column(updated_at);size(11)"`
// 	DeviceCount int    `json:"device_count" orm:"column(device_count);size(64);default(0)`
// }

// // func init() {
// // 	orm.NewOrm().QueryTable(new(AuthUser))
// // }

// func AuthUsers() orm.QuerySeter {
// 	return orm.NewOrm().QueryTable(new(AuthUser))
// }

// func CheckUserId(userId int) bool {
// 	exist := AuthUsers().Filter("Id", userId).Exist()
// 	return exist
// }

// func CheckUserName(username string) bool {
// 	exist := AuthUsers().Filter("Username", username).Exist()
// 	return exist
// }

// func CheckUserIdandToken(userId int, token string) bool{
// 	exist := AuthUsers().Filter("Id", userId).Filter("Token", token).Exist()
// 	return exist
// }

// func CheckEmail(email string) bool {
// 	exist := AuthUsers().Filter("Email", email).Exist()
// 	return exist
// }

// func (u *AuthUser) CheckPassword(password string) (ok bool, err error) {
// 	hash, err := utils.GeneratePassHash(password, u.Salt)
// 	if err != nil {
// 		return false, err
// 	}
// 	return u.Password == hash, nil
// }

// func GetUserById(id int) (v *AuthUser, err error) {
// 	o := orm.NewOrm()
// 	v = &AuthUser{Id: id}
// 	if err = o.QueryTable(new(AuthUser)).Filter("Id", id).RelatedSel().One(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

// func GetUserByUserName(username string) (v *AuthUser, err error) {
// 	o := orm.NewOrm()
// 	v = &AuthUser{Username: username}
// 	if err = o.QueryTable(new(AuthUser)).Filter("Username", username).RelatedSel().One(v); err == nil {
// 		return v, nil
// 	}
// 	return v, nil
// }

// func GetAllUser(query map[string]string, fields []string, sortby []string, order []string, offset int, limit int) (ml []interface{}, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(AuthUser))

// 	for k, v := range query {
// 		k = strings.Replace(k, ".", "__", -1)
// 		qs = qs.Filter(k, v)
// 	}

// 	var sortFields []string
// 	if len(sortby) != 0 {
// 		if len(sortby) == len(order) {
// 			for i, v := range sortby {
// 				orderby := ""
// 				if order[i] == "desc" {
// 					orderby = "-" + v
// 				} else if order[i] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either asc or desc")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 			qs = qs.OrderBy(sortFields...)
// 		} else if len(sortby) != len(order) && len(order) == 1 {
// 			for _, v := range sortby {
// 				orderby := ""
// 				if order[0] == "desc" {
// 					orderby = "-" + v
// 				} else if order[0] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, errors.New("Error: Invalid order. Must be either asc or desc")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 		} else if len(sortby) != len(order) && len(order) != 1 {
// 			return nil, errors.New("Error: 'sortby', 'order', sizes mismatch or 'order' size is not 1")
// 		}
// 	} else {
// 		if len(order) != 0 {
// 			return nil, errors.New("Error: unused 'order' fields")
// 		}
// 	}

// 	var l []AuthUser
// 	qs = qs.OrderBy(sortFields...).RelatedSel()
// 	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
// 		if len(fields) == 0 {
// 			for _, v := range l {
// 				ml = append(ml, v)
// 			}
// 		} else {
// 			for _, v := range l {
// 				m := make(map[string]interface{})
// 				val := reflect.ValueOf(v)
// 				for _, fname := range fields {
// 					m[fname] = val.FieldByName(fname).Interface()
// 				}
// 				ml = append(ml, m)
// 			}
// 		}
// 		return ml, nil
// 	}
// 	return nil, err
// }

// func UserByToken(token string)(bool, AuthUser){
// 	o := orm.NewOrm()
// 	var authUser AuthUser
// 	err := o.QueryTable(authUser).Filter("Token", token).One(&authUser)
// 	return err != orm.ErrNoRows, authUser
// }

// func AuthLogin(username string, password string) (bool, *AuthUser) {
// 	o := orm.NewOrm()
// 	authUser, err := GetUserByUserName(username)
// 	if err != nil {
// 		return false, nil
// 	}
// 	passwordHash, err := utils.GeneratePassHash(password, authUser.Salt)
// 	if err != nil {
// 		return false, nil
// 	}
// 	err = o.QueryTable(authUser).Filter("Username", username).Filter("Password", passwordHash).One(authUser)
// 	return err != orm.ErrNoRows, authUser
// }

// // func GetUserByUserName(username string) (err error, authUser *AuthUser) {
// // 	o := orm.NewOrm()
// // 	authUser = &AuthUser{Username: username}
// // 	if err := o.QueryTable(authUser).Filter("Username", username).One(&authUser); err == nil {
// // 		return nil, authUser
// // 	}
// // 	return err, nil
// // }

// func AddUser(m *AuthUser)(*AuthUser, error){
// 	o := orm.NewOrm()
// 	salt, err := utils.GenerateSalt()
// 	if err != nil {
// 		return nil, err
// 	}

// 	passwordHash, err := utils.GeneratePassHash(m.Password, salt)
// 	if err != nil{
// 		return nil, err
// 	}
// 	CreatedAt := time.Now().UTC().Unix()
// 	UpdatedAt := CreatedAt
// 	LastLogin := CreatedAt

// 	et := utils.EasyTo
// }

// func AddAuthUser(m *AuthUser) (id int64, err error) {
// 	o := orm.NewOrm()
// 	CreatedAt := time.Now().UTC().Unix()

// 	authUser := AuthUser{
// 		Id:        m.Id,
// 		Name:      m.Name,
// 		Email:     m.Email,
// 		Active:    m.Active,
// 		CreatedAt: CreatedAt,
// 	}

// 	id, err = o.Insert(&authUser)
// 	if err == nil {
// 		return id, err
// 	}
// 	return 0, err
// }

// func GetauthUserById(id int) (v *AuthUser, err error) {
// 	o := orm.NewOrm()
// 	v = &AuthUser{Id: id}
// 	if err = o.QueryTable(new(AuthUser)).Filter("Id", id).RelatedSel().One(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

// func GetAllAuthUsers(query map[string]string, fields []string, sortby []string, order []string, offset int, limit int) (ml []interface{}, totalCount int64, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(AuthUser))

// 	//query k=v
// 	for k, v := range query {
// 		k = strings.Replace(k, ".", "__", -1)
// 		qs = qs.Filter(k, v)
// 	}

// 	//order by
// 	var sortFields []string
// 	if len(sortby) != 0 {
// 		if len(sortby) == len(order) {
// 			for i, v := range sortby {
// 				orderby := ""
// 				if order[i] == "desc" {
// 					orderby = "-" + v
// 				} else if order[i] == "asc" {
// 					orderby = v
// 				} else {
// 					return nil, 0, errors.New("Error: Invalid order. Must be either asc or desc")
// 				}
// 				sortFields = append(sortFields, orderby)
// 			}
// 			qs = qs.OrderBy(sortFields...)
// 		} else if len(sortby) != len(order) && len(order) == 1 {
// 			for _, v := range sortby {
// 				orderBy := ""
// 				if order[0] == "desc" {
// 					orderBy = "-" + v
// 				} else if order[0] == "asc" {
// 					orderBy = v
// 				} else {
// 					return nil, 0, errors.New("Error: Invalid order. Must be either asc or desc")
// 				}
// 				sortFields = append(sortFields, orderBy)
// 			}
// 		} else if len(sortby) != len(order) && len(order) != 1 {
// 			return nil, 0, errors.New("Error: 'sortby', 'order', sizes mismatch or 'order' size is not 1")
// 		}
// 	} else {
// 		if len(order) != 0 {
// 			return nil, 0, errors.New("Error: unused 'order' fields")
// 		}
// 	}

// 	return nil, 0, err
// }

// func UpdateAuthUserById(m *AuthUser) (err error) {
// 	o := orm.NewOrm()
// 	v := AuthUser{Id: m.Id}

// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Update(m); err == nil {
// 			fmt.Println("Number of records updated in database:", num)
// 		}
// 	}
// 	return
// }

// func DeleteAuthUser(id int) (err error) {
// 	o := orm.NewOrm()
// 	v := AuthUser{Id: id}
// 	if err = o.Read(&v); err == nil {
// 		var num int64
// 		if num, err = o.Delete(&AuthUser{Id: id}); err == nil {
// 			fmt.Println("Number of records deleted in database:", num)
// 		}
// 	}
// 	return
// }
