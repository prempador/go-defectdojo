# ReImportScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanDate** | Pointer to **string** | Scan completion date will be used on all findings. | [optional] 
**MinimumSeverity** | Pointer to **string** | Minimum severity level to be imported  * &#x60;Info&#x60; - Info * &#x60;Low&#x60; - Low * &#x60;Medium&#x60; - Medium * &#x60;High&#x60; - High * &#x60;Critical&#x60; - Critical | [optional] [default to "Info"]
**Active** | **bool** | Override the active setting from the tool. | 
**Verified** | **bool** | Override the verified setting from the tool. | 
**DoNotReactivate** | Pointer to **bool** | Select if the import should ignore active findings from the report, useful for triage-less scanners. Will keep existing findings closed, without reactivating them. For more information check the docs. | [optional] [default to false]
**ScanType** | **string** | * &#x60;Acunetix Scan&#x60; - Acunetix Scan * &#x60;Acunetix360 Scan&#x60; - Acunetix360 Scan * &#x60;Anchore Engine Scan&#x60; - Anchore Engine Scan * &#x60;Anchore Enterprise Policy Check&#x60; - Anchore Enterprise Policy Check * &#x60;Anchore Grype&#x60; - Anchore Grype * &#x60;AnchoreCTL Policies Report&#x60; - AnchoreCTL Policies Report * &#x60;AnchoreCTL Vuln Report&#x60; - AnchoreCTL Vuln Report * &#x60;AppSpider Scan&#x60; - AppSpider Scan * &#x60;Aqua Scan&#x60; - Aqua Scan * &#x60;Arachni Scan&#x60; - Arachni Scan * &#x60;AuditJS Scan&#x60; - AuditJS Scan * &#x60;AWS Prowler Scan&#x60; - AWS Prowler Scan * &#x60;AWS Prowler V3&#x60; - AWS Prowler V3 * &#x60;AWS Scout2 Scan&#x60; - AWS Scout2 Scan * &#x60;AWS Security Finding Format (ASFF) Scan&#x60; - AWS Security Finding Format (ASFF) Scan * &#x60;AWS Security Hub Scan&#x60; - AWS Security Hub Scan * &#x60;Azure Security Center Recommendations Scan&#x60; - Azure Security Center Recommendations Scan * &#x60;Bandit Scan&#x60; - Bandit Scan * &#x60;BlackDuck API&#x60; - BlackDuck API * &#x60;Blackduck Binary Analysis&#x60; - Blackduck Binary Analysis * &#x60;Blackduck Component Risk&#x60; - Blackduck Component Risk * &#x60;Blackduck Hub Scan&#x60; - Blackduck Hub Scan * &#x60;Brakeman Scan&#x60; - Brakeman Scan * &#x60;Bugcrowd API Import&#x60; - Bugcrowd API Import * &#x60;BugCrowd Scan&#x60; - BugCrowd Scan * &#x60;Bundler-Audit Scan&#x60; - Bundler-Audit Scan * &#x60;Burp Enterprise Scan&#x60; - Burp Enterprise Scan * &#x60;Burp GraphQL API&#x60; - Burp GraphQL API * &#x60;Burp REST API&#x60; - Burp REST API * &#x60;Burp Scan&#x60; - Burp Scan * &#x60;CargoAudit Scan&#x60; - CargoAudit Scan * &#x60;Checkmarx OSA&#x60; - Checkmarx OSA * &#x60;Checkmarx Scan&#x60; - Checkmarx Scan * &#x60;Checkmarx Scan detailed&#x60; - Checkmarx Scan detailed * &#x60;Checkov Scan&#x60; - Checkov Scan * &#x60;Chef Inspect Log&#x60; - Chef Inspect Log * &#x60;Clair Scan&#x60; - Clair Scan * &#x60;Cloudsploit Scan&#x60; - Cloudsploit Scan * &#x60;Cobalt.io API Import&#x60; - Cobalt.io API Import * &#x60;Cobalt.io Scan&#x60; - Cobalt.io Scan * &#x60;Codechecker Report native&#x60; - Codechecker Report native * &#x60;Contrast Scan&#x60; - Contrast Scan * &#x60;Coverity API&#x60; - Coverity API * &#x60;Crashtest Security JSON File&#x60; - Crashtest Security JSON File * &#x60;Crashtest Security XML File&#x60; - Crashtest Security XML File * &#x60;CredScan Scan&#x60; - CredScan Scan * &#x60;CycloneDX Scan&#x60; - CycloneDX Scan * &#x60;DawnScanner Scan&#x60; - DawnScanner Scan * &#x60;Dependency Check Scan&#x60; - Dependency Check Scan * &#x60;Dependency Track Finding Packaging Format (FPF) Export&#x60; - Dependency Track Finding Packaging Format (FPF) Export * &#x60;Detect-secrets Scan&#x60; - Detect-secrets Scan * &#x60;docker-bench-security Scan&#x60; - docker-bench-security Scan * &#x60;Dockle Scan&#x60; - Dockle Scan * &#x60;DrHeader JSON Importer&#x60; - DrHeader JSON Importer * &#x60;DSOP Scan&#x60; - DSOP Scan * &#x60;Edgescan Scan&#x60; - Edgescan Scan * &#x60;ESLint Scan&#x60; - ESLint Scan * &#x60;Fortify Scan&#x60; - Fortify Scan * &#x60;Generic Findings Import&#x60; - Generic Findings Import * &#x60;Ggshield Scan&#x60; - Ggshield Scan * &#x60;Github Vulnerability Scan&#x60; - Github Vulnerability Scan * &#x60;GitLab API Fuzzing Report Scan&#x60; - GitLab API Fuzzing Report Scan * &#x60;GitLab Container Scan&#x60; - GitLab Container Scan * &#x60;GitLab DAST Report&#x60; - GitLab DAST Report * &#x60;GitLab Dependency Scanning Report&#x60; - GitLab Dependency Scanning Report * &#x60;GitLab SAST Report&#x60; - GitLab SAST Report * &#x60;GitLab Secret Detection Report&#x60; - GitLab Secret Detection Report * &#x60;Gitleaks Scan&#x60; - Gitleaks Scan * &#x60;Google Cloud Artifact Vulnerability Scan&#x60; - Google Cloud Artifact Vulnerability Scan * &#x60;Gosec Scanner&#x60; - Gosec Scanner * &#x60;Govulncheck Scanner&#x60; - Govulncheck Scanner * &#x60;HackerOne Cases&#x60; - HackerOne Cases * &#x60;Hadolint Dockerfile check&#x60; - Hadolint Dockerfile check * &#x60;Harbor Vulnerability Scan&#x60; - Harbor Vulnerability Scan * &#x60;HCLAppScan XML&#x60; - HCLAppScan XML * &#x60;Horusec Scan&#x60; - Horusec Scan * &#x60;Humble Json Importer&#x60; - Humble Json Importer * &#x60;HuskyCI Report&#x60; - HuskyCI Report * &#x60;Hydra Scan&#x60; - Hydra Scan * &#x60;IBM AppScan DAST&#x60; - IBM AppScan DAST * &#x60;Immuniweb Scan&#x60; - Immuniweb Scan * &#x60;IntSights Report&#x60; - IntSights Report * &#x60;JFrog Xray API Summary Artifact Scan&#x60; - JFrog Xray API Summary Artifact Scan * &#x60;JFrog Xray On Demand Binary Scan&#x60; - JFrog Xray On Demand Binary Scan * &#x60;JFrog Xray Scan&#x60; - JFrog Xray Scan * &#x60;JFrog Xray Unified Scan&#x60; - JFrog Xray Unified Scan * &#x60;KICS Scan&#x60; - KICS Scan * &#x60;Kiuwan Scan&#x60; - Kiuwan Scan * &#x60;kube-bench Scan&#x60; - kube-bench Scan * &#x60;Kubeaudit Scan&#x60; - Kubeaudit Scan * &#x60;KubeHunter Scan&#x60; - KubeHunter Scan * &#x60;Kubescape JSON Importer&#x60; - Kubescape JSON Importer * &#x60;Mend Scan&#x60; - Mend Scan * &#x60;Meterian Scan&#x60; - Meterian Scan * &#x60;Microfocus Webinspect Scan&#x60; - Microfocus Webinspect Scan * &#x60;MobSF Scan&#x60; - MobSF Scan * &#x60;Mobsfscan Scan&#x60; - Mobsfscan Scan * &#x60;Mozilla Observatory Scan&#x60; - Mozilla Observatory Scan * &#x60;MSDefender Parser&#x60; - MSDefender Parser * &#x60;Netsparker Scan&#x60; - Netsparker Scan * &#x60;NeuVector (compliance)&#x60; - NeuVector (compliance) * &#x60;NeuVector (REST)&#x60; - NeuVector (REST) * &#x60;Nexpose Scan&#x60; - Nexpose Scan * &#x60;Nikto Scan&#x60; - Nikto Scan * &#x60;Nmap Scan&#x60; - Nmap Scan * &#x60;Node Security Platform Scan&#x60; - Node Security Platform Scan * &#x60;NPM Audit Scan&#x60; - NPM Audit Scan * &#x60;Nuclei Scan&#x60; - Nuclei Scan * &#x60;Openscap Vulnerability Scan&#x60; - Openscap Vulnerability Scan * &#x60;OpenVAS Parser&#x60; - OpenVAS Parser * &#x60;ORT evaluated model Importer&#x60; - ORT evaluated model Importer * &#x60;OssIndex Devaudit SCA Scan Importer&#x60; - OssIndex Devaudit SCA Scan Importer * &#x60;Outpost24 Scan&#x60; - Outpost24 Scan * &#x60;PHP Security Audit v2&#x60; - PHP Security Audit v2 * &#x60;PHP Symfony Security Check&#x60; - PHP Symfony Security Check * &#x60;pip-audit Scan&#x60; - pip-audit Scan * &#x60;PMD Scan&#x60; - PMD Scan * &#x60;Popeye Scan&#x60; - Popeye Scan * &#x60;PWN SAST&#x60; - PWN SAST * &#x60;Qualys Infrastructure Scan (WebGUI XML)&#x60; - Qualys Infrastructure Scan (WebGUI XML) * &#x60;Qualys Scan&#x60; - Qualys Scan * &#x60;Qualys Webapp Scan&#x60; - Qualys Webapp Scan * &#x60;Red Hat Satellite&#x60; - Red Hat Satellite * &#x60;Retire.js Scan&#x60; - Retire.js Scan * &#x60;Risk Recon API Importer&#x60; - Risk Recon API Importer * &#x60;Rubocop Scan&#x60; - Rubocop Scan * &#x60;Rusty Hog Scan&#x60; - Rusty Hog Scan * &#x60;SARIF&#x60; - SARIF * &#x60;Scantist Scan&#x60; - Scantist Scan * &#x60;Scout Suite Scan&#x60; - Scout Suite Scan * &#x60;Semgrep JSON Report&#x60; - Semgrep JSON Report * &#x60;SKF Scan&#x60; - SKF Scan * &#x60;Snyk Scan&#x60; - Snyk Scan * &#x60;Solar Appscreener Scan&#x60; - Solar Appscreener Scan * &#x60;SonarQube API Import&#x60; - SonarQube API Import * &#x60;SonarQube Scan&#x60; - SonarQube Scan * &#x60;SonarQube Scan detailed&#x60; - SonarQube Scan detailed * &#x60;Sonatype Application Scan&#x60; - Sonatype Application Scan * &#x60;SpotBugs Scan&#x60; - SpotBugs Scan * &#x60;SSH Audit Importer&#x60; - SSH Audit Importer * &#x60;SSL Labs Scan&#x60; - SSL Labs Scan * &#x60;Sslscan&#x60; - Sslscan * &#x60;Sslyze Scan&#x60; - Sslyze Scan * &#x60;SSLyze Scan (JSON)&#x60; - SSLyze Scan (JSON) * &#x60;StackHawk HawkScan&#x60; - StackHawk HawkScan * &#x60;Sysdig Vulnerability Report&#x60; - Sysdig Vulnerability Report * &#x60;Talisman Scan&#x60; - Talisman Scan * &#x60;Tenable Scan&#x60; - Tenable Scan * &#x60;Terrascan Scan&#x60; - Terrascan Scan * &#x60;Testssl Scan&#x60; - Testssl Scan * &#x60;TFSec Scan&#x60; - TFSec Scan * &#x60;Threagile risks report&#x60; - Threagile risks report * &#x60;Trivy Operator Scan&#x60; - Trivy Operator Scan * &#x60;Trivy Scan&#x60; - Trivy Scan * &#x60;Trufflehog Scan&#x60; - Trufflehog Scan * &#x60;Trufflehog3 Scan&#x60; - Trufflehog3 Scan * &#x60;Trustwave Fusion API Scan&#x60; - Trustwave Fusion API Scan * &#x60;Trustwave Scan (CSV)&#x60; - Trustwave Scan (CSV) * &#x60;Twistlock Image Scan&#x60; - Twistlock Image Scan * &#x60;VCG Scan&#x60; - VCG Scan * &#x60;Veracode Scan&#x60; - Veracode Scan * &#x60;Veracode SourceClear Scan&#x60; - Veracode SourceClear Scan * &#x60;Vulners&#x60; - Vulners * &#x60;Wapiti Scan&#x60; - Wapiti Scan * &#x60;Wazuh&#x60; - Wazuh * &#x60;WFuzz JSON report&#x60; - WFuzz JSON report * &#x60;Whispers Scan&#x60; - Whispers Scan * &#x60;WhiteHat Sentinel&#x60; - WhiteHat Sentinel * &#x60;Wpscan&#x60; - Wpscan * &#x60;Xanitizer Scan&#x60; - Xanitizer Scan * &#x60;Yarn Audit Scan&#x60; - Yarn Audit Scan * &#x60;ZAP Scan&#x60; - ZAP Scan | 
**EndpointToAdd** | Pointer to **int32** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**ProductTypeName** | Pointer to **string** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**EngagementName** | Pointer to **string** |  | [optional] 
**EngagementEndDate** | Pointer to **string** | End Date for Engagement. Default is current time + 365 days. Required format year-month-day | [optional] 
**SourceCodeManagementUri** | Pointer to **string** | Resource link to source code | [optional] 
**Test** | Pointer to **int32** |  | [optional] 
**TestTitle** | Pointer to **string** |  | [optional] 
**AutoCreateContext** | Pointer to **bool** |  | [optional] 
**DeduplicationOnEngagement** | Pointer to **bool** |  | [optional] 
**PushToJira** | Pointer to **bool** |  | [optional] [default to false]
**CloseOldFindings** | Pointer to **bool** | Select if old findings no longer present in the report get closed as mitigated when importing. | [optional] [default to true]
**CloseOldFindingsProductScope** | Pointer to **bool** | Select if close_old_findings applies to all findings of the same type in the product. By default, it is false meaning that only old findings of the same type in the engagement are in scope. Note that this only applies on the first call to reimport-scan. | [optional] [default to false]
**Version** | Pointer to **string** | Version that will be set on existing Test object. Leave empty to leave existing value in place. | [optional] 
**BuildId** | Pointer to **string** | ID of the build that was scanned. | [optional] 
**BranchTag** | Pointer to **string** | Branch or Tag that was scanned. | [optional] 
**CommitHash** | Pointer to **string** | Commit that was scanned. | [optional] 
**ApiScanConfiguration** | Pointer to **NullableInt32** |  | [optional] 
**Service** | Pointer to **string** | A service is a self-contained piece of functionality within a Product. This is an optional field which is used in deduplication and closing of old findings when set. This affects the whole engagement/product depending on your deduplication scope. | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Lead** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]string** | Modify existing tags that help describe this scan. (Existing test tags will be overwritten) | [optional] 
**GroupBy** | Pointer to **string** | Choose an option to automatically group new findings by the chosen option.  * &#x60;component_name&#x60; - Component Name * &#x60;component_name+component_version&#x60; - Component Name + Version * &#x60;file_path&#x60; - File path * &#x60;finding_title&#x60; - Finding Title | [optional] 
**CreateFindingGroupsForAllFindings** | Pointer to **bool** | If set to false, finding groups will only be created when there is more than one grouped finding | [optional] [default to true]
**TestId** | **int32** |  | [readonly] 
**EngagementId** | **int32** |  | [readonly] 
**ProductId** | **int32** |  | [readonly] 
**ProductTypeId** | **int32** |  | [readonly] 
**Statistics** | [**ImportStatistics**](ImportStatistics.md) |  | [readonly] 
**ApplyTagsToFindings** | Pointer to **bool** | If set to True, the tags will be applied to the findings | [optional] 

## Methods

### NewReImportScan

`func NewReImportScan(active bool, verified bool, scanType string, testId int32, engagementId int32, productId int32, productTypeId int32, statistics ImportStatistics, ) *ReImportScan`

NewReImportScan instantiates a new ReImportScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReImportScanWithDefaults

`func NewReImportScanWithDefaults() *ReImportScan`

NewReImportScanWithDefaults instantiates a new ReImportScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanDate

`func (o *ReImportScan) GetScanDate() string`

GetScanDate returns the ScanDate field if non-nil, zero value otherwise.

### GetScanDateOk

`func (o *ReImportScan) GetScanDateOk() (*string, bool)`

GetScanDateOk returns a tuple with the ScanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDate

`func (o *ReImportScan) SetScanDate(v string)`

SetScanDate sets ScanDate field to given value.

### HasScanDate

`func (o *ReImportScan) HasScanDate() bool`

HasScanDate returns a boolean if a field has been set.

### GetMinimumSeverity

`func (o *ReImportScan) GetMinimumSeverity() string`

GetMinimumSeverity returns the MinimumSeverity field if non-nil, zero value otherwise.

### GetMinimumSeverityOk

`func (o *ReImportScan) GetMinimumSeverityOk() (*string, bool)`

GetMinimumSeverityOk returns a tuple with the MinimumSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumSeverity

`func (o *ReImportScan) SetMinimumSeverity(v string)`

SetMinimumSeverity sets MinimumSeverity field to given value.

### HasMinimumSeverity

`func (o *ReImportScan) HasMinimumSeverity() bool`

HasMinimumSeverity returns a boolean if a field has been set.

### GetActive

`func (o *ReImportScan) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ReImportScan) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ReImportScan) SetActive(v bool)`

SetActive sets Active field to given value.


### GetVerified

`func (o *ReImportScan) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ReImportScan) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ReImportScan) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetDoNotReactivate

`func (o *ReImportScan) GetDoNotReactivate() bool`

GetDoNotReactivate returns the DoNotReactivate field if non-nil, zero value otherwise.

### GetDoNotReactivateOk

`func (o *ReImportScan) GetDoNotReactivateOk() (*bool, bool)`

GetDoNotReactivateOk returns a tuple with the DoNotReactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotReactivate

`func (o *ReImportScan) SetDoNotReactivate(v bool)`

SetDoNotReactivate sets DoNotReactivate field to given value.

### HasDoNotReactivate

`func (o *ReImportScan) HasDoNotReactivate() bool`

HasDoNotReactivate returns a boolean if a field has been set.

### GetScanType

`func (o *ReImportScan) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ReImportScan) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ReImportScan) SetScanType(v string)`

SetScanType sets ScanType field to given value.


### GetEndpointToAdd

`func (o *ReImportScan) GetEndpointToAdd() int32`

GetEndpointToAdd returns the EndpointToAdd field if non-nil, zero value otherwise.

### GetEndpointToAddOk

`func (o *ReImportScan) GetEndpointToAddOk() (*int32, bool)`

GetEndpointToAddOk returns a tuple with the EndpointToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointToAdd

`func (o *ReImportScan) SetEndpointToAdd(v int32)`

SetEndpointToAdd sets EndpointToAdd field to given value.

### HasEndpointToAdd

`func (o *ReImportScan) HasEndpointToAdd() bool`

HasEndpointToAdd returns a boolean if a field has been set.

### GetFile

`func (o *ReImportScan) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ReImportScan) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ReImportScan) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *ReImportScan) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetProductTypeName

`func (o *ReImportScan) GetProductTypeName() string`

GetProductTypeName returns the ProductTypeName field if non-nil, zero value otherwise.

### GetProductTypeNameOk

`func (o *ReImportScan) GetProductTypeNameOk() (*string, bool)`

GetProductTypeNameOk returns a tuple with the ProductTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeName

`func (o *ReImportScan) SetProductTypeName(v string)`

SetProductTypeName sets ProductTypeName field to given value.

### HasProductTypeName

`func (o *ReImportScan) HasProductTypeName() bool`

HasProductTypeName returns a boolean if a field has been set.

### GetProductName

`func (o *ReImportScan) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ReImportScan) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ReImportScan) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *ReImportScan) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetEngagementName

`func (o *ReImportScan) GetEngagementName() string`

GetEngagementName returns the EngagementName field if non-nil, zero value otherwise.

### GetEngagementNameOk

`func (o *ReImportScan) GetEngagementNameOk() (*string, bool)`

GetEngagementNameOk returns a tuple with the EngagementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementName

`func (o *ReImportScan) SetEngagementName(v string)`

SetEngagementName sets EngagementName field to given value.

### HasEngagementName

`func (o *ReImportScan) HasEngagementName() bool`

HasEngagementName returns a boolean if a field has been set.

### GetEngagementEndDate

`func (o *ReImportScan) GetEngagementEndDate() string`

GetEngagementEndDate returns the EngagementEndDate field if non-nil, zero value otherwise.

### GetEngagementEndDateOk

`func (o *ReImportScan) GetEngagementEndDateOk() (*string, bool)`

GetEngagementEndDateOk returns a tuple with the EngagementEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementEndDate

`func (o *ReImportScan) SetEngagementEndDate(v string)`

SetEngagementEndDate sets EngagementEndDate field to given value.

### HasEngagementEndDate

`func (o *ReImportScan) HasEngagementEndDate() bool`

HasEngagementEndDate returns a boolean if a field has been set.

### GetSourceCodeManagementUri

`func (o *ReImportScan) GetSourceCodeManagementUri() string`

GetSourceCodeManagementUri returns the SourceCodeManagementUri field if non-nil, zero value otherwise.

### GetSourceCodeManagementUriOk

`func (o *ReImportScan) GetSourceCodeManagementUriOk() (*string, bool)`

GetSourceCodeManagementUriOk returns a tuple with the SourceCodeManagementUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeManagementUri

`func (o *ReImportScan) SetSourceCodeManagementUri(v string)`

SetSourceCodeManagementUri sets SourceCodeManagementUri field to given value.

### HasSourceCodeManagementUri

`func (o *ReImportScan) HasSourceCodeManagementUri() bool`

HasSourceCodeManagementUri returns a boolean if a field has been set.

### GetTest

`func (o *ReImportScan) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *ReImportScan) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *ReImportScan) SetTest(v int32)`

SetTest sets Test field to given value.

### HasTest

`func (o *ReImportScan) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTestTitle

`func (o *ReImportScan) GetTestTitle() string`

GetTestTitle returns the TestTitle field if non-nil, zero value otherwise.

### GetTestTitleOk

`func (o *ReImportScan) GetTestTitleOk() (*string, bool)`

GetTestTitleOk returns a tuple with the TestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTitle

`func (o *ReImportScan) SetTestTitle(v string)`

SetTestTitle sets TestTitle field to given value.

### HasTestTitle

`func (o *ReImportScan) HasTestTitle() bool`

HasTestTitle returns a boolean if a field has been set.

### GetAutoCreateContext

`func (o *ReImportScan) GetAutoCreateContext() bool`

GetAutoCreateContext returns the AutoCreateContext field if non-nil, zero value otherwise.

### GetAutoCreateContextOk

`func (o *ReImportScan) GetAutoCreateContextOk() (*bool, bool)`

GetAutoCreateContextOk returns a tuple with the AutoCreateContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateContext

`func (o *ReImportScan) SetAutoCreateContext(v bool)`

SetAutoCreateContext sets AutoCreateContext field to given value.

### HasAutoCreateContext

`func (o *ReImportScan) HasAutoCreateContext() bool`

HasAutoCreateContext returns a boolean if a field has been set.

### GetDeduplicationOnEngagement

`func (o *ReImportScan) GetDeduplicationOnEngagement() bool`

GetDeduplicationOnEngagement returns the DeduplicationOnEngagement field if non-nil, zero value otherwise.

### GetDeduplicationOnEngagementOk

`func (o *ReImportScan) GetDeduplicationOnEngagementOk() (*bool, bool)`

GetDeduplicationOnEngagementOk returns a tuple with the DeduplicationOnEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationOnEngagement

`func (o *ReImportScan) SetDeduplicationOnEngagement(v bool)`

SetDeduplicationOnEngagement sets DeduplicationOnEngagement field to given value.

### HasDeduplicationOnEngagement

`func (o *ReImportScan) HasDeduplicationOnEngagement() bool`

HasDeduplicationOnEngagement returns a boolean if a field has been set.

### GetPushToJira

`func (o *ReImportScan) GetPushToJira() bool`

GetPushToJira returns the PushToJira field if non-nil, zero value otherwise.

### GetPushToJiraOk

`func (o *ReImportScan) GetPushToJiraOk() (*bool, bool)`

GetPushToJiraOk returns a tuple with the PushToJira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToJira

`func (o *ReImportScan) SetPushToJira(v bool)`

SetPushToJira sets PushToJira field to given value.

### HasPushToJira

`func (o *ReImportScan) HasPushToJira() bool`

HasPushToJira returns a boolean if a field has been set.

### GetCloseOldFindings

`func (o *ReImportScan) GetCloseOldFindings() bool`

GetCloseOldFindings returns the CloseOldFindings field if non-nil, zero value otherwise.

### GetCloseOldFindingsOk

`func (o *ReImportScan) GetCloseOldFindingsOk() (*bool, bool)`

GetCloseOldFindingsOk returns a tuple with the CloseOldFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseOldFindings

`func (o *ReImportScan) SetCloseOldFindings(v bool)`

SetCloseOldFindings sets CloseOldFindings field to given value.

### HasCloseOldFindings

`func (o *ReImportScan) HasCloseOldFindings() bool`

HasCloseOldFindings returns a boolean if a field has been set.

### GetCloseOldFindingsProductScope

`func (o *ReImportScan) GetCloseOldFindingsProductScope() bool`

GetCloseOldFindingsProductScope returns the CloseOldFindingsProductScope field if non-nil, zero value otherwise.

### GetCloseOldFindingsProductScopeOk

`func (o *ReImportScan) GetCloseOldFindingsProductScopeOk() (*bool, bool)`

GetCloseOldFindingsProductScopeOk returns a tuple with the CloseOldFindingsProductScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseOldFindingsProductScope

`func (o *ReImportScan) SetCloseOldFindingsProductScope(v bool)`

SetCloseOldFindingsProductScope sets CloseOldFindingsProductScope field to given value.

### HasCloseOldFindingsProductScope

`func (o *ReImportScan) HasCloseOldFindingsProductScope() bool`

HasCloseOldFindingsProductScope returns a boolean if a field has been set.

### GetVersion

`func (o *ReImportScan) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReImportScan) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReImportScan) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReImportScan) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBuildId

`func (o *ReImportScan) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *ReImportScan) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *ReImportScan) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *ReImportScan) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetBranchTag

`func (o *ReImportScan) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *ReImportScan) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *ReImportScan) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *ReImportScan) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### GetCommitHash

`func (o *ReImportScan) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ReImportScan) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ReImportScan) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ReImportScan) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetApiScanConfiguration

`func (o *ReImportScan) GetApiScanConfiguration() int32`

GetApiScanConfiguration returns the ApiScanConfiguration field if non-nil, zero value otherwise.

### GetApiScanConfigurationOk

`func (o *ReImportScan) GetApiScanConfigurationOk() (*int32, bool)`

GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiScanConfiguration

`func (o *ReImportScan) SetApiScanConfiguration(v int32)`

SetApiScanConfiguration sets ApiScanConfiguration field to given value.

### HasApiScanConfiguration

`func (o *ReImportScan) HasApiScanConfiguration() bool`

HasApiScanConfiguration returns a boolean if a field has been set.

### SetApiScanConfigurationNil

`func (o *ReImportScan) SetApiScanConfigurationNil(b bool)`

 SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil

### UnsetApiScanConfiguration
`func (o *ReImportScan) UnsetApiScanConfiguration()`

UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil
### GetService

`func (o *ReImportScan) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ReImportScan) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ReImportScan) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ReImportScan) HasService() bool`

HasService returns a boolean if a field has been set.

### GetEnvironment

`func (o *ReImportScan) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ReImportScan) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ReImportScan) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ReImportScan) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLead

`func (o *ReImportScan) GetLead() int32`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ReImportScan) GetLeadOk() (*int32, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ReImportScan) SetLead(v int32)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ReImportScan) HasLead() bool`

HasLead returns a boolean if a field has been set.

### SetLeadNil

`func (o *ReImportScan) SetLeadNil(b bool)`

 SetLeadNil sets the value for Lead to be an explicit nil

### UnsetLead
`func (o *ReImportScan) UnsetLead()`

UnsetLead ensures that no value is present for Lead, not even an explicit nil
### GetTags

`func (o *ReImportScan) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReImportScan) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReImportScan) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReImportScan) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGroupBy

`func (o *ReImportScan) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *ReImportScan) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *ReImportScan) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *ReImportScan) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetCreateFindingGroupsForAllFindings

`func (o *ReImportScan) GetCreateFindingGroupsForAllFindings() bool`

GetCreateFindingGroupsForAllFindings returns the CreateFindingGroupsForAllFindings field if non-nil, zero value otherwise.

### GetCreateFindingGroupsForAllFindingsOk

`func (o *ReImportScan) GetCreateFindingGroupsForAllFindingsOk() (*bool, bool)`

GetCreateFindingGroupsForAllFindingsOk returns a tuple with the CreateFindingGroupsForAllFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFindingGroupsForAllFindings

`func (o *ReImportScan) SetCreateFindingGroupsForAllFindings(v bool)`

SetCreateFindingGroupsForAllFindings sets CreateFindingGroupsForAllFindings field to given value.

### HasCreateFindingGroupsForAllFindings

`func (o *ReImportScan) HasCreateFindingGroupsForAllFindings() bool`

HasCreateFindingGroupsForAllFindings returns a boolean if a field has been set.

### GetTestId

`func (o *ReImportScan) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *ReImportScan) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *ReImportScan) SetTestId(v int32)`

SetTestId sets TestId field to given value.


### GetEngagementId

`func (o *ReImportScan) GetEngagementId() int32`

GetEngagementId returns the EngagementId field if non-nil, zero value otherwise.

### GetEngagementIdOk

`func (o *ReImportScan) GetEngagementIdOk() (*int32, bool)`

GetEngagementIdOk returns a tuple with the EngagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementId

`func (o *ReImportScan) SetEngagementId(v int32)`

SetEngagementId sets EngagementId field to given value.


### GetProductId

`func (o *ReImportScan) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ReImportScan) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ReImportScan) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetProductTypeId

`func (o *ReImportScan) GetProductTypeId() int32`

GetProductTypeId returns the ProductTypeId field if non-nil, zero value otherwise.

### GetProductTypeIdOk

`func (o *ReImportScan) GetProductTypeIdOk() (*int32, bool)`

GetProductTypeIdOk returns a tuple with the ProductTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeId

`func (o *ReImportScan) SetProductTypeId(v int32)`

SetProductTypeId sets ProductTypeId field to given value.


### GetStatistics

`func (o *ReImportScan) GetStatistics() ImportStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ReImportScan) GetStatisticsOk() (*ImportStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ReImportScan) SetStatistics(v ImportStatistics)`

SetStatistics sets Statistics field to given value.


### GetApplyTagsToFindings

`func (o *ReImportScan) GetApplyTagsToFindings() bool`

GetApplyTagsToFindings returns the ApplyTagsToFindings field if non-nil, zero value otherwise.

### GetApplyTagsToFindingsOk

`func (o *ReImportScan) GetApplyTagsToFindingsOk() (*bool, bool)`

GetApplyTagsToFindingsOk returns a tuple with the ApplyTagsToFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTagsToFindings

`func (o *ReImportScan) SetApplyTagsToFindings(v bool)`

SetApplyTagsToFindings sets ApplyTagsToFindings field to given value.

### HasApplyTagsToFindings

`func (o *ReImportScan) HasApplyTagsToFindings() bool`

HasApplyTagsToFindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


