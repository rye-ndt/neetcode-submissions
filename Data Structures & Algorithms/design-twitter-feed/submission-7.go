type tweet struct {
	id int
	poster int
	time int
}

type tweetH []*tweet

func (h tweetH) Len() int { return len(h) }
func (h tweetH) Less(i, j int) bool { return h[i].time > h[j].time }
func (h tweetH) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *tweetH) Push(v any) { *h = append(*h, v.(*tweet)) }
func (h *tweetH) Pop() any {
	clone := *h 
	last := clone[len(clone)-1]
	*h = clone[:len(clone)-1]
	return last
}

type Twitter struct {
	clock int
	tweets *tweetH
	followings map[int]map[int]bool
}

func Constructor() Twitter {
	h := &tweetH{}
	heap.Init(h)

	return Twitter{
		tweets: h,
		followings: map[int]map[int]bool{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	t := &tweet{
		time: this.clock+1,
		id: tweetId,
		poster: userId, 
	}

	this.clock++

	heap.Push(this.tweets, t)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	redundant := []*tweet{}
	result := []int{}
	
	for i := 0; i < 10 && this.tweets.Len() > 0; i++ {
		t := heap.Pop(this.tweets).(*tweet)

		_, follow := this.followings[userId][t.poster]
		if follow || t.poster == userId {
			result = append(result, t.id)
		} else { i-- }

		redundant = append(redundant, t)
	}

	for _, r := range redundant {
		heap.Push(this.tweets, r)
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int)  {	
	if _, found := this.followings[followerId]; !found {
		this.followings[followerId] = map[int]bool{}
	}

	this.followings[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	delete(this.followings[followerId], followeeId)
}