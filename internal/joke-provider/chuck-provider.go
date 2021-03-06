package jokeprovider

import (
	"io/ioutil"
	"net/http"

	"github.com/spore2102/joker/internal/config"
	"github.com/spore2102/joker/internal/utils"
)

type chuckJokeProvider struct {
	apiUrl string
}

type ChuckJokeProvider interface {
	GetJoke() (string, error)
}

func initChuckJokeProvider(cfg config.ChuckJokesConfig) ChuckJokeProvider {
	return &chuckJokeProvider{
		apiUrl: cfg.Url,
	}
}

func (provider *chuckJokeProvider) GetJoke() (string, error) {
	resp, err := http.Get(provider.apiUrl)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	joke, err := utils.GetByKeyFromJson(body, "value")

	if err != nil {
		return "", err
	}

	return joke, nil
}
