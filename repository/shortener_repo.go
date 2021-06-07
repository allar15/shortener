package repository

import (

	"github.com/ssdb/gossdb/ssdb"
)

type ShortenerRepo struct {
	client  *ssdb.Client
}


func NewShortenerRepo(client *ssdb.Client) *ShortenerRepo{
	return &ShortenerRepo{client}
}


func (s *ShortenerRepo) SetPair(long, short string) (err error){
	_, err = s.client.Do("multi_set", long, short, short, long)
	return
}

func (s *ShortenerRepo) GetValue(key string) (value string, err error){
	values, err  := s.client.Do("get", key)
	if values[0] == "not_found"{
		return
	}else{
		value = values[1]
	}
	return
}
