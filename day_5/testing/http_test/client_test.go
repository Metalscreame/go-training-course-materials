package iterable

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

const (
	path = "https://google.com"
)

func TestClient_SendEmail(t *testing.T) {

	registerError := func() {
		httpmock.RegisterResponder("GET", path,
			httpmock.NewBytesResponder(500, nil))
	}
	registerErrorBadReuqest := func() {
		httpmock.RegisterResponder("GET", path,
			httpmock.NewBytesResponder(400, nil))
	}
	registerOk := func() {
		httpmock.RegisterResponder("GET", path,
			httpmock.NewBytesResponder(200, nil))
	}
	testCases := []struct {
		name       string
		wantError  bool
		templateID string
		httpMock   func()
	}{
		{
			name:       "not int in template id",
			wantError:  true,
			templateID: "not int",
			httpMock:   func() {},
		},
		{
			name:       "got error from Send email",
			wantError:  true,
			templateID: "123",
			httpMock:   registerError,
		},
		{
			name:       "got error from Send email",
			wantError:  true,
			templateID: "123",
			httpMock:   registerErrorBadReuqest,
		},
		{
			name:       "ok",
			wantError:  false,
			templateID: "123",
			httpMock:   registerOk,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			tc.httpMock()

			cli := Client{client: http.DefaultClient}

			err := cli.GetData()
			if tc.wantError {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
		})
	}
}
