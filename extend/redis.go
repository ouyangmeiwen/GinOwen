package extend

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func StringExample(client *redis.Client) {
	// 设置键值
	err := client.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值
	val, err := client.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1:", val)
}
func HashExample(client *redis.Client) {
	// 设置哈希表字段
	err := client.HSet(ctx, "user:1000", "name", "John", "age", "30").Err()
	if err != nil {
		panic(err)
	}

	// 获取哈希表字段
	name, err := client.HGet(ctx, "user:1000", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", name)

	// 获取所有字段
	user, err := client.HGetAll(ctx, "user:1000").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("User:", user)
}
func ListExample(client *redis.Client) {
	// 从列表左侧插入元素
	err := client.LPush(ctx, "mylist", "item1", "item2", "item3").Err()
	if err != nil {
		panic(err)
	}

	// 从列表右侧弹出元素
	item, err := client.RPop(ctx, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Popped item:", item)

	// 获取列表中的所有元素
	items, err := client.LRange(ctx, "mylist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("List items:", items)
}
func SetExample(client *redis.Client) {
	// 添加元素到集合
	err := client.SAdd(ctx, "myset", "member1", "member2", "member3").Err()
	if err != nil {
		panic(err)
	}

	// 判断元素是否在集合中
	isMember, err := client.SIsMember(ctx, "myset", "member1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is member1 in myset?", isMember)

	// 获取集合中的所有元素
	members, err := client.SMembers(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set members:", members)
}
func SortedSetExample(client *redis.Client) {
	// 添加元素到有序集合
	err := client.ZAdd(ctx, "leaderboard", redis.Z{Score: 100, Member: "player1"}, redis.Z{Score: 200, Member: "player2"}).Err()
	if err != nil {
		panic(err)
	}

	// 获取有序集合中的元素，按分数从低到高排序
	players, err := client.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Leaderboard:")
	for _, p := range players {
		fmt.Printf("  %s: %f\n", p.Member, p.Score)
	}

	// 获取有序集合中的元素，按分数从高到低排序
	topPlayers, err := client.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Top players:")
	for _, p := range topPlayers {
		fmt.Printf("  %s: %f\n", p.Member, p.Score)
	}
}
func HyperLogLogExample(client *redis.Client) {
	// 添加元素到 HyperLogLog
	err := client.PFAdd(ctx, "hll", "element1", "element2", "element3").Err()
	if err != nil {
		panic(err)
	}

	// 获取 HyperLogLog 的基数估计
	count, err := client.PFCount(ctx, "hll").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HyperLogLog count:", count)
}
func BitmapExample(client *redis.Client) {
	// 设置某个位置的位
	err := client.SetBit(ctx, "user_signin", 5, 1).Err() // 用户在第5天签到
	if err != nil {
		panic(err)
	}

	// 获取某个位置的位
	bit, err := client.GetBit(ctx, "user_signin", 5).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("User signed in on day 5:", bit == 1)

	// 统计位图中值为1的总数（签到天数）
	count, err := client.BitCount(ctx, "user_signin", nil).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Total signed-in days:", count)
}
