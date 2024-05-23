package models

import "time"

type SwipeResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Swipe  `json:"swipe"`
}

type Swipe struct {
	Username   string    `json:"username"`
	IsPremium  bool      `json:"isPremmium"`
	LastSwipe  time.Time `json:"lastSwipe"`
	SwipeCount int       `json:"swipeCount"`
}
