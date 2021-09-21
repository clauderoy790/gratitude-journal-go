package tests

// import (
// 	"github.com/clauderoy790/gratitude-journal/config"
// 	"github.com/clauderoy790/gratitude-journal/helper"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"testing"
// )

// var email = "test@gmail.com"
// var password = "12345678"

// func TestDeleteUser(t *testing.T) {
// 	helper.MongoHelper.Connect()
// 	defer helper.MongoHelper.Disconnect()
// 	uh := helper.UserHelper;

// 	uh.Register(email,password,password);
// 	uh.DeleteUser(email)
// 	if user, _ := uh.GetUser(email);user.ID != primitive.NilObjectID {
// 		t.Fatalf("failed to delete user %v",email)
// 	}
// }

// func TestRegister(t *testing.T) {
// 	helper.MongoHelper.Connect()
// 	defer helper.MongoHelper.Disconnect()
// 	uh := helper.UserHelper;

// 	uh.DeleteUser(email)
// 	if res := uh.Register("test","11111111","11111111");  res.Error == "" {
// 		t.Fatalf("email needs to be in a valid format")
// 	}

// 	if res := uh.Register(email,"123","1"); res.Error == "" {
// 		t.Fatalf("password needs to be at least %v characters",config.Get().App.MinPasswordLength)
// 	}

// 	if res := uh.Register(email,"12323423423423","14432423423423445"); res.Error == "" {
// 		t.Fatalf("password needs to match")
// 	}

// 	if res := uh.Register(email,password,password); res.Error != "" {
// 		t.Fatalf("tried to create a valid user but it failed: %v",res.Error)
// 	}

// 	if res := uh.Register(email,password,password); res.Error == "" {
// 		t.Fatalf("tried to register a user that already exists and did not get an error.")
// 	}

// }

// func TestLogin(t *testing.T) {
// 	helper.MongoHelper.Connect()
// 	defer helper.MongoHelper.Disconnect()
// 	uh := helper.UserHelper;
// 	uh.DeleteUser(email)

// 	if res := uh.Login(email,"312412");res.Success || res.UserID != "" || res.Error == "" {
// 		t.Fatalf("tried to login to a non registered account.")
// 	}

// 	uh.Register(email,"12345678","12345678")

// 	if res := uh.Login(email,"312412");res.Success || res.UserID != "" || res.Error == "" {
// 		t.Fatalf("tried to login to a registred account with an invalid password.")
// 	}

// 	res := uh.Login(email,"12345678")
// 	if !res.Success || res.UserID == "" || res.Error != "" {
// 		t.Fatalf("tried to login to a registred account with a valid password without success. %v %v %v",res.Success,res.UserID,res.Error)
// 	}

// 	objId, err := primitive.ObjectIDFromHex(res.UserID)
// 	if err != nil {
// 		t.Fatalf("failed to convert %v to object id: ",err)
// 	}

// 	user,_ := uh.GetUser(email)
// 	if objId != user.ID {
// 		t.Fatalf("login returned the wrong user id.")
// 	}
// }

//todo here
