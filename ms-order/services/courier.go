package services

// func (o *OrderService) AddCourierInfo(ctx context.Context, in *pb.AddCourierInfoRequest) (*pb.CourierResponse, error) {
// 	// get resi

// 	info, err := helpers.TrackPackage(in.Awb, strings.ToLower(in.Company))
// 	if err != nil {
// 		return nil, err
// 	}

// 	input := model.Courier{
// 		AWB:         info.Data.Summary.AWB,
// 		Company:     info.Data.Summary.Courier,
// 		Status:      info.Data.Summary.Status,
// 		Date:        info.Data.Summary.Date,
// 		Fee:         info.Data.Summary.Amount,
// 		Origin:      info.Data.Detail.Origin,
// 		Destination: info.Data.Detail.Destination,
// 		History:     info.Data.History,
// 	}

// 	res, err := o.repo.AddCourierInfo(&input)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, nil
// }
