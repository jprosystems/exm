package model

type Utilisateur struct {
	UTILISATEURID    string `json:"UTILISATEURID"`
	USERNAME         string `json:"USERNAME"`
	MOT_DE_PASSE     string `json:"MOT_DE_PASSE"`
	TYPE_UTILISATEUR string `json:"TYPE_UTILISATEUR"`
	MAIL             string `json:"MAIL"`
}
