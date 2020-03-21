package server

import (
	"net/http"
	"os"
	"fmt"
	"strings"
)

type Source struct {
	Name	string
	URL		string
}

type Subscriber string

type Article struct {
	Title			string
	Publisher	Source
	URL				string
}

var subscribers = make(map[Source][]Subscriber)

func (src *Source) addSubscriber(sub Subscriber) error {
	subscribers[src] = append(subscribers[src], sub)
}

func (src *Source) removeSubscriber(sub Subscriber) error {

}

func (art *Article) updateSubscribers() error {
	src := art.Publisher
	subs := subscribers[src]
	for _, sub := range subs {
		sub.notify(art)
	}
}

func (sub *Subscriber) notify(art Article) error {
	
}

