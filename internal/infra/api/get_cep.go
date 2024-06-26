package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/darleilopes/google-cloud-run/config"
	"github.com/darleilopes/google-cloud-run/internal/dto"
	"github.com/darleilopes/google-cloud-run/internal/entity"
)

var createCepEndpoint = func(baseUrl, cep string) string {
	return strings.Join([]string{baseUrl, "ws", cep, "json"}, "/")
}

type CEPFromAPI struct {
	cnf *config.Config
}

func NewCEPFromAPI(cnf *config.Config) *CEPFromAPI {
	return &CEPFromAPI{
		cnf: cnf,
	}
}

func (cap *CEPFromAPI) Get(ctx context.Context, cep string) (entity.Location, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req, reqErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		createCepEndpoint(cap.cnf.CEP.URL, cep),
		nil,
	)
	if reqErr != nil {
		return entity.Location{}, reqErr
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, doErr := client.Do(req)
	if doErr != nil {
		return entity.Location{}, doErr
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return entity.Location{}, readErr
	}

	var location dto.LocationOut
	if unmErr := json.Unmarshal(bodyBytes, &location); unmErr != nil {
		return entity.Location{}, unmErr
	}

	if location.Erro {
		return entity.Location{}, entity.ErrCEPNotFound
	}

	return entity.Location{
		Cep:        location.CEP,
		Localidade: location.Localidade,
	}, nil
}
