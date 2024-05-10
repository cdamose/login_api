package conversion

import (
	"login_api/internal/auth/model/dao"
	"login_api/internal/auth/model/dto"
)

func ConvertToUpdatedUserProfile(profile dao.UserProfile) dto.UserProfile {
	return dto.UserProfile{
		UserId:      profile.UserID,
		IsVerified:  profile.IsVerified,
		CreatedAt:   profile.CreatedAt.String(),
		VerfiedAt:   "",
		PhoneNumber: profile.MobileNumber,
	}
}
