package model

type User struct {
	Id   int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"type:varchar(10);index:name_index" json:"name"`
	Sex  int    `json:"sex"`
}

func init() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic("AutoMigrate run auto migration for given models fila")
	}
}

func (User) TableName() string {
	return "user"
}

func GetUserList(users *[]*User) {
	DB.Find(&users)
}

func GetUserById(user *User, id int) {
	DB.First(&user, id)
}

func GenerateRandomUserData() {
	DB.Create(&User{
		//Name: util.RandString(6),
		//Sex:  util.RandomInt(0, 2),
	})
}
