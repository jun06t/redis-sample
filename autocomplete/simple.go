package autocomplete

import (
	"strings"

	"github.com/go-redis/redis"
)

func AddUpdateContact(client *redis.Client, list string, contact string) error {
	pipe := client.Pipeline()
	pipe.LRem(list, 0, contact)
	pipe.LPush(list, contact)
	pipe.LTrim(list, 0, 99)
	_, err := pipe.Exec()

	return err
}

func RemoveContact(client *redis.Client, list string, contact string) error {
	pipe := client.Pipeline()
	pipe.LRem(list, 0, contact)
	_, err := pipe.Exec()

	return err
}

func FetchAutocompleteList(client *redis.Client, list string, prefix string) []string {
	candidates := client.LRange(list, 0, -1).Val()
	result := make([]string, 0, len(candidates))
	for i := range candidates {
		if strings.HasPrefix(candidates[i], prefix) {
			result = append(result, candidates[i])
		}
	}

	return result
}
