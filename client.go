struct User {
	Username string
    Salt []byte
    EncPrivKey []byte
    EncSignKey []byte
    FileMapUUID uuid.UUID
}


func InitUser(username string, password string) {

	var user User

	if len(username) < 1 || len(password) > 32 {
		return nil, errors.New("username and password must be between 1 and 32 characters")
	}

	user.Username = username

	// uuid will be the hash of the username
	var uuid = hash.Sha256(username)

	// check if the user already exists
	var user = GetUser(uuid)
	if user != nil {
		return user
	}

	// create a new user
	var masterKey = Argon2Key(password, uuid)
	newUser.Save()

	return newUser
}