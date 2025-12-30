ข้อมูลที่ต้องการบันทึก 
service 
    - fixture
    - odds asia handicap 0.25 / 0.5 / 0.75
    - predictions


----------- *-* -----------
Strategies
    make Strategies condition (สำหรับ คู่ที่ใช่) สร้างกฏเกฏ การฟิวเตอร์ ให้เข้าสูตร
service
    api 
        - filter data follow Strategies
        - fixture by date
        - group country league 

----------- *-* -----------
1 . golang แยกโดยเฉพาะ 
post method   ใส่ fixture Id  เพื่อบันทึก ใน big query