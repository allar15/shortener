package logic

import (
	"context"
	"math/rand"
	"time"
	"shortener/repository"
)

type ShortenerLogic struct {
	repo repository.ShortenerRepo
}

func NewClientLogic(r repository.ShortenerRepo) *ShortenerLogic{
	return &ShortenerLogic{
		repo: r,
		}
	}

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func (l *ShortenerLogic) CreatePair(ctx context.Context, long, addr string) (short string, err error){
	rand.Seed(time.Now().UnixNano())
	result, err := l.repo.GetValue(long)
	if err != nil{
		return
	}
	if result == ""{
		result = long
		for result == long{
			short = createKey(8)
			result, err = l.repo.GetValue(short) // проверка на занятость укороченного варианта URL
		}
		err = l.repo.SetPair(long, short)
		if err != nil {
			return
		}
	}else{
		short, err = l.repo.GetValue(long)
	}
	return
}

func (l *ShortenerLogic) GetPair(ctx context.Context, url string) (parent string, err error){
	parent, err = l.repo.GetValue(url)
	return
}

func createKey(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = symbols[rand.Intn(len(symbols))]
    }
    return string(b)
}


