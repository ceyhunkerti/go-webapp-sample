package controller

const (
	API           = "/api"
	APIBooks      = API + "/books"
	APIBooksID    = APIBooks + "/:id"
	APICategories = API + "/categories"
	APIFormats    = API + "/formats"
	APIDatasets   = API + "/datasets"
	APILineages   = API + "/lineages"
)

const (
	// APIAccount represents the group of auth management API.
	APIAccount = API + "/auth"
	// APIAccountLoginStatus represents the API to get the status of logged in account.
	APIAccountLoginStatus = APIAccount + "/loginStatus"
	// APIAccountLoginAccount represents the API to get the logged in account.
	APIAccountLoginAccount = APIAccount + "/loginAccount"
	// APIAccountLogin represents the API to login by session authentication.
	APIAccountLogin = APIAccount + "/login"
	// APIAccountLogout represents the API to logout.
	APIAccountLogout = APIAccount + "/logout"
)

const (
	// APIHealth represents the API to get the status of this application.
	APIHealth = API + "/health"
)
