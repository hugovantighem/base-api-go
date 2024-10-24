package infra

import (
	"cmp"
	"context"
	"myproject/app"
	"slices"
	"sync"

	"github.com/sirupsen/logrus"
)

type InMemoryDelegationStorage struct {
	mu    sync.Mutex
	items []app.Delegation
}

func NewInMemoryDelegationStorage() *InMemoryDelegationStorage {
	return &InMemoryDelegationStorage{
		mu:    sync.Mutex{},
		items: make([]app.Delegation, 0),
	}
}

func (x *InMemoryDelegationStorage) Search(ctx context.Context, criteria app.SearchCriteria) ([]app.Delegation, error) {
	logrus.Debugf("search with criteria: %s", criteria)
	x.mu.Lock()
	defer x.mu.Unlock()
	result := []app.Delegation{}
	for _, item := range x.items {
		if criteria.Year != nil {
			if item.Timestamp.Year() == *criteria.Year {
				result = append(result, item)
			}
		} else {
			result = append(result, item)
		}
	}
	return result, nil
}

func (x *InMemoryDelegationStorage) GetLast(ctx context.Context) (app.Delegation, error) {
	x.mu.Lock()
	defer x.mu.Unlock()
	if len(x.items) == 0 {
		return app.Delegation{}, app.ErrNotFound
	}
	return slices.MaxFunc(x.items, func(a, b app.Delegation) int { return cmp.Compare(a.Id, b.Id) }), nil
}

func (x *InMemoryDelegationStorage) Save(ctx context.Context, items []app.Delegation) error {
	logrus.Debugf("saving: %+v", items)
	x.mu.Lock()
	defer x.mu.Unlock()
	slices.SortFunc(items, func(a, b app.Delegation) int { return cmp.Compare(a.Id, b.Id) })
	x.items = append(x.items, items...)
	return nil
}
