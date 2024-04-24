# \FindingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindingsAcceptRisksCreate**](FindingsAPI.md#FindingsAcceptRisksCreate) | **Post** /api/v2/findings/accept_risks/ | 
[**FindingsCloseCreate**](FindingsAPI.md#FindingsCloseCreate) | **Post** /api/v2/findings/{id}/close/ | 
[**FindingsCreate**](FindingsAPI.md#FindingsCreate) | **Post** /api/v2/findings/ | 
[**FindingsDeletePreviewList**](FindingsAPI.md#FindingsDeletePreviewList) | **Get** /api/v2/findings/{id}/delete_preview/ | 
[**FindingsDestroy**](FindingsAPI.md#FindingsDestroy) | **Delete** /api/v2/findings/{id}/ | 
[**FindingsDuplicateList**](FindingsAPI.md#FindingsDuplicateList) | **Get** /api/v2/findings/{id}/duplicate/ | 
[**FindingsDuplicateResetCreate**](FindingsAPI.md#FindingsDuplicateResetCreate) | **Post** /api/v2/findings/{id}/duplicate/reset/ | 
[**FindingsFilesCreate**](FindingsAPI.md#FindingsFilesCreate) | **Post** /api/v2/findings/{id}/files/ | 
[**FindingsFilesDownloadRetrieve**](FindingsAPI.md#FindingsFilesDownloadRetrieve) | **Get** /api/v2/findings/{id}/files/download/{file_id}/ | 
[**FindingsFilesRetrieve**](FindingsAPI.md#FindingsFilesRetrieve) | **Get** /api/v2/findings/{id}/files/ | 
[**FindingsGenerateReportCreate**](FindingsAPI.md#FindingsGenerateReportCreate) | **Post** /api/v2/findings/generate_report/ | 
[**FindingsList**](FindingsAPI.md#FindingsList) | **Get** /api/v2/findings/ | 
[**FindingsMetadataCreate**](FindingsAPI.md#FindingsMetadataCreate) | **Post** /api/v2/findings/{id}/metadata/ | 
[**FindingsMetadataDestroy**](FindingsAPI.md#FindingsMetadataDestroy) | **Delete** /api/v2/findings/{id}/metadata/ | 
[**FindingsMetadataList**](FindingsAPI.md#FindingsMetadataList) | **Get** /api/v2/findings/{id}/metadata/ | 
[**FindingsMetadataUpdate**](FindingsAPI.md#FindingsMetadataUpdate) | **Put** /api/v2/findings/{id}/metadata/ | 
[**FindingsNotesCreate**](FindingsAPI.md#FindingsNotesCreate) | **Post** /api/v2/findings/{id}/notes/ | 
[**FindingsNotesRetrieve**](FindingsAPI.md#FindingsNotesRetrieve) | **Get** /api/v2/findings/{id}/notes/ | 
[**FindingsOriginalCreate**](FindingsAPI.md#FindingsOriginalCreate) | **Post** /api/v2/findings/{id}/original/{new_fid}/ | 
[**FindingsPartialUpdate**](FindingsAPI.md#FindingsPartialUpdate) | **Patch** /api/v2/findings/{id}/ | 
[**FindingsRemoveNotePartialUpdate**](FindingsAPI.md#FindingsRemoveNotePartialUpdate) | **Patch** /api/v2/findings/{id}/remove_note/ | 
[**FindingsRemoveTagsPartialUpdate**](FindingsAPI.md#FindingsRemoveTagsPartialUpdate) | **Patch** /api/v2/findings/{id}/remove_tags/ | 
[**FindingsRemoveTagsUpdate**](FindingsAPI.md#FindingsRemoveTagsUpdate) | **Put** /api/v2/findings/{id}/remove_tags/ | 
[**FindingsRequestResponseCreate**](FindingsAPI.md#FindingsRequestResponseCreate) | **Post** /api/v2/findings/{id}/request_response/ | 
[**FindingsRequestResponseRetrieve**](FindingsAPI.md#FindingsRequestResponseRetrieve) | **Get** /api/v2/findings/{id}/request_response/ | 
[**FindingsRetrieve**](FindingsAPI.md#FindingsRetrieve) | **Get** /api/v2/findings/{id}/ | 
[**FindingsTagsCreate**](FindingsAPI.md#FindingsTagsCreate) | **Post** /api/v2/findings/{id}/tags/ | 
[**FindingsTagsRetrieve**](FindingsAPI.md#FindingsTagsRetrieve) | **Get** /api/v2/findings/{id}/tags/ | 
[**FindingsUpdate**](FindingsAPI.md#FindingsUpdate) | **Put** /api/v2/findings/{id}/ | 



## FindingsAcceptRisksCreate

> PaginatedRiskAcceptanceList FindingsAcceptRisksCreate(ctx).AcceptedRiskRequest(acceptedRiskRequest).Active(active).After(after).Before(before).ComponentName(componentName).ComponentVersion(componentVersion).Created(created).Cvssv3(cvssv3).Cvssv3Score(cvssv3Score).Cwe(cwe).Date(date).DefectReviewRequestedBy(defectReviewRequestedBy).Description(description).Duplicate(duplicate).DuplicateFinding(duplicateFinding).DynamicFinding(dynamicFinding).EffortForFixing(effortForFixing).Endpoints(endpoints).EpssPercentile(epssPercentile).EpssScore(epssScore).FalseP(falseP).FilePath(filePath).FindingGroup(findingGroup).FoundBy(foundBy).HasJira(hasJira).HasTags(hasTags).HashCode(hashCode).Id(id).Impact(impact).InheritedTags(inheritedTags).IsMitigated(isMitigated).JiraChange(jiraChange).JiraCreation(jiraCreation).LastReviewed(lastReviewed).LastReviewedBy(lastReviewedBy).LastStatusUpdate(lastStatusUpdate).Limit(limit).Mitigated(mitigated).MitigatedBy(mitigatedBy).Mitigation(mitigation).NbOccurences(nbOccurences).NotTag(notTag).NotTags(notTags).NotTestEngagementProductTags(notTestEngagementProductTags).NotTestEngagementTags(notTestEngagementTags).NotTestTags(notTestTags).NumericalSeverity(numericalSeverity).O(o).Offset(offset).On(on).OutOfScope(outOfScope).OutsideOfSla(outsideOfSla).Param(param).Payload(payload).PlannedRemediationDate(plannedRemediationDate).PlannedRemediationVersion(plannedRemediationVersion).ProductLifecycle(productLifecycle).ProductName(productName).ProductNameContains(productNameContains).PublishDate(publishDate).References(references).Reporter(reporter).ReviewRequestedBy(reviewRequestedBy).Reviewers(reviewers).RiskAcceptance(riskAcceptance).RiskAccepted(riskAccepted).SastSinkObject(sastSinkObject).SastSourceFilePath(sastSourceFilePath).SastSourceLine(sastSourceLine).SastSourceObject(sastSourceObject).ScannerConfidence(scannerConfidence).Service(service).Severity(severity).SeverityJustification(severityJustification).SlaExpirationDate(slaExpirationDate).SlaStartDate(slaStartDate).SonarqubeIssue(sonarqubeIssue).StaticFinding(staticFinding).StepsToReproduce(stepsToReproduce).Tag(tag).Tags(tags).Test(test).TestEngagement(testEngagement).TestEngagementProduct(testEngagementProduct).TestEngagementProductProdType(testEngagementProductProdType).TestEngagementProductTags(testEngagementProductTags).TestEngagementTags(testEngagementTags).TestTags(testTags).TestTestType(testTestType).Title(title).UnderDefectReview(underDefectReview).UnderReview(underReview).UniqueIdFromTool(uniqueIdFromTool).Verified(verified).VulnIdFromTool(vulnIdFromTool).VulnerabilityId(vulnerabilityId).Execute()



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
	acceptedRiskRequest := []openapiclient.AcceptedRiskRequest{*openapiclient.NewAcceptedRiskRequest("VulnerabilityId_example", "Justification_example", "AcceptedBy_example")} // []AcceptedRiskRequest | 
	active := true // bool |  (optional)
	after := time.Now() // string |  (optional)
	before := time.Now() // string |  (optional)
	componentName := "componentName_example" // string |  (optional)
	componentVersion := "componentVersion_example" // string |  (optional)
	created := time.Now() // time.Time | The date the finding was created inside DefectDojo.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	cvssv3 := "cvssv3_example" // string |  (optional)
	cvssv3Score := float32(3.4) // float32 |  (optional)
	cwe := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	date := time.Now() // string | The date the flaw was discovered.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	defectReviewRequestedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	description := "description_example" // string |  (optional)
	duplicate := true // bool |  (optional)
	duplicateFinding := int32(56) // int32 |  (optional)
	dynamicFinding := true // bool |  (optional)
	effortForFixing := "effortForFixing_example" // string |  (optional)
	endpoints := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	epssPercentile := float32(3.4) // float32 |  (optional)
	epssScore := float32(3.4) // float32 |  (optional)
	falseP := true // bool |  (optional)
	filePath := "filePath_example" // string |  (optional)
	findingGroup := []float32{float32(123)} // []float32 | Multiple values may be separated by commas. (optional)
	foundBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	hasJira := true // bool |  (optional)
	hasTags := true // bool | Has tags (optional)
	hashCode := "hashCode_example" // string |  (optional)
	id := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	impact := "impact_example" // string |  (optional)
	inheritedTags := [][]int32{[]int32{int32(123)}} // [][]int32 | Internal use tags sepcifically for maintaining parity with product. This field will be present as a subset in the tags field (optional)
	isMitigated := true // bool |  (optional)
	jiraChange := time.Now() // time.Time | The date the linked Jira issue was last modified.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	jiraCreation := time.Now() // time.Time | The date a Jira issue was created from this finding.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	lastReviewed := time.Now() // time.Time | Provides the date the flaw was last 'touched' by a tester.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	lastReviewedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	lastStatusUpdate := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	mitigated := time.Now() // time.Time | Denotes if this flaw has been fixed by storing the date it was fixed.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	mitigatedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	mitigation := "mitigation_example" // string |  (optional)
	nbOccurences := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	notTestEngagementProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on product (optional)
	notTestEngagementTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on engagement (optional)
	notTestTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on test (optional)
	numericalSeverity := "numericalSeverity_example" // string |  (optional)
	o := []string{"O_example"} // []string | Ordering  * `active` - Active * `-active` - Active (descending) * `component_name` - Component name * `-component_name` - Component name (descending) * `component_version` - Component version * `-component_version` - Component version (descending) * `created` - Created * `-created` - Created (descending) * `last_status_update` - Last status update * `-last_status_update` - Last status update (descending) * `last_reviewed` - Last reviewed * `-last_reviewed` - Last reviewed (descending) * `cwe` - Cwe * `-cwe` - Cwe (descending) * `date` - Date * `-date` - Date (descending) * `duplicate` - Duplicate * `-duplicate` - Duplicate (descending) * `dynamic_finding` - Dynamic finding * `-dynamic_finding` - Dynamic finding (descending) * `false_p` - False p * `-false_p` - False p (descending) * `found_by` - Found by * `-found_by` - Found by (descending) * `id` - Id * `-id` - Id (descending) * `is_mitigated` - Is mitigated * `-is_mitigated` - Is mitigated (descending) * `numerical_severity` - Numerical severity * `-numerical_severity` - Numerical severity (descending) * `out_of_scope` - Out of scope * `-out_of_scope` - Out of scope (descending) * `severity` - Severity * `-severity` - Severity (descending) * `reviewers` - Reviewers * `-reviewers` - Reviewers (descending) * `static_finding` - Static finding * `-static_finding` - Static finding (descending) * `test__engagement__product__name` - Test  engagement  product  name * `-test__engagement__product__name` - Test  engagement  product  name (descending) * `title` - Title * `-title` - Title (descending) * `under_defect_review` - Under defect review * `-under_defect_review` - Under defect review (descending) * `under_review` - Under review * `-under_review` - Under review (descending) * `verified` - Verified * `-verified` - Verified (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	on := time.Now() // string |  (optional)
	outOfScope := true // bool |  (optional)
	outsideOfSla := float32(8.14) // float32 |  (optional)
	param := "param_example" // string |  (optional)
	payload := "payload_example" // string |  (optional)
	plannedRemediationDate := time.Now() // string |  (optional)
	plannedRemediationVersion := "plannedRemediationVersion_example" // string |  (optional)
	productLifecycle := "productLifecycle_example" // string | Comma separated list of exact product lifecycles (optional)
	productName := "productName_example" // string | exact product name (optional)
	productNameContains := "productNameContains_example" // string | exact product name (optional)
	publishDate := time.Now() // string |  (optional)
	references := "references_example" // string |  (optional)
	reporter := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	reviewRequestedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	reviewers := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	riskAcceptance := float32(8.14) // float32 |  (optional)
	riskAccepted := true // bool |  (optional)
	sastSinkObject := "sastSinkObject_example" // string |  (optional)
	sastSourceFilePath := "sastSourceFilePath_example" // string |  (optional)
	sastSourceLine := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	sastSourceObject := "sastSourceObject_example" // string |  (optional)
	scannerConfidence := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	service := "service_example" // string |  (optional)
	severity := "severity_example" // string |  (optional)
	severityJustification := "severityJustification_example" // string |  (optional)
	slaExpirationDate := time.Now() // string |  (optional)
	slaStartDate := time.Now() // string |  (optional)
	sonarqubeIssue := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	staticFinding := true // bool |  (optional)
	stepsToReproduce := "stepsToReproduce_example" // string |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	test := int32(56) // int32 |  (optional)
	testEngagement := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	testEngagementProduct := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	testEngagementProductProdType := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	testEngagementProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on product (optional)
	testEngagementTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on engagement (optional)
	testTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on test (optional)
	testTestType := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	title := "title_example" // string |  (optional)
	underDefectReview := true // bool |  (optional)
	underReview := true // bool |  (optional)
	uniqueIdFromTool := "uniqueIdFromTool_example" // string |  (optional)
	verified := true // bool |  (optional)
	vulnIdFromTool := "vulnIdFromTool_example" // string |  (optional)
	vulnerabilityId := "vulnerabilityId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsAcceptRisksCreate(context.Background()).AcceptedRiskRequest(acceptedRiskRequest).Active(active).After(after).Before(before).ComponentName(componentName).ComponentVersion(componentVersion).Created(created).Cvssv3(cvssv3).Cvssv3Score(cvssv3Score).Cwe(cwe).Date(date).DefectReviewRequestedBy(defectReviewRequestedBy).Description(description).Duplicate(duplicate).DuplicateFinding(duplicateFinding).DynamicFinding(dynamicFinding).EffortForFixing(effortForFixing).Endpoints(endpoints).EpssPercentile(epssPercentile).EpssScore(epssScore).FalseP(falseP).FilePath(filePath).FindingGroup(findingGroup).FoundBy(foundBy).HasJira(hasJira).HasTags(hasTags).HashCode(hashCode).Id(id).Impact(impact).InheritedTags(inheritedTags).IsMitigated(isMitigated).JiraChange(jiraChange).JiraCreation(jiraCreation).LastReviewed(lastReviewed).LastReviewedBy(lastReviewedBy).LastStatusUpdate(lastStatusUpdate).Limit(limit).Mitigated(mitigated).MitigatedBy(mitigatedBy).Mitigation(mitigation).NbOccurences(nbOccurences).NotTag(notTag).NotTags(notTags).NotTestEngagementProductTags(notTestEngagementProductTags).NotTestEngagementTags(notTestEngagementTags).NotTestTags(notTestTags).NumericalSeverity(numericalSeverity).O(o).Offset(offset).On(on).OutOfScope(outOfScope).OutsideOfSla(outsideOfSla).Param(param).Payload(payload).PlannedRemediationDate(plannedRemediationDate).PlannedRemediationVersion(plannedRemediationVersion).ProductLifecycle(productLifecycle).ProductName(productName).ProductNameContains(productNameContains).PublishDate(publishDate).References(references).Reporter(reporter).ReviewRequestedBy(reviewRequestedBy).Reviewers(reviewers).RiskAcceptance(riskAcceptance).RiskAccepted(riskAccepted).SastSinkObject(sastSinkObject).SastSourceFilePath(sastSourceFilePath).SastSourceLine(sastSourceLine).SastSourceObject(sastSourceObject).ScannerConfidence(scannerConfidence).Service(service).Severity(severity).SeverityJustification(severityJustification).SlaExpirationDate(slaExpirationDate).SlaStartDate(slaStartDate).SonarqubeIssue(sonarqubeIssue).StaticFinding(staticFinding).StepsToReproduce(stepsToReproduce).Tag(tag).Tags(tags).Test(test).TestEngagement(testEngagement).TestEngagementProduct(testEngagementProduct).TestEngagementProductProdType(testEngagementProductProdType).TestEngagementProductTags(testEngagementProductTags).TestEngagementTags(testEngagementTags).TestTags(testTags).TestTestType(testTestType).Title(title).UnderDefectReview(underDefectReview).UnderReview(underReview).UniqueIdFromTool(uniqueIdFromTool).Verified(verified).VulnIdFromTool(vulnIdFromTool).VulnerabilityId(vulnerabilityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsAcceptRisksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsAcceptRisksCreate`: PaginatedRiskAcceptanceList
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsAcceptRisksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindingsAcceptRisksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptedRiskRequest** | [**[]AcceptedRiskRequest**](AcceptedRiskRequest.md) |  | 
 **active** | **bool** |  | 
 **after** | **string** |  | 
 **before** | **string** |  | 
 **componentName** | **string** |  | 
 **componentVersion** | **string** |  | 
 **created** | **time.Time** | The date the finding was created inside DefectDojo.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **cvssv3** | **string** |  | 
 **cvssv3Score** | **float32** |  | 
 **cwe** | **[]int32** | Multiple values may be separated by commas. | 
 **date** | **string** | The date the flaw was discovered.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **defectReviewRequestedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **description** | **string** |  | 
 **duplicate** | **bool** |  | 
 **duplicateFinding** | **int32** |  | 
 **dynamicFinding** | **bool** |  | 
 **effortForFixing** | **string** |  | 
 **endpoints** | **[]int32** | Multiple values may be separated by commas. | 
 **epssPercentile** | **float32** |  | 
 **epssScore** | **float32** |  | 
 **falseP** | **bool** |  | 
 **filePath** | **string** |  | 
 **findingGroup** | **[]float32** | Multiple values may be separated by commas. | 
 **foundBy** | **[]int32** | Multiple values may be separated by commas. | 
 **hasJira** | **bool** |  | 
 **hasTags** | **bool** | Has tags | 
 **hashCode** | **string** |  | 
 **id** | **[]int32** | Multiple values may be separated by commas. | 
 **impact** | **string** |  | 
 **inheritedTags** | **[][]int32** | Internal use tags sepcifically for maintaining parity with product. This field will be present as a subset in the tags field | 
 **isMitigated** | **bool** |  | 
 **jiraChange** | **time.Time** | The date the linked Jira issue was last modified.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **jiraCreation** | **time.Time** | The date a Jira issue was created from this finding.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **lastReviewed** | **time.Time** | Provides the date the flaw was last &#39;touched&#39; by a tester.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **lastReviewedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **lastStatusUpdate** | **time.Time** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **mitigated** | **time.Time** | Denotes if this flaw has been fixed by storing the date it was fixed.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **mitigatedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **mitigation** | **string** |  | 
 **nbOccurences** | **[]int32** | Multiple values may be separated by commas. | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **notTestEngagementProductTags** | **[]string** | Comma separated list of exact tags not present on product | 
 **notTestEngagementTags** | **[]string** | Comma separated list of exact tags not present on engagement | 
 **notTestTags** | **[]string** | Comma separated list of exact tags present on test | 
 **numericalSeverity** | **string** |  | 
 **o** | **[]string** | Ordering  * &#x60;active&#x60; - Active * &#x60;-active&#x60; - Active (descending) * &#x60;component_name&#x60; - Component name * &#x60;-component_name&#x60; - Component name (descending) * &#x60;component_version&#x60; - Component version * &#x60;-component_version&#x60; - Component version (descending) * &#x60;created&#x60; - Created * &#x60;-created&#x60; - Created (descending) * &#x60;last_status_update&#x60; - Last status update * &#x60;-last_status_update&#x60; - Last status update (descending) * &#x60;last_reviewed&#x60; - Last reviewed * &#x60;-last_reviewed&#x60; - Last reviewed (descending) * &#x60;cwe&#x60; - Cwe * &#x60;-cwe&#x60; - Cwe (descending) * &#x60;date&#x60; - Date * &#x60;-date&#x60; - Date (descending) * &#x60;duplicate&#x60; - Duplicate * &#x60;-duplicate&#x60; - Duplicate (descending) * &#x60;dynamic_finding&#x60; - Dynamic finding * &#x60;-dynamic_finding&#x60; - Dynamic finding (descending) * &#x60;false_p&#x60; - False p * &#x60;-false_p&#x60; - False p (descending) * &#x60;found_by&#x60; - Found by * &#x60;-found_by&#x60; - Found by (descending) * &#x60;id&#x60; - Id * &#x60;-id&#x60; - Id (descending) * &#x60;is_mitigated&#x60; - Is mitigated * &#x60;-is_mitigated&#x60; - Is mitigated (descending) * &#x60;numerical_severity&#x60; - Numerical severity * &#x60;-numerical_severity&#x60; - Numerical severity (descending) * &#x60;out_of_scope&#x60; - Out of scope * &#x60;-out_of_scope&#x60; - Out of scope (descending) * &#x60;severity&#x60; - Severity * &#x60;-severity&#x60; - Severity (descending) * &#x60;reviewers&#x60; - Reviewers * &#x60;-reviewers&#x60; - Reviewers (descending) * &#x60;static_finding&#x60; - Static finding * &#x60;-static_finding&#x60; - Static finding (descending) * &#x60;test__engagement__product__name&#x60; - Test  engagement  product  name * &#x60;-test__engagement__product__name&#x60; - Test  engagement  product  name (descending) * &#x60;title&#x60; - Title * &#x60;-title&#x60; - Title (descending) * &#x60;under_defect_review&#x60; - Under defect review * &#x60;-under_defect_review&#x60; - Under defect review (descending) * &#x60;under_review&#x60; - Under review * &#x60;-under_review&#x60; - Under review (descending) * &#x60;verified&#x60; - Verified * &#x60;-verified&#x60; - Verified (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **on** | **string** |  | 
 **outOfScope** | **bool** |  | 
 **outsideOfSla** | **float32** |  | 
 **param** | **string** |  | 
 **payload** | **string** |  | 
 **plannedRemediationDate** | **string** |  | 
 **plannedRemediationVersion** | **string** |  | 
 **productLifecycle** | **string** | Comma separated list of exact product lifecycles | 
 **productName** | **string** | exact product name | 
 **productNameContains** | **string** | exact product name | 
 **publishDate** | **string** |  | 
 **references** | **string** |  | 
 **reporter** | **[]int32** | Multiple values may be separated by commas. | 
 **reviewRequestedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **reviewers** | **[]int32** | Multiple values may be separated by commas. | 
 **riskAcceptance** | **float32** |  | 
 **riskAccepted** | **bool** |  | 
 **sastSinkObject** | **string** |  | 
 **sastSourceFilePath** | **string** |  | 
 **sastSourceLine** | **[]int32** | Multiple values may be separated by commas. | 
 **sastSourceObject** | **string** |  | 
 **scannerConfidence** | **[]int32** | Multiple values may be separated by commas. | 
 **service** | **string** |  | 
 **severity** | **string** |  | 
 **severityJustification** | **string** |  | 
 **slaExpirationDate** | **string** |  | 
 **slaStartDate** | **string** |  | 
 **sonarqubeIssue** | **[]int32** | Multiple values may be separated by commas. | 
 **staticFinding** | **bool** |  | 
 **stepsToReproduce** | **string** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **test** | **int32** |  | 
 **testEngagement** | **[]int32** | Multiple values may be separated by commas. | 
 **testEngagementProduct** | **[]int32** | Multiple values may be separated by commas. | 
 **testEngagementProductProdType** | **[]int32** | Multiple values may be separated by commas. | 
 **testEngagementProductTags** | **[]string** | Comma separated list of exact tags present on product | 
 **testEngagementTags** | **[]string** | Comma separated list of exact tags present on engagement | 
 **testTags** | **[]string** | Comma separated list of exact tags present on test | 
 **testTestType** | **[]int32** | Multiple values may be separated by commas. | 
 **title** | **string** |  | 
 **underDefectReview** | **bool** |  | 
 **underReview** | **bool** |  | 
 **uniqueIdFromTool** | **string** |  | 
 **verified** | **bool** |  | 
 **vulnIdFromTool** | **string** |  | 
 **vulnerabilityId** | **string** |  | 

### Return type

[**PaginatedRiskAcceptanceList**](PaginatedRiskAcceptanceList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsCloseCreate

> FindingClose FindingsCloseCreate(ctx, id).FindingCloseRequest(findingCloseRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	findingCloseRequest := *openapiclient.NewFindingCloseRequest() // FindingCloseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsCloseCreate(context.Background(), id).FindingCloseRequest(findingCloseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsCloseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsCloseCreate`: FindingClose
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsCloseCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsCloseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findingCloseRequest** | [**FindingCloseRequest**](FindingCloseRequest.md) |  | 

### Return type

[**FindingClose**](FindingClose.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsCreate

> FindingCreate FindingsCreate(ctx).FindingCreateRequest(findingCreateRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	findingCreateRequest := *openapiclient.NewFindingCreateRequest(int32(123), []int32{int32(123)}, "Title_example", "Severity_example", "Description_example", false, false, "NumericalSeverity_example") // FindingCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsCreate(context.Background()).FindingCreateRequest(findingCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsCreate`: FindingCreate
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findingCreateRequest** | [**FindingCreateRequest**](FindingCreateRequest.md) |  | 

### Return type

[**FindingCreate**](FindingCreate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsDeletePreviewList

> PaginatedDeletePreviewList FindingsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsDeletePreviewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDeletePreviewList**](PaginatedDeletePreviewList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsDestroy

> FindingsDestroy(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsDuplicateList

> []Finding FindingsDuplicateList(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsDuplicateList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsDuplicateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsDuplicateList`: []Finding
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsDuplicateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsDuplicateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Finding**](Finding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsDuplicateResetCreate

> FindingsDuplicateResetCreate(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsDuplicateResetCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsDuplicateResetCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsDuplicateResetCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsFilesCreate

> File FindingsFilesCreate(ctx, id).Title(title).File(file).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	title := "title_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsFilesCreate(context.Background(), id).Title(title).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsFilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsFilesCreate`: File
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsFilesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

[**File**](File.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsFilesDownloadRetrieve

> RawFile FindingsFilesDownloadRetrieve(ctx, fileId, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	fileId := "fileId_example" // string | 
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsFilesDownloadRetrieve(context.Background(), fileId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsFilesDownloadRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsFilesDownloadRetrieve`: RawFile
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsFilesDownloadRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsFilesDownloadRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RawFile**](RawFile.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsFilesRetrieve

> FindingToFiles FindingsFilesRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsFilesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsFilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsFilesRetrieve`: FindingToFiles
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsFilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsFilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindingToFiles**](FindingToFiles.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsGenerateReportCreate

> ReportGenerate FindingsGenerateReportCreate(ctx).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	reportGenerateOptionRequest := *openapiclient.NewReportGenerateOptionRequest() // ReportGenerateOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsGenerateReportCreate(context.Background()).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsGenerateReportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsGenerateReportCreate`: ReportGenerate
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsGenerateReportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindingsGenerateReportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportGenerateOptionRequest** | [**ReportGenerateOptionRequest**](ReportGenerateOptionRequest.md) |  | 

### Return type

[**ReportGenerate**](ReportGenerate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsList

> PaginatedFindingList FindingsList(ctx).Active(active).After(after).Before(before).ComponentName(componentName).ComponentVersion(componentVersion).Created(created).Cvssv3(cvssv3).Cvssv3Score(cvssv3Score).Cwe(cwe).Date(date).DefectReviewRequestedBy(defectReviewRequestedBy).Description(description).Duplicate(duplicate).DuplicateFinding(duplicateFinding).DynamicFinding(dynamicFinding).EffortForFixing(effortForFixing).Endpoints(endpoints).EpssPercentile(epssPercentile).EpssScore(epssScore).FalseP(falseP).FilePath(filePath).FindingGroup(findingGroup).FoundBy(foundBy).HasJira(hasJira).HasTags(hasTags).HashCode(hashCode).Id(id).Impact(impact).InheritedTags(inheritedTags).IsMitigated(isMitigated).JiraChange(jiraChange).JiraCreation(jiraCreation).LastReviewed(lastReviewed).LastReviewedBy(lastReviewedBy).LastStatusUpdate(lastStatusUpdate).Limit(limit).Mitigated(mitigated).MitigatedBy(mitigatedBy).Mitigation(mitigation).NbOccurences(nbOccurences).NotTag(notTag).NotTags(notTags).NotTestEngagementProductTags(notTestEngagementProductTags).NotTestEngagementTags(notTestEngagementTags).NotTestTags(notTestTags).NumericalSeverity(numericalSeverity).O(o).Offset(offset).On(on).OutOfScope(outOfScope).OutsideOfSla(outsideOfSla).Param(param).Payload(payload).PlannedRemediationDate(plannedRemediationDate).PlannedRemediationVersion(plannedRemediationVersion).Prefetch(prefetch).ProductLifecycle(productLifecycle).ProductName(productName).ProductNameContains(productNameContains).PublishDate(publishDate).References(references).RelatedFields(relatedFields).Reporter(reporter).ReviewRequestedBy(reviewRequestedBy).Reviewers(reviewers).RiskAcceptance(riskAcceptance).RiskAccepted(riskAccepted).SastSinkObject(sastSinkObject).SastSourceFilePath(sastSourceFilePath).SastSourceLine(sastSourceLine).SastSourceObject(sastSourceObject).ScannerConfidence(scannerConfidence).Service(service).Severity(severity).SeverityJustification(severityJustification).SlaExpirationDate(slaExpirationDate).SlaStartDate(slaStartDate).SonarqubeIssue(sonarqubeIssue).StaticFinding(staticFinding).StepsToReproduce(stepsToReproduce).Tag(tag).Tags(tags).Test(test).TestEngagement(testEngagement).TestEngagementProduct(testEngagementProduct).TestEngagementProductProdType(testEngagementProductProdType).TestEngagementProductTags(testEngagementProductTags).TestEngagementTags(testEngagementTags).TestTags(testTags).TestTestType(testTestType).Title(title).UnderDefectReview(underDefectReview).UnderReview(underReview).UniqueIdFromTool(uniqueIdFromTool).Verified(verified).VulnIdFromTool(vulnIdFromTool).VulnerabilityId(vulnerabilityId).Execute()



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
	active := true // bool |  (optional)
	after := time.Now() // string |  (optional)
	before := time.Now() // string |  (optional)
	componentName := "componentName_example" // string |  (optional)
	componentVersion := "componentVersion_example" // string |  (optional)
	created := time.Now() // time.Time | The date the finding was created inside DefectDojo.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	cvssv3 := "cvssv3_example" // string |  (optional)
	cvssv3Score := float32(3.4) // float32 |  (optional)
	cwe := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	date := time.Now() // string | The date the flaw was discovered.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	defectReviewRequestedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	description := "description_example" // string |  (optional)
	duplicate := true // bool |  (optional)
	duplicateFinding := int32(56) // int32 |  (optional)
	dynamicFinding := true // bool |  (optional)
	effortForFixing := "effortForFixing_example" // string |  (optional)
	endpoints := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	epssPercentile := float32(3.4) // float32 |  (optional)
	epssScore := float32(3.4) // float32 |  (optional)
	falseP := true // bool |  (optional)
	filePath := "filePath_example" // string |  (optional)
	findingGroup := []float32{float32(123)} // []float32 | Multiple values may be separated by commas. (optional)
	foundBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	hasJira := true // bool |  (optional)
	hasTags := true // bool | Has tags (optional)
	hashCode := "hashCode_example" // string |  (optional)
	id := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	impact := "impact_example" // string |  (optional)
	inheritedTags := [][]int32{[]int32{int32(123)}} // [][]int32 | Internal use tags sepcifically for maintaining parity with product. This field will be present as a subset in the tags field (optional)
	isMitigated := true // bool |  (optional)
	jiraChange := time.Now() // time.Time | The date the linked Jira issue was last modified.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	jiraCreation := time.Now() // time.Time | The date a Jira issue was created from this finding.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	lastReviewed := time.Now() // time.Time | Provides the date the flaw was last 'touched' by a tester.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	lastReviewedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	lastStatusUpdate := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	mitigated := time.Now() // time.Time | Denotes if this flaw has been fixed by storing the date it was fixed.  * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	mitigatedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	mitigation := "mitigation_example" // string |  (optional)
	nbOccurences := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	notTestEngagementProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on product (optional)
	notTestEngagementTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on engagement (optional)
	notTestTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on test (optional)
	numericalSeverity := "numericalSeverity_example" // string |  (optional)
	o := []string{"O_example"} // []string | Ordering  * `active` - Active * `-active` - Active (descending) * `component_name` - Component name * `-component_name` - Component name (descending) * `component_version` - Component version * `-component_version` - Component version (descending) * `created` - Created * `-created` - Created (descending) * `last_status_update` - Last status update * `-last_status_update` - Last status update (descending) * `last_reviewed` - Last reviewed * `-last_reviewed` - Last reviewed (descending) * `cwe` - Cwe * `-cwe` - Cwe (descending) * `date` - Date * `-date` - Date (descending) * `duplicate` - Duplicate * `-duplicate` - Duplicate (descending) * `dynamic_finding` - Dynamic finding * `-dynamic_finding` - Dynamic finding (descending) * `false_p` - False p * `-false_p` - False p (descending) * `found_by` - Found by * `-found_by` - Found by (descending) * `id` - Id * `-id` - Id (descending) * `is_mitigated` - Is mitigated * `-is_mitigated` - Is mitigated (descending) * `numerical_severity` - Numerical severity * `-numerical_severity` - Numerical severity (descending) * `out_of_scope` - Out of scope * `-out_of_scope` - Out of scope (descending) * `severity` - Severity * `-severity` - Severity (descending) * `reviewers` - Reviewers * `-reviewers` - Reviewers (descending) * `static_finding` - Static finding * `-static_finding` - Static finding (descending) * `test__engagement__product__name` - Test  engagement  product  name * `-test__engagement__product__name` - Test  engagement  product  name (descending) * `title` - Title * `-title` - Title (descending) * `under_defect_review` - Under defect review * `-under_defect_review` - Under defect review (descending) * `under_review` - Under review * `-under_review` - Under review (descending) * `verified` - Verified * `-verified` - Verified (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	on := time.Now() // string |  (optional)
	outOfScope := true // bool |  (optional)
	outsideOfSla := float32(8.14) // float32 |  (optional)
	param := "param_example" // string |  (optional)
	payload := "payload_example" // string |  (optional)
	plannedRemediationDate := time.Now() // string |  (optional)
	plannedRemediationVersion := "plannedRemediationVersion_example" // string |  (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	productLifecycle := "productLifecycle_example" // string | Comma separated list of exact product lifecycles (optional)
	productName := "productName_example" // string | exact product name (optional)
	productNameContains := "productNameContains_example" // string | exact product name (optional)
	publishDate := time.Now() // string |  (optional)
	references := "references_example" // string |  (optional)
	relatedFields := true // bool | Expand finding external relations (engagement, environment, product,                                             product_type, test, test_type) (optional)
	reporter := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	reviewRequestedBy := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	reviewers := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	riskAcceptance := float32(8.14) // float32 |  (optional)
	riskAccepted := true // bool |  (optional)
	sastSinkObject := "sastSinkObject_example" // string |  (optional)
	sastSourceFilePath := "sastSourceFilePath_example" // string |  (optional)
	sastSourceLine := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	sastSourceObject := "sastSourceObject_example" // string |  (optional)
	scannerConfidence := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	service := "service_example" // string |  (optional)
	severity := "severity_example" // string |  (optional)
	severityJustification := "severityJustification_example" // string |  (optional)
	slaExpirationDate := time.Now() // string |  (optional)
	slaStartDate := time.Now() // string |  (optional)
	sonarqubeIssue := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	staticFinding := true // bool |  (optional)
	stepsToReproduce := "stepsToReproduce_example" // string |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	test := int32(56) // int32 |  (optional)
	testEngagement := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	testEngagementProduct := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	testEngagementProductProdType := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	testEngagementProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on product (optional)
	testEngagementTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on engagement (optional)
	testTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on test (optional)
	testTestType := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	title := "title_example" // string |  (optional)
	underDefectReview := true // bool |  (optional)
	underReview := true // bool |  (optional)
	uniqueIdFromTool := "uniqueIdFromTool_example" // string |  (optional)
	verified := true // bool |  (optional)
	vulnIdFromTool := "vulnIdFromTool_example" // string |  (optional)
	vulnerabilityId := "vulnerabilityId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsList(context.Background()).Active(active).After(after).Before(before).ComponentName(componentName).ComponentVersion(componentVersion).Created(created).Cvssv3(cvssv3).Cvssv3Score(cvssv3Score).Cwe(cwe).Date(date).DefectReviewRequestedBy(defectReviewRequestedBy).Description(description).Duplicate(duplicate).DuplicateFinding(duplicateFinding).DynamicFinding(dynamicFinding).EffortForFixing(effortForFixing).Endpoints(endpoints).EpssPercentile(epssPercentile).EpssScore(epssScore).FalseP(falseP).FilePath(filePath).FindingGroup(findingGroup).FoundBy(foundBy).HasJira(hasJira).HasTags(hasTags).HashCode(hashCode).Id(id).Impact(impact).InheritedTags(inheritedTags).IsMitigated(isMitigated).JiraChange(jiraChange).JiraCreation(jiraCreation).LastReviewed(lastReviewed).LastReviewedBy(lastReviewedBy).LastStatusUpdate(lastStatusUpdate).Limit(limit).Mitigated(mitigated).MitigatedBy(mitigatedBy).Mitigation(mitigation).NbOccurences(nbOccurences).NotTag(notTag).NotTags(notTags).NotTestEngagementProductTags(notTestEngagementProductTags).NotTestEngagementTags(notTestEngagementTags).NotTestTags(notTestTags).NumericalSeverity(numericalSeverity).O(o).Offset(offset).On(on).OutOfScope(outOfScope).OutsideOfSla(outsideOfSla).Param(param).Payload(payload).PlannedRemediationDate(plannedRemediationDate).PlannedRemediationVersion(plannedRemediationVersion).Prefetch(prefetch).ProductLifecycle(productLifecycle).ProductName(productName).ProductNameContains(productNameContains).PublishDate(publishDate).References(references).RelatedFields(relatedFields).Reporter(reporter).ReviewRequestedBy(reviewRequestedBy).Reviewers(reviewers).RiskAcceptance(riskAcceptance).RiskAccepted(riskAccepted).SastSinkObject(sastSinkObject).SastSourceFilePath(sastSourceFilePath).SastSourceLine(sastSourceLine).SastSourceObject(sastSourceObject).ScannerConfidence(scannerConfidence).Service(service).Severity(severity).SeverityJustification(severityJustification).SlaExpirationDate(slaExpirationDate).SlaStartDate(slaStartDate).SonarqubeIssue(sonarqubeIssue).StaticFinding(staticFinding).StepsToReproduce(stepsToReproduce).Tag(tag).Tags(tags).Test(test).TestEngagement(testEngagement).TestEngagementProduct(testEngagementProduct).TestEngagementProductProdType(testEngagementProductProdType).TestEngagementProductTags(testEngagementProductTags).TestEngagementTags(testEngagementTags).TestTags(testTags).TestTestType(testTestType).Title(title).UnderDefectReview(underDefectReview).UnderReview(underReview).UniqueIdFromTool(uniqueIdFromTool).Verified(verified).VulnIdFromTool(vulnIdFromTool).VulnerabilityId(vulnerabilityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsList`: PaginatedFindingList
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** |  | 
 **after** | **string** |  | 
 **before** | **string** |  | 
 **componentName** | **string** |  | 
 **componentVersion** | **string** |  | 
 **created** | **time.Time** | The date the finding was created inside DefectDojo.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **cvssv3** | **string** |  | 
 **cvssv3Score** | **float32** |  | 
 **cwe** | **[]int32** | Multiple values may be separated by commas. | 
 **date** | **string** | The date the flaw was discovered.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **defectReviewRequestedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **description** | **string** |  | 
 **duplicate** | **bool** |  | 
 **duplicateFinding** | **int32** |  | 
 **dynamicFinding** | **bool** |  | 
 **effortForFixing** | **string** |  | 
 **endpoints** | **[]int32** | Multiple values may be separated by commas. | 
 **epssPercentile** | **float32** |  | 
 **epssScore** | **float32** |  | 
 **falseP** | **bool** |  | 
 **filePath** | **string** |  | 
 **findingGroup** | **[]float32** | Multiple values may be separated by commas. | 
 **foundBy** | **[]int32** | Multiple values may be separated by commas. | 
 **hasJira** | **bool** |  | 
 **hasTags** | **bool** | Has tags | 
 **hashCode** | **string** |  | 
 **id** | **[]int32** | Multiple values may be separated by commas. | 
 **impact** | **string** |  | 
 **inheritedTags** | **[][]int32** | Internal use tags sepcifically for maintaining parity with product. This field will be present as a subset in the tags field | 
 **isMitigated** | **bool** |  | 
 **jiraChange** | **time.Time** | The date the linked Jira issue was last modified.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **jiraCreation** | **time.Time** | The date a Jira issue was created from this finding.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **lastReviewed** | **time.Time** | Provides the date the flaw was last &#39;touched&#39; by a tester.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **lastReviewedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **lastStatusUpdate** | **time.Time** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **mitigated** | **time.Time** | Denotes if this flaw has been fixed by storing the date it was fixed.  * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **mitigatedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **mitigation** | **string** |  | 
 **nbOccurences** | **[]int32** | Multiple values may be separated by commas. | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **notTestEngagementProductTags** | **[]string** | Comma separated list of exact tags not present on product | 
 **notTestEngagementTags** | **[]string** | Comma separated list of exact tags not present on engagement | 
 **notTestTags** | **[]string** | Comma separated list of exact tags present on test | 
 **numericalSeverity** | **string** |  | 
 **o** | **[]string** | Ordering  * &#x60;active&#x60; - Active * &#x60;-active&#x60; - Active (descending) * &#x60;component_name&#x60; - Component name * &#x60;-component_name&#x60; - Component name (descending) * &#x60;component_version&#x60; - Component version * &#x60;-component_version&#x60; - Component version (descending) * &#x60;created&#x60; - Created * &#x60;-created&#x60; - Created (descending) * &#x60;last_status_update&#x60; - Last status update * &#x60;-last_status_update&#x60; - Last status update (descending) * &#x60;last_reviewed&#x60; - Last reviewed * &#x60;-last_reviewed&#x60; - Last reviewed (descending) * &#x60;cwe&#x60; - Cwe * &#x60;-cwe&#x60; - Cwe (descending) * &#x60;date&#x60; - Date * &#x60;-date&#x60; - Date (descending) * &#x60;duplicate&#x60; - Duplicate * &#x60;-duplicate&#x60; - Duplicate (descending) * &#x60;dynamic_finding&#x60; - Dynamic finding * &#x60;-dynamic_finding&#x60; - Dynamic finding (descending) * &#x60;false_p&#x60; - False p * &#x60;-false_p&#x60; - False p (descending) * &#x60;found_by&#x60; - Found by * &#x60;-found_by&#x60; - Found by (descending) * &#x60;id&#x60; - Id * &#x60;-id&#x60; - Id (descending) * &#x60;is_mitigated&#x60; - Is mitigated * &#x60;-is_mitigated&#x60; - Is mitigated (descending) * &#x60;numerical_severity&#x60; - Numerical severity * &#x60;-numerical_severity&#x60; - Numerical severity (descending) * &#x60;out_of_scope&#x60; - Out of scope * &#x60;-out_of_scope&#x60; - Out of scope (descending) * &#x60;severity&#x60; - Severity * &#x60;-severity&#x60; - Severity (descending) * &#x60;reviewers&#x60; - Reviewers * &#x60;-reviewers&#x60; - Reviewers (descending) * &#x60;static_finding&#x60; - Static finding * &#x60;-static_finding&#x60; - Static finding (descending) * &#x60;test__engagement__product__name&#x60; - Test  engagement  product  name * &#x60;-test__engagement__product__name&#x60; - Test  engagement  product  name (descending) * &#x60;title&#x60; - Title * &#x60;-title&#x60; - Title (descending) * &#x60;under_defect_review&#x60; - Under defect review * &#x60;-under_defect_review&#x60; - Under defect review (descending) * &#x60;under_review&#x60; - Under review * &#x60;-under_review&#x60; - Under review (descending) * &#x60;verified&#x60; - Verified * &#x60;-verified&#x60; - Verified (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **on** | **string** |  | 
 **outOfScope** | **bool** |  | 
 **outsideOfSla** | **float32** |  | 
 **param** | **string** |  | 
 **payload** | **string** |  | 
 **plannedRemediationDate** | **string** |  | 
 **plannedRemediationVersion** | **string** |  | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **productLifecycle** | **string** | Comma separated list of exact product lifecycles | 
 **productName** | **string** | exact product name | 
 **productNameContains** | **string** | exact product name | 
 **publishDate** | **string** |  | 
 **references** | **string** |  | 
 **relatedFields** | **bool** | Expand finding external relations (engagement, environment, product,                                             product_type, test, test_type) | 
 **reporter** | **[]int32** | Multiple values may be separated by commas. | 
 **reviewRequestedBy** | **[]int32** | Multiple values may be separated by commas. | 
 **reviewers** | **[]int32** | Multiple values may be separated by commas. | 
 **riskAcceptance** | **float32** |  | 
 **riskAccepted** | **bool** |  | 
 **sastSinkObject** | **string** |  | 
 **sastSourceFilePath** | **string** |  | 
 **sastSourceLine** | **[]int32** | Multiple values may be separated by commas. | 
 **sastSourceObject** | **string** |  | 
 **scannerConfidence** | **[]int32** | Multiple values may be separated by commas. | 
 **service** | **string** |  | 
 **severity** | **string** |  | 
 **severityJustification** | **string** |  | 
 **slaExpirationDate** | **string** |  | 
 **slaStartDate** | **string** |  | 
 **sonarqubeIssue** | **[]int32** | Multiple values may be separated by commas. | 
 **staticFinding** | **bool** |  | 
 **stepsToReproduce** | **string** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **test** | **int32** |  | 
 **testEngagement** | **[]int32** | Multiple values may be separated by commas. | 
 **testEngagementProduct** | **[]int32** | Multiple values may be separated by commas. | 
 **testEngagementProductProdType** | **[]int32** | Multiple values may be separated by commas. | 
 **testEngagementProductTags** | **[]string** | Comma separated list of exact tags present on product | 
 **testEngagementTags** | **[]string** | Comma separated list of exact tags present on engagement | 
 **testTags** | **[]string** | Comma separated list of exact tags present on test | 
 **testTestType** | **[]int32** | Multiple values may be separated by commas. | 
 **title** | **string** |  | 
 **underDefectReview** | **bool** |  | 
 **underReview** | **bool** |  | 
 **uniqueIdFromTool** | **string** |  | 
 **verified** | **bool** |  | 
 **vulnIdFromTool** | **string** |  | 
 **vulnerabilityId** | **string** |  | 

### Return type

[**PaginatedFindingList**](PaginatedFindingList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsMetadataCreate

> FindingMeta FindingsMetadataCreate(ctx, id).FindingMetaRequest(findingMetaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	findingMetaRequest := *openapiclient.NewFindingMetaRequest("Name_example", "Value_example") // FindingMetaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsMetadataCreate(context.Background(), id).FindingMetaRequest(findingMetaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsMetadataCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsMetadataCreate`: FindingMeta
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsMetadataCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsMetadataCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findingMetaRequest** | [**FindingMetaRequest**](FindingMetaRequest.md) |  | 

### Return type

[**FindingMeta**](FindingMeta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsMetadataDestroy

> FindingsMetadataDestroy(ctx, id).Name(name).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	name := int32(56) // int32 | name of the metadata to retrieve. If name is empty, return all the                                     metadata associated with the finding

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsMetadataDestroy(context.Background(), id).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsMetadataDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsMetadataDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **int32** | name of the metadata to retrieve. If name is empty, return all the                                     metadata associated with the finding | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsMetadataList

> []FindingMeta FindingsMetadataList(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsMetadataList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsMetadataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsMetadataList`: []FindingMeta
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsMetadataList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsMetadataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]FindingMeta**](FindingMeta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsMetadataUpdate

> FindingMeta FindingsMetadataUpdate(ctx, id).FindingMetaRequest(findingMetaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	findingMetaRequest := *openapiclient.NewFindingMetaRequest("Name_example", "Value_example") // FindingMetaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsMetadataUpdate(context.Background(), id).FindingMetaRequest(findingMetaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsMetadataUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsMetadataUpdate`: FindingMeta
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsMetadataUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsMetadataUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findingMetaRequest** | [**FindingMetaRequest**](FindingMetaRequest.md) |  | 

### Return type

[**FindingMeta**](FindingMeta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsNotesCreate

> Note FindingsNotesCreate(ctx, id).AddNewNoteOptionRequest(addNewNoteOptionRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	addNewNoteOptionRequest := *openapiclient.NewAddNewNoteOptionRequest("Entry_example") // AddNewNoteOptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsNotesCreate(context.Background(), id).AddNewNoteOptionRequest(addNewNoteOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsNotesCreate`: Note
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addNewNoteOptionRequest** | [**AddNewNoteOptionRequest**](AddNewNoteOptionRequest.md) |  | 

### Return type

[**Note**](Note.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsNotesRetrieve

> FindingToNotes FindingsNotesRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsNotesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsNotesRetrieve`: FindingToNotes
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsNotesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindingToNotes**](FindingToNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsOriginalCreate

> FindingsOriginalCreate(ctx, id, newFid).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	newFid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsOriginalCreate(context.Background(), id, newFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsOriginalCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 
**newFid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsOriginalCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsPartialUpdate

> Finding FindingsPartialUpdate(ctx, id).PatchedFindingRequest(patchedFindingRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	patchedFindingRequest := *openapiclient.NewPatchedFindingRequest() // PatchedFindingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsPartialUpdate(context.Background(), id).PatchedFindingRequest(patchedFindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsPartialUpdate`: Finding
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFindingRequest** | [**PatchedFindingRequest**](PatchedFindingRequest.md) |  | 

### Return type

[**Finding**](Finding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsRemoveNotePartialUpdate

> FindingsRemoveNotePartialUpdate(ctx, id).PatchedFindingNoteRequest(patchedFindingNoteRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	patchedFindingNoteRequest := *openapiclient.NewPatchedFindingNoteRequest() // PatchedFindingNoteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsRemoveNotePartialUpdate(context.Background(), id).PatchedFindingNoteRequest(patchedFindingNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsRemoveNotePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsRemoveNotePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFindingNoteRequest** | [**PatchedFindingNoteRequest**](PatchedFindingNoteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsRemoveTagsPartialUpdate

> FindingsRemoveTagsPartialUpdate(ctx, id).PatchedTagRequest(patchedTagRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	patchedTagRequest := *openapiclient.NewPatchedTagRequest() // PatchedTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsRemoveTagsPartialUpdate(context.Background(), id).PatchedTagRequest(patchedTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsRemoveTagsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsRemoveTagsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTagRequest** | [**PatchedTagRequest**](PatchedTagRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsRemoveTagsUpdate

> FindingsRemoveTagsUpdate(ctx, id).TagRequest(tagRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	tagRequest := *openapiclient.NewTagRequest([]string{"Tags_example"}) // TagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingsAPI.FindingsRemoveTagsUpdate(context.Background(), id).TagRequest(tagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsRemoveTagsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsRemoveTagsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagRequest** | [**TagRequest**](TagRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsRequestResponseCreate

> BurpRawRequestResponse FindingsRequestResponseCreate(ctx, id).BurpRawRequestResponseRequest(burpRawRequestResponseRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	burpRawRequestResponseRequest := *openapiclient.NewBurpRawRequestResponseRequest([]map[string]string{map[string]string{"key": "Inner_example"}}) // BurpRawRequestResponseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsRequestResponseCreate(context.Background(), id).BurpRawRequestResponseRequest(burpRawRequestResponseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsRequestResponseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsRequestResponseCreate`: BurpRawRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsRequestResponseCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsRequestResponseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **burpRawRequestResponseRequest** | [**BurpRawRequestResponseRequest**](BurpRawRequestResponseRequest.md) |  | 

### Return type

[**BurpRawRequestResponse**](BurpRawRequestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsRequestResponseRetrieve

> BurpRawRequestResponse FindingsRequestResponseRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsRequestResponseRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsRequestResponseRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsRequestResponseRetrieve`: BurpRawRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsRequestResponseRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsRequestResponseRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BurpRawRequestResponse**](BurpRawRequestResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsRetrieve

> Finding FindingsRetrieve(ctx, id).Prefetch(prefetch).RelatedFields(relatedFields).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	relatedFields := true // bool | Expand finding external relations (engagement, environment, product,                                             product_type, test, test_type) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsRetrieve(context.Background(), id).Prefetch(prefetch).RelatedFields(relatedFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsRetrieve`: Finding
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **relatedFields** | **bool** | Expand finding external relations (engagement, environment, product,                                             product_type, test, test_type) | 

### Return type

[**Finding**](Finding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsTagsCreate

> Tag FindingsTagsCreate(ctx, id).TagRequest(tagRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	tagRequest := *openapiclient.NewTagRequest([]string{"Tags_example"}) // TagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsTagsCreate(context.Background(), id).TagRequest(tagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsTagsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsTagsCreate`: Tag
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsTagsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsTagsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagRequest** | [**TagRequest**](TagRequest.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsTagsRetrieve

> Tag FindingsTagsRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsTagsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsTagsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsTagsRetrieve`: Tag
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsTagsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsTagsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingsUpdate

> Finding FindingsUpdate(ctx, id).FindingRequest(findingRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this finding.
	findingRequest := *openapiclient.NewFindingRequest("Title_example", "Severity_example", "Description_example", "NumericalSeverity_example") // FindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingsAPI.FindingsUpdate(context.Background(), id).FindingRequest(findingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingsAPI.FindingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingsUpdate`: Finding
	fmt.Fprintf(os.Stdout, "Response from `FindingsAPI.FindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findingRequest** | [**FindingRequest**](FindingRequest.md) |  | 

### Return type

[**Finding**](Finding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

