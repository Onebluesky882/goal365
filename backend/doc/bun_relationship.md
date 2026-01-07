✅ Parent (One)
type MyAnalytics struct {
bun.BaseModel `bun:"table:my_analytics,alias:ma"`

    ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`

    // 🔥 One → Many
    NaWinTaTips []*NaWinTatips `bun:"rel:has-many,join:id=tips_analytics_id"`

}

✅ Child (Many)

// models/nawinta.go
type NaWinTatips struct {
bun.BaseModel `bun:"table:nawinta,alias:nw"`

    ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`

    TipsAnalyticsID *uuid.UUID `bun:"tips_analytics_id,type:uuid,nullzero"`

    // 🔥 Many → One
    TipsAnalytics *MyAnalytics `bun:"rel:belongs-to,join:tips_analytics_id=id"`

    Payload json.RawMessage `bun:"type:jsonb,notnull"`

}

Query Example (Preload children)
analytics := new(MyAnalytics)

err := db.NewSelect().
Model(analytics).
Where("ma.id = ?", id).
Relation("NaWinTaTips").
Scan(ctx)

----------------------------- _-_ --------------------------------
Many-to-Many (Parent ↔ Parent)

✅ Fixture (Many-to-Many)
type Fixture struct {
bun.BaseModel `bun:"table:fixtures,alias:f"`

    ID int `bun:",pk"`

    Leagues []*League `bun:"m2m:fixture_leagues,join:Fixture=League"`

}

3️⃣ Parent / Child Rule (จำง่ายมาก)

✅ Parent
• ไม่มี FK
• มี slice []\*Child
• ใช้ has-many

✅ Child
• มี FK
• มี pointer \*Parent
• ใช้ belongs-to

4️⃣ Cheat Sheet (จำอันนี้พอ)
Relation Parent tag Child tag
One → Many has-many belongs-to
Optional FK slice \*uuid.UUID + nullzero
Many ↔ Many m2m:table m2m:table
