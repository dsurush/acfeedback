package services

import (
	"acfeedback/pkg/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type FeedbackSvc struct {
	pool *pgxpool.Pool
}

func NewFeedbackSvc(pool *pgxpool.Pool) *FeedbackSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil")) // <- be accurate
	}
	return &FeedbackSvc{pool: pool}
}

func (receiver *FeedbackSvc) DbInit() (err error) {
	ddls := []string{feedBacksDDL}
	for _, ddl := range ddls {
		_, err := receiver.pool.Exec(context.Background(), ddl)
		if err != nil {
			log.Printf("err, %e\n", err)
			return err
		}
	}
	return nil
}

func (receiver *FeedbackSvc) Save(feedback models.FeedBack) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), userSaveDml, feedback.Feedback, feedback.Who, feedback.Whom)
	if err != nil {
		log.Print("can't add to db")
		return err
	}
	return
}

func (receiver *FeedbackSvc) GetFeedbackById(id string) (FeedBack models.FeedBack, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	err = conn.QueryRow(context.Background(), getFeedbackByIdDml, id).Scan(
		&FeedBack.Id,
		&FeedBack.Feedback,
		&FeedBack.Who,
		&FeedBack.Whom,
	//	&FeedBack.FeedbackTime,
		&FeedBack.Remove,
	)
	if err != nil {
		fmt.Printf("can't read from db %e", err)
		return
	}
	return FeedBack, nil
}

func (receiver *FeedbackSvc) RemoveFeedbackByID(id string) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e\n", err)
		return
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), removeFeedbackByIdDml, id)
	if err != nil {
		fmt.Printf("can't update user %e\n", err)
		return
	}
	return nil
}
func (receiver *FeedbackSvc) GetFeedbackList() (feedbacks []models.FeedBack, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), getFeedBackListDml)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		feedback := models.FeedBack{}
		rows.Scan(
			&feedback.Id,
			&feedback.Feedback,
			&feedback.Who,
			&feedback.Whom,
			//&FeedBack.FeedbackTime,
			&feedback.Remove,
		)
		feedbacks = append(feedbacks, feedback)
	}
	if rows.Err() != nil {
		log.Printf("rows err %s", err)
		return nil, rows.Err()
	}
	return
}