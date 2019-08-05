package template

type templateUsecase struct {
	templateRepo TemplateRepository
}

func NewTemplateUsecase() TemplateUsecase {
	return &templateUsecase{
		templateRepo: NewTemplateRepository(),
	}
}
