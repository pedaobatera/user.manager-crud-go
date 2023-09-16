package model

import (
	"encoding/json"
	"fmt"

	"user.manager-crud-go/src/configuration/rest_err"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int

	GetID() string

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()

	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(email, password, name string, age int,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserLoginDomain(email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(name string, age int,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}
