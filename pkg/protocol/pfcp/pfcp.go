package pfcp

func GetCause(id int) string {
	var pfcp_cause = map[int]string{
		1:  "Request accepted (success)",
		2:  "More Usage Report to send",
		3:  "Request partially accepted",
		64: "Request rejected (reason not specified)",
		65: "Session context not found",
		66: "Mandatory IE missing",
		67: "Conditional IE missing",
		68: "Invalid length",
		69: "Mandatory IE incorrect",
		70: "Invalid Forwarding Policy",
		71: "Invalid F-TEID allocation option",
		72: "No established PFCP Association ",
		73: "Rule creation/modification Failure ",
		74: "PFCP entity in congestion",
		75: "No resources available",
		76: "Service not supported",
		77: "System failure",
		78: "Redirection Requested",
		79: "All dynamic addresses are occupied",
		80: "Unknown Pre-defined Rule",
		81: "Unknown Application ID",
		82: "L2TP tunnel Establishment failure",
		83: "L2TP session Establishment failure",
		84: "L2TP tunnel release",
		85: "L2TP session release",		
	}
	if id == 0 {
		return ""
	}
	return "Cause: " + pfcp_cause[id] + "\\n"
}

