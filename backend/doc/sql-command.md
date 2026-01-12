```sql


SELECT * FROM "my-analytics" WHERE fixture_id = '1382788';
SELECT * from "my-analytics" WHERE team_home = 'Merthyr Town';

SELECT * FROM "my-analytics" WHERE id = '4518919a-ea04-4019-975c-2eefaa17595e';

SELECT COUNT (*) FROM "my-bets"

CREATE TABLE sportsbooks (
	id uuid DEFAULT gen_random_uuid(),
	match_date date NOT NULL,
	fx_id int NOT NULL,
	home text NOT NULL,
	away text NOT NULL,
	league text NOT NULL,
	country text NOT NULL,
	bet jsonb,
	created_at timestamptz DEFAULT now(),
	updated_at timestamptz,

	PRIMARY KEY (match_date, fx_id)
) PARTITION BY RANGE (match_date);





SELECT
  id,
  transaction_id::text,
  fixture_id,
  market,
  selection,
  odds,
  amount,
  status
FROM bets
WHERE transaction_id = 'd42151a6-a958-4a9a-880d-8bbd47c1e87d';


-- //////////////////////////////////////////////  cascade
--   ALTER TABLE bets ลบ
 DROP CONSTRAINT bets_transaction_id_fkey;


-- ถ้าลบ transaction ก็ลบ bets auto
ALTER TABLE bets
ADD CONSTRAINT bets_transaction_id_fkey
FOREIGN KEY (transaction_id)
REFERENCES transactions(id)
ON DELETE CASCADE;


-- make index
-- CREATE INDEX idx_bets_transaction_created_at
-- ON bets (transaction_id, created_at);

-- ตรวจสอบ index working
-- SELECT indexname, indexdef
-- FROM pg_indexes
-- WHERE tablename = 'bets';


-- EXPLAIN ANALYZE
-- SELECT *
-- FROM bets
-- WHERE transaction_id = '948b9da3-5f57-4a9e-9bc6-3b794acda519'
-- ORDER BY created_at ASC;


# performace index
-- 2.1 Index สำหรับรายการแข่งตามวัน (ใช้บ่อยที่สุด)
CREATE INDEX idx_sportsbooks_match_date
ON sportsbooks (match_date);

-- 2.2 Composite index สำหรับ league + date (แนะนำ)
CREATE INDEX idx_sportsbooks_league_date
ON sportsbooks (league, match_date);

-- 2.3 Composite index สำหรับ country + date
CREATE INDEX idx_sportsbooks_country_date
ON sportsbooks (country, match_date);

--  2.4 Index สำหรับ fx_id (ถ้ามาจาก feed ภายนอก)
CREATE UNIQUE INDEX idx_sportsbooks_fx_id_match_date
ON sportsbooks (fx_id, match_date);

```
