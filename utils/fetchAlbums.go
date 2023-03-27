package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Phaseant/MusicTUI"
)

func FetchAlbums(url string) ([]MusicTUI.Album, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var albums []MusicTUI.Album

	if err = json.Unmarshal(body, &albums); err != nil {
		return nil, errors.New("unable to unmarshall JSON")
	}

	return albums, nil
}
