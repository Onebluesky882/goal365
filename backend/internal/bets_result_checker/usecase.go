package betresultchecker

import "fmt"

/*
	Match Winner

Home/Away
Second Half Winner
Asian Handicap
Goals Over/Under
Goals Over/Under First Half
HT/FT Double
Exact Score
Highest Scoring Half
Correct Score - First Half
Double Chance
First Half Winner
Asian Handicap First Half
Double Chance - First Half
Odd/Even
Draw No Bet (1st Half)
Draw No Bet (2nd Half)
*/

// Strategy Pattern
// Market Name  ──▶ Strategy (Function) ──▶ Result
// Table-driven Design
// “This uses a function-based Strategy Pattern with a table-driven dispatcher.”
/*
จำ 3 ข้อ
- code ต้อง scale ได้ เพิ่มหรือลดได้ในอนาคตไม่พัง
- test unit ได้ง่าย
- เพิ่ม market ใหม่ = เพิ่ม function

*/

//  Function Registry — map[string]BetSettleFunc 
type BetSettleFunc func(ctx MatchContext, betValue any) SettleResult


var settleHandler = map[string]BetSettleFunc{
	"Home/Away": func(ctx MatchContext, betValue any) SettleResult {
		return settleHomeAway(ctx, betValue.(string))
	},
}

func RefactorBetSettle(name string, ctx MatchContext, bet any) (SettleResult, error) {

	handler, ok := settleHandler[name]
	if !ok {
		return "", fmt.Errorf("un suport market %s", name)
	}
	return handler(ctx, bet), nil
}

const (
	Win  SettleResult = "WIN"
	Lose SettleResult = "LOSE"
	Void SettleResult = "Void"
)

// ------- *-* ----------
func settleHomeAway(ctx MatchContext, bet string) SettleResult {
	switch bet {
	case "Hone":
		if ctx.HomeScore > ctx.AwayScore {
			return Win
		}
	case "Away":
		if ctx.AwayScore > ctx.HomeScore {
			return Win

		}
	case "Draw":
		if ctx.HomeScore == ctx.AwayScore {
			return Win
		}
	}
	return Lose
}
