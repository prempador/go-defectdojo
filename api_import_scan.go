/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


type ImportScanAPI interface {

	/*
	ImportScanCreate Method for ImportScanCreate

	Imports a scan report into an engagement or product.

By ID:
- Create a Product (or use an existing product)
- Create an Engagement inside the product
- Provide the id of the engagement in the `engagement` parameter

In this scenario a new Test will be created inside the engagement.

By Names:
- Create a Product (or use an existing product)
- Create an Engagement inside the product
- Provide `product_name`
- Provide `engagement_name`
- Optionally provide `product_type_name`

In this scenario Defect Dojo will look up the Engagement by the provided details.

When using names you can let the importer automatically create Engagements, Products and Product_Types
by using `auto_create_context=True`.

When `auto_create_context` is set to `True` you can use `deduplication_on_engagement` to restrict deduplication for
imported Findings to the newly created Engagement.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiImportScanCreateRequest
	*/
	ImportScanCreate(ctx context.Context) ApiImportScanCreateRequest

	// ImportScanCreateExecute executes the request
	//  @return ImportScan
	ImportScanCreateExecute(r ApiImportScanCreateRequest) (*ImportScan, *http.Response, error)
}

// ImportScanAPIService ImportScanAPI service
type ImportScanAPIService service

type ApiImportScanCreateRequest struct {
	ctx context.Context
	ApiService ImportScanAPI
	active *bool
	verified *bool
	scanType *string
	scanDate *string
	minimumSeverity *string
	endpointToAdd *int32
	file *os.File
	productTypeName *string
	productName *string
	engagementName *string
	engagementEndDate *string
	sourceCodeManagementUri *string
	engagement *int32
	testTitle *string
	autoCreateContext *bool
	deduplicationOnEngagement *bool
	lead *int32
	tags *[]string
	closeOldFindings *bool
	closeOldFindingsProductScope *bool
	pushToJira *bool
	environment *string
	version *string
	buildId *string
	branchTag *string
	commitHash *string
	apiScanConfiguration *int32
	service *string
	groupBy *string
	createFindingGroupsForAllFindings *bool
}

// Override the active setting from the tool.
func (r ApiImportScanCreateRequest) Active(active bool) ApiImportScanCreateRequest {
	r.active = &active
	return r
}

// Override the verified setting from the tool.
func (r ApiImportScanCreateRequest) Verified(verified bool) ApiImportScanCreateRequest {
	r.verified = &verified
	return r
}

// * &#x60;Acunetix Scan&#x60; - Acunetix Scan * &#x60;Acunetix360 Scan&#x60; - Acunetix360 Scan * &#x60;Anchore Engine Scan&#x60; - Anchore Engine Scan * &#x60;Anchore Enterprise Policy Check&#x60; - Anchore Enterprise Policy Check * &#x60;Anchore Grype&#x60; - Anchore Grype * &#x60;AnchoreCTL Policies Report&#x60; - AnchoreCTL Policies Report * &#x60;AnchoreCTL Vuln Report&#x60; - AnchoreCTL Vuln Report * &#x60;AppSpider Scan&#x60; - AppSpider Scan * &#x60;Aqua Scan&#x60; - Aqua Scan * &#x60;Arachni Scan&#x60; - Arachni Scan * &#x60;AuditJS Scan&#x60; - AuditJS Scan * &#x60;AWS Prowler Scan&#x60; - AWS Prowler Scan * &#x60;AWS Prowler V3&#x60; - AWS Prowler V3 * &#x60;AWS Scout2 Scan&#x60; - AWS Scout2 Scan * &#x60;AWS Security Finding Format (ASFF) Scan&#x60; - AWS Security Finding Format (ASFF) Scan * &#x60;AWS Security Hub Scan&#x60; - AWS Security Hub Scan * &#x60;Azure Security Center Recommendations Scan&#x60; - Azure Security Center Recommendations Scan * &#x60;Bandit Scan&#x60; - Bandit Scan * &#x60;BlackDuck API&#x60; - BlackDuck API * &#x60;Blackduck Binary Analysis&#x60; - Blackduck Binary Analysis * &#x60;Blackduck Component Risk&#x60; - Blackduck Component Risk * &#x60;Blackduck Hub Scan&#x60; - Blackduck Hub Scan * &#x60;Brakeman Scan&#x60; - Brakeman Scan * &#x60;Bugcrowd API Import&#x60; - Bugcrowd API Import * &#x60;BugCrowd Scan&#x60; - BugCrowd Scan * &#x60;Bundler-Audit Scan&#x60; - Bundler-Audit Scan * &#x60;Burp Enterprise Scan&#x60; - Burp Enterprise Scan * &#x60;Burp GraphQL API&#x60; - Burp GraphQL API * &#x60;Burp REST API&#x60; - Burp REST API * &#x60;Burp Scan&#x60; - Burp Scan * &#x60;CargoAudit Scan&#x60; - CargoAudit Scan * &#x60;Checkmarx OSA&#x60; - Checkmarx OSA * &#x60;Checkmarx Scan&#x60; - Checkmarx Scan * &#x60;Checkmarx Scan detailed&#x60; - Checkmarx Scan detailed * &#x60;Checkov Scan&#x60; - Checkov Scan * &#x60;Clair Klar Scan&#x60; - Clair Klar Scan * &#x60;Clair Scan&#x60; - Clair Scan * &#x60;Cloudsploit Scan&#x60; - Cloudsploit Scan * &#x60;Cobalt.io API Import&#x60; - Cobalt.io API Import * &#x60;Cobalt.io Scan&#x60; - Cobalt.io Scan * &#x60;Codechecker Report native&#x60; - Codechecker Report native * &#x60;Contrast Scan&#x60; - Contrast Scan * &#x60;Coverity API&#x60; - Coverity API * &#x60;Crashtest Security JSON File&#x60; - Crashtest Security JSON File * &#x60;Crashtest Security XML File&#x60; - Crashtest Security XML File * &#x60;CredScan Scan&#x60; - CredScan Scan * &#x60;CycloneDX Scan&#x60; - CycloneDX Scan * &#x60;DawnScanner Scan&#x60; - DawnScanner Scan * &#x60;Dependency Check Scan&#x60; - Dependency Check Scan * &#x60;Dependency Track Finding Packaging Format (FPF) Export&#x60; - Dependency Track Finding Packaging Format (FPF) Export * &#x60;Detect-secrets Scan&#x60; - Detect-secrets Scan * &#x60;docker-bench-security Scan&#x60; - docker-bench-security Scan * &#x60;Dockle Scan&#x60; - Dockle Scan * &#x60;DrHeader JSON Importer&#x60; - DrHeader JSON Importer * &#x60;DSOP Scan&#x60; - DSOP Scan * &#x60;Edgescan Scan&#x60; - Edgescan Scan * &#x60;ESLint Scan&#x60; - ESLint Scan * &#x60;Fortify Scan&#x60; - Fortify Scan * &#x60;Generic Findings Import&#x60; - Generic Findings Import * &#x60;Ggshield Scan&#x60; - Ggshield Scan * &#x60;Github Vulnerability Scan&#x60; - Github Vulnerability Scan * &#x60;GitLab API Fuzzing Report Scan&#x60; - GitLab API Fuzzing Report Scan * &#x60;GitLab Container Scan&#x60; - GitLab Container Scan * &#x60;GitLab DAST Report&#x60; - GitLab DAST Report * &#x60;GitLab Dependency Scanning Report&#x60; - GitLab Dependency Scanning Report * &#x60;GitLab SAST Report&#x60; - GitLab SAST Report * &#x60;GitLab Secret Detection Report&#x60; - GitLab Secret Detection Report * &#x60;Gitleaks Scan&#x60; - Gitleaks Scan * &#x60;Gosec Scanner&#x60; - Gosec Scanner * &#x60;Govulncheck Scanner&#x60; - Govulncheck Scanner * &#x60;HackerOne Cases&#x60; - HackerOne Cases * &#x60;Hadolint Dockerfile check&#x60; - Hadolint Dockerfile check * &#x60;Harbor Vulnerability Scan&#x60; - Harbor Vulnerability Scan * &#x60;HCLAppScan XML&#x60; - HCLAppScan XML * &#x60;Horusec Scan&#x60; - Horusec Scan * &#x60;Humble Json Importer&#x60; - Humble Json Importer * &#x60;HuskyCI Report&#x60; - HuskyCI Report * &#x60;Hydra Scan&#x60; - Hydra Scan * &#x60;IBM AppScan DAST&#x60; - IBM AppScan DAST * &#x60;Immuniweb Scan&#x60; - Immuniweb Scan * &#x60;IntSights Report&#x60; - IntSights Report * &#x60;JFrog Xray API Summary Artifact Scan&#x60; - JFrog Xray API Summary Artifact Scan * &#x60;JFrog Xray On Demand Binary Scan&#x60; - JFrog Xray On Demand Binary Scan * &#x60;JFrog Xray Scan&#x60; - JFrog Xray Scan * &#x60;JFrog Xray Unified Scan&#x60; - JFrog Xray Unified Scan * &#x60;KICS Scan&#x60; - KICS Scan * &#x60;Kiuwan Scan&#x60; - Kiuwan Scan * &#x60;kube-bench Scan&#x60; - kube-bench Scan * &#x60;KubeHunter Scan&#x60; - KubeHunter Scan * &#x60;Meterian Scan&#x60; - Meterian Scan * &#x60;Microfocus Webinspect Scan&#x60; - Microfocus Webinspect Scan * &#x60;MobSF Scan&#x60; - MobSF Scan * &#x60;Mobsfscan Scan&#x60; - Mobsfscan Scan * &#x60;Mozilla Observatory Scan&#x60; - Mozilla Observatory Scan * &#x60;MSDefender Parser&#x60; - MSDefender Parser * &#x60;Netsparker Scan&#x60; - Netsparker Scan * &#x60;NeuVector (compliance)&#x60; - NeuVector (compliance) * &#x60;NeuVector (REST)&#x60; - NeuVector (REST) * &#x60;Nexpose Scan&#x60; - Nexpose Scan * &#x60;Nikto Scan&#x60; - Nikto Scan * &#x60;Nmap Scan&#x60; - Nmap Scan * &#x60;Node Security Platform Scan&#x60; - Node Security Platform Scan * &#x60;NPM Audit Scan&#x60; - NPM Audit Scan * &#x60;Nuclei Scan&#x60; - Nuclei Scan * &#x60;Openscap Vulnerability Scan&#x60; - Openscap Vulnerability Scan * &#x60;OpenVAS CSV&#x60; - OpenVAS CSV * &#x60;OpenVAS XML&#x60; - OpenVAS XML * &#x60;ORT evaluated model Importer&#x60; - ORT evaluated model Importer * &#x60;OssIndex Devaudit SCA Scan Importer&#x60; - OssIndex Devaudit SCA Scan Importer * &#x60;Outpost24 Scan&#x60; - Outpost24 Scan * &#x60;PHP Security Audit v2&#x60; - PHP Security Audit v2 * &#x60;PHP Symfony Security Check&#x60; - PHP Symfony Security Check * &#x60;pip-audit Scan&#x60; - pip-audit Scan * &#x60;PMD Scan&#x60; - PMD Scan * &#x60;Popeye Scan&#x60; - Popeye Scan * &#x60;PWN SAST&#x60; - PWN SAST * &#x60;Qualys Infrastructure Scan (WebGUI XML)&#x60; - Qualys Infrastructure Scan (WebGUI XML) * &#x60;Qualys Scan&#x60; - Qualys Scan * &#x60;Qualys Webapp Scan&#x60; - Qualys Webapp Scan * &#x60;Retire.js Scan&#x60; - Retire.js Scan * &#x60;Risk Recon API Importer&#x60; - Risk Recon API Importer * &#x60;Rubocop Scan&#x60; - Rubocop Scan * &#x60;Rusty Hog Scan&#x60; - Rusty Hog Scan * &#x60;SARIF&#x60; - SARIF * &#x60;Scantist Scan&#x60; - Scantist Scan * &#x60;Scout Suite Scan&#x60; - Scout Suite Scan * &#x60;Semgrep JSON Report&#x60; - Semgrep JSON Report * &#x60;SKF Scan&#x60; - SKF Scan * &#x60;Snyk Scan&#x60; - Snyk Scan * &#x60;Solar Appscreener Scan&#x60; - Solar Appscreener Scan * &#x60;SonarQube API Import&#x60; - SonarQube API Import * &#x60;SonarQube Scan&#x60; - SonarQube Scan * &#x60;SonarQube Scan detailed&#x60; - SonarQube Scan detailed * &#x60;Sonatype Application Scan&#x60; - Sonatype Application Scan * &#x60;SpotBugs Scan&#x60; - SpotBugs Scan * &#x60;SSH Audit Importer&#x60; - SSH Audit Importer * &#x60;SSL Labs Scan&#x60; - SSL Labs Scan * &#x60;Sslscan&#x60; - Sslscan * &#x60;Sslyze Scan&#x60; - Sslyze Scan * &#x60;SSLyze Scan (JSON)&#x60; - SSLyze Scan (JSON) * &#x60;StackHawk HawkScan&#x60; - StackHawk HawkScan * &#x60;Sysdig Vulnerability Report - Pipeline, Registry and Runtime (CSV)&#x60; - Sysdig Vulnerability Report - Pipeline, Registry and Runtime (CSV) * &#x60;Talisman Scan&#x60; - Talisman Scan * &#x60;Tenable Scan&#x60; - Tenable Scan * &#x60;Terrascan Scan&#x60; - Terrascan Scan * &#x60;Testssl Scan&#x60; - Testssl Scan * &#x60;TFSec Scan&#x60; - TFSec Scan * &#x60;Threagile risks report&#x60; - Threagile risks report * &#x60;Trivy Operator Scan&#x60; - Trivy Operator Scan * &#x60;Trivy Scan&#x60; - Trivy Scan * &#x60;Trufflehog Scan&#x60; - Trufflehog Scan * &#x60;Trufflehog3 Scan&#x60; - Trufflehog3 Scan * &#x60;Trustwave Fusion API Scan&#x60; - Trustwave Fusion API Scan * &#x60;Trustwave Scan (CSV)&#x60; - Trustwave Scan (CSV) * &#x60;Twistlock Image Scan&#x60; - Twistlock Image Scan * &#x60;VCG Scan&#x60; - VCG Scan * &#x60;Veracode Scan&#x60; - Veracode Scan * &#x60;Veracode SourceClear Scan&#x60; - Veracode SourceClear Scan * &#x60;Vulners&#x60; - Vulners * &#x60;Wapiti Scan&#x60; - Wapiti Scan * &#x60;Wazuh&#x60; - Wazuh * &#x60;WFuzz JSON report&#x60; - WFuzz JSON report * &#x60;Whispers Scan&#x60; - Whispers Scan * &#x60;WhiteHat Sentinel&#x60; - WhiteHat Sentinel * &#x60;Whitesource Scan&#x60; - Whitesource Scan * &#x60;Wpscan&#x60; - Wpscan * &#x60;Xanitizer Scan&#x60; - Xanitizer Scan * &#x60;Yarn Audit Scan&#x60; - Yarn Audit Scan * &#x60;ZAP Scan&#x60; - ZAP Scan
func (r ApiImportScanCreateRequest) ScanType(scanType string) ApiImportScanCreateRequest {
	r.scanType = &scanType
	return r
}

// Scan completion date will be used on all findings.
func (r ApiImportScanCreateRequest) ScanDate(scanDate string) ApiImportScanCreateRequest {
	r.scanDate = &scanDate
	return r
}

// Minimum severity level to be imported  * &#x60;Info&#x60; - Info * &#x60;Low&#x60; - Low * &#x60;Medium&#x60; - Medium * &#x60;High&#x60; - High * &#x60;Critical&#x60; - Critical
func (r ApiImportScanCreateRequest) MinimumSeverity(minimumSeverity string) ApiImportScanCreateRequest {
	r.minimumSeverity = &minimumSeverity
	return r
}

// The IP address, host name or full URL. It must be valid
func (r ApiImportScanCreateRequest) EndpointToAdd(endpointToAdd int32) ApiImportScanCreateRequest {
	r.endpointToAdd = &endpointToAdd
	return r
}

func (r ApiImportScanCreateRequest) File(file *os.File) ApiImportScanCreateRequest {
	r.file = file
	return r
}

func (r ApiImportScanCreateRequest) ProductTypeName(productTypeName string) ApiImportScanCreateRequest {
	r.productTypeName = &productTypeName
	return r
}

func (r ApiImportScanCreateRequest) ProductName(productName string) ApiImportScanCreateRequest {
	r.productName = &productName
	return r
}

func (r ApiImportScanCreateRequest) EngagementName(engagementName string) ApiImportScanCreateRequest {
	r.engagementName = &engagementName
	return r
}

// End Date for Engagement. Default is current time + 365 days. Required format year-month-day
func (r ApiImportScanCreateRequest) EngagementEndDate(engagementEndDate string) ApiImportScanCreateRequest {
	r.engagementEndDate = &engagementEndDate
	return r
}

// Resource link to source code
func (r ApiImportScanCreateRequest) SourceCodeManagementUri(sourceCodeManagementUri string) ApiImportScanCreateRequest {
	r.sourceCodeManagementUri = &sourceCodeManagementUri
	return r
}

func (r ApiImportScanCreateRequest) Engagement(engagement int32) ApiImportScanCreateRequest {
	r.engagement = &engagement
	return r
}

func (r ApiImportScanCreateRequest) TestTitle(testTitle string) ApiImportScanCreateRequest {
	r.testTitle = &testTitle
	return r
}

func (r ApiImportScanCreateRequest) AutoCreateContext(autoCreateContext bool) ApiImportScanCreateRequest {
	r.autoCreateContext = &autoCreateContext
	return r
}

func (r ApiImportScanCreateRequest) DeduplicationOnEngagement(deduplicationOnEngagement bool) ApiImportScanCreateRequest {
	r.deduplicationOnEngagement = &deduplicationOnEngagement
	return r
}

func (r ApiImportScanCreateRequest) Lead(lead int32) ApiImportScanCreateRequest {
	r.lead = &lead
	return r
}

// Add tags that help describe this scan.
func (r ApiImportScanCreateRequest) Tags(tags []string) ApiImportScanCreateRequest {
	r.tags = &tags
	return r
}

// Select if old findings no longer present in the report get closed as mitigated when importing. If service has been set, only the findings for this service will be closed.
func (r ApiImportScanCreateRequest) CloseOldFindings(closeOldFindings bool) ApiImportScanCreateRequest {
	r.closeOldFindings = &closeOldFindings
	return r
}

// Select if close_old_findings applies to all findings of the same type in the product. By default, it is false meaning that only old findings of the same type in the engagement are in scope.
func (r ApiImportScanCreateRequest) CloseOldFindingsProductScope(closeOldFindingsProductScope bool) ApiImportScanCreateRequest {
	r.closeOldFindingsProductScope = &closeOldFindingsProductScope
	return r
}

func (r ApiImportScanCreateRequest) PushToJira(pushToJira bool) ApiImportScanCreateRequest {
	r.pushToJira = &pushToJira
	return r
}

func (r ApiImportScanCreateRequest) Environment(environment string) ApiImportScanCreateRequest {
	r.environment = &environment
	return r
}

// Version that was scanned.
func (r ApiImportScanCreateRequest) Version(version string) ApiImportScanCreateRequest {
	r.version = &version
	return r
}

// ID of the build that was scanned.
func (r ApiImportScanCreateRequest) BuildId(buildId string) ApiImportScanCreateRequest {
	r.buildId = &buildId
	return r
}

// Branch or Tag that was scanned.
func (r ApiImportScanCreateRequest) BranchTag(branchTag string) ApiImportScanCreateRequest {
	r.branchTag = &branchTag
	return r
}

// Commit that was scanned.
func (r ApiImportScanCreateRequest) CommitHash(commitHash string) ApiImportScanCreateRequest {
	r.commitHash = &commitHash
	return r
}

func (r ApiImportScanCreateRequest) ApiScanConfiguration(apiScanConfiguration int32) ApiImportScanCreateRequest {
	r.apiScanConfiguration = &apiScanConfiguration
	return r
}

// A service is a self-contained piece of functionality within a Product. This is an optional field which is used in deduplication and closing of old findings when set. This affects the whole engagement/product depending on your deduplication scope.
func (r ApiImportScanCreateRequest) Service(service string) ApiImportScanCreateRequest {
	r.service = &service
	return r
}

// Choose an option to automatically group new findings by the chosen option.  * &#x60;component_name&#x60; - Component Name * &#x60;component_name+component_version&#x60; - Component Name + Version * &#x60;file_path&#x60; - File path * &#x60;finding_title&#x60; - Finding Title
func (r ApiImportScanCreateRequest) GroupBy(groupBy string) ApiImportScanCreateRequest {
	r.groupBy = &groupBy
	return r
}

// If set to false, finding groups will only be created when there is more than one grouped finding
func (r ApiImportScanCreateRequest) CreateFindingGroupsForAllFindings(createFindingGroupsForAllFindings bool) ApiImportScanCreateRequest {
	r.createFindingGroupsForAllFindings = &createFindingGroupsForAllFindings
	return r
}

func (r ApiImportScanCreateRequest) Execute() (*ImportScan, *http.Response, error) {
	return r.ApiService.ImportScanCreateExecute(r)
}

/*
ImportScanCreate Method for ImportScanCreate

Imports a scan report into an engagement or product.

By ID:
- Create a Product (or use an existing product)
- Create an Engagement inside the product
- Provide the id of the engagement in the `engagement` parameter

In this scenario a new Test will be created inside the engagement.

By Names:
- Create a Product (or use an existing product)
- Create an Engagement inside the product
- Provide `product_name`
- Provide `engagement_name`
- Optionally provide `product_type_name`

In this scenario Defect Dojo will look up the Engagement by the provided details.

When using names you can let the importer automatically create Engagements, Products and Product_Types
by using `auto_create_context=True`.

When `auto_create_context` is set to `True` you can use `deduplication_on_engagement` to restrict deduplication for
imported Findings to the newly created Engagement.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImportScanCreateRequest
*/
func (a *ImportScanAPIService) ImportScanCreate(ctx context.Context) ApiImportScanCreateRequest {
	return ApiImportScanCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ImportScan
func (a *ImportScanAPIService) ImportScanCreateExecute(r ApiImportScanCreateRequest) (*ImportScan, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ImportScan
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportScanAPIService.ImportScanCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/import-scan/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.active == nil {
		return localVarReturnValue, nil, reportError("active is required and must be specified")
	}
	if r.verified == nil {
		return localVarReturnValue, nil, reportError("verified is required and must be specified")
	}
	if r.scanType == nil {
		return localVarReturnValue, nil, reportError("scanType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.scanDate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scan_date", r.scanDate, "")
	}
	if r.minimumSeverity != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "minimum_severity", r.minimumSeverity, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "active", r.active, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "verified", r.verified, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "scan_type", r.scanType, "")
	if r.endpointToAdd != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "endpoint_to_add", r.endpointToAdd, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.productTypeName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "product_type_name", r.productTypeName, "")
	}
	if r.productName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "product_name", r.productName, "")
	}
	if r.engagementName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "engagement_name", r.engagementName, "")
	}
	if r.engagementEndDate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "engagement_end_date", r.engagementEndDate, "")
	}
	if r.sourceCodeManagementUri != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "source_code_management_uri", r.sourceCodeManagementUri, "")
	}
	if r.engagement != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "engagement", r.engagement, "")
	}
	if r.testTitle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "test_title", r.testTitle, "")
	}
	if r.autoCreateContext != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "auto_create_context", r.autoCreateContext, "")
	}
	if r.deduplicationOnEngagement != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "deduplication_on_engagement", r.deduplicationOnEngagement, "")
	}
	if r.lead != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "lead", r.lead, "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tags", r.tags, "csv")
	}
	if r.closeOldFindings != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "close_old_findings", r.closeOldFindings, "")
	}
	if r.closeOldFindingsProductScope != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "close_old_findings_product_scope", r.closeOldFindingsProductScope, "")
	}
	if r.pushToJira != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "push_to_jira", r.pushToJira, "")
	}
	if r.environment != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "environment", r.environment, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "version", r.version, "")
	}
	if r.buildId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "build_id", r.buildId, "")
	}
	if r.branchTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "branch_tag", r.branchTag, "")
	}
	if r.commitHash != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "commit_hash", r.commitHash, "")
	}
	if r.apiScanConfiguration != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "api_scan_configuration", r.apiScanConfiguration, "")
	}
	if r.service != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "service", r.service, "")
	}
	if r.groupBy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "group_by", r.groupBy, "")
	}
	if r.createFindingGroupsForAllFindings != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "create_finding_groups_for_all_findings", r.createFindingGroupsForAllFindings, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
