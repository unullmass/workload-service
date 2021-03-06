/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package mock

import (
	"intel/isecl/workload-service/model"
	"intel/isecl/workload-service/repository"
)

type MockReport struct {
	CreateFn                   func(*model.Report) error
	RetrieveByFilterCriteriaFn func(repository.ReportFilter) ([]model.Report, error)
	DeleteByReportIDFn         func(string) error
}

func (m *MockReport) Create(r *model.Report) error {
	if m.CreateFn != nil {
		return m.CreateFn(r)
	}
	return nil
}

func (m *MockReport) RetrieveByFilterCriteria(filter repository.ReportFilter) ([]model.Report, error) {
	if m.RetrieveByFilterCriteriaFn != nil {
		return m.RetrieveByFilterCriteriaFn(filter)
	}
	return []model.Report{r}, nil
}

func (m *MockReport) DeleteByReportID(reportID string) error {
	if m.DeleteByReportIDFn != nil {
		return m.DeleteByReportIDFn(reportID)
	}
	return nil
}
