package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (h *Handler) ViewJobByID(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	jid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	jobData, err := h.service.ViewJobById(ctx, jid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *Handler) ViewAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	jobDatas, err := h.service.ViewAllJobs(ctx)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobDatas)

}

func (h *Handler) ViewJobByCompanyID(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("cid")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.ViewJobByCompanyID(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *Handler) CreateJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("cid")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	var jobData models.Jobs

	err = json.NewDecoder(c.Request.Body).Decode(&jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name, location and field",
		})
		return
	}

	jobData, err = h.service.AddJobDetails(ctx, jobData, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

// func (h *Handler) JobApplication(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
// 	if !ok {
// 		log.Error().Msg("traceid missing from context")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error": http.StatusText(http.StatusInternalServerError),
// 		})
// 		return
// 	}

// 	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
// 	if !ok {
// 		log.Error().Str("Trace Id", traceid).Msg("login first")
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
// 		return
// 	}

// 	var companyData models.Company

// 	err := json.NewDecoder(c.Request.Body).Decode(&companyData)
// 	if err != nil {
// 		log.Error().Err(err).Str("trace id", traceid)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": "please provide valid name, location and field",
// 		})
// 		return
// 	}

// 	validate := validator.New()
// 	err = validate.Struct(companyData)
// 	if err != nil {
// 		log.Error().Err(err).Str("trace id", traceid)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": "please provide valid name, location and field",
// 		})
// 		return
// 	}

// 	applicationData, err = h.service.ApplicationsForJob(ctx, companyData)
// 	if err != nil {
// 		log.Error().Err(err).Str("trace id", traceid)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, applicationData)

// }
