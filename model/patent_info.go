package model

type PatentInfo struct {
	Id              int    `gorm:"int(60) not null; primary_key;AUTO_INCREMENT" json:"id"`
	PatentId        string `gorm:"varchar(254) not null " json:"patent_id"`
	PatentName      string `gorm:"varchar(254) not null " json:"patent_name"`
	PatentType      string `gorm:"varchar(254) not null " json:"applicant"`
	PatentField     string `gorm:"varchar(254) not null " json:"patent_type"`
	ApplicationDate string `gorm:"varchar(254) not null " json:"patent_field"`
	PublicationDate string `gorm:"varchar(254) not null " json:"application_date"`
	Inventor        string `gorm:"varchar(254) not null " json:"publication_date"`
	PatentAgency    string `gorm:"varchar(254) not null " json:"patent_agency"`
	Abstract        string `gorm:"longtext(254) not null " json:"abstract"`
	SummaryByGdp    string `gorm:"longtext(254) not null " json:"summary_by_gdp"`
	PdfLocation     string `gorm:"varchar(254) not null " json:"pdf_location"`
}
