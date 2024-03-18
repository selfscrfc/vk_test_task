package redis

import (
	"fmt"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"vk_test_task/config"
)

func TestRepository_CreateAccessToken(t *testing.T) {
	client, _ := redismock.NewClientMock()
	defer client.Close()

	r := Repository{DB: client, Cfg: &config.Config{
		Server: config.Server{
			AccessSecret:    "s",
			AccessLifetime:  1,
			RefreshSecret:   "r",
			RefreshLifetime: 2,
		},
	}}

	testTable := []struct {
		name    string
		wantErr bool
		userId  string
		isAdmin bool
	}{
		{
			name:    "default",
			wantErr: false,
			userId:  "id1",
			isAdmin: true,
		},
		{
			name:    "default",
			wantErr: true,
			userId:  "",
			isAdmin: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, _, err := r.CreateAccessToken(testCase.userId, testCase.isAdmin)

			if testCase.wantErr == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRepository_CreateRefreshToken(t *testing.T) {
	client, _ := redismock.NewClientMock()
	defer client.Close()

	r := Repository{DB: client, Cfg: &config.Config{
		Server: config.Server{
			AccessSecret:    "s",
			AccessLifetime:  1,
			RefreshSecret:   "r",
			RefreshLifetime: 2,
		},
	}}

	testTable := []struct {
		name    string
		wantErr bool
		isAdmin bool
	}{
		{
			name:    "default",
			wantErr: false,
			isAdmin: true,
		},
		{
			name:    "default",
			wantErr: false,
			isAdmin: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, _, err := r.CreateRefreshToken(testCase.isAdmin)

			if testCase.wantErr == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRepository_VerifyRefreshToken(t *testing.T) {
	client, mock := redismock.NewClientMock()
	defer client.Close()

	r := Repository{DB: client, Cfg: &config.Config{
		Server: config.Server{
			AccessSecret:    "secret",
			AccessLifetime:  1,
			RefreshSecret:   "secret",
			RefreshLifetime: 10000,
		},
	}}

	rtoken, _, _ := r.CreateRefreshToken(true)

	type mockBehaviour func(userId string, token string)

	testTable := []struct {
		name          string
		wantErr       bool
		userId        string
		token         string
		mockBehaviour mockBehaviour
	}{
		{
			name:    "no userId",
			wantErr: true,
			userId:  "",
			mockBehaviour: func(userId string, token string) {

			},
			token: rtoken,
		},
		{
			name:    "no token",
			wantErr: true,
			userId:  "nmae",
			mockBehaviour: func(userId string, token string) {

			},
			token: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := r.VerifyRefreshToken(testCase.userId, testCase.token)

			testCase.mockBehaviour(testCase.userId, testCase.token)

			if testCase.wantErr == true {
				assert.Error(t, err)
			} else {
				if err = mock.ExpectationsWereMet(); err != nil {
					t.Fatal(err)
				}
				assert.NoError(t, err)
			}
		})
	}
}

func TestRepository_UpdateAccessToken(t *testing.T) {
	client, mock := redismock.NewClientMock()
	defer client.Close()

	r := Repository{DB: client, Cfg: &config.Config{
		Server: config.Server{
			AccessSecret:    "secret",
			AccessLifetime:  1,
			RefreshSecret:   "secret",
			RefreshLifetime: 10000,
		},
	}}

	rtoken, _, _ := r.CreateRefreshToken(true)

	type mockBehaviour func(userId string, token string)

	testTable := []struct {
		name          string
		wantErr       bool
		userId        string
		token         string
		mockBehaviour mockBehaviour
	}{
		{
			name:    "no userId",
			wantErr: true,
			userId:  "",
			mockBehaviour: func(userId string, token string) {

			},
			token: rtoken,
		},
		{
			name:    "no token",
			wantErr: true,
			userId:  "nmae",
			mockBehaviour: func(userId string, token string) {

			},
			token: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := r.VerifyRefreshToken(testCase.userId, testCase.token)

			testCase.mockBehaviour(testCase.userId, testCase.token)

			if testCase.wantErr == true {
				assert.Error(t, err)
			} else {
				if err = mock.ExpectationsWereMet(); err != nil {
					t.Fatal(err)
				}
				assert.NoError(t, err)
			}
		})
	}
}

func TestRepository_CreateTokensPair(t *testing.T) {
	client, mock := redismock.NewClientMock()
	defer client.Close()

	r := Repository{DB: client, Cfg: &config.Config{
		Server: config.Server{
			AccessSecret:    "secret",
			AccessLifetime:  1,
			RefreshSecret:   "secret",
			RefreshLifetime: 10000,
		},
	}}

	type mockBehaviour func(userId string, isAdmin bool)

	testTable := []struct {
		name          string
		wantErr       bool
		userId        string
		isAdmin       bool
		mockBehaviour mockBehaviour
	}{
		{
			name:    "default",
			wantErr: true,
			userId:  "userid",
			mockBehaviour: func(userId string, isAdmin bool) {
				mock.ExpectSet(fmt.Sprintf("refresh-token-%s", userId), "", 0)
			},
			isAdmin: true,
		},
		{
			name:    "default",
			wantErr: true,
			userId:  "",
			mockBehaviour: func(userId string, isAdmin bool) {
				mock.ExpectSet(fmt.Sprintf("refresh-token-%s", userId), "", 0)
			},
			isAdmin: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, _, _, err := r.CreateTokensPair(testCase.userId, testCase.isAdmin)

			testCase.mockBehaviour(testCase.userId, testCase.isAdmin)

			if testCase.wantErr == true {
				assert.Error(t, err)
			} else {
				if err = mock.ExpectationsWereMet(); err != nil {
					t.Fatal(err)
				}
				assert.NoError(t, err)
			}
		})
	}
}
