package module

import (
	"github.com/golang/mock/gomock"
	"github.com/wilgun/joy-technologies-be/internal/test/mockapi"
	"github.com/wilgun/joy-technologies-be/internal/test/mockstore"
	"testing"
)

type MockComponent struct {
	mockAPI    *mockapi.MockContract
	bookModule *bookModule
	bookStore  *mockstore.MockBookStore
}

func InitMock(t *testing.T) *MockComponent {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockContract(mockCtrl)
	mockStore := mockstore.NewMockBookStore(mockCtrl)

	bookModule := NewBookModule(BookModuleParam{
		OpenLibrary: mockAPI,
		BookStore:   mockStore,
	})

	return &MockComponent{
		mockAPI:    mockAPI,
		bookModule: bookModule,
		bookStore:  mockStore,
	}
}
