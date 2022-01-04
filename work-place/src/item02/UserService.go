package item02

type UserService struct {

	customers []User
	cosNumber int

}

func (service *UserService) List() []User {

	return service.customers
}

func (service *UserService) AddUser(id int, name string, address string, email string) {

	user := UserFactory(id,name,address,email)

	service.customers = append(service.customers,user)
}

func (service *UserService) DelUser(id int) bool{

	for i,_ := range service.customers{

		if service.customers[i].id==id{

			service.customers = append(service.customers[:i],service.customers[i+1:]...)
			return true
		}

	}

	return false
}

func (service *UserService) UpdateUser(id int,name string,address string,email string) bool{

	for i,_ := range service.customers{

		if service.customers[i].id == id{

			service.customers[i].id = id
			service.customers[i].name = name
			service.customers[i].email = email
			service.customers[i].address = address

			return true
		}
	}

	return false
}

func UserServiceFactory() *UserService{

	user := UserFactory(1,"gyh","beijing","172.com")

	userService := &UserService{}

	userService.customers = append(userService.customers,user)
	userService.cosNumber = 1

	return userService
}

func (service *UserService) GetCustomers() []User{

	return service.customers
}