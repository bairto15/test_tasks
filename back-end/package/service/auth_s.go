package service

import (
	"fmt"	
)

var RedisAuthUser = map[string]bool{}

func (s *Service) Auth(login, pasword string) (string, error) {
	user, err := s.repository.GetUser(login, pasword)
	if err != nil {
		return user.Id, err
	}

	if user.Id == "" {
		return user.Id, fmt.Errorf("Логин или пароль не найдены")
	}

	err = s.repository.Auth(login)
	if err != nil {
		s.Logger.Error(err)
	}

	RedisAuthUser[login] = true

	return user.Id, err
}

func (s *Service) Out(login string) error {
	err := s.repository.Out(login)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	RedisAuthUser[login] = false

	return nil
}