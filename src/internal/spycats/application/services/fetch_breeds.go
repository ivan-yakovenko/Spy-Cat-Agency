package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-faster/errors"
)

var (
	cachedBreeds   []dtos.BreedName
	cacheTimestamp time.Time
	cacheTTL       = time.Hour * 24
	cacheMu        sync.RWMutex
	httpClient     = &http.Client{Timeout: time.Second * 5}
)

func fetchBreedsNames() ([]dtos.BreedName, error) {

	cacheMu.RLock()

	if cachedBreeds != nil && time.Since(cacheTimestamp) < cacheTTL {
		breeds := make([]dtos.BreedName, len(cachedBreeds))
		copy(breeds, cachedBreeds)
		cacheMu.RUnlock()
		return breeds, nil
	}

	cacheMu.RUnlock()

	response, err := httpClient.Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error connecting/fetching data from the catapi", err)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	var breedNames []dtos.BreedName

	if err = json.NewDecoder(response.Body).Decode(&breedNames); err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error decoding data body into dto", err)
	}

	cacheMu.Lock()
	cachedBreeds = breedNames
	cacheTimestamp = time.Now()
	cacheMu.Unlock()

	return breedNames, nil

}

func isValidBreed(breedName string, breedNames []dtos.BreedName) bool {
	for _, breed := range breedNames {
		if breed.Name == breedName {
			return true
		}
	}

	return false
}

var ErrorInvalidBreed = errors.New("Invalid breed name")
