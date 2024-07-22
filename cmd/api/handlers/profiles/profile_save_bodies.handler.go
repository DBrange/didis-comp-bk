package handlers

// func UpdateUserSaveBody(c *gin.Context) (*profile_dto.UpdateProfileDTOReq, *location_dto.UpdateLocationDTOReq, error) {
// 	var attributesToUpdate api_dto.UpdateUserDTOReq
// 	if err := c.ShouldBindJSON(&attributesToUpdate); err != nil {
// 		return nil, nil, err
// 	}

// 	err := utils.Validate.StructExcept(attributesToUpdate, "Location")

// 	if err != nil {
// 		err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
// 		if errors.Is(err, customerrors.ErrValidationFailed) {
// 			appErr := customerrors.AppError{
// 				Code: customerrors.ErrCodeValidationFailed,
// 				Msg:  fmt.Sprintf("error validation: %v", err),
// 			}
// 			return nil, nil, appErr
// 		}

// 		return nil, nil, fmt.Errorf("error validation: %w", err)
// 	}

// 	var onlyUpdateLocation *location_dto.UpdateLocationDTOReq
// 	if attributesToUpdate.Location != nil {
// 		err = utils.Validate.Struct(attributesToUpdate.Location)
// 		if err != nil {
// 			err = fmt.Errorf("%w: validation failed: %v", customerrors.ErrValidationFailed, err.Error())
// 			if errors.Is(err, customerrors.ErrValidationFailed) {
// 				appErr := customerrors.AppError{
// 					Code: customerrors.ErrCodeValidationFailed,
// 					Msg:  fmt.Sprintf("error validation: %v", err),
// 				}
// 				return nil, nil, appErr
// 			}
// 			return nil, nil, fmt.Errorf("error validation: %w", err)
// 		}
// 		onlyUpdateLocation = mappers.OnlyUpdateLocation(&attributesToUpdate)
// 	}

// 	onlyUpdateUser := mappers.OnlyUpdateUser(&attributesToUpdate)

// 	return onlyUpdateUser, onlyUpdateLocation, nil
// }
