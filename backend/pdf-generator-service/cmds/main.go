package main

import (
	"github.com/gin-gonic/gin"
	"pdf-generator-service/builder"
	"pdf-generator-service/environment"
)

func main() {
	services := builder.BuildServices()

	_ = struct {
		Name            string
		Title           string
		Phone           string
		Email           string
		Website         string
		LinkedinProfile string
		GithubProfile   string
		Educations      []any
	}{
		Name:            "Borislav Petric",
		Title:           "Software Engineer",
		Phone:           "+381695557525",
		Email:           "borislavpetric66@gmail.com",
		Website:         "borislavpetric.com",
		LinkedinProfile: "borislav.linkedin.com",
		GithubProfile:   "borislav.github.com",
		Educations: []any{struct {
			Years   string
			Title   string
			Faculty string
		}{Years: "2021-2023", Title: "Software Engineer", Faculty: "Faculty of Organisational Sciences"}},
	}

	app := gin.Default()
	builder.BuildAPILayer(app, services)

	go builder.ServeGRPCServer(services)
	app.Run(environment.HTTP_SERVER_HOST.Value())
}
