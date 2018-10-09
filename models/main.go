package models 

type loginPayload struct {
    Email  string        `json:"email" bson:"email"`
    Password   string    `json:"password" bson:"password"`
}

type signupPayload struct {
    Email string        `json:"email" bson:"email"`
    Password string      `json:"password" bson:"password"`
    Username   string `json:"username" bson:"username"`
}

type uploadPayload struct {
    Filename string        `json:"email" bson:"email"`
}

type uploadDBPayload struct {
    image []byte `json:"email" bson:"email"`
}

