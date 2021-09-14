package s1ap

// 3GPP TS 24.301 9.9.3.12
func GetTypeOfId(id int) string {
	var eps_type_of_identity = map[int]string{
		0: "",
		1: "IMSI",
		2: "",
		3: "IMEI",
		4: "",
		5: "",
		6: "GUTI",
		7: "",
	}

	if id != 0 {
		return "Attach type: " + eps_type_of_identity[id] + "\\n"
		
	}
	return ""
}

// 3GPP TS 24.301 
func GetDcnr(id int) string {
	var nas_epc_emm_dcnr_cap = map[int]string{
		0: "not supported",
		1: "supported",
	}

	if id == 1 {
		return "DCNR: " + nas_epc_emm_dcnr_cap[id] + "\\n"
	}
	
	return ""
}
