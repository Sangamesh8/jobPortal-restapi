package service

import (
	"context"
	"job-portal-api/internal/models"

	"gorm.io/gorm"
)

func (s *Service) AddJobDetails(ctx context.Context, jobData models.CreateJobs, CompanyID uint64) (models.ResponseForJobs, error) {
	job := models.Jobs{
		CompanyID:       uint(CompanyID),
		Name:            jobData.Name,
		MinNoticePeriod: jobData.MinNoticePeriod,
		MaxNoticePeriod: jobData.MaxNoticePeriod,
		Budget:          jobData.Budget,
		JobDescription:  jobData.JobDescription,
	}

	for _, v := range jobData.JobLocation {
		jobLoc := models.JobLocation{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.JobLocation = append(job.JobLocation, jobLoc)
	}
	for _, v := range jobData.Technology {
		jobTec := models.Technology{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Technology = append(job.Technology, jobTec)
	}
	for _, v := range jobData.Qualification {
		jobQual := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Qualification = append(job.Qualification, jobQual)
	}
	for _, v := range jobData.Shift {
		jobShift := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.Shift = append(job.Shift, jobShift)
	}
	for _, v := range jobData.JobType {
		jobJobType := models.JobType{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.JobType = append(job.JobType, jobJobType)
	}
	for _, v := range jobData.WorkMode {
		jobworkMode := models.WorkMode{
			Model: gorm.Model{
				ID: v,
			},
		}
		job.WorkMode = append(job.WorkMode, jobworkMode)
	}
	createdJob, err := s.UserRepo.CreateJob(ctx, job)
	if err != nil {
		return models.ResponseForJobs{}, err
	}
	return createdJob, nil
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

func ProcessJobApplication(ctx context.Context, jobData []models.JobApplicant) ([]models.JobApplicant, error) {
}

func (s *Service) compareAndCheck(jobData models.JobApplicant) (bool, models.JobApplicant, error) {

}
