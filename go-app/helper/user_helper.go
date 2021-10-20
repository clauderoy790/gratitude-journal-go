package helper

import (
	"net/mail"
	"time"

	"github.com/clauderoy790/gratitude-journal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var UserHelper = UserHelp{}

type UserHelp struct{}

func (UserHelp) Register(email, password, verifiedPassword string) repository.RegisterResult {
	// if !isValidEmail(email) {
	// 	return repository.RegisterResult{"","Email must be valid."}
	// }
	// if len(password) < config.Get().App.MinPasswordLength {
	// 	return repository.RegisterResult{UserID:"",Error:fmt.Sprintf("Password must be at least %v characters.",config.Get().App.MinPasswordLength)}
	// }
	// if password != verifiedPassword {
	// 	return repository.RegisterResult{"","Passwords are not identical."}
	// }
	// if user, _ := UserHelper.GetUser(email); user.ID != primitive.NilObjectID {
	// 	return repository.RegisterResult{UserID:"",Error:email + " is already registered"}
	// }
	//todo here
	return registerUser(email, password)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (UserHelp) GetUser(email string) (repository.User, error) {
	user := repository.User{}
	if err := MongoHelper.UsersCollection.FindOne(MongoHelper.Context, bson.D{{"email", email}}).Decode(&user); err != nil {
		return repository.User{}, err
	}

	return user, nil
}

func registerUser(email, password string) repository.RegisterResult {
	var err error
	var hashed string

	hashed, err = hashPass(password)
	if err != nil {
		return repository.RegisterResult{"", GetErrorMessage(err)}
	}

	user := repository.User{
		Email:        email,
		PasswordHash: hashed,
		DateCreated:  time.Now(),
	}
	res, err := MongoHelper.UsersCollection.InsertOne(MongoHelper.Context, user)
	insertedId, ok := res.InsertedID.(primitive.ObjectID)

	if err != nil || !ok {
		return repository.RegisterResult{"", "A server error occurred, try again later"}
	}

	return repository.RegisterResult{UserID: insertedId.Hex(), Error: GetErrorMessage(err)}
}

func hashPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (UserHelp) Login(email, password string) repository.LoginResult {
	// var err error
	// var user repository.User

	// user, err = UserHelper.GetUser(email)
	// if err == mongo.ErrNoDocuments {
	// 	return repository.LoginResult{"", false, "This user is not registered."}
	// } else if err != nil {
	// 	fmt.Println("error: ", err)
	// 	return repository.LoginResult{Error: "A server error occurred, try again later."}
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	// if err != nil {
	// 	return repository.LoginResult{"", false, "Password is not valid"}
	// }

	return repository.LoginResult{}
	//return repository.LoginResult{UserID: user.ID.Hex(), Success: true}
	//todo here
}

func (UserHelp) DeleteUser(email string) (bool, error) {
	// if user, err := UserHelper.GetUser(email); err != nil {
	// 	return false, err
	// } else if user.ID != primitive.NilObjectID {
	// 	res, err := MongoHelper.UsersCollection.DeleteOne(MongoHelper.Context, user)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	return res.DeletedCount == 1, err
	// }
	//todo here
	return false, nil
}
