package model

var PatentInfoRepo = new(PatentInfo)

func (u PatentInfo) Save(param *PatentInfo) (*int, error) {
	er := DB.Debug().Table("patent_info").Save(param).Error
	if er != nil {
		return nil, er
	}
	return &param.Id, nil
}
