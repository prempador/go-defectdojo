# SystemSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDeduplication** | Pointer to **bool** | With this setting turned on, DefectDojo deduplicates findings by comparing endpoints, cwe fields, and titles. If two findings share a URL and have the same CWE or title, DefectDojo marks the recent finding as a duplicate. When deduplication is enabled, a list of deduplicated findings is added to the engagement view. | [optional] 
**DeleteDuplicates** | Pointer to **bool** | Requires next setting: maximum number of duplicates to retain. | [optional] 
**MaxDupes** | Pointer to **NullableInt32** | When enabled, if a single issue reaches the maximum number of duplicates, the oldest will be deleted. Duplicate will not be deleted when left empty. A value of 0 will remove all duplicates. | [optional] 
**EmailFrom** | Pointer to **string** |  | [optional] 
**EnableJira** | Pointer to **bool** |  | [optional] 
**EnableJiraWebHook** | Pointer to **bool** | Please note: It is strongly recommended to use a secret below and / or IP whitelist the JIRA server using a proxy such as Nginx. | [optional] 
**DisableJiraWebhookSecret** | Pointer to **bool** | Allows incoming requests without a secret (discouraged legacy behaviour) | [optional] 
**JiraWebhookSecret** | Pointer to **NullableString** | Secret needed in URL for incoming JIRA Webhook | [optional] 
**JiraMinimumSeverity** | Pointer to **NullableString** | * &#x60;Critical&#x60; - Critical * &#x60;High&#x60; - High * &#x60;Medium&#x60; - Medium * &#x60;Low&#x60; - Low * &#x60;Info&#x60; - Info | [optional] 
**JiraLabels** | Pointer to **NullableString** | JIRA issue labels space seperated | [optional] 
**AddVulnerabilityIdToJiraLabel** | Pointer to **bool** |  | [optional] 
**EnableGithub** | Pointer to **bool** |  | [optional] 
**EnableSlackNotifications** | Pointer to **bool** |  | [optional] 
**SlackChannel** | Pointer to **string** | Optional. Needed if you want to send global notifications. | [optional] 
**SlackToken** | Pointer to **string** | Token required for interacting with Slack. Get one at https://api.slack.com/tokens | [optional] 
**SlackUsername** | Pointer to **string** | Optional. Will take your bot name otherwise. | [optional] 
**EnableMsteamsNotifications** | Pointer to **bool** |  | [optional] 
**MsteamsUrl** | Pointer to **string** | The full URL of the incoming webhook | [optional] 
**EnableMailNotifications** | Pointer to **bool** |  | [optional] 
**MailNotificationsTo** | Pointer to **string** |  | [optional] 
**FalsePositiveHistory** | Pointer to **bool** | (EXPERIMENTAL) DefectDojo will automatically mark the finding as a false positive if an equal finding (according to its dedupe algorithm) has been previously marked as a false positive on the same product. ATTENTION: Although the deduplication algorithm is used to determine if a finding should be marked as a false positive, this feature will not work if deduplication is enabled since it doesn&#39;t make sense to use both. | [optional] 
**RetroactiveFalsePositiveHistory** | Pointer to **bool** | (EXPERIMENTAL) FP History will also retroactively mark/unmark all existing equal findings in the same product as a false positives. Only works if the False Positive History feature is also enabled. | [optional] 
**UrlPrefix** | Pointer to **string** | URL prefix if DefectDojo is installed in it&#39;s own virtual subdirectory. | [optional] 
**TeamName** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** | * &#x60;Africa/Abidjan&#x60; - Africa/Abidjan * &#x60;Africa/Accra&#x60; - Africa/Accra * &#x60;Africa/Addis_Ababa&#x60; - Africa/Addis_Ababa * &#x60;Africa/Algiers&#x60; - Africa/Algiers * &#x60;Africa/Asmara&#x60; - Africa/Asmara * &#x60;Africa/Asmera&#x60; - Africa/Asmera * &#x60;Africa/Bamako&#x60; - Africa/Bamako * &#x60;Africa/Bangui&#x60; - Africa/Bangui * &#x60;Africa/Banjul&#x60; - Africa/Banjul * &#x60;Africa/Bissau&#x60; - Africa/Bissau * &#x60;Africa/Blantyre&#x60; - Africa/Blantyre * &#x60;Africa/Brazzaville&#x60; - Africa/Brazzaville * &#x60;Africa/Bujumbura&#x60; - Africa/Bujumbura * &#x60;Africa/Cairo&#x60; - Africa/Cairo * &#x60;Africa/Casablanca&#x60; - Africa/Casablanca * &#x60;Africa/Ceuta&#x60; - Africa/Ceuta * &#x60;Africa/Conakry&#x60; - Africa/Conakry * &#x60;Africa/Dakar&#x60; - Africa/Dakar * &#x60;Africa/Dar_es_Salaam&#x60; - Africa/Dar_es_Salaam * &#x60;Africa/Djibouti&#x60; - Africa/Djibouti * &#x60;Africa/Douala&#x60; - Africa/Douala * &#x60;Africa/El_Aaiun&#x60; - Africa/El_Aaiun * &#x60;Africa/Freetown&#x60; - Africa/Freetown * &#x60;Africa/Gaborone&#x60; - Africa/Gaborone * &#x60;Africa/Harare&#x60; - Africa/Harare * &#x60;Africa/Johannesburg&#x60; - Africa/Johannesburg * &#x60;Africa/Juba&#x60; - Africa/Juba * &#x60;Africa/Kampala&#x60; - Africa/Kampala * &#x60;Africa/Khartoum&#x60; - Africa/Khartoum * &#x60;Africa/Kigali&#x60; - Africa/Kigali * &#x60;Africa/Kinshasa&#x60; - Africa/Kinshasa * &#x60;Africa/Lagos&#x60; - Africa/Lagos * &#x60;Africa/Libreville&#x60; - Africa/Libreville * &#x60;Africa/Lome&#x60; - Africa/Lome * &#x60;Africa/Luanda&#x60; - Africa/Luanda * &#x60;Africa/Lubumbashi&#x60; - Africa/Lubumbashi * &#x60;Africa/Lusaka&#x60; - Africa/Lusaka * &#x60;Africa/Malabo&#x60; - Africa/Malabo * &#x60;Africa/Maputo&#x60; - Africa/Maputo * &#x60;Africa/Maseru&#x60; - Africa/Maseru * &#x60;Africa/Mbabane&#x60; - Africa/Mbabane * &#x60;Africa/Mogadishu&#x60; - Africa/Mogadishu * &#x60;Africa/Monrovia&#x60; - Africa/Monrovia * &#x60;Africa/Nairobi&#x60; - Africa/Nairobi * &#x60;Africa/Ndjamena&#x60; - Africa/Ndjamena * &#x60;Africa/Niamey&#x60; - Africa/Niamey * &#x60;Africa/Nouakchott&#x60; - Africa/Nouakchott * &#x60;Africa/Ouagadougou&#x60; - Africa/Ouagadougou * &#x60;Africa/Porto-Novo&#x60; - Africa/Porto-Novo * &#x60;Africa/Sao_Tome&#x60; - Africa/Sao_Tome * &#x60;Africa/Timbuktu&#x60; - Africa/Timbuktu * &#x60;Africa/Tripoli&#x60; - Africa/Tripoli * &#x60;Africa/Tunis&#x60; - Africa/Tunis * &#x60;Africa/Windhoek&#x60; - Africa/Windhoek * &#x60;America/Adak&#x60; - America/Adak * &#x60;America/Anchorage&#x60; - America/Anchorage * &#x60;America/Anguilla&#x60; - America/Anguilla * &#x60;America/Antigua&#x60; - America/Antigua * &#x60;America/Araguaina&#x60; - America/Araguaina * &#x60;America/Argentina/Buenos_Aires&#x60; - America/Argentina/Buenos_Aires * &#x60;America/Argentina/Catamarca&#x60; - America/Argentina/Catamarca * &#x60;America/Argentina/ComodRivadavia&#x60; - America/Argentina/ComodRivadavia * &#x60;America/Argentina/Cordoba&#x60; - America/Argentina/Cordoba * &#x60;America/Argentina/Jujuy&#x60; - America/Argentina/Jujuy * &#x60;America/Argentina/La_Rioja&#x60; - America/Argentina/La_Rioja * &#x60;America/Argentina/Mendoza&#x60; - America/Argentina/Mendoza * &#x60;America/Argentina/Rio_Gallegos&#x60; - America/Argentina/Rio_Gallegos * &#x60;America/Argentina/Salta&#x60; - America/Argentina/Salta * &#x60;America/Argentina/San_Juan&#x60; - America/Argentina/San_Juan * &#x60;America/Argentina/San_Luis&#x60; - America/Argentina/San_Luis * &#x60;America/Argentina/Tucuman&#x60; - America/Argentina/Tucuman * &#x60;America/Argentina/Ushuaia&#x60; - America/Argentina/Ushuaia * &#x60;America/Aruba&#x60; - America/Aruba * &#x60;America/Asuncion&#x60; - America/Asuncion * &#x60;America/Atikokan&#x60; - America/Atikokan * &#x60;America/Atka&#x60; - America/Atka * &#x60;America/Bahia&#x60; - America/Bahia * &#x60;America/Bahia_Banderas&#x60; - America/Bahia_Banderas * &#x60;America/Barbados&#x60; - America/Barbados * &#x60;America/Belem&#x60; - America/Belem * &#x60;America/Belize&#x60; - America/Belize * &#x60;America/Blanc-Sablon&#x60; - America/Blanc-Sablon * &#x60;America/Boa_Vista&#x60; - America/Boa_Vista * &#x60;America/Bogota&#x60; - America/Bogota * &#x60;America/Boise&#x60; - America/Boise * &#x60;America/Buenos_Aires&#x60; - America/Buenos_Aires * &#x60;America/Cambridge_Bay&#x60; - America/Cambridge_Bay * &#x60;America/Campo_Grande&#x60; - America/Campo_Grande * &#x60;America/Cancun&#x60; - America/Cancun * &#x60;America/Caracas&#x60; - America/Caracas * &#x60;America/Catamarca&#x60; - America/Catamarca * &#x60;America/Cayenne&#x60; - America/Cayenne * &#x60;America/Cayman&#x60; - America/Cayman * &#x60;America/Chicago&#x60; - America/Chicago * &#x60;America/Chihuahua&#x60; - America/Chihuahua * &#x60;America/Ciudad_Juarez&#x60; - America/Ciudad_Juarez * &#x60;America/Coral_Harbour&#x60; - America/Coral_Harbour * &#x60;America/Cordoba&#x60; - America/Cordoba * &#x60;America/Costa_Rica&#x60; - America/Costa_Rica * &#x60;America/Creston&#x60; - America/Creston * &#x60;America/Cuiaba&#x60; - America/Cuiaba * &#x60;America/Curacao&#x60; - America/Curacao * &#x60;America/Danmarkshavn&#x60; - America/Danmarkshavn * &#x60;America/Dawson&#x60; - America/Dawson * &#x60;America/Dawson_Creek&#x60; - America/Dawson_Creek * &#x60;America/Denver&#x60; - America/Denver * &#x60;America/Detroit&#x60; - America/Detroit * &#x60;America/Dominica&#x60; - America/Dominica * &#x60;America/Edmonton&#x60; - America/Edmonton * &#x60;America/Eirunepe&#x60; - America/Eirunepe * &#x60;America/El_Salvador&#x60; - America/El_Salvador * &#x60;America/Ensenada&#x60; - America/Ensenada * &#x60;America/Fort_Nelson&#x60; - America/Fort_Nelson * &#x60;America/Fort_Wayne&#x60; - America/Fort_Wayne * &#x60;America/Fortaleza&#x60; - America/Fortaleza * &#x60;America/Glace_Bay&#x60; - America/Glace_Bay * &#x60;America/Godthab&#x60; - America/Godthab * &#x60;America/Goose_Bay&#x60; - America/Goose_Bay * &#x60;America/Grand_Turk&#x60; - America/Grand_Turk * &#x60;America/Grenada&#x60; - America/Grenada * &#x60;America/Guadeloupe&#x60; - America/Guadeloupe * &#x60;America/Guatemala&#x60; - America/Guatemala * &#x60;America/Guayaquil&#x60; - America/Guayaquil * &#x60;America/Guyana&#x60; - America/Guyana * &#x60;America/Halifax&#x60; - America/Halifax * &#x60;America/Havana&#x60; - America/Havana * &#x60;America/Hermosillo&#x60; - America/Hermosillo * &#x60;America/Indiana/Indianapolis&#x60; - America/Indiana/Indianapolis * &#x60;America/Indiana/Knox&#x60; - America/Indiana/Knox * &#x60;America/Indiana/Marengo&#x60; - America/Indiana/Marengo * &#x60;America/Indiana/Petersburg&#x60; - America/Indiana/Petersburg * &#x60;America/Indiana/Tell_City&#x60; - America/Indiana/Tell_City * &#x60;America/Indiana/Vevay&#x60; - America/Indiana/Vevay * &#x60;America/Indiana/Vincennes&#x60; - America/Indiana/Vincennes * &#x60;America/Indiana/Winamac&#x60; - America/Indiana/Winamac * &#x60;America/Indianapolis&#x60; - America/Indianapolis * &#x60;America/Inuvik&#x60; - America/Inuvik * &#x60;America/Iqaluit&#x60; - America/Iqaluit * &#x60;America/Jamaica&#x60; - America/Jamaica * &#x60;America/Jujuy&#x60; - America/Jujuy * &#x60;America/Juneau&#x60; - America/Juneau * &#x60;America/Kentucky/Louisville&#x60; - America/Kentucky/Louisville * &#x60;America/Kentucky/Monticello&#x60; - America/Kentucky/Monticello * &#x60;America/Knox_IN&#x60; - America/Knox_IN * &#x60;America/Kralendijk&#x60; - America/Kralendijk * &#x60;America/La_Paz&#x60; - America/La_Paz * &#x60;America/Lima&#x60; - America/Lima * &#x60;America/Los_Angeles&#x60; - America/Los_Angeles * &#x60;America/Louisville&#x60; - America/Louisville * &#x60;America/Lower_Princes&#x60; - America/Lower_Princes * &#x60;America/Maceio&#x60; - America/Maceio * &#x60;America/Managua&#x60; - America/Managua * &#x60;America/Manaus&#x60; - America/Manaus * &#x60;America/Marigot&#x60; - America/Marigot * &#x60;America/Martinique&#x60; - America/Martinique * &#x60;America/Matamoros&#x60; - America/Matamoros * &#x60;America/Mazatlan&#x60; - America/Mazatlan * &#x60;America/Mendoza&#x60; - America/Mendoza * &#x60;America/Menominee&#x60; - America/Menominee * &#x60;America/Merida&#x60; - America/Merida * &#x60;America/Metlakatla&#x60; - America/Metlakatla * &#x60;America/Mexico_City&#x60; - America/Mexico_City * &#x60;America/Miquelon&#x60; - America/Miquelon * &#x60;America/Moncton&#x60; - America/Moncton * &#x60;America/Monterrey&#x60; - America/Monterrey * &#x60;America/Montevideo&#x60; - America/Montevideo * &#x60;America/Montreal&#x60; - America/Montreal * &#x60;America/Montserrat&#x60; - America/Montserrat * &#x60;America/Nassau&#x60; - America/Nassau * &#x60;America/New_York&#x60; - America/New_York * &#x60;America/Nipigon&#x60; - America/Nipigon * &#x60;America/Nome&#x60; - America/Nome * &#x60;America/Noronha&#x60; - America/Noronha * &#x60;America/North_Dakota/Beulah&#x60; - America/North_Dakota/Beulah * &#x60;America/North_Dakota/Center&#x60; - America/North_Dakota/Center * &#x60;America/North_Dakota/New_Salem&#x60; - America/North_Dakota/New_Salem * &#x60;America/Nuuk&#x60; - America/Nuuk * &#x60;America/Ojinaga&#x60; - America/Ojinaga * &#x60;America/Panama&#x60; - America/Panama * &#x60;America/Pangnirtung&#x60; - America/Pangnirtung * &#x60;America/Paramaribo&#x60; - America/Paramaribo * &#x60;America/Phoenix&#x60; - America/Phoenix * &#x60;America/Port-au-Prince&#x60; - America/Port-au-Prince * &#x60;America/Port_of_Spain&#x60; - America/Port_of_Spain * &#x60;America/Porto_Acre&#x60; - America/Porto_Acre * &#x60;America/Porto_Velho&#x60; - America/Porto_Velho * &#x60;America/Puerto_Rico&#x60; - America/Puerto_Rico * &#x60;America/Punta_Arenas&#x60; - America/Punta_Arenas * &#x60;America/Rainy_River&#x60; - America/Rainy_River * &#x60;America/Rankin_Inlet&#x60; - America/Rankin_Inlet * &#x60;America/Recife&#x60; - America/Recife * &#x60;America/Regina&#x60; - America/Regina * &#x60;America/Resolute&#x60; - America/Resolute * &#x60;America/Rio_Branco&#x60; - America/Rio_Branco * &#x60;America/Rosario&#x60; - America/Rosario * &#x60;America/Santa_Isabel&#x60; - America/Santa_Isabel * &#x60;America/Santarem&#x60; - America/Santarem * &#x60;America/Santiago&#x60; - America/Santiago * &#x60;America/Santo_Domingo&#x60; - America/Santo_Domingo * &#x60;America/Sao_Paulo&#x60; - America/Sao_Paulo * &#x60;America/Scoresbysund&#x60; - America/Scoresbysund * &#x60;America/Shiprock&#x60; - America/Shiprock * &#x60;America/Sitka&#x60; - America/Sitka * &#x60;America/St_Barthelemy&#x60; - America/St_Barthelemy * &#x60;America/St_Johns&#x60; - America/St_Johns * &#x60;America/St_Kitts&#x60; - America/St_Kitts * &#x60;America/St_Lucia&#x60; - America/St_Lucia * &#x60;America/St_Thomas&#x60; - America/St_Thomas * &#x60;America/St_Vincent&#x60; - America/St_Vincent * &#x60;America/Swift_Current&#x60; - America/Swift_Current * &#x60;America/Tegucigalpa&#x60; - America/Tegucigalpa * &#x60;America/Thule&#x60; - America/Thule * &#x60;America/Thunder_Bay&#x60; - America/Thunder_Bay * &#x60;America/Tijuana&#x60; - America/Tijuana * &#x60;America/Toronto&#x60; - America/Toronto * &#x60;America/Tortola&#x60; - America/Tortola * &#x60;America/Vancouver&#x60; - America/Vancouver * &#x60;America/Virgin&#x60; - America/Virgin * &#x60;America/Whitehorse&#x60; - America/Whitehorse * &#x60;America/Winnipeg&#x60; - America/Winnipeg * &#x60;America/Yakutat&#x60; - America/Yakutat * &#x60;America/Yellowknife&#x60; - America/Yellowknife * &#x60;Antarctica/Casey&#x60; - Antarctica/Casey * &#x60;Antarctica/Davis&#x60; - Antarctica/Davis * &#x60;Antarctica/DumontDUrville&#x60; - Antarctica/DumontDUrville * &#x60;Antarctica/Macquarie&#x60; - Antarctica/Macquarie * &#x60;Antarctica/Mawson&#x60; - Antarctica/Mawson * &#x60;Antarctica/McMurdo&#x60; - Antarctica/McMurdo * &#x60;Antarctica/Palmer&#x60; - Antarctica/Palmer * &#x60;Antarctica/Rothera&#x60; - Antarctica/Rothera * &#x60;Antarctica/South_Pole&#x60; - Antarctica/South_Pole * &#x60;Antarctica/Syowa&#x60; - Antarctica/Syowa * &#x60;Antarctica/Troll&#x60; - Antarctica/Troll * &#x60;Antarctica/Vostok&#x60; - Antarctica/Vostok * &#x60;Arctic/Longyearbyen&#x60; - Arctic/Longyearbyen * &#x60;Asia/Aden&#x60; - Asia/Aden * &#x60;Asia/Almaty&#x60; - Asia/Almaty * &#x60;Asia/Amman&#x60; - Asia/Amman * &#x60;Asia/Anadyr&#x60; - Asia/Anadyr * &#x60;Asia/Aqtau&#x60; - Asia/Aqtau * &#x60;Asia/Aqtobe&#x60; - Asia/Aqtobe * &#x60;Asia/Ashgabat&#x60; - Asia/Ashgabat * &#x60;Asia/Ashkhabad&#x60; - Asia/Ashkhabad * &#x60;Asia/Atyrau&#x60; - Asia/Atyrau * &#x60;Asia/Baghdad&#x60; - Asia/Baghdad * &#x60;Asia/Bahrain&#x60; - Asia/Bahrain * &#x60;Asia/Baku&#x60; - Asia/Baku * &#x60;Asia/Bangkok&#x60; - Asia/Bangkok * &#x60;Asia/Barnaul&#x60; - Asia/Barnaul * &#x60;Asia/Beirut&#x60; - Asia/Beirut * &#x60;Asia/Bishkek&#x60; - Asia/Bishkek * &#x60;Asia/Brunei&#x60; - Asia/Brunei * &#x60;Asia/Calcutta&#x60; - Asia/Calcutta * &#x60;Asia/Chita&#x60; - Asia/Chita * &#x60;Asia/Choibalsan&#x60; - Asia/Choibalsan * &#x60;Asia/Chongqing&#x60; - Asia/Chongqing * &#x60;Asia/Chungking&#x60; - Asia/Chungking * &#x60;Asia/Colombo&#x60; - Asia/Colombo * &#x60;Asia/Dacca&#x60; - Asia/Dacca * &#x60;Asia/Damascus&#x60; - Asia/Damascus * &#x60;Asia/Dhaka&#x60; - Asia/Dhaka * &#x60;Asia/Dili&#x60; - Asia/Dili * &#x60;Asia/Dubai&#x60; - Asia/Dubai * &#x60;Asia/Dushanbe&#x60; - Asia/Dushanbe * &#x60;Asia/Famagusta&#x60; - Asia/Famagusta * &#x60;Asia/Gaza&#x60; - Asia/Gaza * &#x60;Asia/Harbin&#x60; - Asia/Harbin * &#x60;Asia/Hebron&#x60; - Asia/Hebron * &#x60;Asia/Ho_Chi_Minh&#x60; - Asia/Ho_Chi_Minh * &#x60;Asia/Hong_Kong&#x60; - Asia/Hong_Kong * &#x60;Asia/Hovd&#x60; - Asia/Hovd * &#x60;Asia/Irkutsk&#x60; - Asia/Irkutsk * &#x60;Asia/Istanbul&#x60; - Asia/Istanbul * &#x60;Asia/Jakarta&#x60; - Asia/Jakarta * &#x60;Asia/Jayapura&#x60; - Asia/Jayapura * &#x60;Asia/Jerusalem&#x60; - Asia/Jerusalem * &#x60;Asia/Kabul&#x60; - Asia/Kabul * &#x60;Asia/Kamchatka&#x60; - Asia/Kamchatka * &#x60;Asia/Karachi&#x60; - Asia/Karachi * &#x60;Asia/Kashgar&#x60; - Asia/Kashgar * &#x60;Asia/Kathmandu&#x60; - Asia/Kathmandu * &#x60;Asia/Katmandu&#x60; - Asia/Katmandu * &#x60;Asia/Khandyga&#x60; - Asia/Khandyga * &#x60;Asia/Kolkata&#x60; - Asia/Kolkata * &#x60;Asia/Krasnoyarsk&#x60; - Asia/Krasnoyarsk * &#x60;Asia/Kuala_Lumpur&#x60; - Asia/Kuala_Lumpur * &#x60;Asia/Kuching&#x60; - Asia/Kuching * &#x60;Asia/Kuwait&#x60; - Asia/Kuwait * &#x60;Asia/Macao&#x60; - Asia/Macao * &#x60;Asia/Macau&#x60; - Asia/Macau * &#x60;Asia/Magadan&#x60; - Asia/Magadan * &#x60;Asia/Makassar&#x60; - Asia/Makassar * &#x60;Asia/Manila&#x60; - Asia/Manila * &#x60;Asia/Muscat&#x60; - Asia/Muscat * &#x60;Asia/Nicosia&#x60; - Asia/Nicosia * &#x60;Asia/Novokuznetsk&#x60; - Asia/Novokuznetsk * &#x60;Asia/Novosibirsk&#x60; - Asia/Novosibirsk * &#x60;Asia/Omsk&#x60; - Asia/Omsk * &#x60;Asia/Oral&#x60; - Asia/Oral * &#x60;Asia/Phnom_Penh&#x60; - Asia/Phnom_Penh * &#x60;Asia/Pontianak&#x60; - Asia/Pontianak * &#x60;Asia/Pyongyang&#x60; - Asia/Pyongyang * &#x60;Asia/Qatar&#x60; - Asia/Qatar * &#x60;Asia/Qostanay&#x60; - Asia/Qostanay * &#x60;Asia/Qyzylorda&#x60; - Asia/Qyzylorda * &#x60;Asia/Rangoon&#x60; - Asia/Rangoon * &#x60;Asia/Riyadh&#x60; - Asia/Riyadh * &#x60;Asia/Saigon&#x60; - Asia/Saigon * &#x60;Asia/Sakhalin&#x60; - Asia/Sakhalin * &#x60;Asia/Samarkand&#x60; - Asia/Samarkand * &#x60;Asia/Seoul&#x60; - Asia/Seoul * &#x60;Asia/Shanghai&#x60; - Asia/Shanghai * &#x60;Asia/Singapore&#x60; - Asia/Singapore * &#x60;Asia/Srednekolymsk&#x60; - Asia/Srednekolymsk * &#x60;Asia/Taipei&#x60; - Asia/Taipei * &#x60;Asia/Tashkent&#x60; - Asia/Tashkent * &#x60;Asia/Tbilisi&#x60; - Asia/Tbilisi * &#x60;Asia/Tehran&#x60; - Asia/Tehran * &#x60;Asia/Tel_Aviv&#x60; - Asia/Tel_Aviv * &#x60;Asia/Thimbu&#x60; - Asia/Thimbu * &#x60;Asia/Thimphu&#x60; - Asia/Thimphu * &#x60;Asia/Tokyo&#x60; - Asia/Tokyo * &#x60;Asia/Tomsk&#x60; - Asia/Tomsk * &#x60;Asia/Ujung_Pandang&#x60; - Asia/Ujung_Pandang * &#x60;Asia/Ulaanbaatar&#x60; - Asia/Ulaanbaatar * &#x60;Asia/Ulan_Bator&#x60; - Asia/Ulan_Bator * &#x60;Asia/Urumqi&#x60; - Asia/Urumqi * &#x60;Asia/Ust-Nera&#x60; - Asia/Ust-Nera * &#x60;Asia/Vientiane&#x60; - Asia/Vientiane * &#x60;Asia/Vladivostok&#x60; - Asia/Vladivostok * &#x60;Asia/Yakutsk&#x60; - Asia/Yakutsk * &#x60;Asia/Yangon&#x60; - Asia/Yangon * &#x60;Asia/Yekaterinburg&#x60; - Asia/Yekaterinburg * &#x60;Asia/Yerevan&#x60; - Asia/Yerevan * &#x60;Atlantic/Azores&#x60; - Atlantic/Azores * &#x60;Atlantic/Bermuda&#x60; - Atlantic/Bermuda * &#x60;Atlantic/Canary&#x60; - Atlantic/Canary * &#x60;Atlantic/Cape_Verde&#x60; - Atlantic/Cape_Verde * &#x60;Atlantic/Faeroe&#x60; - Atlantic/Faeroe * &#x60;Atlantic/Faroe&#x60; - Atlantic/Faroe * &#x60;Atlantic/Jan_Mayen&#x60; - Atlantic/Jan_Mayen * &#x60;Atlantic/Madeira&#x60; - Atlantic/Madeira * &#x60;Atlantic/Reykjavik&#x60; - Atlantic/Reykjavik * &#x60;Atlantic/South_Georgia&#x60; - Atlantic/South_Georgia * &#x60;Atlantic/St_Helena&#x60; - Atlantic/St_Helena * &#x60;Atlantic/Stanley&#x60; - Atlantic/Stanley * &#x60;Australia/ACT&#x60; - Australia/ACT * &#x60;Australia/Adelaide&#x60; - Australia/Adelaide * &#x60;Australia/Brisbane&#x60; - Australia/Brisbane * &#x60;Australia/Broken_Hill&#x60; - Australia/Broken_Hill * &#x60;Australia/Canberra&#x60; - Australia/Canberra * &#x60;Australia/Currie&#x60; - Australia/Currie * &#x60;Australia/Darwin&#x60; - Australia/Darwin * &#x60;Australia/Eucla&#x60; - Australia/Eucla * &#x60;Australia/Hobart&#x60; - Australia/Hobart * &#x60;Australia/LHI&#x60; - Australia/LHI * &#x60;Australia/Lindeman&#x60; - Australia/Lindeman * &#x60;Australia/Lord_Howe&#x60; - Australia/Lord_Howe * &#x60;Australia/Melbourne&#x60; - Australia/Melbourne * &#x60;Australia/NSW&#x60; - Australia/NSW * &#x60;Australia/North&#x60; - Australia/North * &#x60;Australia/Perth&#x60; - Australia/Perth * &#x60;Australia/Queensland&#x60; - Australia/Queensland * &#x60;Australia/South&#x60; - Australia/South * &#x60;Australia/Sydney&#x60; - Australia/Sydney * &#x60;Australia/Tasmania&#x60; - Australia/Tasmania * &#x60;Australia/Victoria&#x60; - Australia/Victoria * &#x60;Australia/West&#x60; - Australia/West * &#x60;Australia/Yancowinna&#x60; - Australia/Yancowinna * &#x60;Brazil/Acre&#x60; - Brazil/Acre * &#x60;Brazil/DeNoronha&#x60; - Brazil/DeNoronha * &#x60;Brazil/East&#x60; - Brazil/East * &#x60;Brazil/West&#x60; - Brazil/West * &#x60;CET&#x60; - CET * &#x60;CST6CDT&#x60; - CST6CDT * &#x60;Canada/Atlantic&#x60; - Canada/Atlantic * &#x60;Canada/Central&#x60; - Canada/Central * &#x60;Canada/Eastern&#x60; - Canada/Eastern * &#x60;Canada/Mountain&#x60; - Canada/Mountain * &#x60;Canada/Newfoundland&#x60; - Canada/Newfoundland * &#x60;Canada/Pacific&#x60; - Canada/Pacific * &#x60;Canada/Saskatchewan&#x60; - Canada/Saskatchewan * &#x60;Canada/Yukon&#x60; - Canada/Yukon * &#x60;Chile/Continental&#x60; - Chile/Continental * &#x60;Chile/EasterIsland&#x60; - Chile/EasterIsland * &#x60;Cuba&#x60; - Cuba * &#x60;EET&#x60; - EET * &#x60;EST&#x60; - EST * &#x60;EST5EDT&#x60; - EST5EDT * &#x60;Egypt&#x60; - Egypt * &#x60;Eire&#x60; - Eire * &#x60;Etc/GMT&#x60; - Etc/GMT * &#x60;Etc/GMT+0&#x60; - Etc/GMT+0 * &#x60;Etc/GMT+1&#x60; - Etc/GMT+1 * &#x60;Etc/GMT+10&#x60; - Etc/GMT+10 * &#x60;Etc/GMT+11&#x60; - Etc/GMT+11 * &#x60;Etc/GMT+12&#x60; - Etc/GMT+12 * &#x60;Etc/GMT+2&#x60; - Etc/GMT+2 * &#x60;Etc/GMT+3&#x60; - Etc/GMT+3 * &#x60;Etc/GMT+4&#x60; - Etc/GMT+4 * &#x60;Etc/GMT+5&#x60; - Etc/GMT+5 * &#x60;Etc/GMT+6&#x60; - Etc/GMT+6 * &#x60;Etc/GMT+7&#x60; - Etc/GMT+7 * &#x60;Etc/GMT+8&#x60; - Etc/GMT+8 * &#x60;Etc/GMT+9&#x60; - Etc/GMT+9 * &#x60;Etc/GMT-0&#x60; - Etc/GMT-0 * &#x60;Etc/GMT-1&#x60; - Etc/GMT-1 * &#x60;Etc/GMT-10&#x60; - Etc/GMT-10 * &#x60;Etc/GMT-11&#x60; - Etc/GMT-11 * &#x60;Etc/GMT-12&#x60; - Etc/GMT-12 * &#x60;Etc/GMT-13&#x60; - Etc/GMT-13 * &#x60;Etc/GMT-14&#x60; - Etc/GMT-14 * &#x60;Etc/GMT-2&#x60; - Etc/GMT-2 * &#x60;Etc/GMT-3&#x60; - Etc/GMT-3 * &#x60;Etc/GMT-4&#x60; - Etc/GMT-4 * &#x60;Etc/GMT-5&#x60; - Etc/GMT-5 * &#x60;Etc/GMT-6&#x60; - Etc/GMT-6 * &#x60;Etc/GMT-7&#x60; - Etc/GMT-7 * &#x60;Etc/GMT-8&#x60; - Etc/GMT-8 * &#x60;Etc/GMT-9&#x60; - Etc/GMT-9 * &#x60;Etc/GMT0&#x60; - Etc/GMT0 * &#x60;Etc/Greenwich&#x60; - Etc/Greenwich * &#x60;Etc/UCT&#x60; - Etc/UCT * &#x60;Etc/UTC&#x60; - Etc/UTC * &#x60;Etc/Universal&#x60; - Etc/Universal * &#x60;Etc/Zulu&#x60; - Etc/Zulu * &#x60;Europe/Amsterdam&#x60; - Europe/Amsterdam * &#x60;Europe/Andorra&#x60; - Europe/Andorra * &#x60;Europe/Astrakhan&#x60; - Europe/Astrakhan * &#x60;Europe/Athens&#x60; - Europe/Athens * &#x60;Europe/Belfast&#x60; - Europe/Belfast * &#x60;Europe/Belgrade&#x60; - Europe/Belgrade * &#x60;Europe/Berlin&#x60; - Europe/Berlin * &#x60;Europe/Bratislava&#x60; - Europe/Bratislava * &#x60;Europe/Brussels&#x60; - Europe/Brussels * &#x60;Europe/Bucharest&#x60; - Europe/Bucharest * &#x60;Europe/Budapest&#x60; - Europe/Budapest * &#x60;Europe/Busingen&#x60; - Europe/Busingen * &#x60;Europe/Chisinau&#x60; - Europe/Chisinau * &#x60;Europe/Copenhagen&#x60; - Europe/Copenhagen * &#x60;Europe/Dublin&#x60; - Europe/Dublin * &#x60;Europe/Gibraltar&#x60; - Europe/Gibraltar * &#x60;Europe/Guernsey&#x60; - Europe/Guernsey * &#x60;Europe/Helsinki&#x60; - Europe/Helsinki * &#x60;Europe/Isle_of_Man&#x60; - Europe/Isle_of_Man * &#x60;Europe/Istanbul&#x60; - Europe/Istanbul * &#x60;Europe/Jersey&#x60; - Europe/Jersey * &#x60;Europe/Kaliningrad&#x60; - Europe/Kaliningrad * &#x60;Europe/Kiev&#x60; - Europe/Kiev * &#x60;Europe/Kirov&#x60; - Europe/Kirov * &#x60;Europe/Kyiv&#x60; - Europe/Kyiv * &#x60;Europe/Lisbon&#x60; - Europe/Lisbon * &#x60;Europe/Ljubljana&#x60; - Europe/Ljubljana * &#x60;Europe/London&#x60; - Europe/London * &#x60;Europe/Luxembourg&#x60; - Europe/Luxembourg * &#x60;Europe/Madrid&#x60; - Europe/Madrid * &#x60;Europe/Malta&#x60; - Europe/Malta * &#x60;Europe/Mariehamn&#x60; - Europe/Mariehamn * &#x60;Europe/Minsk&#x60; - Europe/Minsk * &#x60;Europe/Monaco&#x60; - Europe/Monaco * &#x60;Europe/Moscow&#x60; - Europe/Moscow * &#x60;Europe/Nicosia&#x60; - Europe/Nicosia * &#x60;Europe/Oslo&#x60; - Europe/Oslo * &#x60;Europe/Paris&#x60; - Europe/Paris * &#x60;Europe/Podgorica&#x60; - Europe/Podgorica * &#x60;Europe/Prague&#x60; - Europe/Prague * &#x60;Europe/Riga&#x60; - Europe/Riga * &#x60;Europe/Rome&#x60; - Europe/Rome * &#x60;Europe/Samara&#x60; - Europe/Samara * &#x60;Europe/San_Marino&#x60; - Europe/San_Marino * &#x60;Europe/Sarajevo&#x60; - Europe/Sarajevo * &#x60;Europe/Saratov&#x60; - Europe/Saratov * &#x60;Europe/Simferopol&#x60; - Europe/Simferopol * &#x60;Europe/Skopje&#x60; - Europe/Skopje * &#x60;Europe/Sofia&#x60; - Europe/Sofia * &#x60;Europe/Stockholm&#x60; - Europe/Stockholm * &#x60;Europe/Tallinn&#x60; - Europe/Tallinn * &#x60;Europe/Tirane&#x60; - Europe/Tirane * &#x60;Europe/Tiraspol&#x60; - Europe/Tiraspol * &#x60;Europe/Ulyanovsk&#x60; - Europe/Ulyanovsk * &#x60;Europe/Uzhgorod&#x60; - Europe/Uzhgorod * &#x60;Europe/Vaduz&#x60; - Europe/Vaduz * &#x60;Europe/Vatican&#x60; - Europe/Vatican * &#x60;Europe/Vienna&#x60; - Europe/Vienna * &#x60;Europe/Vilnius&#x60; - Europe/Vilnius * &#x60;Europe/Volgograd&#x60; - Europe/Volgograd * &#x60;Europe/Warsaw&#x60; - Europe/Warsaw * &#x60;Europe/Zagreb&#x60; - Europe/Zagreb * &#x60;Europe/Zaporozhye&#x60; - Europe/Zaporozhye * &#x60;Europe/Zurich&#x60; - Europe/Zurich * &#x60;GB&#x60; - GB * &#x60;GB-Eire&#x60; - GB-Eire * &#x60;GMT&#x60; - GMT * &#x60;GMT+0&#x60; - GMT+0 * &#x60;GMT-0&#x60; - GMT-0 * &#x60;GMT0&#x60; - GMT0 * &#x60;Greenwich&#x60; - Greenwich * &#x60;HST&#x60; - HST * &#x60;Hongkong&#x60; - Hongkong * &#x60;Iceland&#x60; - Iceland * &#x60;Indian/Antananarivo&#x60; - Indian/Antananarivo * &#x60;Indian/Chagos&#x60; - Indian/Chagos * &#x60;Indian/Christmas&#x60; - Indian/Christmas * &#x60;Indian/Cocos&#x60; - Indian/Cocos * &#x60;Indian/Comoro&#x60; - Indian/Comoro * &#x60;Indian/Kerguelen&#x60; - Indian/Kerguelen * &#x60;Indian/Mahe&#x60; - Indian/Mahe * &#x60;Indian/Maldives&#x60; - Indian/Maldives * &#x60;Indian/Mauritius&#x60; - Indian/Mauritius * &#x60;Indian/Mayotte&#x60; - Indian/Mayotte * &#x60;Indian/Reunion&#x60; - Indian/Reunion * &#x60;Iran&#x60; - Iran * &#x60;Israel&#x60; - Israel * &#x60;Jamaica&#x60; - Jamaica * &#x60;Japan&#x60; - Japan * &#x60;Kwajalein&#x60; - Kwajalein * &#x60;Libya&#x60; - Libya * &#x60;MET&#x60; - MET * &#x60;MST&#x60; - MST * &#x60;MST7MDT&#x60; - MST7MDT * &#x60;Mexico/BajaNorte&#x60; - Mexico/BajaNorte * &#x60;Mexico/BajaSur&#x60; - Mexico/BajaSur * &#x60;Mexico/General&#x60; - Mexico/General * &#x60;NZ&#x60; - NZ * &#x60;NZ-CHAT&#x60; - NZ-CHAT * &#x60;Navajo&#x60; - Navajo * &#x60;PRC&#x60; - PRC * &#x60;PST8PDT&#x60; - PST8PDT * &#x60;Pacific/Apia&#x60; - Pacific/Apia * &#x60;Pacific/Auckland&#x60; - Pacific/Auckland * &#x60;Pacific/Bougainville&#x60; - Pacific/Bougainville * &#x60;Pacific/Chatham&#x60; - Pacific/Chatham * &#x60;Pacific/Chuuk&#x60; - Pacific/Chuuk * &#x60;Pacific/Easter&#x60; - Pacific/Easter * &#x60;Pacific/Efate&#x60; - Pacific/Efate * &#x60;Pacific/Enderbury&#x60; - Pacific/Enderbury * &#x60;Pacific/Fakaofo&#x60; - Pacific/Fakaofo * &#x60;Pacific/Fiji&#x60; - Pacific/Fiji * &#x60;Pacific/Funafuti&#x60; - Pacific/Funafuti * &#x60;Pacific/Galapagos&#x60; - Pacific/Galapagos * &#x60;Pacific/Gambier&#x60; - Pacific/Gambier * &#x60;Pacific/Guadalcanal&#x60; - Pacific/Guadalcanal * &#x60;Pacific/Guam&#x60; - Pacific/Guam * &#x60;Pacific/Honolulu&#x60; - Pacific/Honolulu * &#x60;Pacific/Johnston&#x60; - Pacific/Johnston * &#x60;Pacific/Kanton&#x60; - Pacific/Kanton * &#x60;Pacific/Kiritimati&#x60; - Pacific/Kiritimati * &#x60;Pacific/Kosrae&#x60; - Pacific/Kosrae * &#x60;Pacific/Kwajalein&#x60; - Pacific/Kwajalein * &#x60;Pacific/Majuro&#x60; - Pacific/Majuro * &#x60;Pacific/Marquesas&#x60; - Pacific/Marquesas * &#x60;Pacific/Midway&#x60; - Pacific/Midway * &#x60;Pacific/Nauru&#x60; - Pacific/Nauru * &#x60;Pacific/Niue&#x60; - Pacific/Niue * &#x60;Pacific/Norfolk&#x60; - Pacific/Norfolk * &#x60;Pacific/Noumea&#x60; - Pacific/Noumea * &#x60;Pacific/Pago_Pago&#x60; - Pacific/Pago_Pago * &#x60;Pacific/Palau&#x60; - Pacific/Palau * &#x60;Pacific/Pitcairn&#x60; - Pacific/Pitcairn * &#x60;Pacific/Pohnpei&#x60; - Pacific/Pohnpei * &#x60;Pacific/Ponape&#x60; - Pacific/Ponape * &#x60;Pacific/Port_Moresby&#x60; - Pacific/Port_Moresby * &#x60;Pacific/Rarotonga&#x60; - Pacific/Rarotonga * &#x60;Pacific/Saipan&#x60; - Pacific/Saipan * &#x60;Pacific/Samoa&#x60; - Pacific/Samoa * &#x60;Pacific/Tahiti&#x60; - Pacific/Tahiti * &#x60;Pacific/Tarawa&#x60; - Pacific/Tarawa * &#x60;Pacific/Tongatapu&#x60; - Pacific/Tongatapu * &#x60;Pacific/Truk&#x60; - Pacific/Truk * &#x60;Pacific/Wake&#x60; - Pacific/Wake * &#x60;Pacific/Wallis&#x60; - Pacific/Wallis * &#x60;Pacific/Yap&#x60; - Pacific/Yap * &#x60;Poland&#x60; - Poland * &#x60;Portugal&#x60; - Portugal * &#x60;ROC&#x60; - ROC * &#x60;ROK&#x60; - ROK * &#x60;Singapore&#x60; - Singapore * &#x60;Turkey&#x60; - Turkey * &#x60;UCT&#x60; - UCT * &#x60;US/Alaska&#x60; - US/Alaska * &#x60;US/Aleutian&#x60; - US/Aleutian * &#x60;US/Arizona&#x60; - US/Arizona * &#x60;US/Central&#x60; - US/Central * &#x60;US/East-Indiana&#x60; - US/East-Indiana * &#x60;US/Eastern&#x60; - US/Eastern * &#x60;US/Hawaii&#x60; - US/Hawaii * &#x60;US/Indiana-Starke&#x60; - US/Indiana-Starke * &#x60;US/Michigan&#x60; - US/Michigan * &#x60;US/Mountain&#x60; - US/Mountain * &#x60;US/Pacific&#x60; - US/Pacific * &#x60;US/Samoa&#x60; - US/Samoa * &#x60;UTC&#x60; - UTC * &#x60;Universal&#x60; - Universal * &#x60;W-SU&#x60; - W-SU * &#x60;WET&#x60; - WET * &#x60;Zulu&#x60; - Zulu | [optional] 
**EnableProductGrade** | Pointer to **bool** | Displays a grade letter next to a product to show the overall health. | [optional] 
**ProductGrade** | Pointer to **string** |  | [optional] 
**ProductGradeA** | Pointer to **int32** | Percentage score for an &#39;A&#39; &gt;&#x3D; | [optional] 
**ProductGradeB** | Pointer to **int32** | Percentage score for a &#39;B&#39; &gt;&#x3D; | [optional] 
**ProductGradeC** | Pointer to **int32** | Percentage score for a &#39;C&#39; &gt;&#x3D; | [optional] 
**ProductGradeD** | Pointer to **int32** | Percentage score for a &#39;D&#39; &gt;&#x3D; | [optional] 
**ProductGradeF** | Pointer to **int32** | Percentage score for an &#39;F&#39; &lt;&#x3D; | [optional] 
**EnableProductTagInheritance** | Pointer to **bool** | Enables product tag inheritance globally for all products. Any tags added on a product will automatically be added to all Engagements, Tests, and Findings | [optional] 
**EnableBenchmark** | Pointer to **bool** | Enables Benchmarks such as the OWASP ASVS (Application Security Verification Standard) | [optional] 
**EnableTemplateMatch** | Pointer to **bool** | Enables global remediation advice and matching on CWE and Title. The text will be replaced for mitigation, impact and references on a finding. Useful for providing consistent impact and remediation advice regardless of the scanner. | [optional] 
**EngagementAutoClose** | Pointer to **bool** | Closes an engagement after 3 days (default) past due date including last update. | [optional] 
**EngagementAutoCloseDays** | Pointer to **int32** | Closes an engagement after the specified number of days past due date including last update. | [optional] 
**EnableFindingSla** | Pointer to **bool** | Enables Finding SLA&#39;s for time to remediate. | [optional] 
**EnableNotifySlaActive** | Pointer to **bool** | Enables Notify when time to remediate according to Finding SLA&#39;s is breached for active Findings. | [optional] 
**EnableNotifySlaActiveVerified** | Pointer to **bool** | Enables Notify when time to remediate according to Finding SLA&#39;s is breached for active, verified Findings. | [optional] 
**EnableNotifySlaJiraOnly** | Pointer to **bool** | Enables Notify when time to remediate according to Finding SLA&#39;s is breached for Findings that are linked to JIRA issues. Notification is disabled for Findings not linked to JIRA issues | [optional] 
**EnableNotifySlaExponentialBackoff** | Pointer to **bool** | Enable an exponential backoff strategy for SLA breach notifications, e.g. 1, 2, 4, 8, etc. Otherwise it alerts every day | [optional] 
**AllowAnonymousSurveyRepsonse** | Pointer to **bool** | Enable anyone with a link to the survey to answer a survey | [optional] 
**Credentials** | Pointer to **string** |  | [optional] 
**Disclaimer** | Pointer to **string** | Include this custom disclaimer on all notifications and generated reports | [optional] 
**RiskAcceptanceFormDefaultDays** | Pointer to **NullableInt32** | Default expiry period for risk acceptance form. | [optional] 
**RiskAcceptanceNotifyBeforeExpiration** | Pointer to **NullableInt32** | Notify X days before risk acceptance expires. Leave empty to disable. | [optional] 
**EnableCredentials** | Pointer to **bool** | With this setting turned off, credentials will be disabled in the user interface. | [optional] 
**EnableQuestionnaires** | Pointer to **bool** | With this setting turned off, questionnaires will be disabled in the user interface. | [optional] 
**EnableChecklists** | Pointer to **bool** | With this setting turned off, checklists will be disabled in the user interface. | [optional] 
**EnableEndpointMetadataImport** | Pointer to **bool** | With this setting turned off, endpoint metadata import will be disabled in the user interface. | [optional] 
**EnableUserProfileEditable** | Pointer to **bool** | When turned on users can edit their profiles | [optional] 
**EnableProductTrackingFiles** | Pointer to **bool** | With this setting turned off, the product tracking files will be disabled in the user interface. | [optional] 
**EnableFindingGroups** | Pointer to **bool** | With this setting turned off, the Finding Groups will be disabled. | [optional] 
**EnableCalendar** | Pointer to **bool** | With this setting turned off, the Calendar will be disabled in the user interface. | [optional] 
**DefaultGroupEmailPattern** | Pointer to **string** | New users will only be assigned to the default group, when their email address matches this regex pattern. This is optional condition. | [optional] 
**MinimumPasswordLength** | Pointer to **int32** | Requires user to set passwords greater than minimum length. | [optional] 
**MaximumPasswordLength** | Pointer to **int32** | Requires user to set passwords less than maximum length. | [optional] 
**NumberCharacterRequired** | Pointer to **bool** | Requires user passwords to contain at least one digit (0-9). | [optional] 
**SpecialCharacterRequired** | Pointer to **bool** | Requires user passwords to contain at least one special character (()[]{}|\\&#x60;~!@#$%^&amp;*_-+&#x3D;;:&#39;\&quot;,&lt;&gt;./?). | [optional] 
**LowercaseCharacterRequired** | Pointer to **bool** | Requires user passwords to contain at least one lowercase letter (a-z). | [optional] 
**UppercaseCharacterRequired** | Pointer to **bool** | Requires user passwords to contain at least one uppercase letter (A-Z). | [optional] 
**NonCommonPasswordRequired** | Pointer to **bool** | Requires user passwords to not be part of list of common passwords. | [optional] 
**ApiExposeErrorDetails** | Pointer to **bool** | When turned on, the API will expose error details in the response. | [optional] 
**DefaultGroup** | Pointer to **NullableInt32** | New users will be assigned to this group. | [optional] 
**DefaultGroupRole** | Pointer to **NullableInt32** | New users will be assigned to their default group with this role. | [optional] 

## Methods

### NewSystemSettingsRequest

`func NewSystemSettingsRequest() *SystemSettingsRequest`

NewSystemSettingsRequest instantiates a new SystemSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemSettingsRequestWithDefaults

`func NewSystemSettingsRequestWithDefaults() *SystemSettingsRequest`

NewSystemSettingsRequestWithDefaults instantiates a new SystemSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableDeduplication

`func (o *SystemSettingsRequest) GetEnableDeduplication() bool`

GetEnableDeduplication returns the EnableDeduplication field if non-nil, zero value otherwise.

### GetEnableDeduplicationOk

`func (o *SystemSettingsRequest) GetEnableDeduplicationOk() (*bool, bool)`

GetEnableDeduplicationOk returns a tuple with the EnableDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeduplication

`func (o *SystemSettingsRequest) SetEnableDeduplication(v bool)`

SetEnableDeduplication sets EnableDeduplication field to given value.

### HasEnableDeduplication

`func (o *SystemSettingsRequest) HasEnableDeduplication() bool`

HasEnableDeduplication returns a boolean if a field has been set.

### GetDeleteDuplicates

`func (o *SystemSettingsRequest) GetDeleteDuplicates() bool`

GetDeleteDuplicates returns the DeleteDuplicates field if non-nil, zero value otherwise.

### GetDeleteDuplicatesOk

`func (o *SystemSettingsRequest) GetDeleteDuplicatesOk() (*bool, bool)`

GetDeleteDuplicatesOk returns a tuple with the DeleteDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteDuplicates

`func (o *SystemSettingsRequest) SetDeleteDuplicates(v bool)`

SetDeleteDuplicates sets DeleteDuplicates field to given value.

### HasDeleteDuplicates

`func (o *SystemSettingsRequest) HasDeleteDuplicates() bool`

HasDeleteDuplicates returns a boolean if a field has been set.

### GetMaxDupes

`func (o *SystemSettingsRequest) GetMaxDupes() int32`

GetMaxDupes returns the MaxDupes field if non-nil, zero value otherwise.

### GetMaxDupesOk

`func (o *SystemSettingsRequest) GetMaxDupesOk() (*int32, bool)`

GetMaxDupesOk returns a tuple with the MaxDupes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDupes

`func (o *SystemSettingsRequest) SetMaxDupes(v int32)`

SetMaxDupes sets MaxDupes field to given value.

### HasMaxDupes

`func (o *SystemSettingsRequest) HasMaxDupes() bool`

HasMaxDupes returns a boolean if a field has been set.

### SetMaxDupesNil

`func (o *SystemSettingsRequest) SetMaxDupesNil(b bool)`

 SetMaxDupesNil sets the value for MaxDupes to be an explicit nil

### UnsetMaxDupes
`func (o *SystemSettingsRequest) UnsetMaxDupes()`

UnsetMaxDupes ensures that no value is present for MaxDupes, not even an explicit nil
### GetEmailFrom

`func (o *SystemSettingsRequest) GetEmailFrom() string`

GetEmailFrom returns the EmailFrom field if non-nil, zero value otherwise.

### GetEmailFromOk

`func (o *SystemSettingsRequest) GetEmailFromOk() (*string, bool)`

GetEmailFromOk returns a tuple with the EmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrom

`func (o *SystemSettingsRequest) SetEmailFrom(v string)`

SetEmailFrom sets EmailFrom field to given value.

### HasEmailFrom

`func (o *SystemSettingsRequest) HasEmailFrom() bool`

HasEmailFrom returns a boolean if a field has been set.

### GetEnableJira

`func (o *SystemSettingsRequest) GetEnableJira() bool`

GetEnableJira returns the EnableJira field if non-nil, zero value otherwise.

### GetEnableJiraOk

`func (o *SystemSettingsRequest) GetEnableJiraOk() (*bool, bool)`

GetEnableJiraOk returns a tuple with the EnableJira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableJira

`func (o *SystemSettingsRequest) SetEnableJira(v bool)`

SetEnableJira sets EnableJira field to given value.

### HasEnableJira

`func (o *SystemSettingsRequest) HasEnableJira() bool`

HasEnableJira returns a boolean if a field has been set.

### GetEnableJiraWebHook

`func (o *SystemSettingsRequest) GetEnableJiraWebHook() bool`

GetEnableJiraWebHook returns the EnableJiraWebHook field if non-nil, zero value otherwise.

### GetEnableJiraWebHookOk

`func (o *SystemSettingsRequest) GetEnableJiraWebHookOk() (*bool, bool)`

GetEnableJiraWebHookOk returns a tuple with the EnableJiraWebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableJiraWebHook

`func (o *SystemSettingsRequest) SetEnableJiraWebHook(v bool)`

SetEnableJiraWebHook sets EnableJiraWebHook field to given value.

### HasEnableJiraWebHook

`func (o *SystemSettingsRequest) HasEnableJiraWebHook() bool`

HasEnableJiraWebHook returns a boolean if a field has been set.

### GetDisableJiraWebhookSecret

`func (o *SystemSettingsRequest) GetDisableJiraWebhookSecret() bool`

GetDisableJiraWebhookSecret returns the DisableJiraWebhookSecret field if non-nil, zero value otherwise.

### GetDisableJiraWebhookSecretOk

`func (o *SystemSettingsRequest) GetDisableJiraWebhookSecretOk() (*bool, bool)`

GetDisableJiraWebhookSecretOk returns a tuple with the DisableJiraWebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableJiraWebhookSecret

`func (o *SystemSettingsRequest) SetDisableJiraWebhookSecret(v bool)`

SetDisableJiraWebhookSecret sets DisableJiraWebhookSecret field to given value.

### HasDisableJiraWebhookSecret

`func (o *SystemSettingsRequest) HasDisableJiraWebhookSecret() bool`

HasDisableJiraWebhookSecret returns a boolean if a field has been set.

### GetJiraWebhookSecret

`func (o *SystemSettingsRequest) GetJiraWebhookSecret() string`

GetJiraWebhookSecret returns the JiraWebhookSecret field if non-nil, zero value otherwise.

### GetJiraWebhookSecretOk

`func (o *SystemSettingsRequest) GetJiraWebhookSecretOk() (*string, bool)`

GetJiraWebhookSecretOk returns a tuple with the JiraWebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraWebhookSecret

`func (o *SystemSettingsRequest) SetJiraWebhookSecret(v string)`

SetJiraWebhookSecret sets JiraWebhookSecret field to given value.

### HasJiraWebhookSecret

`func (o *SystemSettingsRequest) HasJiraWebhookSecret() bool`

HasJiraWebhookSecret returns a boolean if a field has been set.

### SetJiraWebhookSecretNil

`func (o *SystemSettingsRequest) SetJiraWebhookSecretNil(b bool)`

 SetJiraWebhookSecretNil sets the value for JiraWebhookSecret to be an explicit nil

### UnsetJiraWebhookSecret
`func (o *SystemSettingsRequest) UnsetJiraWebhookSecret()`

UnsetJiraWebhookSecret ensures that no value is present for JiraWebhookSecret, not even an explicit nil
### GetJiraMinimumSeverity

`func (o *SystemSettingsRequest) GetJiraMinimumSeverity() string`

GetJiraMinimumSeverity returns the JiraMinimumSeverity field if non-nil, zero value otherwise.

### GetJiraMinimumSeverityOk

`func (o *SystemSettingsRequest) GetJiraMinimumSeverityOk() (*string, bool)`

GetJiraMinimumSeverityOk returns a tuple with the JiraMinimumSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraMinimumSeverity

`func (o *SystemSettingsRequest) SetJiraMinimumSeverity(v string)`

SetJiraMinimumSeverity sets JiraMinimumSeverity field to given value.

### HasJiraMinimumSeverity

`func (o *SystemSettingsRequest) HasJiraMinimumSeverity() bool`

HasJiraMinimumSeverity returns a boolean if a field has been set.

### SetJiraMinimumSeverityNil

`func (o *SystemSettingsRequest) SetJiraMinimumSeverityNil(b bool)`

 SetJiraMinimumSeverityNil sets the value for JiraMinimumSeverity to be an explicit nil

### UnsetJiraMinimumSeverity
`func (o *SystemSettingsRequest) UnsetJiraMinimumSeverity()`

UnsetJiraMinimumSeverity ensures that no value is present for JiraMinimumSeverity, not even an explicit nil
### GetJiraLabels

`func (o *SystemSettingsRequest) GetJiraLabels() string`

GetJiraLabels returns the JiraLabels field if non-nil, zero value otherwise.

### GetJiraLabelsOk

`func (o *SystemSettingsRequest) GetJiraLabelsOk() (*string, bool)`

GetJiraLabelsOk returns a tuple with the JiraLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraLabels

`func (o *SystemSettingsRequest) SetJiraLabels(v string)`

SetJiraLabels sets JiraLabels field to given value.

### HasJiraLabels

`func (o *SystemSettingsRequest) HasJiraLabels() bool`

HasJiraLabels returns a boolean if a field has been set.

### SetJiraLabelsNil

`func (o *SystemSettingsRequest) SetJiraLabelsNil(b bool)`

 SetJiraLabelsNil sets the value for JiraLabels to be an explicit nil

### UnsetJiraLabels
`func (o *SystemSettingsRequest) UnsetJiraLabels()`

UnsetJiraLabels ensures that no value is present for JiraLabels, not even an explicit nil
### GetAddVulnerabilityIdToJiraLabel

`func (o *SystemSettingsRequest) GetAddVulnerabilityIdToJiraLabel() bool`

GetAddVulnerabilityIdToJiraLabel returns the AddVulnerabilityIdToJiraLabel field if non-nil, zero value otherwise.

### GetAddVulnerabilityIdToJiraLabelOk

`func (o *SystemSettingsRequest) GetAddVulnerabilityIdToJiraLabelOk() (*bool, bool)`

GetAddVulnerabilityIdToJiraLabelOk returns a tuple with the AddVulnerabilityIdToJiraLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVulnerabilityIdToJiraLabel

`func (o *SystemSettingsRequest) SetAddVulnerabilityIdToJiraLabel(v bool)`

SetAddVulnerabilityIdToJiraLabel sets AddVulnerabilityIdToJiraLabel field to given value.

### HasAddVulnerabilityIdToJiraLabel

`func (o *SystemSettingsRequest) HasAddVulnerabilityIdToJiraLabel() bool`

HasAddVulnerabilityIdToJiraLabel returns a boolean if a field has been set.

### GetEnableGithub

`func (o *SystemSettingsRequest) GetEnableGithub() bool`

GetEnableGithub returns the EnableGithub field if non-nil, zero value otherwise.

### GetEnableGithubOk

`func (o *SystemSettingsRequest) GetEnableGithubOk() (*bool, bool)`

GetEnableGithubOk returns a tuple with the EnableGithub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGithub

`func (o *SystemSettingsRequest) SetEnableGithub(v bool)`

SetEnableGithub sets EnableGithub field to given value.

### HasEnableGithub

`func (o *SystemSettingsRequest) HasEnableGithub() bool`

HasEnableGithub returns a boolean if a field has been set.

### GetEnableSlackNotifications

`func (o *SystemSettingsRequest) GetEnableSlackNotifications() bool`

GetEnableSlackNotifications returns the EnableSlackNotifications field if non-nil, zero value otherwise.

### GetEnableSlackNotificationsOk

`func (o *SystemSettingsRequest) GetEnableSlackNotificationsOk() (*bool, bool)`

GetEnableSlackNotificationsOk returns a tuple with the EnableSlackNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSlackNotifications

`func (o *SystemSettingsRequest) SetEnableSlackNotifications(v bool)`

SetEnableSlackNotifications sets EnableSlackNotifications field to given value.

### HasEnableSlackNotifications

`func (o *SystemSettingsRequest) HasEnableSlackNotifications() bool`

HasEnableSlackNotifications returns a boolean if a field has been set.

### GetSlackChannel

`func (o *SystemSettingsRequest) GetSlackChannel() string`

GetSlackChannel returns the SlackChannel field if non-nil, zero value otherwise.

### GetSlackChannelOk

`func (o *SystemSettingsRequest) GetSlackChannelOk() (*string, bool)`

GetSlackChannelOk returns a tuple with the SlackChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackChannel

`func (o *SystemSettingsRequest) SetSlackChannel(v string)`

SetSlackChannel sets SlackChannel field to given value.

### HasSlackChannel

`func (o *SystemSettingsRequest) HasSlackChannel() bool`

HasSlackChannel returns a boolean if a field has been set.

### GetSlackToken

`func (o *SystemSettingsRequest) GetSlackToken() string`

GetSlackToken returns the SlackToken field if non-nil, zero value otherwise.

### GetSlackTokenOk

`func (o *SystemSettingsRequest) GetSlackTokenOk() (*string, bool)`

GetSlackTokenOk returns a tuple with the SlackToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackToken

`func (o *SystemSettingsRequest) SetSlackToken(v string)`

SetSlackToken sets SlackToken field to given value.

### HasSlackToken

`func (o *SystemSettingsRequest) HasSlackToken() bool`

HasSlackToken returns a boolean if a field has been set.

### GetSlackUsername

`func (o *SystemSettingsRequest) GetSlackUsername() string`

GetSlackUsername returns the SlackUsername field if non-nil, zero value otherwise.

### GetSlackUsernameOk

`func (o *SystemSettingsRequest) GetSlackUsernameOk() (*string, bool)`

GetSlackUsernameOk returns a tuple with the SlackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUsername

`func (o *SystemSettingsRequest) SetSlackUsername(v string)`

SetSlackUsername sets SlackUsername field to given value.

### HasSlackUsername

`func (o *SystemSettingsRequest) HasSlackUsername() bool`

HasSlackUsername returns a boolean if a field has been set.

### GetEnableMsteamsNotifications

`func (o *SystemSettingsRequest) GetEnableMsteamsNotifications() bool`

GetEnableMsteamsNotifications returns the EnableMsteamsNotifications field if non-nil, zero value otherwise.

### GetEnableMsteamsNotificationsOk

`func (o *SystemSettingsRequest) GetEnableMsteamsNotificationsOk() (*bool, bool)`

GetEnableMsteamsNotificationsOk returns a tuple with the EnableMsteamsNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMsteamsNotifications

`func (o *SystemSettingsRequest) SetEnableMsteamsNotifications(v bool)`

SetEnableMsteamsNotifications sets EnableMsteamsNotifications field to given value.

### HasEnableMsteamsNotifications

`func (o *SystemSettingsRequest) HasEnableMsteamsNotifications() bool`

HasEnableMsteamsNotifications returns a boolean if a field has been set.

### GetMsteamsUrl

`func (o *SystemSettingsRequest) GetMsteamsUrl() string`

GetMsteamsUrl returns the MsteamsUrl field if non-nil, zero value otherwise.

### GetMsteamsUrlOk

`func (o *SystemSettingsRequest) GetMsteamsUrlOk() (*string, bool)`

GetMsteamsUrlOk returns a tuple with the MsteamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsteamsUrl

`func (o *SystemSettingsRequest) SetMsteamsUrl(v string)`

SetMsteamsUrl sets MsteamsUrl field to given value.

### HasMsteamsUrl

`func (o *SystemSettingsRequest) HasMsteamsUrl() bool`

HasMsteamsUrl returns a boolean if a field has been set.

### GetEnableMailNotifications

`func (o *SystemSettingsRequest) GetEnableMailNotifications() bool`

GetEnableMailNotifications returns the EnableMailNotifications field if non-nil, zero value otherwise.

### GetEnableMailNotificationsOk

`func (o *SystemSettingsRequest) GetEnableMailNotificationsOk() (*bool, bool)`

GetEnableMailNotificationsOk returns a tuple with the EnableMailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMailNotifications

`func (o *SystemSettingsRequest) SetEnableMailNotifications(v bool)`

SetEnableMailNotifications sets EnableMailNotifications field to given value.

### HasEnableMailNotifications

`func (o *SystemSettingsRequest) HasEnableMailNotifications() bool`

HasEnableMailNotifications returns a boolean if a field has been set.

### GetMailNotificationsTo

`func (o *SystemSettingsRequest) GetMailNotificationsTo() string`

GetMailNotificationsTo returns the MailNotificationsTo field if non-nil, zero value otherwise.

### GetMailNotificationsToOk

`func (o *SystemSettingsRequest) GetMailNotificationsToOk() (*string, bool)`

GetMailNotificationsToOk returns a tuple with the MailNotificationsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailNotificationsTo

`func (o *SystemSettingsRequest) SetMailNotificationsTo(v string)`

SetMailNotificationsTo sets MailNotificationsTo field to given value.

### HasMailNotificationsTo

`func (o *SystemSettingsRequest) HasMailNotificationsTo() bool`

HasMailNotificationsTo returns a boolean if a field has been set.

### GetFalsePositiveHistory

`func (o *SystemSettingsRequest) GetFalsePositiveHistory() bool`

GetFalsePositiveHistory returns the FalsePositiveHistory field if non-nil, zero value otherwise.

### GetFalsePositiveHistoryOk

`func (o *SystemSettingsRequest) GetFalsePositiveHistoryOk() (*bool, bool)`

GetFalsePositiveHistoryOk returns a tuple with the FalsePositiveHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositiveHistory

`func (o *SystemSettingsRequest) SetFalsePositiveHistory(v bool)`

SetFalsePositiveHistory sets FalsePositiveHistory field to given value.

### HasFalsePositiveHistory

`func (o *SystemSettingsRequest) HasFalsePositiveHistory() bool`

HasFalsePositiveHistory returns a boolean if a field has been set.

### GetRetroactiveFalsePositiveHistory

`func (o *SystemSettingsRequest) GetRetroactiveFalsePositiveHistory() bool`

GetRetroactiveFalsePositiveHistory returns the RetroactiveFalsePositiveHistory field if non-nil, zero value otherwise.

### GetRetroactiveFalsePositiveHistoryOk

`func (o *SystemSettingsRequest) GetRetroactiveFalsePositiveHistoryOk() (*bool, bool)`

GetRetroactiveFalsePositiveHistoryOk returns a tuple with the RetroactiveFalsePositiveHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetroactiveFalsePositiveHistory

`func (o *SystemSettingsRequest) SetRetroactiveFalsePositiveHistory(v bool)`

SetRetroactiveFalsePositiveHistory sets RetroactiveFalsePositiveHistory field to given value.

### HasRetroactiveFalsePositiveHistory

`func (o *SystemSettingsRequest) HasRetroactiveFalsePositiveHistory() bool`

HasRetroactiveFalsePositiveHistory returns a boolean if a field has been set.

### GetUrlPrefix

`func (o *SystemSettingsRequest) GetUrlPrefix() string`

GetUrlPrefix returns the UrlPrefix field if non-nil, zero value otherwise.

### GetUrlPrefixOk

`func (o *SystemSettingsRequest) GetUrlPrefixOk() (*string, bool)`

GetUrlPrefixOk returns a tuple with the UrlPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPrefix

`func (o *SystemSettingsRequest) SetUrlPrefix(v string)`

SetUrlPrefix sets UrlPrefix field to given value.

### HasUrlPrefix

`func (o *SystemSettingsRequest) HasUrlPrefix() bool`

HasUrlPrefix returns a boolean if a field has been set.

### GetTeamName

`func (o *SystemSettingsRequest) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *SystemSettingsRequest) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *SystemSettingsRequest) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *SystemSettingsRequest) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetTimeZone

`func (o *SystemSettingsRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SystemSettingsRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SystemSettingsRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SystemSettingsRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetEnableProductGrade

`func (o *SystemSettingsRequest) GetEnableProductGrade() bool`

GetEnableProductGrade returns the EnableProductGrade field if non-nil, zero value otherwise.

### GetEnableProductGradeOk

`func (o *SystemSettingsRequest) GetEnableProductGradeOk() (*bool, bool)`

GetEnableProductGradeOk returns a tuple with the EnableProductGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductGrade

`func (o *SystemSettingsRequest) SetEnableProductGrade(v bool)`

SetEnableProductGrade sets EnableProductGrade field to given value.

### HasEnableProductGrade

`func (o *SystemSettingsRequest) HasEnableProductGrade() bool`

HasEnableProductGrade returns a boolean if a field has been set.

### GetProductGrade

`func (o *SystemSettingsRequest) GetProductGrade() string`

GetProductGrade returns the ProductGrade field if non-nil, zero value otherwise.

### GetProductGradeOk

`func (o *SystemSettingsRequest) GetProductGradeOk() (*string, bool)`

GetProductGradeOk returns a tuple with the ProductGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGrade

`func (o *SystemSettingsRequest) SetProductGrade(v string)`

SetProductGrade sets ProductGrade field to given value.

### HasProductGrade

`func (o *SystemSettingsRequest) HasProductGrade() bool`

HasProductGrade returns a boolean if a field has been set.

### GetProductGradeA

`func (o *SystemSettingsRequest) GetProductGradeA() int32`

GetProductGradeA returns the ProductGradeA field if non-nil, zero value otherwise.

### GetProductGradeAOk

`func (o *SystemSettingsRequest) GetProductGradeAOk() (*int32, bool)`

GetProductGradeAOk returns a tuple with the ProductGradeA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGradeA

`func (o *SystemSettingsRequest) SetProductGradeA(v int32)`

SetProductGradeA sets ProductGradeA field to given value.

### HasProductGradeA

`func (o *SystemSettingsRequest) HasProductGradeA() bool`

HasProductGradeA returns a boolean if a field has been set.

### GetProductGradeB

`func (o *SystemSettingsRequest) GetProductGradeB() int32`

GetProductGradeB returns the ProductGradeB field if non-nil, zero value otherwise.

### GetProductGradeBOk

`func (o *SystemSettingsRequest) GetProductGradeBOk() (*int32, bool)`

GetProductGradeBOk returns a tuple with the ProductGradeB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGradeB

`func (o *SystemSettingsRequest) SetProductGradeB(v int32)`

SetProductGradeB sets ProductGradeB field to given value.

### HasProductGradeB

`func (o *SystemSettingsRequest) HasProductGradeB() bool`

HasProductGradeB returns a boolean if a field has been set.

### GetProductGradeC

`func (o *SystemSettingsRequest) GetProductGradeC() int32`

GetProductGradeC returns the ProductGradeC field if non-nil, zero value otherwise.

### GetProductGradeCOk

`func (o *SystemSettingsRequest) GetProductGradeCOk() (*int32, bool)`

GetProductGradeCOk returns a tuple with the ProductGradeC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGradeC

`func (o *SystemSettingsRequest) SetProductGradeC(v int32)`

SetProductGradeC sets ProductGradeC field to given value.

### HasProductGradeC

`func (o *SystemSettingsRequest) HasProductGradeC() bool`

HasProductGradeC returns a boolean if a field has been set.

### GetProductGradeD

`func (o *SystemSettingsRequest) GetProductGradeD() int32`

GetProductGradeD returns the ProductGradeD field if non-nil, zero value otherwise.

### GetProductGradeDOk

`func (o *SystemSettingsRequest) GetProductGradeDOk() (*int32, bool)`

GetProductGradeDOk returns a tuple with the ProductGradeD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGradeD

`func (o *SystemSettingsRequest) SetProductGradeD(v int32)`

SetProductGradeD sets ProductGradeD field to given value.

### HasProductGradeD

`func (o *SystemSettingsRequest) HasProductGradeD() bool`

HasProductGradeD returns a boolean if a field has been set.

### GetProductGradeF

`func (o *SystemSettingsRequest) GetProductGradeF() int32`

GetProductGradeF returns the ProductGradeF field if non-nil, zero value otherwise.

### GetProductGradeFOk

`func (o *SystemSettingsRequest) GetProductGradeFOk() (*int32, bool)`

GetProductGradeFOk returns a tuple with the ProductGradeF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGradeF

`func (o *SystemSettingsRequest) SetProductGradeF(v int32)`

SetProductGradeF sets ProductGradeF field to given value.

### HasProductGradeF

`func (o *SystemSettingsRequest) HasProductGradeF() bool`

HasProductGradeF returns a boolean if a field has been set.

### GetEnableProductTagInheritance

`func (o *SystemSettingsRequest) GetEnableProductTagInheritance() bool`

GetEnableProductTagInheritance returns the EnableProductTagInheritance field if non-nil, zero value otherwise.

### GetEnableProductTagInheritanceOk

`func (o *SystemSettingsRequest) GetEnableProductTagInheritanceOk() (*bool, bool)`

GetEnableProductTagInheritanceOk returns a tuple with the EnableProductTagInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductTagInheritance

`func (o *SystemSettingsRequest) SetEnableProductTagInheritance(v bool)`

SetEnableProductTagInheritance sets EnableProductTagInheritance field to given value.

### HasEnableProductTagInheritance

`func (o *SystemSettingsRequest) HasEnableProductTagInheritance() bool`

HasEnableProductTagInheritance returns a boolean if a field has been set.

### GetEnableBenchmark

`func (o *SystemSettingsRequest) GetEnableBenchmark() bool`

GetEnableBenchmark returns the EnableBenchmark field if non-nil, zero value otherwise.

### GetEnableBenchmarkOk

`func (o *SystemSettingsRequest) GetEnableBenchmarkOk() (*bool, bool)`

GetEnableBenchmarkOk returns a tuple with the EnableBenchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBenchmark

`func (o *SystemSettingsRequest) SetEnableBenchmark(v bool)`

SetEnableBenchmark sets EnableBenchmark field to given value.

### HasEnableBenchmark

`func (o *SystemSettingsRequest) HasEnableBenchmark() bool`

HasEnableBenchmark returns a boolean if a field has been set.

### GetEnableTemplateMatch

`func (o *SystemSettingsRequest) GetEnableTemplateMatch() bool`

GetEnableTemplateMatch returns the EnableTemplateMatch field if non-nil, zero value otherwise.

### GetEnableTemplateMatchOk

`func (o *SystemSettingsRequest) GetEnableTemplateMatchOk() (*bool, bool)`

GetEnableTemplateMatchOk returns a tuple with the EnableTemplateMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTemplateMatch

`func (o *SystemSettingsRequest) SetEnableTemplateMatch(v bool)`

SetEnableTemplateMatch sets EnableTemplateMatch field to given value.

### HasEnableTemplateMatch

`func (o *SystemSettingsRequest) HasEnableTemplateMatch() bool`

HasEnableTemplateMatch returns a boolean if a field has been set.

### GetEngagementAutoClose

`func (o *SystemSettingsRequest) GetEngagementAutoClose() bool`

GetEngagementAutoClose returns the EngagementAutoClose field if non-nil, zero value otherwise.

### GetEngagementAutoCloseOk

`func (o *SystemSettingsRequest) GetEngagementAutoCloseOk() (*bool, bool)`

GetEngagementAutoCloseOk returns a tuple with the EngagementAutoClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementAutoClose

`func (o *SystemSettingsRequest) SetEngagementAutoClose(v bool)`

SetEngagementAutoClose sets EngagementAutoClose field to given value.

### HasEngagementAutoClose

`func (o *SystemSettingsRequest) HasEngagementAutoClose() bool`

HasEngagementAutoClose returns a boolean if a field has been set.

### GetEngagementAutoCloseDays

`func (o *SystemSettingsRequest) GetEngagementAutoCloseDays() int32`

GetEngagementAutoCloseDays returns the EngagementAutoCloseDays field if non-nil, zero value otherwise.

### GetEngagementAutoCloseDaysOk

`func (o *SystemSettingsRequest) GetEngagementAutoCloseDaysOk() (*int32, bool)`

GetEngagementAutoCloseDaysOk returns a tuple with the EngagementAutoCloseDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementAutoCloseDays

`func (o *SystemSettingsRequest) SetEngagementAutoCloseDays(v int32)`

SetEngagementAutoCloseDays sets EngagementAutoCloseDays field to given value.

### HasEngagementAutoCloseDays

`func (o *SystemSettingsRequest) HasEngagementAutoCloseDays() bool`

HasEngagementAutoCloseDays returns a boolean if a field has been set.

### GetEnableFindingSla

`func (o *SystemSettingsRequest) GetEnableFindingSla() bool`

GetEnableFindingSla returns the EnableFindingSla field if non-nil, zero value otherwise.

### GetEnableFindingSlaOk

`func (o *SystemSettingsRequest) GetEnableFindingSlaOk() (*bool, bool)`

GetEnableFindingSlaOk returns a tuple with the EnableFindingSla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFindingSla

`func (o *SystemSettingsRequest) SetEnableFindingSla(v bool)`

SetEnableFindingSla sets EnableFindingSla field to given value.

### HasEnableFindingSla

`func (o *SystemSettingsRequest) HasEnableFindingSla() bool`

HasEnableFindingSla returns a boolean if a field has been set.

### GetEnableNotifySlaActive

`func (o *SystemSettingsRequest) GetEnableNotifySlaActive() bool`

GetEnableNotifySlaActive returns the EnableNotifySlaActive field if non-nil, zero value otherwise.

### GetEnableNotifySlaActiveOk

`func (o *SystemSettingsRequest) GetEnableNotifySlaActiveOk() (*bool, bool)`

GetEnableNotifySlaActiveOk returns a tuple with the EnableNotifySlaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySlaActive

`func (o *SystemSettingsRequest) SetEnableNotifySlaActive(v bool)`

SetEnableNotifySlaActive sets EnableNotifySlaActive field to given value.

### HasEnableNotifySlaActive

`func (o *SystemSettingsRequest) HasEnableNotifySlaActive() bool`

HasEnableNotifySlaActive returns a boolean if a field has been set.

### GetEnableNotifySlaActiveVerified

`func (o *SystemSettingsRequest) GetEnableNotifySlaActiveVerified() bool`

GetEnableNotifySlaActiveVerified returns the EnableNotifySlaActiveVerified field if non-nil, zero value otherwise.

### GetEnableNotifySlaActiveVerifiedOk

`func (o *SystemSettingsRequest) GetEnableNotifySlaActiveVerifiedOk() (*bool, bool)`

GetEnableNotifySlaActiveVerifiedOk returns a tuple with the EnableNotifySlaActiveVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySlaActiveVerified

`func (o *SystemSettingsRequest) SetEnableNotifySlaActiveVerified(v bool)`

SetEnableNotifySlaActiveVerified sets EnableNotifySlaActiveVerified field to given value.

### HasEnableNotifySlaActiveVerified

`func (o *SystemSettingsRequest) HasEnableNotifySlaActiveVerified() bool`

HasEnableNotifySlaActiveVerified returns a boolean if a field has been set.

### GetEnableNotifySlaJiraOnly

`func (o *SystemSettingsRequest) GetEnableNotifySlaJiraOnly() bool`

GetEnableNotifySlaJiraOnly returns the EnableNotifySlaJiraOnly field if non-nil, zero value otherwise.

### GetEnableNotifySlaJiraOnlyOk

`func (o *SystemSettingsRequest) GetEnableNotifySlaJiraOnlyOk() (*bool, bool)`

GetEnableNotifySlaJiraOnlyOk returns a tuple with the EnableNotifySlaJiraOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySlaJiraOnly

`func (o *SystemSettingsRequest) SetEnableNotifySlaJiraOnly(v bool)`

SetEnableNotifySlaJiraOnly sets EnableNotifySlaJiraOnly field to given value.

### HasEnableNotifySlaJiraOnly

`func (o *SystemSettingsRequest) HasEnableNotifySlaJiraOnly() bool`

HasEnableNotifySlaJiraOnly returns a boolean if a field has been set.

### GetEnableNotifySlaExponentialBackoff

`func (o *SystemSettingsRequest) GetEnableNotifySlaExponentialBackoff() bool`

GetEnableNotifySlaExponentialBackoff returns the EnableNotifySlaExponentialBackoff field if non-nil, zero value otherwise.

### GetEnableNotifySlaExponentialBackoffOk

`func (o *SystemSettingsRequest) GetEnableNotifySlaExponentialBackoffOk() (*bool, bool)`

GetEnableNotifySlaExponentialBackoffOk returns a tuple with the EnableNotifySlaExponentialBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySlaExponentialBackoff

`func (o *SystemSettingsRequest) SetEnableNotifySlaExponentialBackoff(v bool)`

SetEnableNotifySlaExponentialBackoff sets EnableNotifySlaExponentialBackoff field to given value.

### HasEnableNotifySlaExponentialBackoff

`func (o *SystemSettingsRequest) HasEnableNotifySlaExponentialBackoff() bool`

HasEnableNotifySlaExponentialBackoff returns a boolean if a field has been set.

### GetAllowAnonymousSurveyRepsonse

`func (o *SystemSettingsRequest) GetAllowAnonymousSurveyRepsonse() bool`

GetAllowAnonymousSurveyRepsonse returns the AllowAnonymousSurveyRepsonse field if non-nil, zero value otherwise.

### GetAllowAnonymousSurveyRepsonseOk

`func (o *SystemSettingsRequest) GetAllowAnonymousSurveyRepsonseOk() (*bool, bool)`

GetAllowAnonymousSurveyRepsonseOk returns a tuple with the AllowAnonymousSurveyRepsonse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymousSurveyRepsonse

`func (o *SystemSettingsRequest) SetAllowAnonymousSurveyRepsonse(v bool)`

SetAllowAnonymousSurveyRepsonse sets AllowAnonymousSurveyRepsonse field to given value.

### HasAllowAnonymousSurveyRepsonse

`func (o *SystemSettingsRequest) HasAllowAnonymousSurveyRepsonse() bool`

HasAllowAnonymousSurveyRepsonse returns a boolean if a field has been set.

### GetCredentials

`func (o *SystemSettingsRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SystemSettingsRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SystemSettingsRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SystemSettingsRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDisclaimer

`func (o *SystemSettingsRequest) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *SystemSettingsRequest) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *SystemSettingsRequest) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *SystemSettingsRequest) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetRiskAcceptanceFormDefaultDays

`func (o *SystemSettingsRequest) GetRiskAcceptanceFormDefaultDays() int32`

GetRiskAcceptanceFormDefaultDays returns the RiskAcceptanceFormDefaultDays field if non-nil, zero value otherwise.

### GetRiskAcceptanceFormDefaultDaysOk

`func (o *SystemSettingsRequest) GetRiskAcceptanceFormDefaultDaysOk() (*int32, bool)`

GetRiskAcceptanceFormDefaultDaysOk returns a tuple with the RiskAcceptanceFormDefaultDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceFormDefaultDays

`func (o *SystemSettingsRequest) SetRiskAcceptanceFormDefaultDays(v int32)`

SetRiskAcceptanceFormDefaultDays sets RiskAcceptanceFormDefaultDays field to given value.

### HasRiskAcceptanceFormDefaultDays

`func (o *SystemSettingsRequest) HasRiskAcceptanceFormDefaultDays() bool`

HasRiskAcceptanceFormDefaultDays returns a boolean if a field has been set.

### SetRiskAcceptanceFormDefaultDaysNil

`func (o *SystemSettingsRequest) SetRiskAcceptanceFormDefaultDaysNil(b bool)`

 SetRiskAcceptanceFormDefaultDaysNil sets the value for RiskAcceptanceFormDefaultDays to be an explicit nil

### UnsetRiskAcceptanceFormDefaultDays
`func (o *SystemSettingsRequest) UnsetRiskAcceptanceFormDefaultDays()`

UnsetRiskAcceptanceFormDefaultDays ensures that no value is present for RiskAcceptanceFormDefaultDays, not even an explicit nil
### GetRiskAcceptanceNotifyBeforeExpiration

`func (o *SystemSettingsRequest) GetRiskAcceptanceNotifyBeforeExpiration() int32`

GetRiskAcceptanceNotifyBeforeExpiration returns the RiskAcceptanceNotifyBeforeExpiration field if non-nil, zero value otherwise.

### GetRiskAcceptanceNotifyBeforeExpirationOk

`func (o *SystemSettingsRequest) GetRiskAcceptanceNotifyBeforeExpirationOk() (*int32, bool)`

GetRiskAcceptanceNotifyBeforeExpirationOk returns a tuple with the RiskAcceptanceNotifyBeforeExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceNotifyBeforeExpiration

`func (o *SystemSettingsRequest) SetRiskAcceptanceNotifyBeforeExpiration(v int32)`

SetRiskAcceptanceNotifyBeforeExpiration sets RiskAcceptanceNotifyBeforeExpiration field to given value.

### HasRiskAcceptanceNotifyBeforeExpiration

`func (o *SystemSettingsRequest) HasRiskAcceptanceNotifyBeforeExpiration() bool`

HasRiskAcceptanceNotifyBeforeExpiration returns a boolean if a field has been set.

### SetRiskAcceptanceNotifyBeforeExpirationNil

`func (o *SystemSettingsRequest) SetRiskAcceptanceNotifyBeforeExpirationNil(b bool)`

 SetRiskAcceptanceNotifyBeforeExpirationNil sets the value for RiskAcceptanceNotifyBeforeExpiration to be an explicit nil

### UnsetRiskAcceptanceNotifyBeforeExpiration
`func (o *SystemSettingsRequest) UnsetRiskAcceptanceNotifyBeforeExpiration()`

UnsetRiskAcceptanceNotifyBeforeExpiration ensures that no value is present for RiskAcceptanceNotifyBeforeExpiration, not even an explicit nil
### GetEnableCredentials

`func (o *SystemSettingsRequest) GetEnableCredentials() bool`

GetEnableCredentials returns the EnableCredentials field if non-nil, zero value otherwise.

### GetEnableCredentialsOk

`func (o *SystemSettingsRequest) GetEnableCredentialsOk() (*bool, bool)`

GetEnableCredentialsOk returns a tuple with the EnableCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCredentials

`func (o *SystemSettingsRequest) SetEnableCredentials(v bool)`

SetEnableCredentials sets EnableCredentials field to given value.

### HasEnableCredentials

`func (o *SystemSettingsRequest) HasEnableCredentials() bool`

HasEnableCredentials returns a boolean if a field has been set.

### GetEnableQuestionnaires

`func (o *SystemSettingsRequest) GetEnableQuestionnaires() bool`

GetEnableQuestionnaires returns the EnableQuestionnaires field if non-nil, zero value otherwise.

### GetEnableQuestionnairesOk

`func (o *SystemSettingsRequest) GetEnableQuestionnairesOk() (*bool, bool)`

GetEnableQuestionnairesOk returns a tuple with the EnableQuestionnaires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuestionnaires

`func (o *SystemSettingsRequest) SetEnableQuestionnaires(v bool)`

SetEnableQuestionnaires sets EnableQuestionnaires field to given value.

### HasEnableQuestionnaires

`func (o *SystemSettingsRequest) HasEnableQuestionnaires() bool`

HasEnableQuestionnaires returns a boolean if a field has been set.

### GetEnableChecklists

`func (o *SystemSettingsRequest) GetEnableChecklists() bool`

GetEnableChecklists returns the EnableChecklists field if non-nil, zero value otherwise.

### GetEnableChecklistsOk

`func (o *SystemSettingsRequest) GetEnableChecklistsOk() (*bool, bool)`

GetEnableChecklistsOk returns a tuple with the EnableChecklists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChecklists

`func (o *SystemSettingsRequest) SetEnableChecklists(v bool)`

SetEnableChecklists sets EnableChecklists field to given value.

### HasEnableChecklists

`func (o *SystemSettingsRequest) HasEnableChecklists() bool`

HasEnableChecklists returns a boolean if a field has been set.

### GetEnableEndpointMetadataImport

`func (o *SystemSettingsRequest) GetEnableEndpointMetadataImport() bool`

GetEnableEndpointMetadataImport returns the EnableEndpointMetadataImport field if non-nil, zero value otherwise.

### GetEnableEndpointMetadataImportOk

`func (o *SystemSettingsRequest) GetEnableEndpointMetadataImportOk() (*bool, bool)`

GetEnableEndpointMetadataImportOk returns a tuple with the EnableEndpointMetadataImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEndpointMetadataImport

`func (o *SystemSettingsRequest) SetEnableEndpointMetadataImport(v bool)`

SetEnableEndpointMetadataImport sets EnableEndpointMetadataImport field to given value.

### HasEnableEndpointMetadataImport

`func (o *SystemSettingsRequest) HasEnableEndpointMetadataImport() bool`

HasEnableEndpointMetadataImport returns a boolean if a field has been set.

### GetEnableUserProfileEditable

`func (o *SystemSettingsRequest) GetEnableUserProfileEditable() bool`

GetEnableUserProfileEditable returns the EnableUserProfileEditable field if non-nil, zero value otherwise.

### GetEnableUserProfileEditableOk

`func (o *SystemSettingsRequest) GetEnableUserProfileEditableOk() (*bool, bool)`

GetEnableUserProfileEditableOk returns a tuple with the EnableUserProfileEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserProfileEditable

`func (o *SystemSettingsRequest) SetEnableUserProfileEditable(v bool)`

SetEnableUserProfileEditable sets EnableUserProfileEditable field to given value.

### HasEnableUserProfileEditable

`func (o *SystemSettingsRequest) HasEnableUserProfileEditable() bool`

HasEnableUserProfileEditable returns a boolean if a field has been set.

### GetEnableProductTrackingFiles

`func (o *SystemSettingsRequest) GetEnableProductTrackingFiles() bool`

GetEnableProductTrackingFiles returns the EnableProductTrackingFiles field if non-nil, zero value otherwise.

### GetEnableProductTrackingFilesOk

`func (o *SystemSettingsRequest) GetEnableProductTrackingFilesOk() (*bool, bool)`

GetEnableProductTrackingFilesOk returns a tuple with the EnableProductTrackingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductTrackingFiles

`func (o *SystemSettingsRequest) SetEnableProductTrackingFiles(v bool)`

SetEnableProductTrackingFiles sets EnableProductTrackingFiles field to given value.

### HasEnableProductTrackingFiles

`func (o *SystemSettingsRequest) HasEnableProductTrackingFiles() bool`

HasEnableProductTrackingFiles returns a boolean if a field has been set.

### GetEnableFindingGroups

`func (o *SystemSettingsRequest) GetEnableFindingGroups() bool`

GetEnableFindingGroups returns the EnableFindingGroups field if non-nil, zero value otherwise.

### GetEnableFindingGroupsOk

`func (o *SystemSettingsRequest) GetEnableFindingGroupsOk() (*bool, bool)`

GetEnableFindingGroupsOk returns a tuple with the EnableFindingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFindingGroups

`func (o *SystemSettingsRequest) SetEnableFindingGroups(v bool)`

SetEnableFindingGroups sets EnableFindingGroups field to given value.

### HasEnableFindingGroups

`func (o *SystemSettingsRequest) HasEnableFindingGroups() bool`

HasEnableFindingGroups returns a boolean if a field has been set.

### GetEnableCalendar

`func (o *SystemSettingsRequest) GetEnableCalendar() bool`

GetEnableCalendar returns the EnableCalendar field if non-nil, zero value otherwise.

### GetEnableCalendarOk

`func (o *SystemSettingsRequest) GetEnableCalendarOk() (*bool, bool)`

GetEnableCalendarOk returns a tuple with the EnableCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCalendar

`func (o *SystemSettingsRequest) SetEnableCalendar(v bool)`

SetEnableCalendar sets EnableCalendar field to given value.

### HasEnableCalendar

`func (o *SystemSettingsRequest) HasEnableCalendar() bool`

HasEnableCalendar returns a boolean if a field has been set.

### GetDefaultGroupEmailPattern

`func (o *SystemSettingsRequest) GetDefaultGroupEmailPattern() string`

GetDefaultGroupEmailPattern returns the DefaultGroupEmailPattern field if non-nil, zero value otherwise.

### GetDefaultGroupEmailPatternOk

`func (o *SystemSettingsRequest) GetDefaultGroupEmailPatternOk() (*string, bool)`

GetDefaultGroupEmailPatternOk returns a tuple with the DefaultGroupEmailPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupEmailPattern

`func (o *SystemSettingsRequest) SetDefaultGroupEmailPattern(v string)`

SetDefaultGroupEmailPattern sets DefaultGroupEmailPattern field to given value.

### HasDefaultGroupEmailPattern

`func (o *SystemSettingsRequest) HasDefaultGroupEmailPattern() bool`

HasDefaultGroupEmailPattern returns a boolean if a field has been set.

### GetMinimumPasswordLength

`func (o *SystemSettingsRequest) GetMinimumPasswordLength() int32`

GetMinimumPasswordLength returns the MinimumPasswordLength field if non-nil, zero value otherwise.

### GetMinimumPasswordLengthOk

`func (o *SystemSettingsRequest) GetMinimumPasswordLengthOk() (*int32, bool)`

GetMinimumPasswordLengthOk returns a tuple with the MinimumPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPasswordLength

`func (o *SystemSettingsRequest) SetMinimumPasswordLength(v int32)`

SetMinimumPasswordLength sets MinimumPasswordLength field to given value.

### HasMinimumPasswordLength

`func (o *SystemSettingsRequest) HasMinimumPasswordLength() bool`

HasMinimumPasswordLength returns a boolean if a field has been set.

### GetMaximumPasswordLength

`func (o *SystemSettingsRequest) GetMaximumPasswordLength() int32`

GetMaximumPasswordLength returns the MaximumPasswordLength field if non-nil, zero value otherwise.

### GetMaximumPasswordLengthOk

`func (o *SystemSettingsRequest) GetMaximumPasswordLengthOk() (*int32, bool)`

GetMaximumPasswordLengthOk returns a tuple with the MaximumPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPasswordLength

`func (o *SystemSettingsRequest) SetMaximumPasswordLength(v int32)`

SetMaximumPasswordLength sets MaximumPasswordLength field to given value.

### HasMaximumPasswordLength

`func (o *SystemSettingsRequest) HasMaximumPasswordLength() bool`

HasMaximumPasswordLength returns a boolean if a field has been set.

### GetNumberCharacterRequired

`func (o *SystemSettingsRequest) GetNumberCharacterRequired() bool`

GetNumberCharacterRequired returns the NumberCharacterRequired field if non-nil, zero value otherwise.

### GetNumberCharacterRequiredOk

`func (o *SystemSettingsRequest) GetNumberCharacterRequiredOk() (*bool, bool)`

GetNumberCharacterRequiredOk returns a tuple with the NumberCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberCharacterRequired

`func (o *SystemSettingsRequest) SetNumberCharacterRequired(v bool)`

SetNumberCharacterRequired sets NumberCharacterRequired field to given value.

### HasNumberCharacterRequired

`func (o *SystemSettingsRequest) HasNumberCharacterRequired() bool`

HasNumberCharacterRequired returns a boolean if a field has been set.

### GetSpecialCharacterRequired

`func (o *SystemSettingsRequest) GetSpecialCharacterRequired() bool`

GetSpecialCharacterRequired returns the SpecialCharacterRequired field if non-nil, zero value otherwise.

### GetSpecialCharacterRequiredOk

`func (o *SystemSettingsRequest) GetSpecialCharacterRequiredOk() (*bool, bool)`

GetSpecialCharacterRequiredOk returns a tuple with the SpecialCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacterRequired

`func (o *SystemSettingsRequest) SetSpecialCharacterRequired(v bool)`

SetSpecialCharacterRequired sets SpecialCharacterRequired field to given value.

### HasSpecialCharacterRequired

`func (o *SystemSettingsRequest) HasSpecialCharacterRequired() bool`

HasSpecialCharacterRequired returns a boolean if a field has been set.

### GetLowercaseCharacterRequired

`func (o *SystemSettingsRequest) GetLowercaseCharacterRequired() bool`

GetLowercaseCharacterRequired returns the LowercaseCharacterRequired field if non-nil, zero value otherwise.

### GetLowercaseCharacterRequiredOk

`func (o *SystemSettingsRequest) GetLowercaseCharacterRequiredOk() (*bool, bool)`

GetLowercaseCharacterRequiredOk returns a tuple with the LowercaseCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowercaseCharacterRequired

`func (o *SystemSettingsRequest) SetLowercaseCharacterRequired(v bool)`

SetLowercaseCharacterRequired sets LowercaseCharacterRequired field to given value.

### HasLowercaseCharacterRequired

`func (o *SystemSettingsRequest) HasLowercaseCharacterRequired() bool`

HasLowercaseCharacterRequired returns a boolean if a field has been set.

### GetUppercaseCharacterRequired

`func (o *SystemSettingsRequest) GetUppercaseCharacterRequired() bool`

GetUppercaseCharacterRequired returns the UppercaseCharacterRequired field if non-nil, zero value otherwise.

### GetUppercaseCharacterRequiredOk

`func (o *SystemSettingsRequest) GetUppercaseCharacterRequiredOk() (*bool, bool)`

GetUppercaseCharacterRequiredOk returns a tuple with the UppercaseCharacterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercaseCharacterRequired

`func (o *SystemSettingsRequest) SetUppercaseCharacterRequired(v bool)`

SetUppercaseCharacterRequired sets UppercaseCharacterRequired field to given value.

### HasUppercaseCharacterRequired

`func (o *SystemSettingsRequest) HasUppercaseCharacterRequired() bool`

HasUppercaseCharacterRequired returns a boolean if a field has been set.

### GetNonCommonPasswordRequired

`func (o *SystemSettingsRequest) GetNonCommonPasswordRequired() bool`

GetNonCommonPasswordRequired returns the NonCommonPasswordRequired field if non-nil, zero value otherwise.

### GetNonCommonPasswordRequiredOk

`func (o *SystemSettingsRequest) GetNonCommonPasswordRequiredOk() (*bool, bool)`

GetNonCommonPasswordRequiredOk returns a tuple with the NonCommonPasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCommonPasswordRequired

`func (o *SystemSettingsRequest) SetNonCommonPasswordRequired(v bool)`

SetNonCommonPasswordRequired sets NonCommonPasswordRequired field to given value.

### HasNonCommonPasswordRequired

`func (o *SystemSettingsRequest) HasNonCommonPasswordRequired() bool`

HasNonCommonPasswordRequired returns a boolean if a field has been set.

### GetApiExposeErrorDetails

`func (o *SystemSettingsRequest) GetApiExposeErrorDetails() bool`

GetApiExposeErrorDetails returns the ApiExposeErrorDetails field if non-nil, zero value otherwise.

### GetApiExposeErrorDetailsOk

`func (o *SystemSettingsRequest) GetApiExposeErrorDetailsOk() (*bool, bool)`

GetApiExposeErrorDetailsOk returns a tuple with the ApiExposeErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiExposeErrorDetails

`func (o *SystemSettingsRequest) SetApiExposeErrorDetails(v bool)`

SetApiExposeErrorDetails sets ApiExposeErrorDetails field to given value.

### HasApiExposeErrorDetails

`func (o *SystemSettingsRequest) HasApiExposeErrorDetails() bool`

HasApiExposeErrorDetails returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *SystemSettingsRequest) GetDefaultGroup() int32`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *SystemSettingsRequest) GetDefaultGroupOk() (*int32, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *SystemSettingsRequest) SetDefaultGroup(v int32)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *SystemSettingsRequest) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### SetDefaultGroupNil

`func (o *SystemSettingsRequest) SetDefaultGroupNil(b bool)`

 SetDefaultGroupNil sets the value for DefaultGroup to be an explicit nil

### UnsetDefaultGroup
`func (o *SystemSettingsRequest) UnsetDefaultGroup()`

UnsetDefaultGroup ensures that no value is present for DefaultGroup, not even an explicit nil
### GetDefaultGroupRole

`func (o *SystemSettingsRequest) GetDefaultGroupRole() int32`

GetDefaultGroupRole returns the DefaultGroupRole field if non-nil, zero value otherwise.

### GetDefaultGroupRoleOk

`func (o *SystemSettingsRequest) GetDefaultGroupRoleOk() (*int32, bool)`

GetDefaultGroupRoleOk returns a tuple with the DefaultGroupRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupRole

`func (o *SystemSettingsRequest) SetDefaultGroupRole(v int32)`

SetDefaultGroupRole sets DefaultGroupRole field to given value.

### HasDefaultGroupRole

`func (o *SystemSettingsRequest) HasDefaultGroupRole() bool`

HasDefaultGroupRole returns a boolean if a field has been set.

### SetDefaultGroupRoleNil

`func (o *SystemSettingsRequest) SetDefaultGroupRoleNil(b bool)`

 SetDefaultGroupRoleNil sets the value for DefaultGroupRole to be an explicit nil

### UnsetDefaultGroupRole
`func (o *SystemSettingsRequest) UnsetDefaultGroupRole()`

UnsetDefaultGroupRole ensures that no value is present for DefaultGroupRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


