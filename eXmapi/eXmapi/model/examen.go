package model

type Examen struct {
	EXAMENID                   string `json:"EXAMENID"`
	SEMESTRE                   string `json:"SEMESTRE"`
	COURS_GROUPE               string `json:"COURS_GROUPE"`
	DESCRIPTION                string `json:"DESCRIPTION"`
	LANGUE                     string `json:"LANGUE"`
	NOMBRE_QUESTIONS           string `json:"NOMBRE_QUESTIONS"`
	NOMBRE_POINTS_TOTAL        string `json:"NOMBRE_POINTS_TOTAL"`
	DUREE_MAXIMALE             string `json:"DUREE_MAXIMALE"`
	DATE_HEURE_ACTIVATION      string `json:"DATE_HEURE_ACTIVATION"`
	DATE_DERNIERE_MODIFICATION string `json:"DATE_DERNIERE_MODIFICATION"`
}
