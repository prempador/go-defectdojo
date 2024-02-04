# \ReimportScanAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReimportScanCreate**](ReimportScanAPI.md#ReimportScanCreate) | **Post** /api/v2/reimport-scan/ | 



## ReimportScanCreate

> ReImportScan ReimportScanCreate(ctx).Active(active).Verified(verified).ScanType(scanType).ScanDate(scanDate).MinimumSeverity(minimumSeverity).DoNotReactivate(doNotReactivate).EndpointToAdd(endpointToAdd).File(file).ProductTypeName(productTypeName).ProductName(productName).EngagementName(engagementName).EngagementEndDate(engagementEndDate).SourceCodeManagementUri(sourceCodeManagementUri).Test(test).TestTitle(testTitle).AutoCreateContext(autoCreateContext).DeduplicationOnEngagement(deduplicationOnEngagement).PushToJira(pushToJira).CloseOldFindings(closeOldFindings).CloseOldFindingsProductScope(closeOldFindingsProductScope).Version(version).BuildId(buildId).BranchTag(branchTag).CommitHash(commitHash).ApiScanConfiguration(apiScanConfiguration).Service(service).Environment(environment).Lead(lead).Tags(tags).GroupBy(groupBy).CreateFindingGroupsForAllFindings(createFindingGroupsForAllFindings).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	active := true // bool | Override the active setting from the tool.
	verified := true // bool | Override the verified setting from the tool.
	scanType := "scanType_example" // string | * `Acunetix Scan` - Acunetix Scan * `Acunetix360 Scan` - Acunetix360 Scan * `Anchore Engine Scan` - Anchore Engine Scan * `Anchore Enterprise Policy Check` - Anchore Enterprise Policy Check * `Anchore Grype` - Anchore Grype * `AnchoreCTL Policies Report` - AnchoreCTL Policies Report * `AnchoreCTL Vuln Report` - AnchoreCTL Vuln Report * `AppSpider Scan` - AppSpider Scan * `Aqua Scan` - Aqua Scan * `Arachni Scan` - Arachni Scan * `AuditJS Scan` - AuditJS Scan * `AWS Prowler Scan` - AWS Prowler Scan * `AWS Prowler V3` - AWS Prowler V3 * `AWS Scout2 Scan` - AWS Scout2 Scan * `AWS Security Finding Format (ASFF) Scan` - AWS Security Finding Format (ASFF) Scan * `AWS Security Hub Scan` - AWS Security Hub Scan * `Azure Security Center Recommendations Scan` - Azure Security Center Recommendations Scan * `Bandit Scan` - Bandit Scan * `BlackDuck API` - BlackDuck API * `Blackduck Binary Analysis` - Blackduck Binary Analysis * `Blackduck Component Risk` - Blackduck Component Risk * `Blackduck Hub Scan` - Blackduck Hub Scan * `Brakeman Scan` - Brakeman Scan * `Bugcrowd API Import` - Bugcrowd API Import * `BugCrowd Scan` - BugCrowd Scan * `Bundler-Audit Scan` - Bundler-Audit Scan * `Burp Enterprise Scan` - Burp Enterprise Scan * `Burp GraphQL API` - Burp GraphQL API * `Burp REST API` - Burp REST API * `Burp Scan` - Burp Scan * `CargoAudit Scan` - CargoAudit Scan * `Checkmarx OSA` - Checkmarx OSA * `Checkmarx Scan` - Checkmarx Scan * `Checkmarx Scan detailed` - Checkmarx Scan detailed * `Checkov Scan` - Checkov Scan * `Clair Klar Scan` - Clair Klar Scan * `Clair Scan` - Clair Scan * `Cloudsploit Scan` - Cloudsploit Scan * `Cobalt.io API Import` - Cobalt.io API Import * `Cobalt.io Scan` - Cobalt.io Scan * `Codechecker Report native` - Codechecker Report native * `Contrast Scan` - Contrast Scan * `Coverity API` - Coverity API * `Crashtest Security JSON File` - Crashtest Security JSON File * `Crashtest Security XML File` - Crashtest Security XML File * `CredScan Scan` - CredScan Scan * `CycloneDX Scan` - CycloneDX Scan * `DawnScanner Scan` - DawnScanner Scan * `Dependency Check Scan` - Dependency Check Scan * `Dependency Track Finding Packaging Format (FPF) Export` - Dependency Track Finding Packaging Format (FPF) Export * `Detect-secrets Scan` - Detect-secrets Scan * `docker-bench-security Scan` - docker-bench-security Scan * `Dockle Scan` - Dockle Scan * `DrHeader JSON Importer` - DrHeader JSON Importer * `DSOP Scan` - DSOP Scan * `Edgescan Scan` - Edgescan Scan * `ESLint Scan` - ESLint Scan * `Fortify Scan` - Fortify Scan * `Generic Findings Import` - Generic Findings Import * `Ggshield Scan` - Ggshield Scan * `Github Vulnerability Scan` - Github Vulnerability Scan * `GitLab API Fuzzing Report Scan` - GitLab API Fuzzing Report Scan * `GitLab Container Scan` - GitLab Container Scan * `GitLab DAST Report` - GitLab DAST Report * `GitLab Dependency Scanning Report` - GitLab Dependency Scanning Report * `GitLab SAST Report` - GitLab SAST Report * `GitLab Secret Detection Report` - GitLab Secret Detection Report * `Gitleaks Scan` - Gitleaks Scan * `Gosec Scanner` - Gosec Scanner * `Govulncheck Scanner` - Govulncheck Scanner * `HackerOne Cases` - HackerOne Cases * `Hadolint Dockerfile check` - Hadolint Dockerfile check * `Harbor Vulnerability Scan` - Harbor Vulnerability Scan * `HCLAppScan XML` - HCLAppScan XML * `Horusec Scan` - Horusec Scan * `Humble Json Importer` - Humble Json Importer * `HuskyCI Report` - HuskyCI Report * `Hydra Scan` - Hydra Scan * `IBM AppScan DAST` - IBM AppScan DAST * `Immuniweb Scan` - Immuniweb Scan * `IntSights Report` - IntSights Report * `JFrog Xray API Summary Artifact Scan` - JFrog Xray API Summary Artifact Scan * `JFrog Xray On Demand Binary Scan` - JFrog Xray On Demand Binary Scan * `JFrog Xray Scan` - JFrog Xray Scan * `JFrog Xray Unified Scan` - JFrog Xray Unified Scan * `KICS Scan` - KICS Scan * `Kiuwan Scan` - Kiuwan Scan * `kube-bench Scan` - kube-bench Scan * `KubeHunter Scan` - KubeHunter Scan * `Meterian Scan` - Meterian Scan * `Microfocus Webinspect Scan` - Microfocus Webinspect Scan * `MobSF Scan` - MobSF Scan * `Mobsfscan Scan` - Mobsfscan Scan * `Mozilla Observatory Scan` - Mozilla Observatory Scan * `MSDefender Parser` - MSDefender Parser * `Netsparker Scan` - Netsparker Scan * `NeuVector (compliance)` - NeuVector (compliance) * `NeuVector (REST)` - NeuVector (REST) * `Nexpose Scan` - Nexpose Scan * `Nikto Scan` - Nikto Scan * `Nmap Scan` - Nmap Scan * `Node Security Platform Scan` - Node Security Platform Scan * `NPM Audit Scan` - NPM Audit Scan * `Nuclei Scan` - Nuclei Scan * `Openscap Vulnerability Scan` - Openscap Vulnerability Scan * `OpenVAS CSV` - OpenVAS CSV * `OpenVAS XML` - OpenVAS XML * `ORT evaluated model Importer` - ORT evaluated model Importer * `OssIndex Devaudit SCA Scan Importer` - OssIndex Devaudit SCA Scan Importer * `Outpost24 Scan` - Outpost24 Scan * `PHP Security Audit v2` - PHP Security Audit v2 * `PHP Symfony Security Check` - PHP Symfony Security Check * `pip-audit Scan` - pip-audit Scan * `PMD Scan` - PMD Scan * `Popeye Scan` - Popeye Scan * `PWN SAST` - PWN SAST * `Qualys Infrastructure Scan (WebGUI XML)` - Qualys Infrastructure Scan (WebGUI XML) * `Qualys Scan` - Qualys Scan * `Qualys Webapp Scan` - Qualys Webapp Scan * `Retire.js Scan` - Retire.js Scan * `Risk Recon API Importer` - Risk Recon API Importer * `Rubocop Scan` - Rubocop Scan * `Rusty Hog Scan` - Rusty Hog Scan * `SARIF` - SARIF * `Scantist Scan` - Scantist Scan * `Scout Suite Scan` - Scout Suite Scan * `Semgrep JSON Report` - Semgrep JSON Report * `SKF Scan` - SKF Scan * `Snyk Scan` - Snyk Scan * `Solar Appscreener Scan` - Solar Appscreener Scan * `SonarQube API Import` - SonarQube API Import * `SonarQube Scan` - SonarQube Scan * `SonarQube Scan detailed` - SonarQube Scan detailed * `Sonatype Application Scan` - Sonatype Application Scan * `SpotBugs Scan` - SpotBugs Scan * `SSH Audit Importer` - SSH Audit Importer * `SSL Labs Scan` - SSL Labs Scan * `Sslscan` - Sslscan * `Sslyze Scan` - Sslyze Scan * `SSLyze Scan (JSON)` - SSLyze Scan (JSON) * `StackHawk HawkScan` - StackHawk HawkScan * `Sysdig Vulnerability Report - Pipeline, Registry and Runtime (CSV)` - Sysdig Vulnerability Report - Pipeline, Registry and Runtime (CSV) * `Talisman Scan` - Talisman Scan * `Tenable Scan` - Tenable Scan * `Terrascan Scan` - Terrascan Scan * `Testssl Scan` - Testssl Scan * `TFSec Scan` - TFSec Scan * `Threagile risks report` - Threagile risks report * `Trivy Operator Scan` - Trivy Operator Scan * `Trivy Scan` - Trivy Scan * `Trufflehog Scan` - Trufflehog Scan * `Trufflehog3 Scan` - Trufflehog3 Scan * `Trustwave Fusion API Scan` - Trustwave Fusion API Scan * `Trustwave Scan (CSV)` - Trustwave Scan (CSV) * `Twistlock Image Scan` - Twistlock Image Scan * `VCG Scan` - VCG Scan * `Veracode Scan` - Veracode Scan * `Veracode SourceClear Scan` - Veracode SourceClear Scan * `Vulners` - Vulners * `Wapiti Scan` - Wapiti Scan * `Wazuh` - Wazuh * `WFuzz JSON report` - WFuzz JSON report * `Whispers Scan` - Whispers Scan * `WhiteHat Sentinel` - WhiteHat Sentinel * `Whitesource Scan` - Whitesource Scan * `Wpscan` - Wpscan * `Xanitizer Scan` - Xanitizer Scan * `Yarn Audit Scan` - Yarn Audit Scan * `ZAP Scan` - ZAP Scan
	scanDate := time.Now() // string | Scan completion date will be used on all findings. (optional)
	minimumSeverity := "minimumSeverity_example" // string | Minimum severity level to be imported  * `Info` - Info * `Low` - Low * `Medium` - Medium * `High` - High * `Critical` - Critical (optional) (default to "Info")
	doNotReactivate := true // bool | Select if the import should ignore active findings from the report, useful for triage-less scanners. Will keep existing findings closed, without reactivating them. For more information check the docs. (optional) (default to false)
	endpointToAdd := int32(56) // int32 |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	productTypeName := "productTypeName_example" // string |  (optional)
	productName := "productName_example" // string |  (optional)
	engagementName := "engagementName_example" // string |  (optional)
	engagementEndDate := time.Now() // string | End Date for Engagement. Default is current time + 365 days. Required format year-month-day (optional)
	sourceCodeManagementUri := "sourceCodeManagementUri_example" // string | Resource link to source code (optional)
	test := int32(56) // int32 |  (optional)
	testTitle := "testTitle_example" // string |  (optional)
	autoCreateContext := true // bool |  (optional)
	deduplicationOnEngagement := true // bool |  (optional)
	pushToJira := true // bool |  (optional) (default to false)
	closeOldFindings := true // bool | Select if old findings no longer present in the report get closed as mitigated when importing. (optional) (default to true)
	closeOldFindingsProductScope := true // bool | Select if close_old_findings applies to all findings of the same type in the product. By default, it is false meaning that only old findings of the same type in the engagement are in scope. Note that this only applies on the first call to reimport-scan. (optional) (default to false)
	version := "version_example" // string | Version that will be set on existing Test object. Leave empty to leave existing value in place. (optional)
	buildId := "buildId_example" // string | ID of the build that was scanned. (optional)
	branchTag := "branchTag_example" // string | Branch or Tag that was scanned. (optional)
	commitHash := "commitHash_example" // string | Commit that was scanned. (optional)
	apiScanConfiguration := int32(56) // int32 |  (optional)
	service := "service_example" // string | A service is a self-contained piece of functionality within a Product. This is an optional field which is used in deduplication and closing of old findings when set. This affects the whole engagement/product depending on your deduplication scope. (optional)
	environment := "environment_example" // string |  (optional)
	lead := int32(56) // int32 |  (optional)
	tags := []string{"Inner_example"} // []string | Modify existing tags that help describe this scan. (Existing test tags will be overwritten) (optional)
	groupBy := "groupBy_example" // string | Choose an option to automatically group new findings by the chosen option.  * `component_name` - Component Name * `component_name+component_version` - Component Name + Version * `file_path` - File path * `finding_title` - Finding Title (optional)
	createFindingGroupsForAllFindings := true // bool | If set to false, finding groups will only be created when there is more than one grouped finding (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReimportScanAPI.ReimportScanCreate(context.Background()).Active(active).Verified(verified).ScanType(scanType).ScanDate(scanDate).MinimumSeverity(minimumSeverity).DoNotReactivate(doNotReactivate).EndpointToAdd(endpointToAdd).File(file).ProductTypeName(productTypeName).ProductName(productName).EngagementName(engagementName).EngagementEndDate(engagementEndDate).SourceCodeManagementUri(sourceCodeManagementUri).Test(test).TestTitle(testTitle).AutoCreateContext(autoCreateContext).DeduplicationOnEngagement(deduplicationOnEngagement).PushToJira(pushToJira).CloseOldFindings(closeOldFindings).CloseOldFindingsProductScope(closeOldFindingsProductScope).Version(version).BuildId(buildId).BranchTag(branchTag).CommitHash(commitHash).ApiScanConfiguration(apiScanConfiguration).Service(service).Environment(environment).Lead(lead).Tags(tags).GroupBy(groupBy).CreateFindingGroupsForAllFindings(createFindingGroupsForAllFindings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReimportScanAPI.ReimportScanCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReimportScanCreate`: ReImportScan
	fmt.Fprintf(os.Stdout, "Response from `ReimportScanAPI.ReimportScanCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReimportScanCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | Override the active setting from the tool. | 
 **verified** | **bool** | Override the verified setting from the tool. | 
 **scanType** | **string** | * &#x60;Acunetix Scan&#x60; - Acunetix Scan * &#x60;Acunetix360 Scan&#x60; - Acunetix360 Scan * &#x60;Anchore Engine Scan&#x60; - Anchore Engine Scan * &#x60;Anchore Enterprise Policy Check&#x60; - Anchore Enterprise Policy Check * &#x60;Anchore Grype&#x60; - Anchore Grype * &#x60;AnchoreCTL Policies Report&#x60; - AnchoreCTL Policies Report * &#x60;AnchoreCTL Vuln Report&#x60; - AnchoreCTL Vuln Report * &#x60;AppSpider Scan&#x60; - AppSpider Scan * &#x60;Aqua Scan&#x60; - Aqua Scan * &#x60;Arachni Scan&#x60; - Arachni Scan * &#x60;AuditJS Scan&#x60; - AuditJS Scan * &#x60;AWS Prowler Scan&#x60; - AWS Prowler Scan * &#x60;AWS Prowler V3&#x60; - AWS Prowler V3 * &#x60;AWS Scout2 Scan&#x60; - AWS Scout2 Scan * &#x60;AWS Security Finding Format (ASFF) Scan&#x60; - AWS Security Finding Format (ASFF) Scan * &#x60;AWS Security Hub Scan&#x60; - AWS Security Hub Scan * &#x60;Azure Security Center Recommendations Scan&#x60; - Azure Security Center Recommendations Scan * &#x60;Bandit Scan&#x60; - Bandit Scan * &#x60;BlackDuck API&#x60; - BlackDuck API * &#x60;Blackduck Binary Analysis&#x60; - Blackduck Binary Analysis * &#x60;Blackduck Component Risk&#x60; - Blackduck Component Risk * &#x60;Blackduck Hub Scan&#x60; - Blackduck Hub Scan * &#x60;Brakeman Scan&#x60; - Brakeman Scan * &#x60;Bugcrowd API Import&#x60; - Bugcrowd API Import * &#x60;BugCrowd Scan&#x60; - BugCrowd Scan * &#x60;Bundler-Audit Scan&#x60; - Bundler-Audit Scan * &#x60;Burp Enterprise Scan&#x60; - Burp Enterprise Scan * &#x60;Burp GraphQL API&#x60; - Burp GraphQL API * &#x60;Burp REST API&#x60; - Burp REST API * &#x60;Burp Scan&#x60; - Burp Scan * &#x60;CargoAudit Scan&#x60; - CargoAudit Scan * &#x60;Checkmarx OSA&#x60; - Checkmarx OSA * &#x60;Checkmarx Scan&#x60; - Checkmarx Scan * &#x60;Checkmarx Scan detailed&#x60; - Checkmarx Scan detailed * &#x60;Checkov Scan&#x60; - Checkov Scan * &#x60;Clair Klar Scan&#x60; - Clair Klar Scan * &#x60;Clair Scan&#x60; - Clair Scan * &#x60;Cloudsploit Scan&#x60; - Cloudsploit Scan * &#x60;Cobalt.io API Import&#x60; - Cobalt.io API Import * &#x60;Cobalt.io Scan&#x60; - Cobalt.io Scan * &#x60;Codechecker Report native&#x60; - Codechecker Report native * &#x60;Contrast Scan&#x60; - Contrast Scan * &#x60;Coverity API&#x60; - Coverity API * &#x60;Crashtest Security JSON File&#x60; - Crashtest Security JSON File * &#x60;Crashtest Security XML File&#x60; - Crashtest Security XML File * &#x60;CredScan Scan&#x60; - CredScan Scan * &#x60;CycloneDX Scan&#x60; - CycloneDX Scan * &#x60;DawnScanner Scan&#x60; - DawnScanner Scan * &#x60;Dependency Check Scan&#x60; - Dependency Check Scan * &#x60;Dependency Track Finding Packaging Format (FPF) Export&#x60; - Dependency Track Finding Packaging Format (FPF) Export * &#x60;Detect-secrets Scan&#x60; - Detect-secrets Scan * &#x60;docker-bench-security Scan&#x60; - docker-bench-security Scan * &#x60;Dockle Scan&#x60; - Dockle Scan * &#x60;DrHeader JSON Importer&#x60; - DrHeader JSON Importer * &#x60;DSOP Scan&#x60; - DSOP Scan * &#x60;Edgescan Scan&#x60; - Edgescan Scan * &#x60;ESLint Scan&#x60; - ESLint Scan * &#x60;Fortify Scan&#x60; - Fortify Scan * &#x60;Generic Findings Import&#x60; - Generic Findings Import * &#x60;Ggshield Scan&#x60; - Ggshield Scan * &#x60;Github Vulnerability Scan&#x60; - Github Vulnerability Scan * &#x60;GitLab API Fuzzing Report Scan&#x60; - GitLab API Fuzzing Report Scan * &#x60;GitLab Container Scan&#x60; - GitLab Container Scan * &#x60;GitLab DAST Report&#x60; - GitLab DAST Report * &#x60;GitLab Dependency Scanning Report&#x60; - GitLab Dependency Scanning Report * &#x60;GitLab SAST Report&#x60; - GitLab SAST Report * &#x60;GitLab Secret Detection Report&#x60; - GitLab Secret Detection Report * &#x60;Gitleaks Scan&#x60; - Gitleaks Scan * &#x60;Gosec Scanner&#x60; - Gosec Scanner * &#x60;Govulncheck Scanner&#x60; - Govulncheck Scanner * &#x60;HackerOne Cases&#x60; - HackerOne Cases * &#x60;Hadolint Dockerfile check&#x60; - Hadolint Dockerfile check * &#x60;Harbor Vulnerability Scan&#x60; - Harbor Vulnerability Scan * &#x60;HCLAppScan XML&#x60; - HCLAppScan XML * &#x60;Horusec Scan&#x60; - Horusec Scan * &#x60;Humble Json Importer&#x60; - Humble Json Importer * &#x60;HuskyCI Report&#x60; - HuskyCI Report * &#x60;Hydra Scan&#x60; - Hydra Scan * &#x60;IBM AppScan DAST&#x60; - IBM AppScan DAST * &#x60;Immuniweb Scan&#x60; - Immuniweb Scan * &#x60;IntSights Report&#x60; - IntSights Report * &#x60;JFrog Xray API Summary Artifact Scan&#x60; - JFrog Xray API Summary Artifact Scan * &#x60;JFrog Xray On Demand Binary Scan&#x60; - JFrog Xray On Demand Binary Scan * &#x60;JFrog Xray Scan&#x60; - JFrog Xray Scan * &#x60;JFrog Xray Unified Scan&#x60; - JFrog Xray Unified Scan * &#x60;KICS Scan&#x60; - KICS Scan * &#x60;Kiuwan Scan&#x60; - Kiuwan Scan * &#x60;kube-bench Scan&#x60; - kube-bench Scan * &#x60;KubeHunter Scan&#x60; - KubeHunter Scan * &#x60;Meterian Scan&#x60; - Meterian Scan * &#x60;Microfocus Webinspect Scan&#x60; - Microfocus Webinspect Scan * &#x60;MobSF Scan&#x60; - MobSF Scan * &#x60;Mobsfscan Scan&#x60; - Mobsfscan Scan * &#x60;Mozilla Observatory Scan&#x60; - Mozilla Observatory Scan * &#x60;MSDefender Parser&#x60; - MSDefender Parser * &#x60;Netsparker Scan&#x60; - Netsparker Scan * &#x60;NeuVector (compliance)&#x60; - NeuVector (compliance) * &#x60;NeuVector (REST)&#x60; - NeuVector (REST) * &#x60;Nexpose Scan&#x60; - Nexpose Scan * &#x60;Nikto Scan&#x60; - Nikto Scan * &#x60;Nmap Scan&#x60; - Nmap Scan * &#x60;Node Security Platform Scan&#x60; - Node Security Platform Scan * &#x60;NPM Audit Scan&#x60; - NPM Audit Scan * &#x60;Nuclei Scan&#x60; - Nuclei Scan * &#x60;Openscap Vulnerability Scan&#x60; - Openscap Vulnerability Scan * &#x60;OpenVAS CSV&#x60; - OpenVAS CSV * &#x60;OpenVAS XML&#x60; - OpenVAS XML * &#x60;ORT evaluated model Importer&#x60; - ORT evaluated model Importer * &#x60;OssIndex Devaudit SCA Scan Importer&#x60; - OssIndex Devaudit SCA Scan Importer * &#x60;Outpost24 Scan&#x60; - Outpost24 Scan * &#x60;PHP Security Audit v2&#x60; - PHP Security Audit v2 * &#x60;PHP Symfony Security Check&#x60; - PHP Symfony Security Check * &#x60;pip-audit Scan&#x60; - pip-audit Scan * &#x60;PMD Scan&#x60; - PMD Scan * &#x60;Popeye Scan&#x60; - Popeye Scan * &#x60;PWN SAST&#x60; - PWN SAST * &#x60;Qualys Infrastructure Scan (WebGUI XML)&#x60; - Qualys Infrastructure Scan (WebGUI XML) * &#x60;Qualys Scan&#x60; - Qualys Scan * &#x60;Qualys Webapp Scan&#x60; - Qualys Webapp Scan * &#x60;Retire.js Scan&#x60; - Retire.js Scan * &#x60;Risk Recon API Importer&#x60; - Risk Recon API Importer * &#x60;Rubocop Scan&#x60; - Rubocop Scan * &#x60;Rusty Hog Scan&#x60; - Rusty Hog Scan * &#x60;SARIF&#x60; - SARIF * &#x60;Scantist Scan&#x60; - Scantist Scan * &#x60;Scout Suite Scan&#x60; - Scout Suite Scan * &#x60;Semgrep JSON Report&#x60; - Semgrep JSON Report * &#x60;SKF Scan&#x60; - SKF Scan * &#x60;Snyk Scan&#x60; - Snyk Scan * &#x60;Solar Appscreener Scan&#x60; - Solar Appscreener Scan * &#x60;SonarQube API Import&#x60; - SonarQube API Import * &#x60;SonarQube Scan&#x60; - SonarQube Scan * &#x60;SonarQube Scan detailed&#x60; - SonarQube Scan detailed * &#x60;Sonatype Application Scan&#x60; - Sonatype Application Scan * &#x60;SpotBugs Scan&#x60; - SpotBugs Scan * &#x60;SSH Audit Importer&#x60; - SSH Audit Importer * &#x60;SSL Labs Scan&#x60; - SSL Labs Scan * &#x60;Sslscan&#x60; - Sslscan * &#x60;Sslyze Scan&#x60; - Sslyze Scan * &#x60;SSLyze Scan (JSON)&#x60; - SSLyze Scan (JSON) * &#x60;StackHawk HawkScan&#x60; - StackHawk HawkScan * &#x60;Sysdig Vulnerability Report - Pipeline, Registry and Runtime (CSV)&#x60; - Sysdig Vulnerability Report - Pipeline, Registry and Runtime (CSV) * &#x60;Talisman Scan&#x60; - Talisman Scan * &#x60;Tenable Scan&#x60; - Tenable Scan * &#x60;Terrascan Scan&#x60; - Terrascan Scan * &#x60;Testssl Scan&#x60; - Testssl Scan * &#x60;TFSec Scan&#x60; - TFSec Scan * &#x60;Threagile risks report&#x60; - Threagile risks report * &#x60;Trivy Operator Scan&#x60; - Trivy Operator Scan * &#x60;Trivy Scan&#x60; - Trivy Scan * &#x60;Trufflehog Scan&#x60; - Trufflehog Scan * &#x60;Trufflehog3 Scan&#x60; - Trufflehog3 Scan * &#x60;Trustwave Fusion API Scan&#x60; - Trustwave Fusion API Scan * &#x60;Trustwave Scan (CSV)&#x60; - Trustwave Scan (CSV) * &#x60;Twistlock Image Scan&#x60; - Twistlock Image Scan * &#x60;VCG Scan&#x60; - VCG Scan * &#x60;Veracode Scan&#x60; - Veracode Scan * &#x60;Veracode SourceClear Scan&#x60; - Veracode SourceClear Scan * &#x60;Vulners&#x60; - Vulners * &#x60;Wapiti Scan&#x60; - Wapiti Scan * &#x60;Wazuh&#x60; - Wazuh * &#x60;WFuzz JSON report&#x60; - WFuzz JSON report * &#x60;Whispers Scan&#x60; - Whispers Scan * &#x60;WhiteHat Sentinel&#x60; - WhiteHat Sentinel * &#x60;Whitesource Scan&#x60; - Whitesource Scan * &#x60;Wpscan&#x60; - Wpscan * &#x60;Xanitizer Scan&#x60; - Xanitizer Scan * &#x60;Yarn Audit Scan&#x60; - Yarn Audit Scan * &#x60;ZAP Scan&#x60; - ZAP Scan | 
 **scanDate** | **string** | Scan completion date will be used on all findings. | 
 **minimumSeverity** | **string** | Minimum severity level to be imported  * &#x60;Info&#x60; - Info * &#x60;Low&#x60; - Low * &#x60;Medium&#x60; - Medium * &#x60;High&#x60; - High * &#x60;Critical&#x60; - Critical | [default to &quot;Info&quot;]
 **doNotReactivate** | **bool** | Select if the import should ignore active findings from the report, useful for triage-less scanners. Will keep existing findings closed, without reactivating them. For more information check the docs. | [default to false]
 **endpointToAdd** | **int32** |  | 
 **file** | ***os.File** |  | 
 **productTypeName** | **string** |  | 
 **productName** | **string** |  | 
 **engagementName** | **string** |  | 
 **engagementEndDate** | **string** | End Date for Engagement. Default is current time + 365 days. Required format year-month-day | 
 **sourceCodeManagementUri** | **string** | Resource link to source code | 
 **test** | **int32** |  | 
 **testTitle** | **string** |  | 
 **autoCreateContext** | **bool** |  | 
 **deduplicationOnEngagement** | **bool** |  | 
 **pushToJira** | **bool** |  | [default to false]
 **closeOldFindings** | **bool** | Select if old findings no longer present in the report get closed as mitigated when importing. | [default to true]
 **closeOldFindingsProductScope** | **bool** | Select if close_old_findings applies to all findings of the same type in the product. By default, it is false meaning that only old findings of the same type in the engagement are in scope. Note that this only applies on the first call to reimport-scan. | [default to false]
 **version** | **string** | Version that will be set on existing Test object. Leave empty to leave existing value in place. | 
 **buildId** | **string** | ID of the build that was scanned. | 
 **branchTag** | **string** | Branch or Tag that was scanned. | 
 **commitHash** | **string** | Commit that was scanned. | 
 **apiScanConfiguration** | **int32** |  | 
 **service** | **string** | A service is a self-contained piece of functionality within a Product. This is an optional field which is used in deduplication and closing of old findings when set. This affects the whole engagement/product depending on your deduplication scope. | 
 **environment** | **string** |  | 
 **lead** | **int32** |  | 
 **tags** | **[]string** | Modify existing tags that help describe this scan. (Existing test tags will be overwritten) | 
 **groupBy** | **string** | Choose an option to automatically group new findings by the chosen option.  * &#x60;component_name&#x60; - Component Name * &#x60;component_name+component_version&#x60; - Component Name + Version * &#x60;file_path&#x60; - File path * &#x60;finding_title&#x60; - Finding Title | 
 **createFindingGroupsForAllFindings** | **bool** | If set to false, finding groups will only be created when there is more than one grouped finding | [default to true]

### Return type

[**ReImportScan**](ReImportScan.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

