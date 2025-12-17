package Initializers

import (
	"time"
	"wan-api-kol-event/Models"

)

func MigrateAndSeed() {
	// Nếu table đã có, chỉ seed khi chưa có dữ liệu.
	// Nếu table chưa có, anh tạo thủ công hoặc sửa lại AutoMigrate sau.

	var count int64
	if err := DB.Model(&Models.Kol{}).Count(&count).Error; err != nil {
		panic(err)
	}
	if count > 0 {
		return
	}

	now := time.Now()
	kols := make([]Models.Kol, 0, 30)
	for i := 1; i <= 30; i++ {
		kols = append(kols, Models.Kol{
			KolID:                int64(1000 + i),
			UserProfileID:        int64(2000 + i),
			Language:             []string{"en", "vn"}[i%2],
			Education:            "Bachelor",
			ExpectedSalary:       int64(30000 + i*1000),
			ExpectedSalaryEnable: i%2 == 0,
			ChannelSettingTypeID: int64((i % 3) + 1),
			IDFrontURL:           "https://example.com/id-front.jpg",
			IDBackURL:            "https://example.com/id-back.jpg",
			PortraitURL:          "https://example.com/portrait.jpg",
			RewardID:             int64(300 + i),
			PaymentMethodID:      int64(400 + i),
			TestimonialsID:       int64(500 + i),
			VerificationStatus:   i%2 == 0,
			Enabled:              true,
			ActiveDate:           now.AddDate(0, 0, -i),
			Active:               true,
			CreatedBy:            "seed",
			CreatedDate:          now,
			ModifiedBy:           "seed",
			ModifiedDate:         now,
			IsRemove:             false,
			IsOnBoarding:         i%2 == 0,
			Code:                 "KOL-SEED-" + time.Now().Format("20060102") + "-" + fmtInt(i),
			PortraitRightURL:     "https://example.com/portrait-right.jpg",
			PortraitLeftURL:      "https://example.com/portrait-left.jpg",
			LivenessStatus:       i%2 == 0,
		})
	}

	if err := DB.Create(&kols).Error; err != nil {
		panic(err)
	}
}

func fmtInt(i int) string {
	// tránh import strconv ở trên cho gọn
	if i < 10 {
		return "0" + string(rune('0'+i))
	}
	// fallback đơn giản
	return time.Now().Format("15") // không quan trọng format, chỉ cần unique nhẹ
}
