package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	validInput := "Mike"

	testCases := []struct {
		name      string
		input     string
		wantError bool
		mock      func(ctl *gomock.Controller) Repository
	}{
		{
			input: validInput, // remove input show output
			name:  "testCase 1 - got error from user repo",
			mock: func(ctl *gomock.Controller) Repository {
				r := NewMockRepository(ctl)
				r.EXPECT().CreateUser(User{
					Name: validInput,
				}).Return(fmt.Errorf("got errror"))
				return r
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			srv := UserService{r: tc.mock(ctl)}

			err := srv.CreateUser(tc.input)
			if tc.wantError {
				require.NotNil(t, err)
				return
			}

			require.Nil(t, err)
		})
	}
}
