/*
 * AWSAPICallViaCloudTrail
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package s3
// LegalHoldInfo struct for LegalHoldInfo
type LegalHoldInfo struct {
	LastModifiedTime int64 `json:"lastModifiedTime,omitempty"`
	IsUnderLegalHold bool `json:"isUnderLegalHold,omitempty"`
}