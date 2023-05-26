package maker

type (
	Experience struct {
		Company     string `json:"company"`
		JobTitle    string `json:"jobTitle"`
		From        string `json:"from"`
		To          string `json:"to"`
		Description string `json:"description"`
	}

	Education struct {
		Faculty string `json:"faculty"`
		Title   string `json:"title"`
		From    string `json:"from"`
		To      string `json:"to"`
	}

	ProgrammingLanguage struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		Percentage int64  `json:"percentage"`
	}

	Language struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	CVData struct {
		Email                string                 `json:"email,omitempty"`
		Description          string                 `json:"description"`
		JobTitle             string                 `json:"jobTitle"`
		PhoneNumber          string                 `json:"phoneNumber"`
		Website              string                 `json:"website"`
		LinkedinProfile      string                 `json:"linkedinProfile"`
		GithubProfile        string                 `json:"githubProfiile"`
		TemplateName         string                 `json:"templateName"`
		ProfilePhotoBytes    string                 `json:"profilePhotoBytes"`
		Experiences          []*Experience          `json:"experiences"`
		Educations           []*Education           `json:"educations"`
		Languages            []*Language            `json:"languages"`
		ProgrammingLanguages []*ProgrammingLanguage `json:"programmingLanguages"`
	}
)
