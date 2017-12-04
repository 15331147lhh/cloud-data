package entities


type UserInfoAtomicService struct{}


var UserInfoService = UserInfoAtomicService{}


func (*UserInfoAtomicService) Save(user *UserInfo) error {
	return Save(user)
}


func (*UserInfoAtomicService) FindAll() []UserInfo {
	return FindAll()
}


func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	return FindByID(id)
}
