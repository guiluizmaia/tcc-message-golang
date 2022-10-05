package structs

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Age         string `json:"age"`
	Document    string `json:"document"`
	Address     string `json:"address"`
	Nationality string `json:"nationality"`
	MotherName  string `json:"motherName"`
	FatherName  string `json:"fatherName"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
}
