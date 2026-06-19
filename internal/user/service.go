package user

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
