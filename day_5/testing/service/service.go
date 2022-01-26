package service


// Lets now try to test

type User struct {
	Name string
}

//go:generate mockgen -destination=repo_mock.go -package=service -build_flags=-mod=mod github.com/Metalscreame/go-training/day_3/testing/service Repository
type Repository interface {
	CreateUser(u User) error
}

type UserService struct {
	r Repository
}

func (u *UserService) CreateUser(name string) error {
	// validation and stuff


	// saving to db
	us := User{Name: name}

	if err := u.r.CreateUser(us); err != nil {
		return err
	}
	return nil
}
