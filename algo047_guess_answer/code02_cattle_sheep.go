package algo047_guess_answer

const (
	FirstMove  = "先手"
	SecondMove = "后手"
)

func WhoWinAwesome(n int) string {
	if n%5 == 0 || n%5 == 2 {
		return SecondMove
	}
	return FirstMove
}

func WhoWin(n int) string {
	switch n {
	case 0, 2:
		return SecondMove
	case 1, 3, 4:
		return FirstMove
	}
	// 进行博弈
	for want := 1; want <= n; want *= 4 {
		if WhoWin(n-want) == SecondMove {
			return FirstMove
		}
	}
	return SecondMove
}
