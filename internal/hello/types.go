package hello

type User struct {
	Name string `json:"name" bson:"name" validate:"required"`
}
