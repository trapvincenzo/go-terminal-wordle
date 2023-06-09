// Code generated by mockery v2.23.1. DO NOT EDIT.

package render

import (
	mock "github.com/stretchr/testify/mock"
	game "github.com/trapvincenzo/go-terminal-wordle/pkg/game"
)

// MockRenderer is an autogenerated mock type for the Renderer type
type MockRenderer struct {
	mock.Mock
}

type MockRenderer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRenderer) EXPECT() *MockRenderer_Expecter {
	return &MockRenderer_Expecter{mock: &_m.Mock}
}

// Print provides a mock function with given fields: _a0
func (_m *MockRenderer) Print(_a0 string) {
	_m.Called(_a0)
}

// MockRenderer_Print_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Print'
type MockRenderer_Print_Call struct {
	*mock.Call
}

// Print is a helper method to define mock.On call
//   - _a0 string
func (_e *MockRenderer_Expecter) Print(_a0 interface{}) *MockRenderer_Print_Call {
	return &MockRenderer_Print_Call{Call: _e.mock.On("Print", _a0)}
}

func (_c *MockRenderer_Print_Call) Run(run func(_a0 string)) *MockRenderer_Print_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRenderer_Print_Call) Return() *MockRenderer_Print_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRenderer_Print_Call) RunAndReturn(run func(string)) *MockRenderer_Print_Call {
	_c.Call.Return(run)
	return _c
}

// Render provides a mock function with given fields: _a0
func (_m *MockRenderer) Render(_a0 *game.Game) {
	_m.Called(_a0)
}

// MockRenderer_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type MockRenderer_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - _a0 *game.Game
func (_e *MockRenderer_Expecter) Render(_a0 interface{}) *MockRenderer_Render_Call {
	return &MockRenderer_Render_Call{Call: _e.mock.On("Render", _a0)}
}

func (_c *MockRenderer_Render_Call) Run(run func(_a0 *game.Game)) *MockRenderer_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*game.Game))
	})
	return _c
}

func (_c *MockRenderer_Render_Call) Return() *MockRenderer_Render_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRenderer_Render_Call) RunAndReturn(run func(*game.Game)) *MockRenderer_Render_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRenderer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRenderer creates a new instance of MockRenderer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRenderer(t mockConstructorTestingTNewMockRenderer) *MockRenderer {
	mock := &MockRenderer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
