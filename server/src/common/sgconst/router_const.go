package sgconst

const (
	_ int8 = iota
	ROUTER_STATUS_NORMAL_IN
	ROUTER_STATUS_NORMAL_OUT
	ROUTER_STATUS_LATE_IN
	ROUTER_STATUS_EARLY_OUT
	ROUTER_STATUS_EXCEPTION_ONLY_IN
	ROUTER_STATUS_EXCEPTION_ONLY_OUT
)

func GetRouterStatusString(status int8) string {
	switch status {
	case ROUTER_STATUS_NORMAL_IN:
		return "已入站"
	case ROUTER_STATUS_NORMAL_OUT:
		return "已出站"
	case ROUTER_STATUS_LATE_IN:
		return "推后入站"
	case ROUTER_STATUS_EARLY_OUT:
		return "提前出站"
	}
	return ""
}