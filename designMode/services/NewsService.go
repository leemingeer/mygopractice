package services

type NewsService struct {
}

func (us *NewsService) Get(id int) string {
	return "single news"
}
