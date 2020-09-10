type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var registerData RegisterData

	//Parse json into register struct
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	//Validate data
	err := emailx.Validate(registerData.Email)
	if err != nil {
		if err == emailx.ErrInvalidFormat {
			c.JSON(http.StatusConflict, gin.H{
				"status": "Wrong email format",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error while validating email",
		})
	}

	//TODO: add username validation

	//Check if username is already taken
	usernameExists := database.CheckIfUsernameExists(registerData.Username)
	if usernameExists == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This username is already taken",
		})

		return
	}

	//Check if email is already taken
	emailExists := database.CheckIfEmailTaken(registerData.Email)
	if emailExists == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This email is already taken",
		})

		return
	}

	//Create password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while hashing the password",
		})

		return
	}

	registerData.Password = string(hashedPassword)

	// Insert into the database
	err = database.InsertUser(registerData.Username, registerData.Password, registerData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error while inserting user into the database",
		})

		return
	}

	c.JSON(http.StatusOK, &registerData)
}

func Login(c *gin.Context) {
	//Parse JSON into login struct
	var loginData LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	usernameExists := database.CheckIfUsernameExists(loginData.Username)
	if usernameExists == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exists",
		})
		return
	}

	//compare passwords
	passwordToCheck := database.GetUsersPassword(loginData.Username)
	if passwordToCheck == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while checking password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(passwordToCheck), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Passwords do not match",
		})
		return
	}

	//get user id
	id := database.GetUsersID(loginData.Username)

	token, err := middleware.GetAuthToken(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

		return
	}
}
