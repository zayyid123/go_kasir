package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
	"time"
)

type ReportService struct {
	repo repository.ReportRepository
}

func NewReportService(r repository.ReportRepository) *ReportService {
	return &ReportService{repo: r}
}

func (s *ReportService) Report(startDate, endDate *time.Time) (*model.Report, error) {
	return s.repo.Report(startDate, endDate)
}

func (s *ReportService) BestSeller(startDate, endDate *time.Time) (*model.BestSeller, error) {
	return s.repo.BestSellingProduct(startDate, endDate)
}
