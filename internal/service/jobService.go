package service

import (
	"context"
	"errors"
	"fmt"
	"job-portal-api/internal/models"
	"strconv"
	"sync"

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

func (s *Service) ProcessJobApplication(ctx context.Context, jobData []models.JobApplicantResponse) ([]models.JobApplicantResponse, error) {
	var ProccessedJobData []models.JobApplicantResponse
	jobDetails, err := s.UserRepo.ViewJobDetailsByJobId(ctx, 15)
	fmt.Println("hello", jobDetails.JobLocation)

	if err != nil {
		return nil, errors.New("failed to fetch job details from database")
	}

	ch := make(chan models.JobApplicantResponse)
	wg := new(sync.WaitGroup)

	for _, v := range jobData {
		wg.Add(1)
		go func(v models.JobApplicantResponse) {
			defer wg.Done()
			b, _ := applicationFilter(v, jobDetails)
			if b {
				ch <- v
			}
		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for data := range ch {
		ProccessedJobData = append(ProccessedJobData, data)
	}
	return ProccessedJobData, nil
}

func applicationFilter(validateApplication models.JobApplicantResponse, jobDetails models.Jobs) (bool, models.JobApplicantResponse) {

	applicantBudget, err := strconv.Atoi(validateApplication.Jobs.Budget)
	if err != nil {
		panic("error while conversion budget data from applicants")
	}
	compBudget, err := strconv.Atoi(jobDetails.Budget)
	if err != nil {
		panic("error while conversion budget data from posting")
	}
	if applicantBudget > compBudget {
		fmt.Println("failed in budget")
		return false, models.JobApplicantResponse{}

	}
	compMinNoticePeriod, err := strconv.Atoi(jobDetails.MinNoticePeriod)
	fmt.Println(compMinNoticePeriod)
	if err != nil {
		panic("error while conversion min notice  period data from hr posting")
	}
	compMaxNoticePeriod, err := strconv.Atoi(jobDetails.MaxNoticePeriod)
	fmt.Println(compMaxNoticePeriod)
	if err != nil {
		panic("error while conversion max notice period data from hr posting")
	}
	fmt.Println(validateApplication.Jobs.NoticePeriod)
	applicantNoticePeriod, err := strconv.Atoi(validateApplication.Jobs.NoticePeriod)
	fmt.Println(applicantNoticePeriod)
	if err != nil {
		panic("error while conversion notice period from applicant")
	}

	if (applicantNoticePeriod < compMinNoticePeriod) || (applicantNoticePeriod > compMaxNoticePeriod) {
		fmt.Println("failed in notice")

		return false, models.JobApplicantResponse{}
	}
	if validateApplication.Jobs.JobDescription != jobDetails.JobDescription {
		fmt.Println("failed in descrpitoim")

		return false, models.JobApplicantResponse{}
	}

	count := 0
	fmt.Println(validateApplication.Jobs.JobLocation)
	for _, v1 := range validateApplication.Jobs.JobLocation {
		count = 0
		for _, v2 := range jobDetails.JobLocation {
			fmt.Println(v2.Name)
			if v1 == v2.ID {
				count++

			}
		}
	}
	if count == 0 {
		fmt.Println("failed location")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.JobType {
		count = 0
		for _, v2 := range jobDetails.JobType {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed Jobtype")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.Qualification {
		count = 0
		for _, v2 := range jobDetails.Qualification {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed qualification")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.Shift {
		count = 0
		for _, v2 := range jobDetails.Shift {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed shift")
		return false, models.JobApplicantResponse{}
	}

	count = 0
	for _, v1 := range validateApplication.Jobs.Technology {
		count = 0
		for _, v2 := range jobDetails.Technology {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed technology")
		return false, models.JobApplicantResponse{}
	}
	count = 0
	for _, v1 := range validateApplication.Jobs.WorkMode {
		count = 0
		for _, v2 := range jobDetails.WorkMode {
			if v1 == v2.ID {
				count++
			}

		}
	}
	if count == 0 {
		fmt.Println("failed workmode")
		return false, models.JobApplicantResponse{}
	}

	return true, validateApplication
}

// if (validateApplication == models.JobApplicantResponse{}) {
// 	log.Error().Err(errors.New("no candidates meet requirments"))
// 	return false, models.JobApplicantResponse{}
// }
