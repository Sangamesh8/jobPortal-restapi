package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
	Type     string `json:"type" validate:"required"`
}
type Jobs struct {
	gorm.Model
	Company         Company         `json:"-" gorm:"ForeignKey:company_id"`
	CompanyID       uint            `json:"company_id"`
	Name            string          `json:"title"`
	MinNoticePeriod string          `json:"min_notice_period"`
	MaxNoticePeriod string          `json:"max_notice_period"`
	Budget          string          `json:"budget"`
	JobLocation     []JobLocation   `gorm:"many2many:job_location;"`
	Technology      []Technology    `gorm:"many2many:technology;"`
	WorkMode        []WorkMode      `gorm:"many2many:workmode;"`
	JobDescription  string          `json : "jobdescription"`
	Qualification   []Qualification `gorm:"many2many:qualification;"`
	Shift           []Shift         `gorm:"many2many:shift;"`
	JobType         []JobType       `gorm:"many2many:jobtype;"`
}

type CreateJobs struct {
	CompanyID       uint            `json:"company_id"`
	Name            string          `json:"title"`
	MinNoticePeriod string          `json:"min_notice_period"`
	MaxNoticePeriod string          `json:"max_notice_period"`
	Budget          string          `json:"budget"`
	JobLocation     []JobLocation   `json : "job_location;"`
	Technology      []Technology    `json : "technology;"`
	WorkMode        []WorkMode      `json : "workmode;"`
	JobDescription  string          `json : "jobdescription"`
	Qualification   []Qualification `json : "qualification;"`
	Shift           []Shift         `json : "shift;"`
	JobType         []JobType       `json : "jobtype;"`
}

type JobApplicant struct {
	CompanyID       uint            `json:"company_id"`
	Name            string          `json:"title"`
	MinNoticePeriod string          `json:"min_notice_period"`
	MaxNoticePeriod string          `json:"max_notice_period"`
	Budget          string          `json:"budget"`
	JobLocation     []JobLocation   `json : "job_location;"`
	Technology      []Technology    `json : "technology;"`
	WorkMode        []WorkMode      `json : "workmode;"`
	JobDescription  string          `json : "jobdescription"`
	Qualification   []Qualification `json : "qualification;"`
	Shift           []Shift         `json : "shift;"`
	JobType         []JobType       `json : "jobtype;"`
}

type JobLocation struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Technology struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type WorkMode struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
type Qualification struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Shift struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type JobType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
