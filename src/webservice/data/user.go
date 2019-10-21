package data

// User は、ユーザの構造体。
type User struct {
	Name string
	Password string
}

// FindByName は、ユーザ名をキーとしてユーザ情報を検索する。
func FindByName(name string) (user User, err error) {

	user = User {}
	err = Db.QueryRow("select name, password from users where name=$1", name).Scan(&user.Name, &user.Password)
	return

}
