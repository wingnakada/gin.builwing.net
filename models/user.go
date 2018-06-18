package models

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"time"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	ID 			int
	Uuid 		string		`gorm:"type:varchar(64);not null;unique"`
	Name 		string		`gorm:"type:varchar(100);not null" binding:"required"`
	Yomi		string		`sql:"type:varchar(100);not null" binding:"required"`
	Pass 		string		`sql:"type:varchar(100)"`
	Password	string
	Email   	string		`gorm:"not null;unique" binding:"required"`
	Role    	string		`gorm:"null"`
	IsSession	bool		`gorm:"null"`
	//Activate	int8
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time
	Profiles	[]Profile	`gorm:"foreignkey:UserID"`
}

type Profile struct {
	ID 			uint
	UserID 		uint		`sql:"index"`
	Name 		string		`gorm:"type:varchar(100);references:users(id)"`
	Content 	string		`sql:"type:text"`
}

/* データの値からUserを作成する */
func (u User) Create(user *User) {
	/* 同姓同名がいないかを検索する */
	db.Where("name = ?",u.Name).FirstOrCreate(&u)
	u.Uuid = createUUID()
	db.Save(&u)
}
/* データの値からUserを修正する */
func (u User) Edit(user *User) {
	u.Name = user.Name
	u.Yomi = user.Yomi
	u.Email = user.Email
	u.Role = user.Role
	db.Model(u).Update(&u)
}
/* データの値からUserをソフトデリートする */
func (u User) Delete(user *User) {
	db.First(&user)
	db.Delete(user)
}
/* データの値からUserを完全削除する */
func (u User) ForceDelete(user *User) {
	db.Unscoped().First(user)
	db.Unscoped().Delete(&user)
}
/* 削除されたUsersを取得する */
func (u User) Trashed() *[]User {
	users := &[]User{}
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(users)
	return users
}
/* 削除されたUserを復元する */
func (u User) Restore(user *User) *User {
	db.Unscoped().First(user).Update("deleted_at",nil)
	return user
}
func (u User) AllUsers() *[]User {
	users := &[]User{}
	db.Find(users)
	return users
}
func (u User) GetUser(id int) *User {
	user := &u
	user.ID = id
	db.Unscoped().First(user)
	return user
}
func (u User) GetUserFromEmail(email string) *User {
	user := &u
	user.Email = email
	db.Unscoped().Where("email = ?",email).First(user)
	return user
}
func SearchUser(name string) *User {
	user := &User{}
	db.First(user.Name,name)
	return user
}
func SearchUsers(name string) *[]User {
	users:= &[]User{}
	db.Where("name LIKE ?", "%"+name+"%").Find(users)
	return users
}
//func CreatePassword(pass string) string {
//	hash, _ := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
//	return string(hash)
//}
/* パスワードの暗号化 */
func HashPass(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
	return string(hash)
}
/* パスワードの検証 */
func comparePass(hash,pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(pass))
	if err == nil {
		return true
	}else{
		return false
	}
}
///* Loginがパスするかどうかの判断 */
//func (u *User) Login(user *User) bool {
//	//Email自体がなければ
//	if db.Where("email=?",user.Email).First(user).RecordNotFound() {
//		return false
//	}else{
//		//passwordが一致すれば
//		if comparePass(u.Password,user.Password) {
//			return true
//		}else{
//			return false
//		}
//	}
//}
/* Loginがパスするかどうかの判断 */
func (u *User) IsLogin(email,pass string) bool {
	//Email自体がなければ
	if db.Where("email=?",email).First(u).RecordNotFound() {
		return false
	}else{
		//passwordが一致すれば
		if comparePass(u.Password,pass) {
			return true
		}else{
			return false
		}
	}
}
// cookieの設定を行う
func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name: "hoge",
		Value: "bar",
	}
	http.SetCookie(w, cookie)

	//fmt.Fprintf(w, "Cookieの設定ができたよ")
}

func (u *User) SessionSet(c *gin.Context) {
	user := u.GetUserFromEmail(u.Email)
	session := sessions.Default(c)
	session.Set("isAlieve",true)
	session.Set("name", user.Name)
	session.Set("email",user.Email)
	session.Set("uid",user.Uuid)
	session.Set("role",user.Role)
	session.Save()
}
func (u *User) SessionGet(c *gin.Context) interface{}{
	session := sessions.Default(c)
	s := map[string]interface{}{
		"name":session.Get("name"),
		"email":session.Get("email"),
		"uid":session.Get("uid"),
		"role":session.Get("role"),
	}
	session.Save()
	return s
}
func SessionClear(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

//func (u *User) SessionDelete(c *gin.Context) {
//	session := sessions.Default(c)
//	//session.Delete("name")
//	session.Clear()
//	session.Save()
//}