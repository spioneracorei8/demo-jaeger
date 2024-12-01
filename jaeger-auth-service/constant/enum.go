package constant

// -----------------
// # account_status status
// ----------------
type AccountStatus struct{}

const (
	ACTIVE   string = "ACTIVE"
	INACTIVE string = "INACTIVE"
)

var ACCOUNT_STATUS = struct {
	ACTIVE   string
	INACTIVE string
}{
	ACTIVE:   ACTIVE,
	INACTIVE: INACTIVE,
}

// -----------------
// # web_access status
// ----------------
type WebAccess struct{}

const (
	APPLICATION    string = "APPLICATION"
	WEB_MANAGEMENT string = "WEB_MANAGEMENT"
)

var WEB_ACCESS = struct {
	APPLICATION    string
	WEB_MANAGEMENT string
}{
	APPLICATION:    APPLICATION,
	WEB_MANAGEMENT: WEB_MANAGEMENT,
}
