package service

type Service struct{}

func (s *Service) GetList() {}
func (s *Service) Create()  {}
func (s *Service) GetOne()  {}
func (s *Service) Update()  {}
func (s *Service) Delete()  {}

type Group struct {
	ShortUrlService
	RedirectRecordService
	AdminService
}

var ServiceGroup = new(Group)
