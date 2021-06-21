package helpers

import (
	"fmt"
	"gitlab.com/claude.roy790/gratitude-journal/config"
	"gitlab.com/claude.roy790/gratitude-journal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"time"
)

var UserHelper = UserHelp{}

type UserHelp struct{}

func (UserHelp) Register(email, password, verifiedPassword string) models.RegisterResult {
	if !isValidEmail(email) {
		return models.RegisterResult{"","Email must be valid."}
	}
	if len(password) < config.Get().App.MinPasswordLength {
		return models.RegisterResult{"",fmt.Sprintf("Password must be at least %v characters.",config.Get().App.MinPasswordLength)}
	}
	if password != verifiedPassword {
		return models.RegisterResult{"","Passwords are not identical."}
	}
	if user, _ := UserHelper.GetUser(email); user.ID != primitive.NilObjectID {
		return models.RegisterResult{"",email + " is already registered"}
	}

	return registerUser(email, password)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (UserHelp) GetUser(email string) (models.User, error) {
	user := models.User{}
	if err := MongoHelper.UsersCollection.FindOne(MongoHelper.Context, bson.D{{"email", email}}).Decode(&user); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func registerUser(email, password string) models.RegisterResult {
	var err error
	var hashed string

	hashed, err = hashPass(password)
	if err != nil {
		return models.RegisterResult{"",GetErrorMessage(err)}
	}

	user := models.User{
		Email:        email,
		PasswordHash: hashed,
		DateCreated:  time.Now(),
	}
	res, err := MongoHelper.UsersCollection.InsertOne(MongoHelper.Context, user)
	insertedId, ok := res.InsertedID.(primitive.ObjectID)

	if err != nil || !ok {
			return models.RegisterResult{"","A server error occurred, try again later"}
	}

	return models.RegisterResult{insertedId.Hex(), GetErrorMessage(err)}
}

func hashPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (UserHelp) Login(email, password string) models.LoginResult {
	var err error
	var user models.User

	user, err = UserHelper.GetUser(email)
	if err == mongo.ErrNoDocuments {
		return models.LoginResult{"", false, "This user is not registered."}
	} else if err != nil {
		fmt.Println("error: ",err)
		return  models.LoginResult{"", false, "A server error occurred, try again later."}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return  models.LoginResult{"", false, "Password is not valid"}
	}

	return models.LoginResult{user.ID.Hex(), true, ""}
}

func (UserHelp) DeleteUser(email string) (bool,error) {
	if user, err := UserHelper.GetUser(email);err != nil {
		return false,err
	} else if user.ID != primitive.NilObjectID {
		res, err := MongoHelper.UsersCollection.DeleteOne(MongoHelper.Context,user)
		if err != nil {
			return false,err
		}
		return res.DeletedCount == 1,err
	}
	return false,nil
}
