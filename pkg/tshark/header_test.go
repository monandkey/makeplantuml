package tshark

import (
	"testing"
	"local.packages/tshark"
)

var expOut = `
"1","2021-06-16 10:34:38.377646","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","InitialUEMessage, Registration request",,"0x527eec87",,"124",
"2","2021-06-16 10:34:38.592668","10.244.166.179",,,"50718",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=AUSF",,,"0x0000f32b","180",
"3","2021-06-16 10:34:38.592685","10.244.166.179",,,"50718",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=AUSF",,,"0x000063c3","180",
"4","2021-06-16 10:34:38.938723","10.244.166.129",,,"29510",,"10.244.166.179",,,"50718",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"5","2021-06-16 10:34:38.938750","10.103.54.119",,,"29510",,"10.244.166.179",,,"50718",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"6","2021-06-16 10:34:38.950965","10.244.166.179",,,"57206",,"10.103.99.231",,,"29509",,"HTTP2","HEADERS[3]: POST /nausf-auth/v1/ue-authentications",,,"0x0000209e","182",
"7","2021-06-16 10:34:38.950970","10.244.166.179",,,"57206",,"10.244.166.157",,,"29509",,"HTTP2","HEADERS[3]: POST /nausf-auth/v1/ue-authentications",,,"0x000063e1","182",
"8","2021-06-16 10:34:39.151550","10.244.166.157",,,"45510",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AUSF&service-names=nudm-ueau&target-nf-type=UDM",,,"0x0000f327","198",
"9","2021-06-16 10:34:39.151555","10.244.166.157",,,"45510",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AUSF&service-names=nudm-ueau&target-nf-type=UDM",,,"0x000063bf","198",
"10","2021-06-16 10:34:39.165688","10.244.166.129",,,"29510",,"10.244.166.157",,,"45510",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006361","104",
"11","2021-06-16 10:34:39.165706","10.103.54.119",,,"29510",,"10.244.166.157",,,"45510",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2c9","104",
"12","2021-06-16 10:34:39.170719","10.244.166.157",,,"50824",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000003/security-information/generate-auth-data",,,"0x0000dee7","194",
"13","2021-06-16 10:34:39.170726","10.244.166.157",,,"50824",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000003/security-information/generate-auth-data",,,"0x000063dd","194",
"14","2021-06-16 10:34:39.306955","10.244.166.157",,,"50826",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000003/security-information/generate-auth-data",,,"0x0000dee7","194",
"15","2021-06-16 10:34:39.306963","10.244.166.157",,,"50826",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000003/security-information/generate-auth-data",,,"0x000063dd","194",
"16","2021-06-16 10:34:39.358037","10.244.166.163",,,"33418",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x0000f31b","180",
"17","2021-06-16 10:34:39.358046","10.244.166.163",,,"33418",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x000063b3","180",
"18","2021-06-16 10:34:39.361171","10.244.166.129",,,"29510",,"10.244.166.163",,,"33418",,"HTTP2","HEADERS[3]: 200 OK",,,"0x00006367","104",
"19","2021-06-16 10:34:39.361189","10.103.54.119",,,"29510",,"10.244.166.163",,,"33418",,"HTTP2","HEADERS[3]: 200 OK",,,"0x0000f2cf","104",
"20","2021-06-16 10:34:39.364875","10.244.166.163",,,"35368",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000065c1","197",
"21","2021-06-16 10:34:39.364880","10.244.166.163",,,"35368",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000063cd","197",
"22","2021-06-16 10:34:39.527613","10.244.166.163",,,"35370",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000065c1","197",
"23","2021-06-16 10:34:39.527621","10.244.166.163",,,"35370",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000063cd","197",
"24","2021-06-16 10:34:39.685146","10.244.166.138",,,"29504",,"10.244.166.163",,,"35370",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006370","104",
"25","2021-06-16 10:34:39.685178","10.111.169.3",,,"29504",,"10.244.166.163",,,"35370",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006564","104",
"26","2021-06-16 10:34:39.704759","10.244.166.163",,,"35374",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000065db","223",
"27","2021-06-16 10:34:39.704764","10.244.166.163",,,"35374",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000063e7","223",
"28","2021-06-16 10:34:39.784969","10.244.166.138",,,"29504",,"10.244.166.163",,,"35374",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000636b","99",
"29","2021-06-16 10:34:39.784988","10.111.169.3",,,"29504",,"10.244.166.163",,,"35374",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000655f","99",
"30","2021-06-16 10:34:39.785929","10.244.166.163",,,"29503",,"10.244.166.157",,,"50826",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006383","104",
"31","2021-06-16 10:34:39.785945","10.109.34.53",,,"29503",,"10.244.166.157",,,"50826",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000de8d","104",
"32","2021-06-16 10:34:39.787154","10.244.166.157",,,"29509",,"10.244.166.179",,,"57206",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063d6","171",
"33","2021-06-16 10:34:39.787167","10.103.99.231",,,"29509",,"10.244.166.179",,,"57206",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x00002093","171",
"34","2021-06-16 10:34:39.803414","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","DownlinkNASTransport, Authentication request",,"0x8add0609",,"116",
"35","2021-06-16 10:34:39.813882","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport, Authentication failure (Synch failure)",,"0x8fcbc10e",,"132",
"36","2021-06-16 10:34:39.824763","10.244.166.179",,,"57224",,"10.103.99.231",,,"29509",,"HTTP2","HEADERS[3]: POST /nausf-auth/v1/ue-authentications",,,"0x0000209e","182",
"37","2021-06-16 10:34:39.824767","10.244.166.179",,,"57224",,"10.244.166.157",,,"29509",,"HTTP2","HEADERS[3]: POST /nausf-auth/v1/ue-authentications",,,"0x000063e1","182",
"38","2021-06-16 10:34:39.828697","10.244.166.157",,,"45528",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AUSF&service-names=nudm-ueau&target-nf-type=UDM",,,"0x0000f327","198",
"39","2021-06-16 10:34:39.828704","10.244.166.157",,,"45528",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AUSF&service-names=nudm-ueau&target-nf-type=UDM",,,"0x000063bf","198",
"40","2021-06-16 10:34:39.834288","10.244.166.129",,,"29510",,"10.244.166.157",,,"45528",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006361","104",
"41","2021-06-16 10:34:39.834300","10.103.54.119",,,"29510",,"10.244.166.157",,,"45528",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2c9","104",
"42","2021-06-16 10:34:39.839312","10.244.166.157",,,"50842",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000003/security-information/generate-auth-data",,,"0x0000dee7","194",
"43","2021-06-16 10:34:39.839316","10.244.166.157",,,"50842",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/suci-0-208-93-0000-0-0-0000000003/security-information/generate-auth-data",,,"0x000063dd","194",
"44","2021-06-16 10:34:39.843762","10.244.166.163",,,"35382",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000065c1","197",
"45","2021-06-16 10:34:39.843767","10.244.166.163",,,"35382",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x000063cd","197",
"46","2021-06-16 10:34:39.846444","10.244.166.138",,,"29504",,"10.244.166.163",,,"35382",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006370","104",
"47","2021-06-16 10:34:39.846453","10.111.169.3",,,"29504",,"10.244.166.163",,,"35382",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006564","104",
"48","2021-06-16 10:34:39.846933","10.244.166.163",,,"35382",,"10.111.169.3",,,"29504",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x00006573","119",
"49","2021-06-16 10:34:39.846950","10.244.166.163",,,"35382",,"10.244.166.138",,,"29504",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",,,"0x0000637f","119",
"50","2021-06-16 10:34:39.850508","10.244.166.138",,,"29504",,"10.244.166.163",,,"35382",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: 204 No Content",,,"0x00006348","64",
"51","2021-06-16 10:34:39.850516","10.111.169.3",,,"29504",,"10.244.166.163",,,"35382",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: 204 No Content",,,"0x0000653c","64",
"52","2021-06-16 10:34:39.850886","10.244.166.163",,,"29503",,"10.244.166.157",,,"50842",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006383","104",
"53","2021-06-16 10:34:39.850896","10.109.34.53",,,"29503",,"10.244.166.157",,,"50842",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000de8d","104",
"54","2021-06-16 10:34:39.852335","10.244.166.157",,,"29509",,"10.244.166.179",,,"57224",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063d6","171",
"55","2021-06-16 10:34:39.852345","10.103.99.231",,,"29509",,"10.244.166.179",,,"57224",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x00002093","171",
"56","2021-06-16 10:34:39.853275","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","DownlinkNASTransport, Authentication request",,"0x09d132f6",,"132",
"57","2021-06-16 10:34:39.853618","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport, Authentication response",,"0x76d28cb6",,"132",
"58","2021-06-16 10:34:39.855262","10.244.166.179",,,"57224",,"10.103.99.231",,,"29509",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: PUT /nausf-auth/v1/ue-authentications/suci-0-208-93-0000-0-0-0000000003/5g-aka-confirmation",,,"0x0000207e","150",
"59","2021-06-16 10:34:39.855271","10.244.166.179",,,"57224",,"10.244.166.157",,,"29509",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: PUT /nausf-auth/v1/ue-authentications/suci-0-208-93-0000-0-0-0000000003/5g-aka-confirmation",,,"0x000063c1","150",
"60","2021-06-16 10:34:39.860794","10.244.166.179",,,"57232",,"10.103.99.231",,,"29509",,"HTTP2","HEADERS[3]: PUT /nausf-auth/v1/ue-authentications/suci-0-208-93-0000-0-0-0000000003/5g-aka-confirmation",,,"0x000020ae","198",
"61","2021-06-16 10:34:39.860798","10.244.166.179",,,"57232",,"10.244.166.157",,,"29509",,"HTTP2","HEADERS[3]: PUT /nausf-auth/v1/ue-authentications/suci-0-208-93-0000-0-0-0000000003/5g-aka-confirmation",,,"0x000063f1","198",
"62","2021-06-16 10:34:39.864672","10.244.166.157",,,"50848",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/imsi-208930000000003/auth-events",,,"0x0000decc","167",
"63","2021-06-16 10:34:39.864676","10.244.166.157",,,"50848",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: POST /nudm-ueau/v1/imsi-208930000000003/auth-events",,,"0x000063c2","167",
"64","2021-06-16 10:34:39.867756","10.244.166.163",,,"35388",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-status",,,"0x000065cd","209",
"65","2021-06-16 10:34:39.867761","10.244.166.163",,,"35388",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-status",,,"0x000063d9","209",
"66","2021-06-16 10:34:39.931388","10.244.166.138",,,"29504",,"10.244.166.163",,,"35388",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000636b","99",
"67","2021-06-16 10:34:39.931401","10.111.169.3",,,"29504",,"10.244.166.163",,,"35388",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000655f","99",
"68","2021-06-16 10:34:39.931715","10.244.166.163",,,"29503",,"10.244.166.157",,,"50848",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x00006384","105",
"69","2021-06-16 10:34:39.931726","10.109.34.53",,,"29503",,"10.244.166.157",,,"50848",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x0000de8e","105",
"70","2021-06-16 10:34:39.932233","10.244.166.157",,,"29509",,"10.244.166.179",,,"57232",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006393","104",
"71","2021-06-16 10:34:39.932244","10.103.99.231",,,"29509",,"10.244.166.179",,,"57232",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00002050","104",
"72","2021-06-16 10:34:39.938052","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","DownlinkNASTransport, Security mode command",,"0xfefa8872",,"112",
"73","2021-06-16 10:34:39.939974","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport",,"0xbbc1c09f",,"168",
"74","2021-06-16 10:34:39.957549","10.244.166.179",,,"50752",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&supi=imsi-208930000000003&target-nf-type=UDM",,,"0x0000f33c","197",
"75","2021-06-16 10:34:39.957553","10.244.166.179",,,"50752",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&supi=imsi-208930000000003&target-nf-type=UDM",,,"0x000063d4","197",
"76","2021-06-16 10:34:39.965327","10.244.166.129",,,"29510",,"10.244.166.179",,,"50752",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"77","2021-06-16 10:34:39.965356","10.103.54.119",,,"29510",,"10.244.166.179",,,"50752",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"78","2021-06-16 10:34:39.969261","10.244.166.179",,,"34946",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: GET /nudm-sdm/v1/imsi-208930000000003/nssai?plmn-id=20893",,,"0x0000dee3","168",
"79","2021-06-16 10:34:39.969264","10.244.166.179",,,"34946",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: GET /nudm-sdm/v1/imsi-208930000000003/nssai?plmn-id=20893",,,"0x000063d9","168",
"80","2021-06-16 10:34:39.973880","10.244.166.163",,,"35394",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/am-data?supported-features=",,,"0x000065c5","201",
"81","2021-06-16 10:34:39.973891","10.244.166.163",,,"35394",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/am-data?supported-features=",,,"0x000063d1","201",
"82","2021-06-16 10:34:39.978195","10.244.166.138",,,"29504",,"10.244.166.163",,,"35394",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006370","104",
"83","2021-06-16 10:34:39.978220","10.111.169.3",,,"29504",,"10.244.166.163",,,"35394",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006564","104",
"84","2021-06-16 10:34:39.978811","10.244.166.163",,,"29503",,"10.244.166.179",,,"34946",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006399","104",
"85","2021-06-16 10:34:39.978827","10.109.34.53",,,"29503",,"10.244.166.179",,,"34946",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000dea3","104",
"86","2021-06-16 10:34:39.981351","10.244.166.179",,,"50758",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&supi=imsi-208930000000003&target-nf-type=UDM",,,"0x0000f33c","197",
"87","2021-06-16 10:34:39.981355","10.244.166.179",,,"50758",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&supi=imsi-208930000000003&target-nf-type=UDM",,,"0x000063d4","197",
"88","2021-06-16 10:34:39.986742","10.244.166.129",,,"29510",,"10.244.166.179",,,"50758",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"89","2021-06-16 10:34:39.986766","10.103.54.119",,,"29510",,"10.244.166.179",,,"50758",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"90","2021-06-16 10:34:39.991383","10.244.166.179",,,"34952",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: PUT /nudm-uecm/v1/imsi-208930000000003/registrations/amf-3gpp-access",,,"0x0000def1","182",
"91","2021-06-16 10:34:39.991394","10.244.166.179",,,"34952",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: PUT /nudm-uecm/v1/imsi-208930000000003/registrations/amf-3gpp-access",,,"0x000063e7","182",
"92","2021-06-16 10:34:39.998461","10.244.166.163",,,"33452",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x0000f31b","180",
"93","2021-06-16 10:34:39.998465","10.244.166.163",,,"33452",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x000063b3","180",
"94","2021-06-16 10:34:40.001886","10.244.166.129",,,"29510",,"10.244.166.163",,,"33452",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006367","104",
"95","2021-06-16 10:34:40.001896","10.103.54.119",,,"29510",,"10.244.166.163",,,"33452",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2cf","104",
"96","2021-06-16 10:34:40.005857","10.244.166.163",,,"35402",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/amf-3gpp-access",,,"0x000065c4","200",
"97","2021-06-16 10:34:40.005862","10.244.166.163",,,"35402",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/amf-3gpp-access",,,"0x000063d0","200",
"98","2021-06-16 10:34:40.027760","10.244.166.138",,,"29504",,"10.244.166.163",,,"35402",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000636b","99",
"99","2021-06-16 10:34:40.027774","10.111.169.3",,,"29504",,"10.244.166.163",,,"35402",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000655f","99",
"100","2021-06-16 10:34:40.028755","10.244.166.163",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063d9","168",
"101","2021-06-16 10:34:40.028770","10.109.34.53",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x0000dee3","168",
"102","2021-06-16 10:34:40.028973","10.244.166.179",,,"34952",,"10.109.34.53",,,"29503",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: GET /nudm-sdm/v1/imsi-208930000000003/am-data?plmn-id=20893",,,"0x0000dea8","109",
"103","2021-06-16 10:34:40.028980","10.244.166.179",,,"34952",,"10.244.166.163",,,"29503",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: GET /nudm-sdm/v1/imsi-208930000000003/am-data?plmn-id=20893",,,"0x0000639e","109",
"104","2021-06-16 10:34:40.029214","10.244.166.163",,,"35402",,"10.111.169.3",,,"29504",,"HTTP2","[TCP Previous segment not captured] , HEADERS[5]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/am-data?supported-features=20893",,,"0x00006598","156",
"105","2021-06-16 10:34:40.029223","10.244.166.163",,,"35402",,"10.244.166.138",,,"29504",,"HTTP2","[TCP Previous segment not captured] , HEADERS[5]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/am-data?supported-features=20893",,,"0x000063a4","156",
"106","2021-06-16 10:34:40.031551","10.244.166.138",,,"29504",,"10.244.166.163",,,"35402",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: 200 OK",,,"0x0000634d","69",
"107","2021-06-16 10:34:40.031559","10.111.169.3",,,"29504",,"10.244.166.163",,,"35402",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: 200 OK",,,"0x00006541","69",
"108","2021-06-16 10:34:40.031833","10.244.166.163",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: 200 OK",,,"0x00006376","69",
"109","2021-06-16 10:34:40.031842","10.109.34.53",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[5]: 200 OK",,,"0x0000de80","69",
"110","2021-06-16 10:34:40.032129","10.244.166.179",,,"34952",,"10.109.34.53",,,"29503",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[7]: GET /nudm-sdm/v1/imsi-208930000000003/smf-select-data?plmn-id=20893",,,"0x0000deae","115",
"111","2021-06-16 10:34:40.032136","10.244.166.179",,,"34952",,"10.244.166.163",,,"29503",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[7]: GET /nudm-sdm/v1/imsi-208930000000003/smf-select-data?plmn-id=20893",,,"0x000063a4","115",
"112","2021-06-16 10:34:40.034568","10.244.166.163",,,"33456",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x0000f31b","180",
"113","2021-06-16 10:34:40.034572","10.244.166.163",,,"33456",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x000063b3","180",
"114","2021-06-16 10:34:40.037731","10.244.166.129",,,"29510",,"10.244.166.163",,,"33456",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006367","104",
"115","2021-06-16 10:34:40.037739","10.103.54.119",,,"29510",,"10.244.166.163",,,"33456",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2cf","104",
"116","2021-06-16 10:34:40.045243","10.244.166.163",,,"35406",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/smf-selection-subscription-data?supported-features=",,,"0x000065d5","217",
"117","2021-06-16 10:34:40.045248","10.244.166.163",,,"35406",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/smf-selection-subscription-data?supported-features=",,,"0x000063e1","217",
"118","2021-06-16 10:34:40.047543","10.244.166.138",,,"29504",,"10.244.166.163",,,"35406",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000636f","103",
"119","2021-06-16 10:34:40.047554","10.111.169.3",,,"29504",,"10.244.166.163",,,"35406",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006563","103",
"120","2021-06-16 10:34:40.047904","10.244.166.163",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP Previous segment not captured] , HEADERS[7]: 200 OK",,,"0x00006375","68",
"121","2021-06-16 10:34:40.047914","10.109.34.53",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP Previous segment not captured] , HEADERS[7]: 200 OK",,,"0x0000de7f","68",
"122","2021-06-16 10:34:40.048342","10.244.166.179",,,"34952",,"10.109.34.53",,,"29503",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[9]: GET /nudm-sdm/v1/imsi-208930000000003/ue-context-in-smf-data",,,"0x0000dea8","109",
"123","2021-06-16 10:34:40.048349","10.244.166.179",,,"34952",,"10.244.166.163",,,"29503",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[9]: GET /nudm-sdm/v1/imsi-208930000000003/ue-context-in-smf-data",,,"0x0000639e","109",
"124","2021-06-16 10:34:40.050156","10.244.166.163",,,"33460",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x0000f31b","180",
"125","2021-06-16 10:34:40.050160","10.244.166.163",,,"33460",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x000063b3","180",
"126","2021-06-16 10:34:40.052510","10.244.166.129",,,"29510",,"10.244.166.163",,,"33460",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006367","104",
"127","2021-06-16 10:34:40.052519","10.103.54.119",,,"29510",,"10.244.166.163",,,"33460",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2cf","104",
"128","2021-06-16 10:34:40.055079","10.244.166.163",,,"35410",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/smf-registrations?supported-features=",,,"0x000065c4","200",
"129","2021-06-16 10:34:40.055085","10.244.166.163",,,"35410",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/smf-registrations?supported-features=",,,"0x000063d0","200",
"130","2021-06-16 10:34:40.057062","10.244.166.138",,,"29504",,"10.244.166.163",,,"35410",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000636e","102",
"131","2021-06-16 10:34:40.057071","10.111.169.3",,,"29504",,"10.244.166.163",,,"35410",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006562","102",
"132","2021-06-16 10:34:40.058093","10.244.166.163",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP Previous segment not captured] , HEADERS[9]: 200 OK",,,"0x00006374","67",
"133","2021-06-16 10:34:40.058103","10.109.34.53",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP Previous segment not captured] , HEADERS[9]: 200 OK",,,"0x0000de7e","67",
"134","2021-06-16 10:34:40.058339","10.244.166.179",,,"34952",,"10.109.34.53",,,"29503",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[11]: POST /nudm-sdm/v1/imsi-208930000000003/sdm-subscriptions",,,"0x0000dea9","110",
"135","2021-06-16 10:34:40.058346","10.244.166.179",,,"34952",,"10.244.166.163",,,"29503",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[11]: POST /nudm-sdm/v1/imsi-208930000000003/sdm-subscriptions",,,"0x0000639f","110",
"136","2021-06-16 10:34:40.060380","10.244.166.163",,,"33464",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x0000f31b","180",
"137","2021-06-16 10:34:40.060385","10.244.166.163",,,"33464",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=UDM&target-nf-type=UDR",,,"0x000063b3","180",
"138","2021-06-16 10:34:40.063118","10.244.166.129",,,"29510",,"10.244.166.163",,,"33464",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006367","104",
"139","2021-06-16 10:34:40.063127","10.103.54.119",,,"29510",,"10.244.166.163",,,"33464",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2cf","104",
"140","2021-06-16 10:34:40.065752","10.244.166.163",,,"35414",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: POST /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/sdm-subscriptions",,,"0x000065ba","190",
"141","2021-06-16 10:34:40.065756","10.244.166.163",,,"35414",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: POST /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/sdm-subscriptions",,,"0x000063c6","190",
"142","2021-06-16 10:34:40.066687","10.244.166.138",,,"29504",,"10.244.166.163",,,"35414",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063bd","181",
"143","2021-06-16 10:34:40.066694","10.111.169.3",,,"29504",,"10.244.166.163",,,"35414",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000065b1","181",
"144","2021-06-16 10:34:40.067408","10.244.166.163",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[11]: 201 Created",,,"0x000063ab","122",
"145","2021-06-16 10:34:40.067443","10.109.34.53",,,"29503",,"10.244.166.179",,,"34952",,"HTTP2","[TCP ACKed unseen segment] [TCP Previous segment not captured] , HEADERS[11]: 201 Created",,,"0x0000deb5","122",
"146","2021-06-16 10:34:40.072442","10.244.166.179",,,"50778",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&supi=imsi-208930000000003&target-nf-type=PCF",,,"0x0000f33c","197",
"147","2021-06-16 10:34:40.072450","10.244.166.179",,,"50778",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&supi=imsi-208930000000003&target-nf-type=PCF",,,"0x000063d4","197",
"148","2021-06-16 10:34:40.074440","10.244.166.129",,,"29510",,"10.244.166.179",,,"50778",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"149","2021-06-16 10:34:40.074449","10.103.54.119",,,"29510",,"10.244.166.179",,,"50778",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"150","2021-06-16 10:34:40.076463","10.244.166.179",,,"32978",,"10.97.216.196",,,"29507",,"HTTP2","HEADERS[3]: POST /npcf-am-policy-control/v1/policies",,,"0x0000955e","159",
"151","2021-06-16 10:34:40.076468","10.244.166.179",,,"32978",,"10.244.166.170",,,"29507",,"HTTP2","HEADERS[3]: POST /npcf-am-policy-control/v1/policies",,,"0x000063d7","159",
"152","2021-06-16 10:34:40.224708","10.244.166.179",,,"32980",,"10.97.216.196",,,"29507",,"HTTP2","HEADERS[3]: POST /npcf-am-policy-control/v1/policies",,,"0x0000955e","159",
"153","2021-06-16 10:34:40.224717","10.244.166.179",,,"32980",,"10.244.166.170",,,"29507",,"HTTP2","HEADERS[3]: POST /npcf-am-policy-control/v1/policies",,,"0x000063d7","159",
"154","2021-06-16 10:34:40.270599","10.244.166.170",,,"33404",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=PCF&target-nf-type=UDR",,,"0x0000f322","180",
"155","2021-06-16 10:34:40.270604","10.244.166.170",,,"33404",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=PCF&target-nf-type=UDR",,,"0x000063ba","180",
"156","2021-06-16 10:34:40.275186","10.244.166.129",,,"29510",,"10.244.166.170",,,"33404",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000636e","104",
"157","2021-06-16 10:34:40.275195","10.103.54.119",,,"29510",,"10.244.166.170",,,"33404",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2d6","104",
"158","2021-06-16 10:34:40.282704","10.244.166.170",,,"46590",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/policy-data/ues/imsi-208930000000003/am-data",,,"0x000065ac","169",
"159","2021-06-16 10:34:40.282709","10.244.166.170",,,"46590",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/policy-data/ues/imsi-208930000000003/am-data",,,"0x000063b8","169",
"160","2021-06-16 10:34:40.284880","10.244.166.138",,,"29504",,"10.244.166.170",,,"46590",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006376","103",
"161","2021-06-16 10:34:40.284892","10.111.169.3",,,"29504",,"10.244.166.170",,,"46590",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000656a","103",
"162","2021-06-16 10:34:40.288377","10.244.166.170",,,"29507",,"10.244.166.179",,,"32980",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063dd","165",
"163","2021-06-16 10:34:40.288402","10.97.216.196",,,"29507",,"10.244.166.179",,,"32980",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x00009564","165",
"164","2021-06-16 10:34:40.299095","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","InitialContextSetupRequest",,"0x11a66f42",,"220",
"165","2021-06-16 10:34:40.300187","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP","InitialContextSetupResponse",,"0xa10d83bc",,"84",
"166","2021-06-16 10:34:40.501141","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport",,"0x0175294e",,"228",
"167","2021-06-16 10:34:40.509070","10.244.166.179",,,"50788",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=NSSF",,,"0x0000f32b","180",
"168","2021-06-16 10:34:40.509075","10.244.166.179",,,"50788",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=NSSF",,,"0x000063c3","180",
"169","2021-06-16 10:34:40.512456","10.244.166.129",,,"29510",,"10.244.166.179",,,"50788",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"170","2021-06-16 10:34:40.512464","10.103.54.119",,,"29510",,"10.244.166.179",,,"50788",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"171","2021-06-16 10:34:40.676858","10.244.166.179",,,"50792",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?dnn=internet&requester-nf-type=AMF&service-names=nsmf-pdusession&snssais=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D&target-nf-type=SMF&target-plmn-list=%7B%22mcc%22%3A%22208%22%2C%22mnc%22%3A%2293%22%7D",,,"0x0000f39f","296",
"172","2021-06-16 10:34:40.676862","10.244.166.179",,,"50792",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?dnn=internet&requester-nf-type=AMF&service-names=nsmf-pdusession&snssais=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D&target-nf-type=SMF&target-plmn-list=%7B%22mcc%22%3A%22208%22%2C%22mnc%22%3A%2293%22%7D",,,"0x00006437","296",
"173","2021-06-16 10:34:40.682382","10.244.166.179",,,"50794",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?dnn=internet&requester-nf-type=AMF&service-names=nsmf-pdusession&snssais=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D&target-nf-type=SMF&target-plmn-list=%7B%22mcc%22%3A%22208%22%2C%22mnc%22%3A%2293%22%7D",,,"0x0000f39f","296",
"174","2021-06-16 10:34:40.682388","10.244.166.179",,,"50794",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?dnn=internet&requester-nf-type=AMF&service-names=nsmf-pdusession&snssais=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D&target-nf-type=SMF&target-plmn-list=%7B%22mcc%22%3A%22208%22%2C%22mnc%22%3A%2293%22%7D",,,"0x00006437","296",
"175","2021-06-16 10:34:40.689134","10.244.166.129",,,"29510",,"10.244.166.179",,,"50794",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006377","104",
"176","2021-06-16 10:34:40.689152","10.103.54.119",,,"29510",,"10.244.166.179",,,"50794",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2df","104",
"177","2021-06-16 10:34:40.696881","10.244.166.179",,,"59080",,"10.101.204.203",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts",,,"0x0000899c","210",
"178","2021-06-16 10:34:40.696885","10.244.166.179",,,"59080",,"10.244.166.158",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts",,,"0x000063fe","210",
"179","2021-06-16 10:34:40.923442","10.244.166.158",,,"51308",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-type=UDM",,,"0x0000f316","180",
"180","2021-06-16 10:34:40.923448","10.244.166.158",,,"51308",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-type=UDM",,,"0x000063ae","180",
"181","2021-06-16 10:34:40.929259","10.244.166.129",,,"29510",,"10.244.166.158",,,"51308",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006362","104",
"182","2021-06-16 10:34:40.929274","10.103.54.119",,,"29510",,"10.244.166.158",,,"51308",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2ca","104",
"183","2021-06-16 10:34:40.935162","10.244.166.158",,,"33560",,"10.109.34.53",,,"29503",,"HTTP2","HEADERS[3]: GET /nudm-sdm/v1/imsi-208930000000003/sm-data?dnn=internet&plmn-id=20893&single-nssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x0000df02","220",
"184","2021-06-16 10:34:40.935168","10.244.166.158",,,"33560",,"10.244.166.163",,,"29503",,"HTTP2","HEADERS[3]: GET /nudm-sdm/v1/imsi-208930000000003/sm-data?dnn=internet&plmn-id=20893&single-nssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000063f8","220",
"185","2021-06-16 10:34:40.946547","10.244.166.163",,,"35440",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/sm-data?single-nssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000065e0","228",
"186","2021-06-16 10:34:40.946555","10.244.166.163",,,"35440",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/sm-data?single-nssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000063ec","228",
"187","2021-06-16 10:34:40.948197","10.244.166.163",,,"35442",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/sm-data?single-nssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000065e0","228",
"188","2021-06-16 10:34:40.948202","10.244.166.163",,,"35442",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/subscription-data/imsi-208930000000003/20893/provisioned-data/sm-data?single-nssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000063ec","228",
"189","2021-06-16 10:34:40.949973","10.244.166.138",,,"29504",,"10.244.166.163",,,"35442",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000636f","103",
"190","2021-06-16 10:34:40.949981","10.111.169.3",,,"29504",,"10.244.166.163",,,"35442",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006563","103",
"191","2021-06-16 10:34:40.950488","10.244.166.163",,,"29503",,"10.244.166.158",,,"33560",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006384","104",
"192","2021-06-16 10:34:40.950501","10.109.34.53",,,"29503",,"10.244.166.158",,,"33560",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000de8e","104",
"193","2021-06-16 10:34:40.966297","10.244.166.158",,,"51316",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-type=PCF",,,"0x0000f316","180",
"194","2021-06-16 10:34:40.966301","10.244.166.158",,,"51316",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-type=PCF",,,"0x000063ae","180",
"195","2021-06-16 10:34:40.998051","10.244.166.129",,,"29510",,"10.244.166.158",,,"51316",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006362","104",
"196","2021-06-16 10:34:40.998067","10.103.54.119",,,"29510",,"10.244.166.158",,,"51316",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2ca","104",
"197","2021-06-16 10:34:41.002891","10.244.166.158",,,"44572",,"10.97.216.196",,,"29507",,"HTTP2","HEADERS[3]: POST /npcf-smpolicycontrol/v1/sm-policies",,,"0x0000954b","161",
"198","2021-06-16 10:34:41.002896","10.244.166.158",,,"44572",,"10.244.166.170",,,"29507",,"HTTP2","HEADERS[3]: POST /npcf-smpolicycontrol/v1/sm-policies",,,"0x000063c4","161",
"199","2021-06-16 10:34:41.009193","10.244.166.170",,,"46614",,"10.111.169.3",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/policy-data/ues/imsi-208930000000003/sm-data?dnn=internet&snssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000065db","216",
"200","2021-06-16 10:34:41.009198","10.244.166.170",,,"46614",,"10.244.166.138",,,"29504",,"HTTP2","HEADERS[3]: GET /nudr-dr/v1/policy-data/ues/imsi-208930000000003/sm-data?dnn=internet&snssai=%7B%22sst%22%3A1%2C%22sd%22%3A%22010203%22%7D",,,"0x000063e7","216",
"201","2021-06-16 10:34:41.014116","10.244.166.138",,,"29504",,"10.244.166.170",,,"46614",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006376","103",
"202","2021-06-16 10:34:41.014126","10.111.169.3",,,"29504",,"10.244.166.170",,,"46614",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000656a","103",
"203","2021-06-16 10:34:41.021559","10.244.166.170",,,"29507",,"10.244.166.158",,,"44572",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063c8","165",
"204","2021-06-16 10:34:41.021576","10.97.216.196",,,"29507",,"10.244.166.158",,,"44572",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x0000954f","165",
"205","2021-06-16 10:34:41.027817","10.244.166.158",,,"51322",,"10.103.54.119",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-instance-id=2e1d4086-ae4b-49ab-a99d-493382292ece&target-nf-type=AMF",,,"0x0000f33f","221",
"206","2021-06-16 10:34:41.027822","10.244.166.158",,,"51322",,"10.244.166.129",,,"29510",,"HTTP2","HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-instance-id=2e1d4086-ae4b-49ab-a99d-493382292ece&target-nf-type=AMF",,,"0x000063d7","221",
"207","2021-06-16 10:34:41.033052","10.244.166.129",,,"29510",,"10.244.166.158",,,"51322",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006362","104",
"208","2021-06-16 10:34:41.033062","10.103.54.119",,,"29510",,"10.244.166.158",,,"51322",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x0000f2ca","104",
"209","2021-06-16 10:34:41.035420","10.244.166.158",,,"29502",,"10.244.166.179",,,"59080",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x000063ee","194",
"210","2021-06-16 10:34:41.035435","10.101.204.203",,,"29502",,"10.244.166.179",,,"59080",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 201 Created",,,"0x0000898c","194",
"211","2021-06-16 10:34:41.035474","172.16.30.20",,"8805",,,"172.16.30.30",,"8805",,,"PFCP","PFCP Session Establishment Request","0x0000954f",,,"255",
"212","2021-06-16 10:34:41.035476","172.16.30.20",,"8805",,,"172.16.30.30",,"8805",,,"PFCP","PFCP Session Establishment Request","0x0000954f",,,"255",
"213","2021-06-16 10:34:41.074638","172.16.30.30",,"8805",,,"172.16.30.20",,"8805",,,"PFCP","PFCP Session Establishment Response","0x000094ab",,,"91",
"214","2021-06-16 10:34:41.074642","172.16.30.30",,"8805",,,"172.16.30.20",,"8805",,,"PFCP","PFCP Session Establishment Response","0x000094ab",,,"91",
"215","2021-06-16 10:34:41.090444","10.244.166.158",,,"52568",,"10.244.166.179",,,"29518",,"HTTP2","HEADERS[3]: POST /namf-comm/v1/ue-contexts/imsi-208930000000003/n1-n2-messages",,,"0x00006413","231",
"216","2021-06-16 10:34:41.090477","10.244.166.158",,,"52568",,"10.244.166.179",,,"29518",,"HTTP2/JSON/NAS-5GS/NGAP","DATA[3], JavaScript Object Notation (application/json), PDU session establishment accept (PDU session type IPv4 only allowed)",,,"0x000066b5","905",
"217","2021-06-16 10:34:41.160946","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","PDUSessionResourceSetupRequest",,"0x85c7e49a",,"200",
"218","2021-06-16 10:34:41.167214","10.244.166.179",,,"29518",,"10.244.166.158",,,"52568",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00006393","103",
"219","2021-06-16 10:34:41.195829","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP","PDUSessionResourceSetupResponse",,"0x928aaae4",,"104",
"220","2021-06-16 10:34:41.214746","10.244.166.179",,,"59100",,"10.101.204.203",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x000089c3","249",
"221","2021-06-16 10:34:41.214754","10.244.166.179",,,"59100",,"10.244.166.158",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x00006425","249",
"222","2021-06-16 10:34:41.215694","10.244.166.179",,,"59100",,"10.101.204.203",,,"29502",,"HTTP2/JSON/NGAP","DATA[3], JavaScript Object Notation (application/json)",,,"0x00008b7a","688",
"223","2021-06-16 10:34:41.215710","10.244.166.179",,,"59100",,"10.244.166.158",,,"29502",,"HTTP2/JSON/NGAP","DATA[3], JavaScript Object Notation (application/json)",,,"0x000065dc","688",
"224","2021-06-16 10:34:41.217572","172.16.30.20",,"8805",,,"172.16.30.30",,"8805",,,"PFCP","PFCP Session Modification Request","0x000094fb",,,"171",
"225","2021-06-16 10:34:41.217575","172.16.30.20",,"8805",,,"172.16.30.30",,"8805",,,"PFCP","PFCP Session Modification Request","0x000094fb",,,"171",
"226","2021-06-16 10:34:41.218962","172.16.30.30",,"8805",,,"172.16.30.20",,"8805",,,"PFCP","PFCP Session Modification Response","0x00009481",,,"49",
"227","2021-06-16 10:34:41.218966","172.16.30.30",,"8805",,,"172.16.30.20",,"8805",,,"PFCP","PFCP Session Modification Response","0x00009481",,,"49",
"228","2021-06-16 10:34:41.222981","10.244.166.158",,,"29502",,"10.244.166.179",,,"59100",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x000063ca","158",
"229","2021-06-16 10:34:41.222995","10.101.204.203",,,"29502",,"10.244.166.179",,,"59100",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00008968","158",
"230","2021-06-16 10:35:09.818348","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport",,"0xb100c683",,"136",
"231","2021-06-16 10:35:10.061840","10.244.166.179",,,"59152",,"10.101.204.203",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x000089c2","248",
"232","2021-06-16 10:35:10.061850","10.244.166.179",,,"59152",,"10.244.166.158",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x00006424","248",
"233","2021-06-16 10:35:10.254749","172.16.30.20",,"8805",,,"172.16.30.30",,"8805",,,"PFCP","PFCP Session Deletion Request","0x0000947c",,,"44",
"234","2021-06-16 10:35:10.254775","172.16.30.20",,"8805",,,"172.16.30.30",,"8805",,,"PFCP","PFCP Session Deletion Request","0x0000947c",,,"44",
"235","2021-06-16 10:35:10.293886","172.16.30.30",,"8805",,,"172.16.30.20",,"8805",,,"PFCP","PFCP Session Deletion Response","0x00009481",,,"49",
"236","2021-06-16 10:35:10.293891","172.16.30.30",,"8805",,,"172.16.30.20",,"8805",,,"PFCP","PFCP Session Deletion Response","0x00009481",,,"49",
"237","2021-06-16 10:35:10.301705","10.244.166.158",,,"29502",,"10.244.166.179",,,"59152",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x000063ca","158",
"238","2021-06-16 10:35:10.301735","10.101.204.203",,,"29502",,"10.244.166.179",,,"59152",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00008968","158",
"239","2021-06-16 10:35:10.311507","10.244.166.158",,,"29502",,"10.244.166.179",,,"59152",,"HTTP2/JSON/NAS-5GS/NGAP","[TCP ACKed unseen segment] , DATA[3], JavaScript Object Notation (application/json), PDU session release command (Unknown)",,,"0x000065c7","667",
"240","2021-06-16 10:35:10.311579","10.101.204.203",,,"29502",,"10.244.166.179",,,"59152",,"HTTP2/JSON/NAS-5GS/NGAP","[TCP ACKed unseen segment] , DATA[3], JavaScript Object Notation (application/json), PDU session release command (Unknown)",,,"0x00008b65","667",
"241","2021-06-16 10:35:10.322050","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","PDUSessionResourceReleaseCommand",,"0xe6f6fa34",,"104",
"242","2021-06-16 10:35:10.327225","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP","PDUSessionResourceReleaseResponse",,"0x805a4ebc",,"116",
"243","2021-06-16 10:35:10.337160","10.244.166.179",,,"59154",,"10.101.204.203",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x000089c3","249",
"244","2021-06-16 10:35:10.337170","10.244.166.179",,,"59154",,"10.244.166.158",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x00006425","249",
"245","2021-06-16 10:35:10.337212","10.244.166.179",,,"59154",,"10.101.204.203",,,"29502",,"HTTP2/JSON/NGAP","DATA[3], JavaScript Object Notation (application/json)",,,"0x00008b6d","675",
"246","2021-06-16 10:35:10.337216","10.244.166.179",,,"59154",,"10.244.166.158",,,"29502",,"HTTP2/JSON/NGAP","DATA[3], JavaScript Object Notation (application/json)",,,"0x000065cf","675",
"247","2021-06-16 10:35:10.339649","10.244.166.158",,,"29502",,"10.244.166.179",,,"59154",,"HTTP2","WINDOW_UPDATE[0], HEADERS[3]: 200 OK",,,"0x000063d7","171",
"248","2021-06-16 10:35:10.339654","10.101.204.203",,,"29502",,"10.244.166.179",,,"59154",,"HTTP2","WINDOW_UPDATE[0], HEADERS[3]: 200 OK",,,"0x00008975","171",
"249","2021-06-16 10:35:10.528945","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport",,"0x4cad2a71",,"132",
"250","2021-06-16 10:35:10.534965","10.244.166.179",,,"59156",,"10.101.204.203",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x000089c2","248",
"251","2021-06-16 10:35:10.534969","10.244.166.179",,,"59156",,"10.244.166.158",,,"29502",,"HTTP2","HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",,,"0x00006424","248",
"252","2021-06-16 10:35:10.558157","10.244.166.158",,,"52628",,"10.244.166.179",,,"29518",,"HTTP2","HEADERS[3]: POST /namf-callback/v1/smContextStatus/20893cafe0000000001/1",,,"0x000063d8","172",
"253","2021-06-16 10:35:10.576007","10.244.166.179",,,"29518",,"10.244.166.158",,,"52628",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000638f","99",
"254","2021-06-16 10:35:10.577567","10.244.166.158",,,"29502",,"10.244.166.179",,,"59156",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x000063c9","157",
"255","2021-06-16 10:35:10.577575","10.101.204.203",,,"29502",,"10.244.166.179",,,"59156",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 200 OK",,,"0x00008967","157",
"256","2021-06-16 10:35:21.579790","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport, Deregistration request (UE originating)",,"0x682f74c8",,"120",
"257","2021-06-16 10:35:21.774682","10.244.166.179",,,"33092",,"10.97.216.196",,,"29507",,"HTTP2","HEADERS[3]: DELETE /npcf-am-policy-control/v1/policies/imsi-208930000000003-1",,,"0x00009577","184",
"258","2021-06-16 10:35:21.774708","10.244.166.179",,,"33092",,"10.244.166.170",,,"29507",,"HTTP2","HEADERS[3]: DELETE /npcf-am-policy-control/v1/policies/imsi-208930000000003-1",,,"0x000063f0","184",
"259","2021-06-16 10:35:21.780613","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP/NAS-5GS","UplinkNASTransport, Registration request",,"0x0b268dcb",,"120",
"260","2021-06-16 10:35:21.919516","10.244.166.170",,,"29507",,"10.244.166.179",,,"33092",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x0000639b","99",
"261","2021-06-16 10:35:21.919545","10.97.216.196",,,"29507",,"10.244.166.179",,,"33092",,"HTTP2","[TCP ACKed unseen segment] , HEADERS[3]: 204 No Content",,,"0x00009522","99",
"262","2021-06-16 10:35:21.931935","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP","UEContextReleaseCommand",,"0xce1c7f47",,"84",
"263","2021-06-16 10:35:21.937576","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP","UEContextReleaseComplete",,"0xd061c859",,"108",
"264","2021-06-16 10:35:21.950713","172.16.10.20",,,,"38412","172.16.10.10",,,,"46571","NGAP/NAS-5GS","DownlinkNASTransport, Registration reject (Protocol error, unspecified)",,"0x8bf08af0",,"96",
"265","2021-06-16 10:35:21.951069","172.16.10.10",,,,"46571","172.16.10.20",,,,"38412","NGAP","ErrorIndication",,"0xaf08a9fe",,"80",
`

func TestSetAddress(t *testing.T) {
	type address struct {
		v4    string
		v6    string
		lenv4 string
		lenv6 string
	}

	type Tests struct {
		name   string
		args   address
		want   string
	}

	tests := []Tests {
		{
			name: "v4 Only",
			args: address{
				v4:    "1.1.1.1",
				v6:    "",
				lenv4: "",
				lenv6: "",
			},
			want: "1.1.1.1",
		},
		{
			name: "v6 Only",
			args: address{
				v4:    "",
				v6:    "2001::0",
				lenv4: "",
				lenv6: "",
			},
			want: "2001::0",
		},
		{
			name: "v4v6 Dual",
			args: address{
				v4:    "1.1.1.1",
				v6:    "2001::0",
				lenv4: "100",
				lenv6: "120",
			},
			want: "1.1.1.1",
		},
		{
			name: "v4 Dual",
			args: address{
				v4:    "1.1.1.1,2.2.2.2",
				v6:    "",
				lenv4: "100,200",
				lenv6: "",
			},
			want: "1.1.1.1",
		},
		{
			name: "v6 Dual",
			args: address{
				v4:    "",
				v6:    "2001::0,2001::1",
				lenv4: "",
				lenv6: "200,100",
			},
			want: "2001::1",
		},
		{
			name: "v4v6 Dual Exp",
			args: address{
				v4:    "1.1.1.1,2.2.2.2",
				v6:    "2001::0,2001::1",
				lenv4: "50,150",
				lenv6: "200,100",
			},
			want: "1.1.1.1",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := tshark.SetAddress(v.args.v4, v.args.v6, v.args.lenv4, v.args.lenv6)
			if res != v.want {
				t.Errorf("The return value is not the expected value.\nres: %s\n", res)
			}
		})
	}
}

func TestSetSetPortAndCheckSum(t *testing.T) {
	type port struct {
		udp  string
		tcp  string
		sctp string
	}

	type Tests struct {
		name   string
		args   port
		want   string
	}

	tests := []Tests{
		{
			name: "UDP Only",
			args: port{
				udp:  "2123",
				tcp:  "",
				sctp: "",
			},
			want: "2123",
		},
		{
			name: "TCP Only",
			args: port{
				udp:  "",
				tcp:  "3868",
				sctp: "",
			},
			want: "3868",
		},
		{
			name: "SCTP Only",
			args: port{
				udp:  "",
				tcp:  "",
				sctp: "36412",
			},
			want: "36412",
		},
		{
			name: "Unanticipated cases",
			args: port{
				udp:  "2123",
				tcp:  "3686",
				sctp: "36412",
			},
			want: "",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := tshark.SetPortAndCheckSum(v.args.udp, v.args.tcp, v.args.sctp)
			if res != v.want {
				t.Errorf("The return value is not the expected value.")
			}
		})
	}
}

func TestSetMessage(t *testing.T) {
	type required struct {
		message  string
		protocol string
	}

	type Tests struct {
		name   string
		args   required
		want   string
	}

	tests := []Tests{
		{
			name: "GTPv2",
			args: required{
				message:  "Create Session Request",
				protocol: "GTPv2",
			},
			want: "Create Session Request",
		},
		{
			name: "PFCP",
			args: required{
				message:  "PFCP Session Establishment Request",
				protocol: "PFCP",
			},
			want: "PFCP Session Establishment Request",
		},
		{
			name: "DIAMETER Request",
			args: required{
				message:  "cmd=Capabilities-Exchange Request(257) flags=R--- appl=Diameter Common Messages(0) h2h=7c8a72c3 e2e=f3d80eea",
				protocol: "DIAMETER",
			},
			want: "Capabilities-Exchange Request",
		},
		{
			name: "DIAMETER Answer",
			args: required{
				message:  "cmd=Capabilities-Exchange Answer(257) flags=---- appl=Diameter Common Messages(0) h2h=7c8a72c3 e2e=f3d80eea",
				protocol: "DIAMETER",
			},
			want: "Capabilities-Exchange Answer",
		},
		{
			name: "3GPP DIAMETER Request",
			args: required{
				message:  "cmd=3GPP-Authentication-Information Request(318) flags=RP-- appl=3GPP S6a/S6d(16777251) h2h=7c8a72c4 e2e=f3d80eeb | ",
				protocol: "DIAMETER",
			},
			want: "Authentication-Information Request",
		},
		{
			name: "3GPP DIAMETER Answer",
			args: required{
				message:  "SACK cmd=3GPP-Authentication-Information Answer(318) flags=-P-- appl=3GPP S6a/S6d(16777251) h2h=7c8a72c4 e2e=f3d80eeb | ",
				protocol: "DIAMETER",
			},
			want: "Authentication-Information Answer",
		},
		{
			name: "S1AP",
			args: required{
				message:  "S1SetupRequest",
				protocol: "S1AP",
			},
			want: "S1SetupRequest",
		},
		{
			name: "HTTP2 Method GET",
			args: required{
				message:  "HEADERS[3]: GET /nnrf-disc/v1/nf-instances?requester-nf-type=SMF&target-nf-type=PCF",
				protocol: "HTTP2",
			},
			want: "GET /nnrf-disc/v1/nf-instances",
		},
		{
			name: "HTTP2 Method HEAD",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method POST",
			args: required{
				message:  "HEADERS[3]: POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",
				protocol: "HTTP2",
			},
			want: "POST /nsmf-pdusession/v1/sm-contexts/urn:uuid:d005b19a-e01e-4198-8a7c-1b68923ef52d/modify",
		},
		{
			name: "HTTP2 Method PUT",
			args: required{
				message:  "HEADERS[3]: PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/amf-3gpp-access",
				protocol: "HTTP2",
			},
			want: "PUT /nudr-dr/v1/subscription-data/imsi-208930000000003/context-data/amf-3gpp-access",
		},
		{
			name: "HTTP2 Method DELETE",
			args: required{
				message:  "HEADERS[3]: DELETE /npcf-am-policy-control/v1/policies/imsi-208930000000003-1",
				protocol: "HTTP2",
			},
			want: "DELETE /npcf-am-policy-control/v1/policies/imsi-208930000000003-1",
		},
		{
			name: "HTTP2 Method CONNECT",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method OPTIONS",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method TRACE",
			args: required{
				message:  "",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "HTTP2 Method PATCH",
			args: required{
				message:  "HEADERS[3]: PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",
				protocol: "HTTP2",
			},
			want: "PATCH /nudr-dr/v1/subscription-data/imsi-208930000000003/authentication-data/authentication-subscription",
		},
		{
			name: "HTTP2 Method Response",
			args: required{
				message:  "HEADERS[3]: 200 OK",
				protocol: "HTTP2",
			},
			want: "200 OK",
		},
		{
			name: "HTTP2 PDU",
			args: required{
				message:  "DATA[3], JavaScript Object Notation (application/json), PDU session establishment accept (PDU session type IPv4 only allowed)",
				protocol: "HTTP2/JSON/NAS-5GS/NGAP",
			},
			want: "PDU session establishment accept (PDU session type IPv4 only allowed)",
		},
		{
			name: "HTTP2 Drop",
			args: required{
				message:  "SETTINGS[0], WINDOW_UPDATE[0]",
				protocol: "HTTP2",
			},
			want: "",
		},
		{
			name: "TCP Drop",
			args: required{
				message:  "8080 → 47970 [ACK] Seq=1 Ack=2 Win=27136 Len=0 TSval=1042549644 TSecr=1042549643",
				protocol: "TCP",
			},
			want: "",
		},
		{
			name: "SCTP Drop",
			args: required{
				message:  "HEARTBEAT",
				protocol: "SCTP",
			},
			want: "",
		},
		{
			name: "ICMP",
			args: required{
				message:  "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
				protocol: "ICMP",
			},
			want: "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
		},
		{
			name: "ICMPv6",
			args: required{
				message:  "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
				protocol: "ICMPv6",
			},
			want: "Echo (ping) request  id=0x7348, seq=27032/39017, ttl=255 (no response found!)",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := tshark.SetMessage(v.args.message, v.args.protocol)
			if res != v.want {
				t.Errorf("The return value is not the expected value.\n %s", res)
			}
		})
	}
}

func TestSetHeader(t *testing.T) {
	var fmtOut tshark.TsharkHeaders

	type Tests struct {
		name   string
		args   string
		want   string
	}

	tests := Tests {
		name: "Normal Case",
		args: expOut,
		want: "",
	}

	t.Run(tests.name, func(t *testing.T) {
		res := fmtOut.SetHeader(tests.args)
		if len(res) == 0 {
			t.Errorf("%s", res)
		}

		for _, v := range res {
			if v.Number == tests.want {
				t.Errorf("Column Number is empty")
			}
			if v.Time == tests.want {
				t.Errorf("Column Time is empty")
			}
			if v.SrcAddr == tests.want {
				t.Errorf("Column SrcAddr is empty")
			}
			if v.SrcPort == tests.want {
				t.Errorf("Column SrcPort is empty")
			}
			if v.DstAddr == tests.want {
				t.Errorf("Column DstAddr is empty")
			}
			if v.DstPort == tests.want {
				t.Errorf("Column DstPort is empty")
			}
			if v.Protocol == tests.want {
				t.Errorf("Column Protocol is empty")
			}
			if v.Message == tests.want {
				t.Errorf("Column Message is empty")
			}
			if v.Checksum == tests.want {
				t.Errorf("Column Checksum is empty")
			}
		}
	})
}
