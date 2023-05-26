package maker

import (
	"context"
	"cv-maker-service/services/authentication"
	"cv-maker-service/services/generator"
	"fmt"
)

func (s SimpleService) getUserDataByEmail(ctx context.Context, email string) (*authentication.FindUserByEmailResponse, error) {
	request := &authentication.FindUserByEmailRequest{
		Email: email,
	}
	res, err := s.AuthenticationService.FindUserByEmail(ctx, request)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s SimpleService) generatePDFDataBytesFromGivenData(ctx context.Context, data *CVData, firstName, lastName string) ([]byte, error) {
	programmingLanguages := make([]*generator.GeneratePDFFromTemplateAndDataProgrammingLanguage, 0)
	for _, pl := range data.ProgrammingLanguages {
		grpcProgrammingLanguage := &generator.GeneratePDFFromTemplateAndDataProgrammingLanguage{
			Id:         pl.ID,
			Percentage: pl.Percentage,
			Name:       pl.Name,
		}
		programmingLanguages = append(programmingLanguages, grpcProgrammingLanguage)
	}

	experiences := make([]*generator.GeneratePDFFromTemplateAndDataExperience, 0)
	for _, exp := range data.Experiences {
		grpcExperience := &generator.GeneratePDFFromTemplateAndDataExperience{
			Company:     exp.Company,
			JobTitle:    exp.JobTitle,
			From:        exp.From,
			To:          exp.To,
			Description: exp.Description,
		}
		experiences = append(experiences, grpcExperience)
	}

	educations := make([]*generator.GeneratePDFFromTemplateAndDataEducation, 0)
	for _, education := range data.Educations {
		grpcEducation := &generator.GeneratePDFFromTemplateAndDataEducation{
			Faculty: education.Faculty,
			Title:   education.Title,
			From:    education.From,
			To:      education.To,
		}
		educations = append(educations, grpcEducation)
	}

	languages := make([]*generator.GeneratePDFFromTemplateAndDataLanguage, 0)
	for _, language := range data.Languages {
		grpcLanguage := &generator.GeneratePDFFromTemplateAndDataLanguage{
			Name: language.Name,
		}
		languages = append(languages, grpcLanguage)
	}

	request := &generator.GeneratePDFFromTemplateAndDataRequest{
		Template: data.TemplateName,
		UserData: &generator.GeneratePDFFromTemplateAndDataUser{
			Name:                 fmt.Sprintf("%s %s", firstName, lastName),
			Email:                data.Email,
			PhoneNumber:          data.PhoneNumber,
			Description:          data.Description,
			ProgrammingLanguages: programmingLanguages,
			Experiences:          experiences,
			Educations:           educations,
			Website:              data.Website,
			LinkedinProfile:      data.LinkedinProfile,
			GithubProfile:        data.GithubProfile,
			ProfilePhotoBytes:    data.ProfilePhotoBytes,
			Languages:            languages,
		},
	}

	//req, err := json.Marshal(request)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = emitter.Push(ctx, req, "cv_maker.create")
	//if err != nil {
	//	return nil, err
	//}

	res, err := s.GeneratorService.GeneratePDFFromTemplateAndData(ctx, request)
	if err != nil {
		return nil, err
	}

	return res.PdfBytes, nil
}
