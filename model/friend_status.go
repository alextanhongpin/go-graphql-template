package model

// FriendStatus represents the enums for FriendStatus.
type FriendStatus string

// String returns the string enum.
func (f FriendStatus) String() string {
	return string(f)
}

const (
	FriendStatusFriend   FriendStatus = "FRIEND"
	FriendStatusBlocking              = "BLOCKING" // You block someone...
	FriendStatusBlocked               = "BLOCKED"  // Someone block you...
	FriendStatusPending               = "PENDING"
	FriendStatusSent                  = "SENT"
	FriendStatusNone                  = "NONE"
	FriendStatusInvalid               = "INVALID"
)
