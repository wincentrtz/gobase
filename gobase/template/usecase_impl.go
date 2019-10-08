package template

type templateUsecase struct {
	templateRepo TemplateRepository
}

func NewTemplateUsecase(tr TemplateRepository) TemplateUsecase {
	return &templateUsecase{
		templateRepo: tr,
	}
}
