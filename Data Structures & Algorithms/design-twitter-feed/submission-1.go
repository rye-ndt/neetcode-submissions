type Tweet struct {
	Id int 
	User int
	Time int
}

type H []Tweet

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i].Time > h[j].Time
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x any) {
	*h = append(*h, x.(Tweet))
}

func (h *H) Pop() any {
	old := *h
	last := old[len(old)-1]
	*h = old[:len(old)-1]

	return last
}

type Twitter struct {
	Tweets *H
	UserToFollowing map[int]map[int]bool
	Time int
}


func Constructor() Twitter {
	h := &H{}
	heap.Init(h)

    return Twitter{
		Tweets: h,
		UserToFollowing: map[int]map[int]bool{},
		Time: 0,
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int) {
	t := Tweet{
		Id: tweetId,
		User: userId,
		Time: this.Time,
	}

    heap.Push(this.Tweets, t)

	this.Time++

	fmt.Println("heap: ", this.Tweets)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    allowList, found := this.UserToFollowing[userId]
	if !found {
		allowList = map[int]bool{}
	}

	allowList[userId] = true

	data := []int{}

	clone := &H{}
	heap.Init(clone)

	for _, item := range *this.Tweets {
		heap.Push(clone, item)
	}

	counter := 0

	for counter < 10 && clone.Len() > 0 {
		item := heap.Pop(clone).(Tweet)

		if exists := allowList[item.User]; exists {
			data = append(data, item.Id)
			counter++
		}
	}

	return data
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if this.UserToFollowing[followerId] == nil {
		this.UserToFollowing[followerId] = map[int]bool{}
	}

	this.UserToFollowing[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.UserToFollowing[followerId] == nil {
		this.UserToFollowing[followerId] = map[int]bool{}
	}
	
	this.UserToFollowing[followerId][followeeId] = false
}
