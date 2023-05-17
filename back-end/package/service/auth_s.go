package service

import (
	"crypto/sha1"
	"fmt"
)

const salt = "kj8932jfgj74thdfgjg78psdfg"

var RedisAuthUser = map[string]bool{}

func (s *Service) Auth(login, pasword string) (string, error) {
	user, err := s.repository.GetUser(login, Hashing(pasword))
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

//Хэширование пароля
func Hashing(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}