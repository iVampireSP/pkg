package pkg

import "golang.org/x/crypto/bcrypt"

// PasswordHash 密码哈希  同PHP函数 password_hash()
func (UtilsStruct) PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// PasswordVerify 密码验证  同PHP函数 password_verify()
func (UtilsStruct) PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))

	return err == nil
}
