package user

import (
	"fmt"
)

type UserService struct {
	users  []User
	nextId int
}

func NewUserService() *UserService {
	return &UserService{
		users: []User{
			{
				Id:    1,
				Name:  "magic",
				Email: "magic@test.com",
			},
			{
				Id:    2,
				Name:  "mike",
				Email: "mike@test.com",
			},
			{
				Id:    3,
				Name:  "john",
				Email: "john@test.com",
			},
		},
		nextId: 4,
	}
}

func (s *UserService) GetUsers() []User {
	return s.users
}

func (s *UserService) GetUserById(id int) (*User, error) {
	for i := range s.users {
		if s.users[i].Id == id {
			return &s.users[i], nil
		}
	}

	return nil, fmt.Errorf("user with id(%d) not found", id)
}

func (s *UserService) CreateUser(req User) (*User, error) {
	req.Id = s.nextId
	s.users = append(s.users, req)
	s.nextId++
	return &s.users[req.Id-1], nil
}

func (s *UserService) UpdateUser(id int, req User) (*User, error) {
	user, err := s.GetUserById(id)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	user.Email = req.Email

	return user, nil
}

func (s *UserService) PatchUser(id int, req map[string]interface{}) (*User, error) {
	user, err := s.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if name, exists := req["name"].(string); exists {
		user.Name = name
	}

	if email, exists := req["email"].(string); exists {
		user.Email = email
	}

	return user, nil

}
