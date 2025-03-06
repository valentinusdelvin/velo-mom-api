package bcrypt

import "golang.org/x/crypto/bcrypt"

type InterBcrypt interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

type Bcrypt struct {
	cost int
}

func Init() InterBcrypt {
	return &Bcrypt{
		cost: 10,
	}
}

func (b *Bcrypt) GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (b *Bcrypt) CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
