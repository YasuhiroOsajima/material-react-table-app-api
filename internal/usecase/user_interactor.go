package usecase

import (
	"encoding/json"
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

func (interactor *UserInteractor) GetUsers() []*model.User {
	jsonFile, err := os.Open("dummy_people.json")
	if err != nil {
		fmt.Println("Can not open dummy data file")
		return nil
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Can not read dummy data file")
		return nil
	}

	users := make([]*model.User, 0)
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		fmt.Println("Can not unmarshal dummy data file")
		return nil
	}

	return users
}
