package main

import (
	"aggregator/internal/database"
	"aggregator/internal/rss"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAgg(s *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("coudn't fetch feed %w", err)
	}
	fmt.Println("Feed:", feed)
	return nil
}

func handlerAddFedd(s *State, cmd Command, user database.User) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.Current_user_name)
	if err != nil {
		return err
	}
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	current_user, err := s.Db.GetUser(context.Background(), s.Cfg.Current_user_name)
	if err != nil {
		return err
	}
	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    current_user.ID,
	})

	if err != nil {
		return err
	}
	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")
	return nil
}
func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}

func printFeeds(feeds []database.GetFeedsRow, s *State) error {
	for _, item := range feeds {
		fmt.Println(item.Name)
		user, err := s.Db.GetUserById(context.Background(), item.UserID)
		if err != nil {
			return err
		}
		fmt.Println(user.Name)
	}
	return nil
}
func handlerGetFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("don't get users %w", err)
	}
	err = printFeeds(feeds, s)
	if err != nil {
		return err
	}
	return nil
}
