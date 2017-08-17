package service

import "github.com/tthanh/yoblog"

// Service represent internal service
type Service struct {
	accountStore yoblog.AccountStore
}

// New create new service
func New(
	accountStore yoblog.AccountStore,
) *Service {
	return &Service{
		accountStore: accountStore,
	}
}
