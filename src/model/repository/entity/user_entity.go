package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int                `bson:"age,omitempty"`
}

// EncryptPassword implements model.UserDomainInterface.
func (*UserEntity) EncryptPassword() {
	panic("unimplemented")
}

// GetAge implements model.UserDomainInterface.
func (*UserEntity) GetAge() int {
	panic("unimplemented")
}

// GetEmail implements model.UserDomainInterface.
func (*UserEntity) GetEmail() string {
	panic("unimplemented")
}

// GetID implements model.UserDomainInterface.
func (*UserEntity) GetID() string {
	panic("unimplemented")
}

// GetJSONValue implements model.UserDomainInterface.
func (*UserEntity) GetJSONValue() (string, error) {
	panic("unimplemented")
}

// GetName implements model.UserDomainInterface.
func (*UserEntity) GetName() string {
	panic("unimplemented")
}

// GetPassword implements model.UserDomainInterface.
func (*UserEntity) GetPassword() string {
	panic("unimplemented")
}

// SetID implements model.UserDomainInterface.
func (*UserEntity) SetID(string) {
	panic("unimplemented")
}
