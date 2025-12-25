export type Root = Root2[]

export interface Root2 {
  FixtureIDs: number[]
  Items: Item[]
}

export interface Item {
  FixtureID: number
  Predictions?: Prediction[]
  Bookmaker: any
  Team: string
  Result: string
  Picked: boolean
}

export interface Prediction {
  predictions: Predictions
  league: League
  teams: Teams
  comparison: Comparison
  h2h: H2h2[]
}

export interface Predictions {
  winner: Winner
  win_or_draw: boolean
  under_over?: string
  goals: Goals
  advice: string
  percent: Percent
}

export interface Winner {
  id: number
  name: string
  comment: string
}

export interface Goals {
  home: string
  away: string
}

export interface Percent {
  home: string
  draw: string
  away: string
}

export interface League {
  id: number
  name: string
  country: string
  logo: string
  flag: string
  season: number
}

export interface Teams {
  home: Home
  away: Away
}

export interface Home {
  id: number
  name: string
  logo: string
  last_5: Last5
  league: League2
}

export interface Last5 {
  played: number
  form: string
  att: string
  def: string
  goals: Goals2
}

export interface Goals2 {
  for: For
  against: Against
}

export interface For {
  total: number
  average: string
}

export interface Against {
  total: number
  average: string
}

export interface League2 {
  form: string
  fixtures: Fixtures
  goals: Goals3
  biggest: Biggest
  clean_sheet: CleanSheet
  failed_to_score: FailedToScore
  penalty: Penalty
  lineups: Lineup[]
  cards: Cards
}

export interface Fixtures {
  played: Played
  wins: Wins
  draws: Draws
  loses: Loses
}

export interface Played {
  home: number
  away: number
  total: number
}

export interface Wins {
  home: number
  away: number
  total: number
}

export interface Draws {
  home: number
  away: number
  total: number
}

export interface Loses {
  home: number
  away: number
  total: number
}

export interface Goals3 {
  for: For2
  against: Against2
}

export interface For2 {
  total: Total
  average: Average
  minute: Minute
  under_over: UnderOver
}

export interface Total {
  home: number
  away: number
  total: number
}

export interface Average {
  home: string
  away: string
  total: string
}

export interface Minute {
  "0-15": N015
  "16-30": N1630
  "31-45": N3145
  "46-60": N4660
  "61-75": N6175
  "76-90": N7690
  "91-105": N91105
  "106-120": N106120
}

export interface N015 {
  total?: number
  percentage?: string
}

export interface N1630 {
  total?: number
  percentage?: string
}

export interface N3145 {
  total?: number
  percentage?: string
}

export interface N4660 {
  total?: number
  percentage?: string
}

export interface N6175 {
  total?: number
  percentage?: string
}

export interface N7690 {
  total: number
  percentage: string
}

export interface N91105 {
  total: any
  percentage: any
}

export interface N106120 {
  total: any
  percentage: any
}

export interface UnderOver {
  "0.5": N05
  "1.5": N15
  "2.5": N25
  "3.5": N35
  "4.5": N45
}

export interface N05 {
  over: number
  under: number
}

export interface N15 {
  over: number
  under: number
}

export interface N25 {
  over: number
  under: number
}

export interface N35 {
  over: number
  under: number
}

export interface N45 {
  over: number
  under: number
}

export interface Against2 {
  total: Total2
  average: Average2
  minute: Minute2
  under_over: UnderOver2
}

export interface Total2 {
  home: number
  away: number
  total: number
}

export interface Average2 {
  home: string
  away: string
  total: string
}

export interface Minute2 {
  "0-15": N0152
  "16-30": N16302
  "31-45": N31452
  "46-60": N46602
  "61-75": N61752
  "76-90": N76902
  "91-105": N911052
  "106-120": N1061202
}

export interface N0152 {
  total?: number
  percentage?: string
}

export interface N16302 {
  total?: number
  percentage?: string
}

export interface N31452 {
  total?: number
  percentage?: string
}

export interface N46602 {
  total?: number
  percentage?: string
}

export interface N61752 {
  total?: number
  percentage?: string
}

export interface N76902 {
  total?: number
  percentage?: string
}

export interface N911052 {
  total: any
  percentage: any
}

export interface N1061202 {
  total: any
  percentage: any
}

export interface UnderOver2 {
  "0.5": N052
  "1.5": N152
  "2.5": N252
  "3.5": N352
  "4.5": N452
}

export interface N052 {
  over: number
  under: number
}

export interface N152 {
  over: number
  under: number
}

export interface N252 {
  over: number
  under: number
}

export interface N352 {
  over: number
  under: number
}

export interface N452 {
  over: number
  under: number
}

export interface Biggest {
  streak: Streak
  wins: Wins2
  loses: Loses2
  goals: Goals4
}

export interface Streak {
  wins: number
  draws: number
  loses: number
}

export interface Wins2 {
  home: string
  away: string
}

export interface Loses2 {
  home: string
  away: string
}

export interface Goals4 {
  for: For3
  against: Against3
}

export interface For3 {
  home: number
  away: number
}

export interface Against3 {
  home: number
  away: number
}

export interface CleanSheet {
  home: number
  away: number
  total: number
}

export interface FailedToScore {
  home: number
  away: number
  total: number
}

export interface Penalty {
  scored: Scored
  missed: Missed
  total: number
}

export interface Scored {
  total: number
  percentage: string
}

export interface Missed {
  total: number
  percentage: string
}

export interface Lineup {
  formation: string
  played: number
}

export interface Cards {
  yellow: Yellow
  red: Red
}

export interface Yellow {
  "0-15": N0153
  "16-30": N16303
  "31-45": N31453
  "46-60": N46603
  "61-75": N61753
  "76-90": N76903
  "91-105": N911053
  "106-120": N1061203
}

export interface N0153 {
  total?: number
  percentage?: string
}

export interface N16303 {
  total?: number
  percentage?: string
}

export interface N31453 {
  total?: number
  percentage?: string
}

export interface N46603 {
  total?: number
  percentage?: string
}

export interface N61753 {
  total?: number
  percentage?: string
}

export interface N76903 {
  total?: number
  percentage?: string
}

export interface N911053 {
  total?: number
  percentage?: string
}

export interface N1061203 {
  total: any
  percentage: any
}

export interface Red {
  "0-15": N0154
  "16-30": N16304
  "31-45": N31454
  "46-60": N46604
  "61-75": N61754
  "76-90": N76904
  "91-105": N911054
  "106-120": N1061204
}

export interface N0154 {
  total: any
  percentage: any
}

export interface N16304 {
  total: any
  percentage: any
}

export interface N31454 {
  total?: number
  percentage?: string
}

export interface N46604 {
  total?: number
  percentage?: string
}

export interface N61754 {
  total?: number
  percentage?: string
}

export interface N76904 {
  total?: number
  percentage?: string
}

export interface N911054 {
  total?: number
  percentage?: string
}

export interface N1061204 {
  total: any
  percentage: any
}

export interface Away {
  id: number
  name: string
  logo: string
  last_5: Last52
  league: League3
}

export interface Last52 {
  played: number
  form: string
  att: string
  def: string
  goals: Goals5
}

export interface Goals5 {
  for: For4
  against: Against4
}

export interface For4 {
  total: number
  average: string
}

export interface Against4 {
  total: number
  average: string
}

export interface League3 {
  form: string
  fixtures: Fixtures2
  goals: Goals6
  biggest: Biggest2
  clean_sheet: CleanSheet2
  failed_to_score: FailedToScore2
  penalty: Penalty2
  lineups: Lineup2[]
  cards: Cards2
}

export interface Fixtures2 {
  played: Played2
  wins: Wins3
  draws: Draws2
  loses: Loses3
}

export interface Played2 {
  home: number
  away: number
  total: number
}

export interface Wins3 {
  home: number
  away: number
  total: number
}

export interface Draws2 {
  home: number
  away: number
  total: number
}

export interface Loses3 {
  home: number
  away: number
  total: number
}

export interface Goals6 {
  for: For5
  against: Against5
}

export interface For5 {
  total: Total3
  average: Average3
  minute: Minute3
  under_over: UnderOver3
}

export interface Total3 {
  home: number
  away: number
  total: number
}

export interface Average3 {
  home: string
  away: string
  total: string
}

export interface Minute3 {
  "0-15": N0155
  "16-30": N16305
  "31-45": N31455
  "46-60": N46605
  "61-75": N61755
  "76-90": N76905
  "91-105": N911055
  "106-120": N1061205
}

export interface N0155 {
  total?: number
  percentage?: string
}

export interface N16305 {
  total?: number
  percentage?: string
}

export interface N31455 {
  total?: number
  percentage?: string
}

export interface N46605 {
  total?: number
  percentage?: string
}

export interface N61755 {
  total?: number
  percentage?: string
}

export interface N76905 {
  total?: number
  percentage?: string
}

export interface N911055 {
  total?: number
  percentage?: string
}

export interface N1061205 {
  total: any
  percentage: any
}

export interface UnderOver3 {
  "0.5": N053
  "1.5": N153
  "2.5": N253
  "3.5": N353
  "4.5": N453
}

export interface N053 {
  over: number
  under: number
}

export interface N153 {
  over: number
  under: number
}

export interface N253 {
  over: number
  under: number
}

export interface N353 {
  over: number
  under: number
}

export interface N453 {
  over: number
  under: number
}

export interface Against5 {
  total: Total4
  average: Average4
  minute: Minute4
  under_over: UnderOver4
}

export interface Total4 {
  home: number
  away: number
  total: number
}

export interface Average4 {
  home: string
  away: string
  total: string
}

export interface Minute4 {
  "0-15": N0156
  "16-30": N16306
  "31-45": N31456
  "46-60": N46606
  "61-75": N61756
  "76-90": N76906
  "91-105": N911056
  "106-120": N1061206
}

export interface N0156 {
  total?: number
  percentage?: string
}

export interface N16306 {
  total?: number
  percentage?: string
}

export interface N31456 {
  total?: number
  percentage?: string
}

export interface N46606 {
  total?: number
  percentage?: string
}

export interface N61756 {
  total?: number
  percentage?: string
}

export interface N76906 {
  total?: number
  percentage?: string
}

export interface N911056 {
  total: any
  percentage: any
}

export interface N1061206 {
  total: any
  percentage: any
}

export interface UnderOver4 {
  "0.5": N054
  "1.5": N154
  "2.5": N254
  "3.5": N354
  "4.5": N454
}

export interface N054 {
  over: number
  under: number
}

export interface N154 {
  over: number
  under: number
}

export interface N254 {
  over: number
  under: number
}

export interface N354 {
  over: number
  under: number
}

export interface N454 {
  over: number
  under: number
}

export interface Biggest2 {
  streak: Streak2
  wins: Wins4
  loses: Loses4
  goals: Goals7
}

export interface Streak2 {
  wins: number
  draws: number
  loses: number
}

export interface Wins4 {
  home: string
  away: string
}

export interface Loses4 {
  home: string
  away: string
}

export interface Goals7 {
  for: For6
  against: Against6
}

export interface For6 {
  home: number
  away: number
}

export interface Against6 {
  home: number
  away: number
}

export interface CleanSheet2 {
  home: number
  away: number
  total: number
}

export interface FailedToScore2 {
  home: number
  away: number
  total: number
}

export interface Penalty2 {
  scored: Scored2
  missed: Missed2
  total: number
}

export interface Scored2 {
  total: number
  percentage: string
}

export interface Missed2 {
  total: number
  percentage: string
}

export interface Lineup2 {
  formation: string
  played: number
}

export interface Cards2 {
  yellow: Yellow2
  red: Red2
}

export interface Yellow2 {
  "0-15": N0157
  "16-30": N16307
  "31-45": N31457
  "46-60": N46607
  "61-75": N61757
  "76-90": N76907
  "91-105": N911057
  "106-120": N1061207
}

export interface N0157 {
  total?: number
  percentage?: string
}

export interface N16307 {
  total?: number
  percentage?: string
}

export interface N31457 {
  total?: number
  percentage?: string
}

export interface N46607 {
  total?: number
  percentage?: string
}

export interface N61757 {
  total?: number
  percentage?: string
}

export interface N76907 {
  total?: number
  percentage?: string
}

export interface N911057 {
  total?: number
  percentage?: string
}

export interface N1061207 {
  total: any
  percentage: any
}

export interface Red2 {
  "0-15": N0158
  "16-30": N16308
  "31-45": N31458
  "46-60": N46608
  "61-75": N61758
  "76-90": N76908
  "91-105": N911058
  "106-120": N1061208
}

export interface N0158 {
  total?: number
  percentage?: string
}

export interface N16308 {
  total?: number
  percentage?: string
}

export interface N31458 {
  total?: number
  percentage?: string
}

export interface N46608 {
  total?: number
  percentage?: string
}

export interface N61758 {
  total?: number
  percentage?: string
}

export interface N76908 {
  total?: number
  percentage?: string
}

export interface N911058 {
  total?: number
  percentage?: string
}

export interface N1061208 {
  total: any
  percentage: any
}

export interface Comparison {
  form: Form
  att: Att
  def: Def
  poisson_distribution: PoissonDistribution
  h2h: H2h
  goals: Goals8
  total: Total5
}

export interface Form {
  home: string
  away: string
}

export interface Att {
  home: string
  away: string
}

export interface Def {
  home: string
  away: string
}

export interface PoissonDistribution {
  home: string
  away: string
}

export interface H2h {
  home: string
  away: string
}

export interface Goals8 {
  home: string
  away: string
}

export interface Total5 {
  home: string
  away: string
}

export interface H2h2 {
  fixture: Fixture
  league: League4
  teams: Teams2
  goals: Goals9
  score: Score
}

export interface Fixture {
  id: number
  referee?: string
  timezone: string
  date: string
  timestamp: number
  periods: Periods
  venue: Venue
  status: Status
}

export interface Periods {
  first: number
  second: number
}

export interface Venue {
  name?: string
  city: string
  id?: number
}

export interface Status {
  long: string
  short: string
  elapsed: number
  extra?: number
}

export interface League4 {
  id: number
  name: string
  country: string
  logo: string
  flag?: string
  season: number
  round: string
  standings: boolean
}

export interface Teams2 {
  home: Home2
  away: Away2
}

export interface Home2 {
  id: number
  name: string
  logo: string
  winner: boolean
}

export interface Away2 {
  id: number
  name: string
  logo: string
  winner: boolean
}

export interface Goals9 {
  home: number
  away: number
}

export interface Score {
  halftime: Halftime
  fulltime: Fulltime
  extratime: Extratime
  penalty: Penalty3
}

export interface Halftime {
  home: number
  away: number
}

export interface Fulltime {
  home: number
  away: number
}

export interface Extratime {
  home?: number
  away?: number
}

export interface Penalty3 {
  home: any
  away: any
}
