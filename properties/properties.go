package properties

import "github.com/alexflint/go-arg"

type args struct {
	AzercellCabinetApiBaseUrl string `arg:"env:AZERCELL_CABINET_API_BASE_URL"`
	AzercellCabinetNumber     string `arg:"env:AZERCELL_CABINET_NUMBER"`
	AzercellCabinetPassword   string `arg:"env:AZERCELL_CABINET_PASSWORD"`
}

var Props args

func LoadProperties() {
	arg.Parse(&Props)
}