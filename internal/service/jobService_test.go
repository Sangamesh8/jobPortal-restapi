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

func TestService_ViewJobById(t *testing.T) {
	type args struct {
		ctx context.Context
		jid uint64
	}
	tests := []struct {
		name string
		//s       *Service
		args             args
		want             models.Jobs
		wantErr          bool
		mockrepoResponse func() (models.Jobs, error)
	}{
		{
			name: "error from db",
			want: models.Jobs{},
			args: args{
				ctx: context.Background(),
				jid: 15,
			},
			wantErr: true,
			mockrepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("error from db")
			},
		},
		{
			name: "success",
			want: models.Jobs{
				Company: models.Company{
					Name: "wework",
				},
				CompanyID: 1,
				Name:      "SDE",
			},
			args: args{
				ctx: context.Background(),
				jid: 15,
			},
			wantErr: false,
			mockrepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Company: models.Company{
						Name: "wework",
					},
					CompanyID: 1,
					Name:      "SDE",
				}, nil
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockrepoResponse != nil {
				mockRepo.EXPECT().ViewJobDetailsByJobId(tt.args.ctx, tt.args.jid).Return(tt.mockrepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJobById(tt.args.ctx, tt.args.jid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx       context.Context
		jobData   models.Jobs
		CompanyID uint64
	}
	tests := []struct {
		name             string
		s                *Service
		args             args
		want             models.Jobs
		wantErr          bool
		mockRepoResponse func() (models.Jobs, error)
	}{
		{
			name: "error case",
			args: args{
				ctx:     context.Background(),
				jobData: models.Jobs{},
			},
			want:    models.Jobs{},
			wantErr: true,
			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{}, errors.New("test error")
			},
		},
		{
			name: "success case",
			args: args{
				ctx: context.Background(),
				jobData: models.Jobs{
					Name:  "Go Developer",
					Salary: "500000",
					CompanyID:    1,
				},
				CompanyID: 1,
			},
			want: models.Jobs{
				Name:  "Go Developer",
				Salary: "500000",
				CompanyID:    1,
			},

			wantErr: false,

			mockRepoResponse: func() (models.Jobs, error) {
				return models.Jobs{
					Name:  "Go Developerr",
					Salary: "500000",
					CompanyID:    1,
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(tt.args.ctx, tt.args.jobData).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData, tt.args.CompanyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewJobByCompanyID(t *testing.T) {
	type args struct {
		ctx context.Context
		cid uint64
	}
	tests := []struct {
		name             string
		s                *Service
		args             args
		want             []models.Jobs
		wantErr          bool
		mockRepoResponse func() ([]models.Jobs, error)
	}{
		{
			name: "error from repository",
			//s:    &Service{},
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return nil, errors.New("error from repository")
			},
		},
		{
			name: "success",
			//s:    &Service{},
			args: args{
				ctx: context.Background(),
				cid: 1,
			},
			want: []models.Jobs{
				{
					Company: models.Company{
						Name: "Example Company",
					},
					CompanyID: 1,
					Name:      "SDE",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return []models.Jobs{
					{
						Company: models.Company{
							Name: "Example Company",
						},
						CompanyID: 1,
						Name:      "SDE",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().FindJobByCompanyID(tt.args.ctx, tt.args.cid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewJobByCompanyID(tt.args.ctx, tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewJobByCompanyID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewJobByCompanyID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ViewAllJobs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		s                *Service
		args             args
		want             []models.Jobs
		wantErr          bool
		mockRepoResponse func() ([]models.Jobs, error)
	}{
		{
			name: "error from db",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return nil, errors.New("error from db")
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Jobs{
				{
					Company: models.Company{
						Name: "wework",
					},
					CompanyID: 1,
					Name:      "SDE",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Jobs, error) {
				return []models.Jobs{
					{
						Company: models.Company{
							Name: "wework",
						},
						CompanyID: 1,
						Name:      "SDE",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().FindAllJobs(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.ViewAllJobs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.ViewAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.ViewAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
