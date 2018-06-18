package models

import (
	//"io"
	"crypto/rand"
	"log"
	"fmt"
)

type Session struct {
	ID 			int			`gorm:"primary_key;AUTO_INCREMENT"`
	Sid 		string		`gorm:"type:varchar(255);not null;unique"`
	Name 		string		`gorm:"type:varchar(100);not null"`
	Email 		string		`gorm:"type:varchar(100);not null"`
	Role 		string		`gorm:"type:varchar(100)"`
	IsAlive		bool
}

//func (s Session) Start() (ss gin.HandlerFunc) {
//	return func(c *gin.Context) {
//		store := cookie.NewStore([]byte("secret"))
//		sessions.Session("GinSession", store)
//		c.Next()
//	}
//}
//
//type Manager struct {
//	cookieName	string
//	lock 		sync.Mutex
//	provider 	Provider
//	maxlifetime int64
//}

//type Provider interface {
//	SessionInit(sid string) (Session, error)
//	SessionRead(sid string) (Session, error)
//	SessionDestroy(sid string) error
//	SessionGC(maxlifeTime int64)
//}

//type Session interface {
//	Set(key, value interface{}) error
//	Get(key interface{}) interface{}
//	Delete(key interface{}) error
//	SessionID() string
//}

//var provides = make(map[string]Provider)

//func SessionStart() gin.HandlerFunc {
//	return func(c *gin.Context){
//		store := cookie.NewStore([]byte("secret"))
//		sessions.Sessions("GinSession", store)
//		c.Next()
//	}
//}

//func Register(name string, provider Provider) {
//	if provider == nil {
//		panic("session:Register provide is nil")
//	}
//	if _, dup := provides[name]; dup {
//		panic("session:Register called twice for provide" + name)
//	}
//	provides[name] = provider
//}

//func (s *Session) sessionId() string {
//	b := make([]byte,32)
//	if _, err := io.ReadFull(rand.Reader, b); err != nil {
//		return ""
//	}
//	return base64.URLEncoding.EncodeToString(b)
//}

func createSid() (sid string) {
	u := new([16]byte)
	_ , err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	sid = fmt.Sprintf("%x-%x-%x-%x-%x",u[0:4],u[4:6],u[6:8],u[8:10],u[10:])
	return
}

//func (mg *Manager) Set(key,value interface{}) {
//	//u := User{}
//	//user := u.GetUserFromEmail(u.Email)
//	//session := sessions.Default(c)
//	//session.Set("isAlieve",true)
//	//session.Set("name", user.Name)
//	//session.Set("email",user.Email)
//	//session.Set("uid",user.Uuid)
//	//session.Save()
//}
//func (mg *Manager) Set(key,value interface{}) {
//	//u := User{}
//	//user := u.GetUserFromEmail(u.Email)
//	//session := sessions.Default(c)
//	//session.Set("isAlieve",true)
//	//session.Set("name", user.Name)
//	//session.Set("email",user.Email)
//	//session.Set("uid",user.Uuid)
//	//session.Save()
//}

//func (s *Session) Get(c *gin.Context) interface{}{
//	session := sessions.Default(c)
//	s := map[string]interface{}{
//		"name":session.Get("name"),
//		"email":session.Get("email"),
//		"uid":session.Get("uid"),
//	}
//	session.Save()
//	return s
//}
//func (m *Manager) Delete(c *gin.Context) {
//	session := sessions.Default(c)
//	session.Clear()
//	session.Save()
//}