package model

type User struct {
	Id             string `db:"id"`
	Username       string `db:"username"`
	Email          string `db:"email"`
	FirstFirstName string `db:"firstname"`
	LastFirstName  string `db:"lastname"`
}

type UserLogin struct {
	Username string `validate:"min=1,max=316"`
	Email    string `valiate:"email"`
}

type InputErrors struct {
	Email    string
	Username string
}

const CreateUserQuery = `INSERT INTO users (id, username, firstname, lastname, email, password) 
VALUES (:id, :username, :firstname, :lastname, :email, :password)`
const UpdateUserQuery = `UPDATE users set name=:name, firstname=:firstname, lastname=:lastname, email=:email, password=:password WHERE id=:id;`
const GetUserByIdQuery = `SELECT * FROM users WHERE id=$1`
const GetUsersQuery = `SELECT * FROM users ORDER BY id ASC`
const DeleteUserQuery = `DELETE FROM users WHERE id=$1`
const GetUserByEmailQuery = `SELECT * FROM users WHERE email=$1`
const GetUserByUserFirstNameQuery = `SELECT * FROM users WHERE name=$1`
const SearchUsersByFirstNameOrEmailQuery = `SELECT * FROM users WHERE name LIKE '%' || $1 || '%' OR email LIKE '%' || $1 || '%';`
