package service

import (
	"context"
	"errors"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestService_UserSignIn(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name string
		//s                *Service
		args             args
		want             string
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		{
			name: "wrong input",
			//s:    &Service{},
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "snu",
					Email:    "",
					Password: "abcd",
				},
			},
			want:    "",
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("error during input")
			},
		},
		{
			name: "Successful input",
			//s:    &Service{},
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "suuu",
					Email:    "suuu@gmail.com",
					Password: "stw",
				},
			},
			want:    "",
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Username:     "suuu",
					Email:        "suuu@gmail.com",
					PasswordHash: "stw",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CheckEmail(tt.args.ctx, tt.args.userData.Email).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.UserSignIn(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.UserSignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UserSignup(t *testing.T) {
	type args struct {
		ctx      context.Context
		userData models.NewUser
	}
	tests := []struct {
		name    string
		//s       *Service
		args    args
		want    models.User
		wantErr bool
		mockRepoResponse func() (models.User, error)
	}{
		{
			name: "wrong input",
			//s:    &Service{},
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "snu",
					Password: "abcd",
				},
			},
			want:    models.User{},
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("insufficient input")
			},
		},
		{
			name: "Successful input",
			//s:    &Service{},
			args: args{
				ctx: context.Background(),
				userData: models.NewUser{
					Username: "suuu",
					Email:    "suuu@gmail.com",
					Password: "stw",
				},
			},
			want: models.User{
				Username:     "suuu",
				Email:        "suuu@gmail.com",
				PasswordHash: "stw",
			},
			wantErr: false,
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Username:     "suuu",
					Email:        "suuu@gmail.com",
					PasswordHash: "stw",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.UserSignup(tt.args.ctx, tt.args.userData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UserSignup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UserSignup() = %v, want %v", got, tt.want)
			}
		})
	}
}
