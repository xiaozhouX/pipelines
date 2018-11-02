// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ReportRunMetricsResponseReportRunMetricResultStatus  - UNSPECIFIED: Default value if not present.
//  - OK: Indicates successful reporting.
//  - INVALID_ARGUMENT: Indicates that the payload of the metric is invalid.
//  - DUPLICATE_REPORTING: Indicates that the metric has been reported before.
//  - INTERNAL_ERROR: Indicates that something went wrong in the server.
// swagger:model ReportRunMetricsResponseReportRunMetricResultStatus
type ReportRunMetricsResponseReportRunMetricResultStatus string

const (

	// ReportRunMetricsResponseReportRunMetricResultStatusUNSPECIFIED captures enum value "UNSPECIFIED"
	ReportRunMetricsResponseReportRunMetricResultStatusUNSPECIFIED ReportRunMetricsResponseReportRunMetricResultStatus = "UNSPECIFIED"

	// ReportRunMetricsResponseReportRunMetricResultStatusOK captures enum value "OK"
	ReportRunMetricsResponseReportRunMetricResultStatusOK ReportRunMetricsResponseReportRunMetricResultStatus = "OK"

	// ReportRunMetricsResponseReportRunMetricResultStatusINVALIDARGUMENT captures enum value "INVALID_ARGUMENT"
	ReportRunMetricsResponseReportRunMetricResultStatusINVALIDARGUMENT ReportRunMetricsResponseReportRunMetricResultStatus = "INVALID_ARGUMENT"

	// ReportRunMetricsResponseReportRunMetricResultStatusDUPLICATEREPORTING captures enum value "DUPLICATE_REPORTING"
	ReportRunMetricsResponseReportRunMetricResultStatusDUPLICATEREPORTING ReportRunMetricsResponseReportRunMetricResultStatus = "DUPLICATE_REPORTING"

	// ReportRunMetricsResponseReportRunMetricResultStatusINTERNALERROR captures enum value "INTERNAL_ERROR"
	ReportRunMetricsResponseReportRunMetricResultStatusINTERNALERROR ReportRunMetricsResponseReportRunMetricResultStatus = "INTERNAL_ERROR"
)

// for schema
var reportRunMetricsResponseReportRunMetricResultStatusEnum []interface{}

func init() {
	var res []ReportRunMetricsResponseReportRunMetricResultStatus
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","OK","INVALID_ARGUMENT","DUPLICATE_REPORTING","INTERNAL_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportRunMetricsResponseReportRunMetricResultStatusEnum = append(reportRunMetricsResponseReportRunMetricResultStatusEnum, v)
	}
}

func (m ReportRunMetricsResponseReportRunMetricResultStatus) validateReportRunMetricsResponseReportRunMetricResultStatusEnum(path, location string, value ReportRunMetricsResponseReportRunMetricResultStatus) error {
	if err := validate.Enum(path, location, value, reportRunMetricsResponseReportRunMetricResultStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this report run metrics response report run metric result status
func (m ReportRunMetricsResponseReportRunMetricResultStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReportRunMetricsResponseReportRunMetricResultStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}