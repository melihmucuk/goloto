# goloto
Milli Piyango client for go.


## Installation

`$ go get github.com/melihmucuk/goloto`

## Usage

* import package

`import "github.com/melihmucuk/goloto"`

* initialize goloto

`sayisal := goloto.Cekilis{goloto.Sayisal} // Or goloto.OnNumara, goloto.SuperLoto, goloto.SansTopu`

* get dates

`sayisalDates, _ := sayisal.GetDates()`

* get result

`sayisalResult, _ := sayisal.GetResults(sayisalDates[0])`

* use result

`fmt.Printf("Sayısal Loto: %s : %s \n", sayisalDates[0].TarihView, sayisalResult.Data.RakamlarNumaraSirasi)`