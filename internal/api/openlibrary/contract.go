package openlibrary

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	baseURLReq = "http://openlibrary.org"

	// Path
	subjectsPath = "/subjects"

	// Extension
	jsonExt = ".json"
)

type Contract interface {
	GetBooksBySubject(ctx context.Context, req UserGetBookRequest) (UserGetBookResponse, error)
}

type openLibrary struct {
	client *http.Client
}

type OpenLibraryParam struct {
	Client *http.Client
}

func NewOpenLibaryApi(param OpenLibraryParam) Contract {
	return &openLibrary{
		client: param.Client,
	}
}

func (o *openLibrary) GetBooksBySubject(ctx context.Context, req UserGetBookRequest) (UserGetBookResponse, error) {
	request, err := http.NewRequest(http.MethodGet, baseURLReq+subjectsPath+"/"+req.Subject+jsonExt, nil)
	if err != nil {
		return UserGetBookResponse{}, err
	}

	resp, err := o.client.Do(request)
	if err != nil {
		return UserGetBookResponse{}, err
	}
	defer resp.Body.Close()

	data := UserGetBookResponse{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return UserGetBookResponse{}, err
	}

	return data, nil
}
