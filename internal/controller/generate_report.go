package controller

func (i *implementation) GenerateReport(path string) error {
	return i.reportUseCase.GenerateReportFromRepository(path)
}
