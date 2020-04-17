package main

import "testing"

var chat int64 = 123456789
var tags []string = []string{"foo", "bar", "spam"}
var result bool

func TestAddSubscribtion(t *testing.T) {
	err := addSubscribtion(chat, tags)
	if err != nil {
		t.Fail()
		t.Logf("Error during insert: %v", err)
	}
}

func TestGetChatByTag(t *testing.T) {
	result = false
	for _, tag := range tags {
		chats, err := getChatByTag(tag)
		if err != nil {
			t.Fatalf("Error %v during test", err)
		}

		for _, c := range chats {
			t.Log(c)
			if c == chat {
				result = true
				break
			}
		}
	}
	if !result {
		t.FailNow()
	}
}

// func TestAlwaisFail(t *testing.T) {
// 	t.Fatal("Check")
// }
func TestDeleteSubscriber(t *testing.T) {
	chats := []int64{chat}
	err := deleteSubscriber(chats)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
