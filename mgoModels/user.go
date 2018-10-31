package models

import (
	"log"

	db "github.com/dangyanglim/go_cnode/database"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//"log"

type User struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string        `json:"name"`
	Loginname   string        `json:"loginname" `
	Pass        string        `json:"pass" `
	Email       string        `json:"email" `
	Avatar      string        `json:"avatar" `
	AccessToken string        `json:"accessToken"`
	Score       uint          `json:"score"`
	Active      bool          `json:"active"`
}
type UserModel struct{}

func (p *UserModel) GetUserById(id string) (user User, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	log.Println(id)
	objectId := bson.ObjectIdHex(id)
	err = mgodb.C("users").Find(bson.M{"_id": objectId}).One(&user)
	return user, err
}
func (p *UserModel) GetUserByName(name string) (user User, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	err = mgodb.C("users").Find(bson.M{"name": name}).One(&user)
	return user, err
}
func (p *UserModel) ActiveUserByName(name string) (err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	err = mgodb.C("users").Update(bson.M{"name": name}, bson.M{"$set": bson.M{"active": true}})
	return err
}
func (p *UserModel) GetUserByNameOrEmail(name string, email string) (user User, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	err = mgodb.C("users").Find(bson.M{"$or": []bson.M{bson.M{"name": name}, bson.M{"email": email}}}).One(&user)
	return user, err
}
func (p *UserModel) NewAndSave(name string, loginname string, email string, pass string, avatar_url string, active bool) (err error) {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	mgodb := db.MogSession.DB("egg_cnode")
	u2, _ := uuid.NewV4()
	user := User{
		Id:          bson.NewObjectId(),
		Name:        name,
		Loginname:   loginname,
		Pass:        string(hashPass),
		Avatar:      "http://www.gravatar.com/avatar/81f36fbf3b658c6a2330ca6840f7cb12?size=48",
		Email:       email,
		Active:      active,
		AccessToken: u2.String(),
	}
	err = mgodb.C("users").Insert(&user)
	log.Println(err)
	return err
}
