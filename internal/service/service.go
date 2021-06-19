package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/odisme0804/go-feature-toggler/lib/toggler"
)

type server struct {
	taggler *toggler.GoflagrToggler
	logger  log.Logger
}

func NewServer(log log.Logger) *server {
	return &server{
		taggler: toggler.NewGoflagrToggler(),
		logger:  log,
	}
}

type header struct {
	MemberID string `header:"X-MEMBER-ID"`
}

func (s *server) SimpleBooleanFlagHandler(ctx *gin.Context) {
	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	enable, err := s.taggler.IsEnable(ctx, toggler.Entity{
		ID:      h.MemberID,
		FlagKey: "test.SimpleBooleanFlag",
	})

	if err != nil {
		s.logger.Log(fmt.Printf("taggler.Evaluation err: %+v\n ", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"SimpleBooleanFlagStatus": enable,
	})
	return
}

func (s *server) MultiVariantsHandler(ctx *gin.Context) {
	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	enable, err := s.taggler.IsEnable(ctx, toggler.Entity{
		ID:      h.MemberID,
		FlagKey: "test.SimpleBooleanFlag",
	})

	if err != nil {
		s.logger.Log(fmt.Printf("taggler.Evaluation err: %+v\n ", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"SimpleBooleanFlagStatus": enable,
	})
	return
}
func (s *server) VariantAttachmentHandler(ctx *gin.Context) {
	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	enable, err := s.taggler.IsEnable(ctx, toggler.Entity{
		ID:      h.MemberID,
		FlagKey: "test.SimpleBooleanFlag",
	})

	if err != nil {
		s.logger.Log(fmt.Printf("taggler.Evaluation err: %+v\n ", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"SimpleBooleanFlagStatus": enable,
	})
	return
}
func (s *server) CustomConstraintHandler(ctx *gin.Context) {
	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	enable, err := s.taggler.IsEnable(ctx, toggler.Entity{
		ID:      h.MemberID,
		FlagKey: "test.SimpleBooleanFlag",
	})

	if err != nil {
		s.logger.Log(fmt.Printf("taggler.Evaluation err: %+v\n ", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"SimpleBooleanFlagStatus": enable,
	})
	return
}
