package storage

import (
	"log"
	"restaurant/model"
)

type UserStorage struct {
	Users []model.User
}

var userStorage *UserStorage

func init() {
	userStorage = &UserStorage{
		Users: make([]model.User, 0),
	}
	log.Println("User storage initialized with empty user list")
}

func NewUserStorage() *UserStorage {
	return userStorage
}

func (s *UserStorage) GetUserByID(id int) (*model.User, bool) {
	for _, user := range s.Users {
		if user.ID == id {
			return &user, true
		}
	}
	return nil, false
}

func (s *UserStorage) GetUserByUsername(username string) (*model.User, bool) {
	for _, user := range s.Users {
		if user.Username == username {
			return &user, true
		}
	}
	return nil, false
}

func (s *UserStorage) AddUser(user model.User) {
	s.Users = append(s.Users, user)
	log.Printf("User added: ID=%d, Username=%s", user.ID, user.Username)
}

func (s *UserStorage) UserExists(username string) bool {
	_, exists := s.GetUserByUsername(username)
	return exists
}

func (s *UserStorage) GetUserCount() int {
	return len(s.Users)
}
