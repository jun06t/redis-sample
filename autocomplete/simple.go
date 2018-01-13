package autocomplete

import "github.com/go-redis/redis"

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
