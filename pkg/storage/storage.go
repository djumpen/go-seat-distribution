package storage

import (
	"github.com/djumpen/go-seat-distribution/pkg/errors"
	"github.com/djumpen/go-seat-distribution/pkg/salon"

	"github.com/bluele/gcache"
)

// Storage represents behavior of entity, that will store aircraft salons
type Storage interface {
	SaveSalon(string, salon.Salon) error
	GetSalon(string) (salon.Salon, error)
}

// CacheStorage stores and returns saved salons
type CacheStorage struct {
	salonCache gcache.Cache
}

// NewCacheStorage creates new cache storage
func NewCacheStorage(cache gcache.Cache) *CacheStorage {
	return &CacheStorage{
		salonCache: cache,
	}
}

// SaveSalon saves salon entity to cache
func (s *CacheStorage) SaveSalon(salonID string, salon salon.Salon) error {
	return s.salonCache.Set(salonID, salon)
}

// GetSalon returns cached salon
func (s *CacheStorage) GetSalon(salonID string) (salon.Salon, error) {
	result, err := s.salonCache.Get(salonID)
	if result == nil {
		return salon.Salon{}, errors.ErrSalonNotFound
	}
	return result.(salon.Salon), err
}
