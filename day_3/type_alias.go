package main
/*
An alias declaration doesn’t create a new distinct type different from the type it’s created from.
It just introduces an alias name T1, an alternate spelling, for the type denoted by T2.

Type aliases are not meant for everyday use. They were introduced to support gradual code repair while moving a type
between packages during large-scale refactoring. Codebase Refactoring (with help from Go) covers this in detail.
 */

type First func(int) int
type Second func(int) First

// TaskInstanceStatus type is used for status definition
type TaskInstanceStatus int

const (
	_ TaskInstanceStatus = iota
	TaskInstanceRunning
	TaskInstanceSuccess
	TaskInstanceFailed
	TaskInstanceScheduled
	TaskInstanceDisabled
	TaskInstanceSomeFailures
	TaskInstancePending
	TaskInstanceStopped
	TaskInstancePostponed
	TaskInstanceCanceled
)

const (
	TaskInstanceRunningText      = "Running"
	TaskInstanceSuccessText      = "Success"
	TaskInstanceFailedText       = "Failed"
	TaskInstanceScheduledText    = "Scheduled"
	TaskInstanceDisabledText     = "Disabled"
	TaskInstanceSomeFailuresText = "Some Failures"
	TaskInstancePendingText      = "Pending"
	TaskInstanceStoppedText      = "Stopped"
	TaskInstancePostponedText    = "Postponed"
	TaskInstanceCanceledText     = "Canceled"
)

// taskInstanceStatusText contains the text representation of the TaskInstanceStatus
var taskInstanceStatusText = map[TaskInstanceStatus]string{
	TaskInstanceStatus(0):    "",
	TaskInstanceRunning:      TaskInstanceRunningText,
	TaskInstanceSuccess:      TaskInstanceSuccessText,
	TaskInstanceFailed:       TaskInstanceFailedText,
	TaskInstanceScheduled:    TaskInstanceScheduledText,
	TaskInstanceDisabled:     TaskInstanceDisabledText,
	TaskInstanceSomeFailures: TaskInstanceSomeFailuresText,
	TaskInstancePending:      TaskInstancePendingText,
	TaskInstanceStopped:      TaskInstanceStoppedText,
	TaskInstancePostponed:    TaskInstancePostponedText,
	TaskInstanceCanceled:     TaskInstanceCanceledText,
}

// TaskInstanceStatusText converts a value of the TaskInstanceStatus to its string representation
// It returns an error if there is no such status.
func TaskInstanceStatusText(status TaskInstanceStatus) (statusText string) {
	statusText, ok := taskInstanceStatusText[status]
	if !ok {
		return ""
	}
	return
}
