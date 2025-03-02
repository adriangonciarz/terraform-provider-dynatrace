package download

type DataSourceClient interface {
	ListInterface() (interface{}, error)
}

type DataSourceDetails struct {
	Values         map[string]interface{}
	UniqueName     string
	ComputedValues map[string]interface{}
}

type DataSource struct {
	RESTMap map[string]*DataSourceDetails
}

type DataSourceData map[string]*DataSource
