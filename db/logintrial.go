package db

import (
	"github.com/rodzy/flash/models"
	"golang.org/x/crypto/bcrypt"
)

//LoginValidation function to try the user login into db
func LoginValidation(email string,password string) (models.User,bool)  {
	u,isReal,_:= UserFound(email)
	if isReal==false {
		return u,false
	}
	//Decryptng input password/ Database password to compare
	passwordBytes := []byte(password)
	passwordUser:=[]byte(u.Password)
	err:=bcrypt.CompareHashAndPassword(passwordUser,passwordBytes)
	if err != nil {
		return u,false
	}
	return u,true
	
}