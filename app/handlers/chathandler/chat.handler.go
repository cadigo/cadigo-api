package chathandler

import (
	"cadigo-api/graph/modelgraph"
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/segmentio/ksuid"
)

type Handler struct {
	redisClient     *redis.Client
	messageChannels map[string]chan *modelgraph.Message
	userChannels    map[string]chan string
	mutex           sync.Mutex
}

// redis-cli -u redis://wK8l1q@34.126.162.158:6379

func NewHandler(redisAddr string, redisPass string) *Handler {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

	return &Handler{
		redisClient:     client,
		messageChannels: map[string]chan *modelgraph.Message{},
		userChannels:    map[string]chan string{},
		mutex:           sync.Mutex{},
	}
}

func (s *Handler) PostMessage(ctx context.Context, input *modelgraph.PostMessageInput) (*modelgraph.Message, error) {
	var roomID string

	if input.RoomID == nil {
		room, err := s.getRoom(input.FromUserID, input.ToUserID)
		if err != nil {
			roomID = ksuid.New().String()
			err := s.createRoom(input.FromUserID, input.ToUserID, roomID)
			if err != nil {
				return nil, err
			}
		} else {
			roomID = *room
		}

	} else {
		roomID = *input.RoomID
	}

	// Create message
	m := &modelgraph.Message{
		ToUserID:   input.ToUserID,
		FromUserID: input.FromUserID,
		CreatedAt:  time.Now().UTC(),
		Message:    input.Message,
		RoomID:     roomID,
	}
	mj, _ := json.Marshal(m)
	if err := s.redisClient.RPush(m.RoomID, mj).Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	// Notify new message
	s.mutex.Lock()
	for k, ch := range s.messageChannels {
		res2 := strings.Split(k, ":")
		if len(res2) == 2 {
			if res2[1] == m.RoomID {
				ch <- m
			}
		}
	}

	s.mutex.Unlock()
	return m, nil
}

func (s *Handler) GetMessages(ctx context.Context, input modelgraph.GetMessagesInput) (*modelgraph.GetMessagesType, error) {
	var roomID string

	getRoom, err := s.getRoom(input.FromUserID, input.ToUserID)

	if err != nil {
		roomID = ksuid.New().String()
		err := s.createRoom(input.FromUserID, input.ToUserID, roomID)
		if err != nil {
			return nil, err
		}
	} else {
		roomID = *getRoom
	}

	cmd := s.redisClient.LRange(roomID, 0, -1)
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	messages := []*modelgraph.Message{}
	for _, mj := range res {
		m := &modelgraph.Message{}
		err = json.Unmarshal([]byte(mj), &m)
		messages = append(messages, m)
	}

	if len(messages) == 0 {

	}

	return &modelgraph.GetMessagesType{
		Data:   messages,
		RoomID: roomID,
	}, nil
}

func (s *Handler) GetOnline(ctx context.Context, input modelgraph.GetOnlineInput) ([]*modelgraph.Online, error) {
	cmd := s.redisClient.SMembers("users")
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	d := []*modelgraph.Online{}

	for _, u := range res {
		d = append(d, &modelgraph.Online{
			UserID: u,
		})
	}

	return d, nil
}

func (s *Handler) SubscriptionChat(ctx context.Context, input modelgraph.ChatInput) (<-chan *modelgraph.Message, error) {
	err := s.createUser(input.CurrentUserID)
	if err != nil {
		return nil, err
	}

	// Create new channel for request
	messages := make(chan *modelgraph.Message, 1)
	s.mutex.Lock()
	s.messageChannels[input.CurrentUserID+":"+input.RoomID] = messages
	s.mutex.Unlock()

	// Delete channel when done
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.messageChannels, input.CurrentUserID+":"+input.RoomID)
		s.mutex.Unlock()
	}()

	return messages, nil
}

func (s *Handler) SubscriptionOnline(ctx context.Context, input modelgraph.OnlineInput) (<-chan string, error) {
	err := s.createUser(input.CurrentUserID)
	if err != nil {
		return nil, err
	}

	// Create new channel for request
	users := make(chan string, 1)
	s.mutex.Lock()
	s.userChannels[input.CurrentUserID] = users
	s.mutex.Unlock()

	// // Delete channel when done
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.userChannels, input.CurrentUserID)
		s.mutex.Unlock()
	}()

	return users, nil
}

func (s *Handler) createUser(user string) error {
	// Upsert user
	if err := s.redisClient.SAdd("users", user).Err(); err != nil {
		return err
	}
	// Notify new user joined
	s.mutex.Lock()
	for _, ch := range s.userChannels {
		ch <- user
	}
	s.mutex.Unlock()
	return nil
}

func (s *Handler) createRoom(fromUserID string, toUserID string, roomID string) error {
	if err := s.redisClient.Set(fromUserID+":"+toUserID, roomID, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (s *Handler) getRoom(fromUserID string, toUserID string) (*string, error) {
	val, err := s.redisClient.Get(fromUserID + ":" + toUserID).Result()

	if err != nil {
		val, err = s.redisClient.Get(toUserID + ":" + fromUserID).Result()

		if err != nil {
			return nil, err
		}
	}

	return &val, nil
}
