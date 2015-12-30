package goloto

type Tur string

const (
	Sayisal   Tur = "sayisal"
	SuperLoto     = "superloto"
	OnNumara      = "onnumara"
	SansTopu      = "sanstopu"
)

const BaseURL string = "http://millipiyango.gov.tr/sonuclar"

type Cekilis struct {
	Tur Tur
}

func (t Tur) String() string {
	switch t {
	case Sayisal:
		return "sayisal"
	case SuperLoto:
		return "superloto"
	case OnNumara:
		return "onnumara"
	case SansTopu:
		return "sanstopu"
	}
	return ""
}

type Date struct {
	Tarih     string `json:"tarih"`
	TarihView string `json:"tarihView"`
}

type Result struct {
	Success bool `json:"success"`
	Data    struct {
		Oid                    string `json:"oid"`
		Hafta                  int    `json:"hafta"`
		BuyukIkramiyeKazananIl string `json:"buyukIkramiyeKazananIl"`
		CekilisTarihi          string `json:"cekilisTarihi"`
		CekilisTuru            string `json:"cekilisTuru"`
		Rakamlar               string `json:"rakamlar"`
		RakamlarNumaraSirasi   string `json:"rakamlarNumaraSirasi"`
		Devretti               bool   `json:"devretti"`
		DevirSayisi            int    `json:"devirSayisi"`
		BilenKisiler           []struct {
			Oid                     string  `json:"oid"`
			KisiBasinaDusenIkramiye float64 `json:"kisiBasinaDusenIkramiye"`
			KisiSayisi              int     `json:"kisiSayisi"`
			Tur                     string  `json:"tur"`
		} `json:"bilenKisiler"`
		BuyukIkrKazananIlIlceler []struct {
			Il       string `json:"il"`
			IlView   string `json:"ilView"`
			Ilce     string `json:"ilce"`
			IlceView string `json:"ilceView"`
		} `json:"buyukIkrKazananIlIlceler"`
		KibrisHasilati       float64 `json:"kibrisHasilati"`
		DevirTutari          float64 `json:"devirTutari"`
		KolonSayisi          float64 `json:"kolonSayisi"`
		Kdv                  float64 `json:"kdv"`
		ToplamHasilat        float64 `json:"toplamHasilat"`
		Hasilat              float64 `json:"hasilat"`
		Sov                  float64 `json:"sov"`
		IkramiyeEH           float64 `json:"ikramiyeEH"`
		BuyukIkramiye        float64 `json:"buyukIkramiye"`
		HaftayaDevredenTutar float64 `json:"haftayaDevredenTutar"`
	} `json:"data"`
}
