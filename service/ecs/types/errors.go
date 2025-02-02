// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have authorization to perform the requested action.
type AccessDeniedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can apply up to 10 custom attributes per resource. You can view the
// attributes of a resource with ListAttributes. You can remove existing attributes
// on a resource with DeleteAttributes.
type AttributeLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AttributeLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AttributeLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AttributeLimitExceededException) ErrorCode() string {
	return "AttributeLimitExceededException"
}
func (e *AttributeLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your Amazon Web Services account has been blocked. For more information, contact
// Amazon Web Services Support (http://aws.amazon.com/contact-us/).
type BlockedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BlockedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BlockedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BlockedException) ErrorCode() string             { return "BlockedException" }
func (e *BlockedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// These errors are usually caused by a client action, such as using an action or
// resource on behalf of a user that doesn't have permissions to use the action or
// resource, or specifying an identifier that is not valid.
type ClientException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientException) ErrorCode() string             { return "ClientException" }
func (e *ClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You cannot delete a cluster that has registered container instances. First,
// deregister the container instances before you can delete the cluster. For more
// information, see DeregisterContainerInstance.
type ClusterContainsContainerInstancesException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ClusterContainsContainerInstancesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterContainsContainerInstancesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterContainsContainerInstancesException) ErrorCode() string {
	return "ClusterContainsContainerInstancesException"
}
func (e *ClusterContainsContainerInstancesException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// You cannot delete a cluster that contains services. First, update the service to
// reduce its desired task count to 0 and then delete the service. For more
// information, see UpdateService and DeleteService.
type ClusterContainsServicesException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ClusterContainsServicesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterContainsServicesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterContainsServicesException) ErrorCode() string {
	return "ClusterContainsServicesException"
}
func (e *ClusterContainsServicesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You cannot delete a cluster that has active tasks.
type ClusterContainsTasksException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ClusterContainsTasksException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterContainsTasksException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterContainsTasksException) ErrorCode() string             { return "ClusterContainsTasksException" }
func (e *ClusterContainsTasksException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified cluster could not be found. You can view your available clusters
// with ListClusters. Amazon ECS clusters are Region-specific.
type ClusterNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ClusterNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterNotFoundException) ErrorCode() string             { return "ClusterNotFoundException" }
func (e *ClusterNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified parameter is invalid. Review the available parameters for the API
// request.
type InvalidParameterException struct {
	Message *string

	noSmithyDocumentSerde
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

// The limit for the resource has been exceeded.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
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

// Amazon ECS is unable to determine the current version of the Amazon ECS
// container agent on the container instance and does not have enough information
// to proceed with an update. This could be because the agent running on the
// container instance is an older or custom version that does not use our version
// information.
type MissingVersionException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MissingVersionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingVersionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingVersionException) ErrorCode() string             { return "MissingVersionException" }
func (e *MissingVersionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is no update available for this Amazon ECS container agent. This could be
// because the agent is already running the latest version, or it is so old that
// there is no update path to the current version.
type NoUpdateAvailableException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NoUpdateAvailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoUpdateAvailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoUpdateAvailableException) ErrorCode() string             { return "NoUpdateAvailableException" }
func (e *NoUpdateAvailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified platform version does not satisfy the task definition's required
// capabilities.
type PlatformTaskDefinitionIncompatibilityException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *PlatformTaskDefinitionIncompatibilityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PlatformTaskDefinitionIncompatibilityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PlatformTaskDefinitionIncompatibilityException) ErrorCode() string {
	return "PlatformTaskDefinitionIncompatibilityException"
}
func (e *PlatformTaskDefinitionIncompatibilityException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified platform version does not exist.
type PlatformUnknownException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *PlatformUnknownException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PlatformUnknownException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PlatformUnknownException) ErrorCode() string             { return "PlatformUnknownException" }
func (e *PlatformUnknownException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource is in-use and cannot be removed.
type ResourceInUseException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource could not be found.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
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

// These errors are usually caused by a server issue.
type ServerException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServerException) ErrorCode() string             { return "ServerException" }
func (e *ServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The specified service is not active. You can't update a service that is
// inactive. If you have previously deleted a service, you can re-create it with
// CreateService.
type ServiceNotActiveException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceNotActiveException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceNotActiveException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceNotActiveException) ErrorCode() string             { return "ServiceNotActiveException" }
func (e *ServiceNotActiveException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified service could not be found. You can view your available services
// with ListServices. Amazon ECS services are cluster-specific and Region-specific.
type ServiceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceNotFoundException) ErrorCode() string             { return "ServiceNotFoundException" }
func (e *ServiceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The target container is not properly configured with the execute command agent
// or the container is no longer active or running.
type TargetNotConnectedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TargetNotConnectedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TargetNotConnectedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TargetNotConnectedException) ErrorCode() string             { return "TargetNotConnectedException" }
func (e *TargetNotConnectedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified target could not be found. You can view your available container
// instances with ListContainerInstances. Amazon ECS container instances are
// cluster-specific and Region-specific.
type TargetNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TargetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TargetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TargetNotFoundException) ErrorCode() string             { return "TargetNotFoundException" }
func (e *TargetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified task set could not be found. You can view your available task sets
// with DescribeTaskSets. Task sets are specific to each cluster, service and
// Region.
type TaskSetNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TaskSetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TaskSetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TaskSetNotFoundException) ErrorCode() string             { return "TaskSetNotFoundException" }
func (e *TaskSetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified task is not supported in this Region.
type UnsupportedFeatureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnsupportedFeatureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedFeatureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedFeatureException) ErrorCode() string             { return "UnsupportedFeatureException" }
func (e *UnsupportedFeatureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is already a current Amazon ECS container agent update in progress on the
// specified container instance. If the container agent becomes disconnected while
// it is in a transitional stage, such as PENDING or STAGING, the update process
// can get stuck in that state. However, when the agent reconnects, it resumes
// where it stopped previously.
type UpdateInProgressException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UpdateInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UpdateInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UpdateInProgressException) ErrorCode() string             { return "UpdateInProgressException" }
func (e *UpdateInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
