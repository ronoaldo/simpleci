package api

type Repository struct {
	URL  string
	Type string
}

type BuildRequest struct {
	Repository Repository
	Branch     string
}
