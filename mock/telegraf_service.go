package mock

import (
	"context"

	platform "github.com/influxdata/influxdb"
)

var _ platform.TelegrafConfigStore = &TelegrafConfigStore{}

// TelegrafConfigStore represents a service for managing telegraf config data.
type TelegrafConfigStore struct {
	*UserResourceMappingService
	FindTelegrafConfigByIDF func(ctx context.Context, id platform.ID) (*platform.TelegrafConfig, error)
	FindTelegrafConfigF     func(ctx context.Context, filter platform.TelegrafConfigFilter) (*platform.TelegrafConfig, error)
	FindTelegrafConfigsF    func(ctx context.Context, filter platform.TelegrafConfigFilter, opt ...platform.FindOptions) ([]*platform.TelegrafConfig, int, error)
	CreateTelegrafConfigF   func(ctx context.Context, tc *platform.TelegrafConfig, userID platform.ID) error
	UpdateTelegrafConfigF   func(ctx context.Context, id platform.ID, tc *platform.TelegrafConfig, userID platform.ID) (*platform.TelegrafConfig, error)
	DeleteTelegrafConfigF   func(ctx context.Context, id platform.ID) error
}

// NewTelegrafConfigStore constructs a new fake TelegrafConfigStore.
func NewTelegrafConfigStore() *TelegrafConfigStore {
	return &TelegrafConfigStore{
		UserResourceMappingService: NewUserResourceMappingService(),
		FindTelegrafConfigByIDF: func(ctx context.Context, id platform.ID) (*platform.TelegrafConfig, error) {
			return nil, nil
		},
		FindTelegrafConfigF: func(_ context.Context, f platform.TelegrafConfigFilter) (*platform.TelegrafConfig, error) {
			return nil, nil
		},
		FindTelegrafConfigsF: func(_ context.Context, f platform.TelegrafConfigFilter, opt ...platform.FindOptions) ([]*platform.TelegrafConfig, int, error) {
			return nil, 0, nil
		},
		CreateTelegrafConfigF: func(_ context.Context, tc *platform.TelegrafConfig, userID platform.ID) error {
			return nil
		},
		UpdateTelegrafConfigF: func(_ context.Context, id platform.ID, tc *platform.TelegrafConfig, userID platform.ID) (*platform.TelegrafConfig, error) {
			return nil, nil
		},
		DeleteTelegrafConfigF: func(_ context.Context, id platform.ID) error {
			return nil
		},
	}
}

// FindTelegrafConfigByID returns a single telegraf config by ID.
func (s *TelegrafConfigStore) FindTelegrafConfigByID(ctx context.Context, id platform.ID) (*platform.TelegrafConfig, error) {
	return s.FindTelegrafConfigByIDF(ctx, id)
}

// FindTelegrafConfigs returns a list of telegraf configs that match filter and the total count of matching telegraf configs.
// Additional options provide pagination & sorting.
func (s *TelegrafConfigStore) FindTelegrafConfigs(ctx context.Context, filter platform.TelegrafConfigFilter, opt ...platform.FindOptions) ([]*platform.TelegrafConfig, int, error) {
	return s.FindTelegrafConfigsF(ctx, filter, opt...)
}

// CreateTelegrafConfig creates a new telegraf config and sets b.ID with the new identifier.
func (s *TelegrafConfigStore) CreateTelegrafConfig(ctx context.Context, tc *platform.TelegrafConfig, userID platform.ID) error {
	return s.CreateTelegrafConfigF(ctx, tc, userID)
}

// UpdateTelegrafConfig updates a single telegraf config.
// Returns the new telegraf config after update.
func (s *TelegrafConfigStore) UpdateTelegrafConfig(ctx context.Context, id platform.ID, tc *platform.TelegrafConfig, userID platform.ID) (*platform.TelegrafConfig, error) {
	return s.UpdateTelegrafConfigF(ctx, id, tc, userID)
}

// DeleteTelegrafConfig removes a telegraf config by ID.
func (s *TelegrafConfigStore) DeleteTelegrafConfig(ctx context.Context, id platform.ID) error {
	return s.DeleteTelegrafConfigF(ctx, id)
}