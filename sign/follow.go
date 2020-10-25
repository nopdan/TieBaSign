package sign

import (
	"encoding/json"
	"fmt"
)

func GetFollow(BDUSS string) ([]string, error) {
	cookie := "BDUSS=" + BDUSS
	b, err := httpget(LIKEURL, cookie)
	if err != nil {
		return nil, fmt.Errorf("GetFollow: %w", err)
	}
	var like Like
	fmt.Println(string(b))
	err = json.Unmarshal(b, &like)
	if err != nil {
		return nil, fmt.Errorf("GetFollow: %w", err)
	}
	list := make([]string, 0, len(like.Data.LikeForum))
	for _, v := range like.Data.LikeForum {
		if v.IsSign != 1 {
			list = append(list, v.ForumName)
		}
	}
	return list, nil
}
