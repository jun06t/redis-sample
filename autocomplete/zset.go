package autocomplete

import (
	"strings"

	"github.com/go-redis/redis"
)

const validChars = "`abcdefghijklmnopqrstqvwxyz{"

func FindPrefixRange(prefix string) (start string, end string) {
	pos := strings.Index(validChars, prefix[len(prefix)-1:])
	suffix := validChars[pos-1 : pos]

	return prefix[:len(prefix)-1] + suffix + "{", prefix + "{"
}

/*
func AutocompleteOnPrefix(client *redis.Client, list string, prefix string) []string {
	start, end := FindPrefixRange(prefix)
	uv4 := uuid.Must(uuid.NewV4())
	client.ZAdd(list, redis.Z{Score: 0, Member: start + uv4}, redis.Z{Score: 0, Member: end + uv4})
	pipe := client.Pipeline()
	sindex := pipeline.zrank(list, start)
}
*/

func JoinList(client *redis.Client, list string, user string) {
	client.ZAdd(list, redis.Z{Score: 0, Member: user})
}

func LeaveList(client *redis.Client, list string, user string) {
	client.ZRem(list, user)
}
