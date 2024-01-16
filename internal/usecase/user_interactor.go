package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/YasuhiroOsajima/material-react-table-app-api/internal/model"
)

type UserInteractor struct {
}

func NewUserInteractor() *UserInteractor {
	return &UserInteractor{}
}

func openDummyFile() ([]*model.User, error) {
	jsonFile, err := os.Open("dummy_people.json")
	if err != nil {
		msg := "Can not open dummy data file"
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		msg := "Can not read dummy data file"
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	users := make([]*model.User, 0)
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		msg := "Can not unmarshal dummy data file"
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	return users, nil
}

func (interactor *UserInteractor) GetUsers() ([]*model.User, error) {
	return openDummyFile()
}

func (interactor *UserInteractor) DeleteUser(username string) error {
	users, err := openDummyFile()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Username == username {
			return nil
		}
	}

	return errors.New("User not found")
}
