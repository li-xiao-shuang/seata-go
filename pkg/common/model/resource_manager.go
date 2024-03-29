package model

import (
	"sync"
)

// Control a branch transaction commit or rollback
type ResourceManagerInbound interface {
	// Commit a branch transaction
	BranchCommit(branchType BranchType, xid string, branchId int64, resourceId string, applicationData []byte) (BranchStatus, error)
	// Rollback a branch transaction
	BranchRollback(branchType BranchType, xid string, branchId int64, resourceId string, applicationData []byte) (BranchStatus, error)
}

// Resource Manager: send outbound request to TC
type ResourceManagerOutbound interface {
	// Branch register long
	BranchRegister(branchType BranchType, resourceId, clientId, xid, applicationData, lockKeys string) (int64, error)
	//  Branch report
	BranchReport(branchType BranchType, xid string, branchId int64, status BranchStatus, applicationData string) error
	// Lock query boolean
	LockQuery(branchType BranchType, resourceId, xid, lockKeys string) (bool, error)
}

//  Resource Manager: common behaviors
type ResourceManager interface {
	ResourceManagerInbound
	ResourceManagerOutbound

	// Register a Resource to be managed by Resource Manager
	RegisterResource(resource Resource) error
	//  Unregister a Resource from the Resource Manager
	UnregisterResource(resource Resource) error
	// Get all resources managed by this manager
	GetManagedResources() sync.Map
	// Get the BranchType
	GetBranchType() BranchType
}

type ResourceManagerGetter interface {
	GetResourceManager() ResourceManager
}
