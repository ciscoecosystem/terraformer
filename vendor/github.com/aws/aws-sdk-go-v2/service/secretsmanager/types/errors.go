// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Secrets Manager can't decrypt the protected secret text using the provided KMS
// key.
type DecryptionFailure struct {
	Message *string
}

func (e *DecryptionFailure) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DecryptionFailure) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DecryptionFailure) ErrorCode() string             { return "DecryptionFailure" }
func (e *DecryptionFailure) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Secrets Manager can't encrypt the protected secret text using the provided KMS
// key. Check that the customer master key (CMK) is available, enabled, and not in
// an invalid state. For more information, see How Key State Affects Use of a
// Customer Master Key
// (http://docs.aws.amazon.com/kms/latest/developerguide/key-state.html).
type EncryptionFailure struct {
	Message *string
}

func (e *EncryptionFailure) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EncryptionFailure) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EncryptionFailure) ErrorCode() string             { return "EncryptionFailure" }
func (e *EncryptionFailure) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred on the server side.
type InternalServiceError struct {
	Message *string
}

func (e *InternalServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceError) ErrorCode() string             { return "InternalServiceError" }
func (e *InternalServiceError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// You provided an invalid NextToken value.
type InvalidNextTokenException struct {
	Message *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You provided an invalid value for a parameter.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You provided a parameter value that is not valid for the current state of the
// resource. Possible causes:
//
// * You tried to perform the operation on a secret
// that's currently marked deleted.
//
// * You tried to enable rotation on a secret
// that doesn't already have a Lambda function ARN configured and you didn't
// include such an ARN as a parameter in this call.
type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because it would exceed one of the Secrets Manager internal
// limits.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You provided a resource-based policy with syntax errors.
type MalformedPolicyDocumentException struct {
	Message *string
}

func (e *MalformedPolicyDocumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedPolicyDocumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedPolicyDocumentException) ErrorCode() string {
	return "MalformedPolicyDocumentException"
}
func (e *MalformedPolicyDocumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request failed because you did not complete all the prerequisite steps.
type PreconditionNotMetException struct {
	Message *string
}

func (e *PreconditionNotMetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionNotMetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionNotMetException) ErrorCode() string             { return "PreconditionNotMetException" }
func (e *PreconditionNotMetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The BlockPublicPolicy parameter is set to true and the resource policy did not
// prevent broad access to the secret.
type PublicPolicyException struct {
	Message *string
}

func (e *PublicPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PublicPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PublicPolicyException) ErrorCode() string             { return "PublicPolicyException" }
func (e *PublicPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource with the ID you requested already exists.
type ResourceExistsException struct {
	Message *string
}

func (e *ResourceExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceExistsException) ErrorCode() string             { return "ResourceExistsException" }
func (e *ResourceExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// We can't find the resource that you asked for.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
