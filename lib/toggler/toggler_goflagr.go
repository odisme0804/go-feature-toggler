package toggler

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/odisme0804/goflagr"
)

type GoflagrToggler struct {
	*goflagr.APIClient
}

func NewGoflagrToggler() *GoflagrToggler { /* fixme: use interface */
	cfg := goflagr.NewConfiguration()
	cfg.BasePath = "http://localhost:18000/api/v1"

	return &GoflagrToggler{
		APIClient: goflagr.NewAPIClient(cfg),
	}
}

type toggleResult struct {
	FlagID            int64        `json:"flagID,omitempty"`
	FlagKey           string       `json:"flagKey,omitempty"`
	FlagSnapshotID    int64        `json:"flagSnapshotID,omitempty"`
	SegmentID         int64        `json:"segmentID,omitempty"`
	VariantID         int64        `json:"variantID,omitempty"`
	VariantKey        string       `json:"variantKey,omitempty"`
	VariantAttachment *interface{} `json:"variantAttachment,omitempty"`
	Timestamp         string       `json:"timestamp,omitempty"`
}

func (s *GoflagrToggler) Evaluation(ctx context.Context, entity Entity) (toggleResult, error) {
	res, _, err := s.EvaluationApi.PostEvaluation(ctx, goflagr.EvalContext{
		EntityID: entity.ID,
		FlagKey:  entity.FlagKey,
	})

	if err != nil {
		return toggleResult{}, err
	}

	if res.VariantAttachment != nil {
		fmt.Printf("evalutation result: %+v\n", *res.VariantAttachment)
	}

	return toggleResult{
		FlagID:            res.FlagID,
		FlagKey:           res.FlagKey,
		FlagSnapshotID:    res.FlagSnapshotID,
		SegmentID:         res.SegmentID,
		VariantID:         res.VariantID,
		VariantKey:        res.VariantKey,
		VariantAttachment: res.VariantAttachment,
		Timestamp:         res.Timestamp,
	}, nil
}

func (s *GoflagrToggler) IsEnable(ctx context.Context, entity Entity) (bool, error) {
	res, _, err := s.EvaluationApi.PostEvaluation(ctx, goflagr.EvalContext{
		EntityID:      entity.ID,
		FlagKey:       entity.FlagKey,
		EntityContext: &entity.Payload,
		EnableDebug:   true,
	})

	if err != nil {
		return false, err
	}

	return res.VariantKey == "on", nil
}

func (s *GoflagrToggler) EvaluateAndLoad(ctx context.Context, entity Entity, val interface{}) (string, error) {
	res, _, err := s.EvaluationApi.PostEvaluation(ctx, goflagr.EvalContext{
		EntityID:      entity.ID,
		FlagKey:       entity.FlagKey,
		EntityContext: &entity.Payload,
	})

	if err != nil {
		return "", err
	}

	if res.VariantAttachment != nil {
		mapstructure.Decode(*res.VariantAttachment, &val)
	}

	return res.VariantKey, nil
}
