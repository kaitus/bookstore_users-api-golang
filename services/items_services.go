package services

var (
	ItemsService itemsServicesInterface = &itemsService{}
)

type itemsServicesInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
