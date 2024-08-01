package service

import (
	context "context"

	types "github.com/Danil-58/weather-app/types"
	mock "github.com/stretchr/testify/mock"
)

// MockWeatherService
type MockWeatherService struct {
	mock.Mock
}

// CreateWeather
func (_m *MockWeatherService) CreateWeather(ctx context.Context, url *types.Api) (*types.StoreData, error) {
	ret := _m.Called(ctx, url)

	var r0 *types.StoreData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Api) (*types.StoreData, error)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Api) *types.StoreData); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StoreData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Api) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWeatherByCity
func (_m *MockWeatherService) GetWeatherByCity(ctx context.Context, city string) (types.StoreData, error) {
	ret := _m.Called(ctx, city)

	var r0 types.StoreData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (types.StoreData, error)); ok {
		return rf(ctx, city)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) types.StoreData); ok {
		r0 = rf(ctx, city)
	} else {
		r0 = ret.Get(0).(types.StoreData)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockWeatherService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWeatherService
func NewMockWeatherService(t mockConstructorTestingTNewMockWeatherService) *MockWeatherService {
	mock := &MockWeatherService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
