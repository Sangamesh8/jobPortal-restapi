package service

import (
	"context"

	"job-portal-api/internal/models"
)

func (s *Service) AddJobDetails(ctx context.Context, jobData models.Jobs, CompanyID uint64) (models.Jobs, error) {
	jobData.CompanyID = uint(CompanyID)
	jobData, err := s.UserRepo.CreateJob(ctx, jobData)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {

	jobData, err := s.UserRepo.ViewJobDetailsByJobId(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil

}

func (s *Service) ViewJobByCompanyID(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FindJobByCompanyID(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
// func (s *Service) ApplicationsForJob(ctx context.Context, applicationData models.Jobs) ([]list, error) {
// 	companyData, err := s.UserRepo.CreateCompany(ctx, applicationData)
// 	if err != nil {
// 		return models.Company{}, err
// 	}
// 	return companyData, nil
// }
