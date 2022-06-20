package vars

var (
	Auth  string
	Round int    = 0
	Url   string = "http://kitkabackend.eastus.cloudapp.azure.com:5010/round/finishv2/%v"
)

type Vars struct {
	Auth  string
	Round int
	Url   string
}
