package module

import (
	"github.com/golang/mock/gomock"
	"github.com/wilgun/joy-technologies-be/internal/test/mockapi"
	"testing"
)

type MockComponent struct {
	mockAPI    *mockapi.MockContract
	bookModule *bookModule
}

func InitMock(t *testing.T) *MockComponent {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockContract(mockCtrl)

	bookModule := NewBookModule(BookModuleParam{
		OpenLibrary: mockAPI,
	})

	return &MockComponent{
		mockAPI:    mockAPI,
		bookModule: bookModule,
	}
}
