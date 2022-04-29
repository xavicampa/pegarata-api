package myapi

import (
	"errors"
	openapi "myapi/go"
)

// ItemAPIService - object implementing interface
type ItemAPIService struct {
	itemStore ItemStore
}

// NewItemAPIService - returns object implementing interface
func NewItemAPIService(itemStore ItemStore) openapi.DefaultApiServicer {
	return &ItemAPIService{itemStore: itemStore}
}

// ItemsGet - Returns a list of items
func (s *ItemAPIService) ItemsGet() (interface{}, error) {
	return s.itemStore.RetrieveItems(), nil
}

// ItemsPost -
func (s *ItemAPIService) ItemsPost(item openapi.Item) (interface{}, error) {
	id := s.itemStore.CreateItem(item.Name)
	result := openapi.PersistedItem{Id: id, Name: item.Name, Done: item.Done}
	return result, nil
}

// ItemsItemIdDelete - Removes an item from the list
func (s *ItemAPIService) ItemsItemIdDelete(itemId string) (interface{}, error) {
	deleted := s.itemStore.DeleteItem(itemId)
	if deleted {
		return nil, nil
	} else {
		return nil, errors.New("NotFound")
	}
}

// ItemsItemIdDelete - Removes an item from the list
func (s *ItemAPIService) ItemsItemIdPut(itemId string) (interface{}, error) {
	toggled := s.itemStore.ToggleItem(itemId)
	if toggled {
		return nil, nil
	} else {
		return nil, errors.New("NotFound")
	}
}

// PingGet - Heartbeat
func (s *ItemAPIService) PingGet() (interface{}, error) {
	result := openapi.HeartBeatResponse{Status: "OK", Message: "All good!"}
	return result, nil
}
