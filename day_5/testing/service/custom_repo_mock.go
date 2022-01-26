package service

type RepoMock struct {

}

func (r RepoMock) CreateUser(u User) error {
	return nil
}
