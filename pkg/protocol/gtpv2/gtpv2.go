package gtpv2

// 3GPP TS 29.274 8.4
func GetCause(id int) string {
	var gtpv2_cause = map[int]string{
		0:   "Reserved",
		1:   "Reserved",
		2:   "Local Detach",
		3:   "Complete Detach",
		4:   "RAT changed from 3GPP to Non-3GPP",
		5:   "ISR deactivation",
		6:   "Error Indication received from RNC/eNodeB/S4-SGSN/MME",
		7:   "IMSI Detach Only",
		8:   "Reactivation Requested",
		9:   "PDN reconnection to this APN disallowed",
		10:  "Access changed from Non-3GPP to 3GPP",
		11:  "PDN connection inactivity timer expires",
		12:  "PGW not responding",
		13:  "Network Failure",
		14:  "QoS parameter mismatch",
		15:  "EPS to 5GS Mobility",
		16:  "Request accepted",
		17:  "Request accepted partially",
		18:  "New PDN type due to network preference",
		19:  "New PDN type due to single address bearer only",
		64:  "Context Not Found",
		65:  "Invalid Message Format",
		66:  "Version not supported by next peer",
		67:  "Invalid length",
		68:  "Service not supported",
		69:  "Mandatory IE incorrect",
		70:  "Mandatory IE missing",
		71:  "Shall not be used",
		72:  "System failure",
		73:  "No resources available",
		74:  "Semantic error in the TFT operation",
		75:  "Syntactic error in the TFT operation",
		76:  "Semantic errors in packet filter(s)",
		77:  "Syntactic errors in packet filter(s)",
		78:  "Missing or unknown APN",
		79:  "Shall not be used",
		80:  "GRE key not found",
		81:  "Relocation failure",
		82:  "Denied in RAT",
		83:  "Preferred PDN type not supported",
		84:  "All dynamic addresses are occupied",
		85:  "UE context without TFT already activated",
		86:  "Protocol type not supported",
		87:  "UE not responding",
		88:  "UE refuses",
		89:  "Service denied",
		90:  "Unable to page UE",
		91:  "No memory available",
		92:  "User authentication failed",
		93:  "APN access denied – no subscription",
		94:  "Request rejected (reason not specified)",
		95:  "P-TMSI Signature mismatch",
		96:  "IMSI/IMEI not known",
		97:  "Semantic error in the TAD operation",
		98:  "Syntactic error in the TAD operation",
		99:  "Shall not be used",
		100: "Remote peer not responding",
		101: "Collision with network initiated request",
		102: "Unable to page UE due to Suspension",
		103: "Conditional IE missing",
		104: "APN Restriction type Incompatible with currently active PDN connection",
		105: "Invalid overall length of the triggered response message and a piggybacked initial message", 
		106: "Data forwarding not supported",
		107: "Invalid reply from remote peer",
		108: "Fallback to GTPv1",
		109: "Invalid peer",
		110: "Temporarily rejected due to handover/TAU/RAU procedure in progress",
		111: "Modifications not limited to S1-U bearers",
		112: "Request rejected for a PMIPv6 reason",
		113: "APN Congestion",
		114: "Bearer handling not supported",
		115: "UE already re-attached",
		116: "Multiple PDN connections for a given APN not allowed",
		117: "Target access restricted for the subscriber",
		118: "Shall not be used",
		119: "MME/SGSN refuses due to VPLMN Policy",
		120: "GTP-C Entity Congestion",
		121: "Late Overlapping Request ",
		122: "Timed out Request",
		123: "UE is temporarily not reachable due to power saving",
		124: "Relocation failure due to NAS message redirection",
		125: "UE not authorised by OCS or external AAA Server",
		126: "Multiple accesses to a PDN connection not allowed",
		127: "Request rejected due to UE capability",
		128: "S1-U Path Failure",
		129: "5GC not allowed",
		130: "PGW mismatch with network slice subscribed by the UE",
		131: "Rejection due to paging restriction",
	}

	if id == 0 || id == 1 {
		return ""
	}

	if _, ok := gtpv2_cause[id]; ok {
		return "Cause: " + gtpv2_cause[id] + "\n"
	}
	return ""
}

// 3GPP TS 29.274
func GetDcnr(id int) string {
	var gtpv2_dcnr = map[int]string{
		0: "not supported",
		1: "supported",
	}

	if id == 1 {
		return "DCNR: " + gtpv2_dcnr[id] + "\n"
	}
	
	return ""
}

// 3GPP TS 29.274
func GetOiIndication(id int) string {
	var gtpv2_oi = map[int]string{
		0: "not supported",
		1: "supported",
	}

	if id == 1 {
		return "OI: " + gtpv2_oi[id] + "\n"
	}
	
	return ""
}

// 3GPP TS 29.274
func GetSiIndication(id int) string {
	var gtpv2_si = map[int]string{
		0: "not supported",
		1: "supported",
	}

	if id == 1 {
		return "OI: " + gtpv2_si[id] + "\n"
	}
	
	return ""
}
