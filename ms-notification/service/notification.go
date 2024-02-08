package services

import (
	"context"
	"errors"
	"ms-notification/model"
	"ms-notification/pb"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Service) CreateNotification(ctx context.Context, in *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	notif := &model.Notification{
		ReceiverID:  in.ReceiverId,
		Subject:     in.Subject,
		Description: in.Description,
		Status:      model.NOTIFICATION_UNREAD,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Notif.Create(notif)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	return &pb.CreateNotificationResponse{
		Message: "success",
	}, nil
}

func (s Service) GetAllNotification(ctx context.Context, in *pb.GetAllNotificationRequest) (*pb.GetAllNotificationResponse, error) {
	notifs, err := s.repo.Notif.GetAll(in.ReceiverId)
	if err != nil {
		return &pb.GetAllNotificationResponse{}, ErrInternal(err, s.log)
	}

	var response = &pb.GetAllNotificationResponse{}
	for _, n := range notifs {
		response.Notifications = append(response.Notifications, &pb.Notification{
			Id:          n.ID.Hex(),
			Subject:     n.Subject,
			Description: n.Description,
			Status:      n.Status,
			CreatedAt:   timestamppb.New(n.CreatedAt),
		})
	}

	return response, nil
}

func (s Service) UpdateNotification(ctx context.Context, in *pb.UpdateNotificationRequest) (*pb.UpdateNotificationResponse, error) {
	notif, err := s.repo.Notif.GetByID(in.Id, in.ReceiverId)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrNotFound("no matching record found")
		default:
			return nil, ErrInternal(err, s.log)
		}
	}

	notif.Status = in.Status

	err = s.repo.Notif.Update(notif)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	return &pb.UpdateNotificationResponse{
		Message: "success",
	}, nil
}

func (s Service) MarkAllAsRead(ctx context.Context, in *pb.MarkAllAsReadRequest) (*pb.MarkAllAsReadResponse, error) {
	notifs, err := s.repo.Notif.GetAllUnread(in.ReceiverId)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	for _, n := range notifs {
		n.Status = model.NOTIFICATION_READ
		err := s.repo.Notif.Update(n)
		if err != nil {
			return nil, ErrInternal(err, s.log)
		}
	}

	return &pb.MarkAllAsReadResponse{
		Message: "success",
	}, nil
}

func (s Service) MarkAsRead(ctx context.Context, in *pb.MarkAsReadRequest) (*pb.MarkAsReadResponse, error) {
	notif, err := s.repo.Notif.GetByID(in.Id, in.ReceiverId)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	notif.Status = model.NOTIFICATION_READ
	err = s.repo.Notif.Update(notif)
	if err != nil {
		return nil, ErrInternal(err, s.log)
	}

	return &pb.MarkAsReadResponse{
		Message: "success",
	}, nil
}
