// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/MostafaGamal/sdk/internal/utils"
	"time"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "approved":
		fallthrough
	case "rejected":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// KYCStatus - KYB status retrieved successfully
type KYCStatus struct {
	Status           *Status    `json:"status,omitempty"`
	SubmissionDate   *time.Time `json:"submissionDate,omitempty"`
	VerificationDate *time.Time `json:"verificationDate,omitempty"`
	Comments         *string    `json:"comments,omitempty"`
}

func (k KYCStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KYCStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KYCStatus) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *KYCStatus) GetSubmissionDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.SubmissionDate
}

func (o *KYCStatus) GetVerificationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.VerificationDate
}

func (o *KYCStatus) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}
