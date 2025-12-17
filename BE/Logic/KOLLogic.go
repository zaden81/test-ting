package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func GetKolLogic(pageIndex, pageSize int64) ([]*DTO.KolDTO, int64, error) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	db := Initializers.DB

	q := db.Model(&Models.Kol{}).
		Where(`"Enabled" = ? AND "Active" = ? AND "IsRemove" = ?`, true, true, false)

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var rows []Models.Kol
	offset := int((pageIndex - 1) * pageSize)
	if err := q.Order(`"KolID" asc`).Limit(int(pageSize)).Offset(offset).Find(&rows).Error; err != nil {
		return nil, 0, err
	}

	out := make([]*DTO.KolDTO, 0, len(rows))
	for _, r := range rows {
		item := DTO.KolDTO{
			KolID:                r.KolID,
			UserProfileID:        r.UserProfileID,
			Language:             r.Language,
			Education:            r.Education,
			ExpectedSalary:       r.ExpectedSalary,
			ExpectedSalaryEnable: r.ExpectedSalaryEnable,
			ChannelSettingTypeID: r.ChannelSettingTypeID,
			IDFrontURL:           r.IDFrontURL,
			IDBackURL:            r.IDBackURL,
			PortraitURL:          r.PortraitURL,
			RewardID:             r.RewardID,
			PaymentMethodID:      r.PaymentMethodID,
			TestimonialsID:       r.TestimonialsID,
			VerificationStatus:   r.VerificationStatus,
			Enabled:              r.Enabled,
			ActiveDate:           r.ActiveDate,
			Active:               r.Active,
			CreatedBy:            r.CreatedBy,
			CreatedDate:          r.CreatedDate,
			ModifiedBy:           r.ModifiedBy,
			ModifiedDate:         r.ModifiedDate,
			IsRemove:             r.IsRemove,
			IsOnBoarding:         r.IsOnBoarding,
			Code:                 r.Code,
			PortraitRightURL:     r.PortraitRightURL,
			PortraitLeftURL:      r.PortraitLeftURL,
			LivenessStatus:       r.LivenessStatus,
		}
		out = append(out, &item)
	}

	return out, total, nil
}
