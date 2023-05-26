package maker

import (
	"context"
	"cv-maker-service/database"
	error2 "cv-maker-service/error"
	"cv-maker-service/repo/cv_data_byte"
	"cv-maker-service/repo/cv_information"
	education2 "cv-maker-service/repo/education"
	experience2 "cv-maker-service/repo/experience"
	"cv-maker-service/repo/user_language"
	"cv-maker-service/repo/user_programming_language"
)

func (s SimpleService) CreateCV(ctx context.Context, dbManager database.GormDbManager, cvData *CVData) error {
	userDataResponse, err := s.getUserDataByEmail(ctx, cvData.Email)
	if err != nil {
		return error2.ServiceLayerError.Wrap(err)
	}

	//err = emitter.Connection.Tx()
	//if err != nil {
	//	return error2.RabbitMQError.Wrap(err)
	//}

	pdfBytes, err := s.generatePDFDataBytesFromGivenData(ctx, cvData, userDataResponse.FirstName, userDataResponse.LastName)
	if err != nil {
		return error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to generate pdf from given data")
	}

	cvDataBytes := &cv_data_byte.CVDataByte{
		UserID:       userDataResponse.Id,
		Bytes:        pdfBytes,
		TemplateName: cvData.TemplateName,
	}

	err = s.MailerService.SendEmailWithParameters(ctx, cvData.Email, userDataResponse.FirstName, userDataResponse.LastName, pdfBytes)
	if err != nil {
		return error2.ServiceLayerError.WrapWithUserMessage(err, "Failed to send email")
	}

	err = dbManager.ExecuteTransaction(func(operations database.Operations) error {
		err := s.CVDataByteRepository.Save(operations, cvDataBytes)
		if err != nil {
			//emitter.Connection.TxRollback()
			return error2.DatabaseError.Wrap(err)
		}

		cvInformation := &cv_information.CVInformation{
			CVDataByteID:    int64(cvDataBytes.ID),
			JobTitle:        cvData.JobTitle,
			PhoneNumber:     cvData.PhoneNumber,
			Website:         cvData.Website,
			LinkedinProfile: cvData.LinkedinProfile,
			GithubProfile:   cvData.GithubProfile,
			Description:     cvData.Description,
		}

		err = s.CVInformationRepository.Save(operations, cvInformation)
		if err != nil {
			//emitter.Connection.TxRollback()
			return error2.DatabaseError.Wrap(err)
		}

		for _, education := range cvData.Educations {
			err = s.EducationRepository.Save(operations, &education2.Education{
				CVDataByteID: int64(cvDataBytes.ID),
				Faculty:      education.Faculty,
				Title:        education.Title,
				From:         education.From,
				To:           education.To,
			})
			if err != nil {
				//emitter.Connection.TxRollback()
				return error2.DatabaseError.Wrap(err)
			}
		}

		for _, experience := range cvData.Experiences {
			err = s.ExperienceRepository.Save(operations, &experience2.Experience{
				Company:      experience.Company,
				JobTitle:     experience.JobTitle,
				From:         experience.From,
				To:           experience.To,
				Description:  experience.Description,
				CVDataByteID: int64(cvDataBytes.ID),
			})
			if err != nil {
				//emitter.Connection.TxRollback()
				return error2.DatabaseError.Wrap(err)
			}
		}

		for _, pl := range cvData.ProgrammingLanguages {
			err = s.UserProgrammingLanguageRepository.Save(operations, &user_programming_language.UserProgrammingLanguage{
				UserID:                userDataResponse.Id,
				ProgrammingLanguageID: pl.ID,
				Percentage:            pl.Percentage,
				CVDataByteID:          int64(cvDataBytes.ID),
			})
			if err != nil {
				//emitter.Connection.TxRollback()
				return error2.DatabaseError.Wrap(err)
			}
		}

		for _, language := range cvData.Languages {
			err = s.UserLanguageRepository.Save(operations, &user_language.UserLanguage{
				UserID:       userDataResponse.Id,
				LanguageID:   language.ID,
				CVDataByteID: int64(cvDataBytes.ID),
			})
			if err != nil {
				//emitter.Connection.TxRollback()
				return error2.DatabaseError.Wrap(err)
			}
		}
		return nil
	})
	if err != nil {
		return error2.DatabaseError.WrapWithUserMessage(err, "Failed to save cv information in transaction for provided information")
	}

	//err = emitter.Connection.TxCommit()
	//if err != nil {
	//	return error2.RabbitMQError.Wrap(err)
	//}

	return nil
}
