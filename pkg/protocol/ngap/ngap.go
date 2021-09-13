package ngap

// 3GPP TS 24.501 9.11.3.2 5GMM cause
// nas_5gs.mm.5gmm_cause
func GetCause(id int) string {
	var fgs_emm_cause = map[int]string{
		3:   "Illegal UE",
		5:   "PEI not accepted",
		6:   "Illegal ME",
		7:   "5GS services not allowed",
		9:   "UE identity cannot be derived by the network",
		10:  "Implicitly de-registered",
		11:  "PLMN not allowed",
		12:  "Tracking area not allowed",
		13:  "Roaming not allowed in this tracking area",
		15:  "No suitable cells in tracking area",
		20:  "MAC failure",
		21:  "Synch failure",
		22:  "Congestion",
		23:  "UE security capabilities mismatch",
		24:  "Security mode rejected, unspecified",
		26:  "Non-5G authentication unacceptable",
		27:  "N1 mode not allowed",
		28:  "Restricted service area",
		31:  "Redirection to EPC required",
		43:  "LADN not available",
		62:  "No network slices available",
		65:  "Maximum number of PDU sessions reached",
		67:  "Insufficient resources for specific slice and DNN",
		69:  "Insufficient resources for specific slice",
		71:  "ngKSI already in use",
		72:  "Non-3GPP access to 5GCN not allowed",
		73:  "Serving network not authorized",
		74:  "Temporarily not authorized for this SNPN",
		75:  "Permanently not authorized for this SNPN",
		76:  "Not authorized for this CAG or authorized for CAG cells only",
		77:  "Wireline access area not allowed",
		78:  "PLMN not allowed to operate at the present UE location",
		79:  "UAS services not allowed",
		90:  "Payload was not forwarded",
		91:  "DNN not supported or not subscribed in the slice",
		92:  "Insufficient user-plane resources for the PDU session",
		95:  "Semantically incorrect message",
		96:  "Invalid mandatory information",
		97:  "Message type non-existent or not implemented",
		98:  "Message type not compatible with the protocol state",
		99:  "Information element non-existent or not implemented",
		100: "Conditional IE error",
		101: "Message not compatible with the protocol state",
		111: "Protocol error, unspecified",
	}

	if _, ok := fgs_emm_cause[id]; ok {
		return "Cause: " + fgs_emm_cause[id] + "\n"
	}
	return ""
}

// 3GPP TS 24.501 9.11.3.3 5GS identity type
// nas_5gs.mm.type_id
func GetTypeOfId(id int) string {
	var fgs_type_of_identity = map[int]string{
		1: "SUCI",
		2: "5G-GUTI",
		3: "IMEI",
		4: "5G-S-TMSI",
		5: "IMEISV",
		6: "MAC address",
		7: "EUI-64",
	}

	if id != 0 {
		return "Attach type: " + fgs_type_of_identity[id] + "\n"
	}
	return ""
}

// nas_5gs.mm.switch_off

// nas_5gs.mm.re_reg_req
