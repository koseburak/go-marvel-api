package marvel

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/koseburak/marvel/config"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestGetCharactersSuccessfully(t *testing.T) {

	enteredCharacterName := "vision"

	// build mock response JSON
	jsonMockResponse := `
	{
		"code": 200,
		"status": "Ok",
		"copyright": "© 2021 MARVEL",
		"data": {
			"total": 1,
			"count": 1,
			"results": [
				{
					"id": 1009697,
					"name": "Vision",
					"description": "The metal monstrosity called Ultron created the synthetic humanoid known as the Vision from the remains of the original android Human Torch of the 1940s to serve as a vehicle of vengeance against the Avengers.",
					"comics": {
						"items": [
							{
								"resourceURI": "http://gateway.marvel.com/v1/public/comics/37406",
								"name": "Age of Ultron (2013) #4"
							}
						],
						"returned": 1
					},
					"series": {
						"items": [
							{
								"resourceURI": "http://gateway.marvel.com/v1/public/series/17318",
								"name": "Age of Ultron (2013)"
							}
						],
						"returned": 1
					},
					"stories": {
						"items": [
							{
								"resourceURI": "http://gateway.marvel.com/v1/public/stories/3484",
								"name": "New Avengers (2004) #18",
								"type": "cover"
							}
						],
						"returned": 1
					},
					"events": {
						"items": [
							{
								"resourceURI": "http://gateway.marvel.com/v1/public/events/116",
								"name": "Acts of Vengeance!"
							}
						],
						"returned": 1
					},
					"urls": [
						{
							"type": "detail",
							"url": "http://marvel.com/comics/characters/1009697/vision?utm_campaign=apiRef&utm_source=10a63c989a3727d0b70aefe93859ef2f"
						}
					]
				}
			]
		}
	}
	`

	// create a new reader using mocked JSON response
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonMockResponse)))

	mockClient := &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	mockConfig := &config.Configuration{}
	marvelClient := NewMarvelClient(mockConfig, mockClient)
	result, err := marvelClient.GetCharacters(enteredCharacterName)
	if err != nil {
		t.Error("TestGetCharactersSuccess failed.")
		return
	}

	if result == nil {
		t.Error("TestGetCharactersSuccess failed, rest api result was empty!")
		return
	}

	if result.Code != 200 {
		t.Error("TestGetCharactersSuccess failed, result code was not 200!")
		return
	}
}

func TestGetCharactersWithUnmarshalData(t *testing.T) {

	enteredCharacterName := "vision"

	// build mock response JSON
	jsonMockResponse := `
	{
		"code": "broked code",
		"status": "Ok",
		"copyright": "© 2021 MARVEL",
		"data": {
			"total": 1,
			"count": 1
		}
	}
	`

	// create a new reader using mocked JSON response
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonMockResponse)))

	mockClient := &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	mockConfig := &config.Configuration{}
	marvelClient := NewMarvelClient(mockConfig, mockClient)
	_, err := marvelClient.GetCharacters(enteredCharacterName)
	if err == nil {
		t.Error("TestGetCharactersSuccess failed.")
		return
	}
}
