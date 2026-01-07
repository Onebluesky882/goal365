"use client";
import React from "react";
import {
  CheckCircle,
  Target,
  TrendingUp,
  Users,
  Shield,
  Trophy,
  BarChart3,
  Star,
} from "lucide-react";

export default function Homepage() {
  const experts = [
    {
      id: 1,
      name: "เซียนเจ๊ก",
      avatar: "เจ",
      specialty: "พรีเมียร์ลีก",
      since: "ตั้งแต่ 2022",
      winRate: "84.3%",
      roi: "12.8%",
      avgMonthlyProfit: "฿8,450",
      joinPrice: "฿299",
      description: "เชี่ยวชาญบอลอังกฤษ วิเคราะห์ลึกทุกแมตช์",
      verified: true,
    },
    {
      id: 2,
      name: "กูรูบอลไทย",
      avatar: "ก",
      specialty: "ไทยลีก",
      since: "ตั้งแต่ 2021",
      winRate: "81.7%",
      roi: "10.5%",
      avgMonthlyProfit: "฿6,230",
      joinPrice: "฿199",
      description: "ติดตามบอลไทยมากว่า 15 ปี",
      verified: true,
    },
    {
      id: 3,
      name: "ลาลีกาคิง",
      avatar: "ล",
      specialty: "ลาลีกา",
      since: "ตั้งแต่ 2020",
      winRate: "86.1%",
      roi: "14.2%",
      avgMonthlyProfit: "฿9,870",
      joinPrice: "฿399",
      description: "เซียนบอลสเปน รู้ลึกทุกทีม",
      featured: true,
      verified: true,
    },
    {
      id: 4,
      name: "อาเซียนเอ็กซ์เปิร์ต ",
      avatar: "อาเซียนเอ็กซ์เปิร์ต",
      specialty: "ฟุตบอลเอเชีย",
      since: "ตั้งแต่ 2019",
      winRate: "79.5%",
      roi: "9.3%",
      avgMonthlyProfit: "฿5,640",
      joinPrice: "฿249",
      description: "เชี่ยวชาญบอลเอเชีย ข้อมูลครบถ้วน",
      verified: true,
    },
    {
      id: 5,
      name: "ยูโรมาสเตอร์",
      avatar: "ย",
      specialty: "บุนเดสลีกา",
      since: "ตั้งแต่ 2018",
      winRate: "82.8%",
      roi: "11.7%",
      avgMonthlyProfit: "฿7,290",
      joinPrice: "฿349",
      description: "กูรูบอลเยอรมัน วิเคราะห์เชิงลึก",
      verified: true,
    },
    {
      id: 6,
      name: "เซเรียเอเลเจนด์",
      avatar: "ซ",
      specialty: "เซเรียอา",
      since: "ตั้งแต่ 2017",
      winRate: "85.4%",
      roi: "13.1%",
      avgMonthlyProfit: "฿8,920",
      joinPrice: "฿379",
      description: "ติดตามบอลอิตาลีมากว่า 20 ปี",
      verified: true,
    },
  ];

  const features = [
    {
      icon: Target,
      title: "เซียนบอลตัวจริง",
      text: "วิเคราะห์โดยกูรูบอลมืออาชีพ ไม่ใช่บอทหรือโปรแกรม",
    },
    {
      icon: TrendingUp,
      title: "วินเรทสูงกว่า 80%",
      text: "สถิติชนะจริง มีหลักฐานย้อนหลังให้ตรวจสอบ",
    },
    {
      icon: Trophy,
      title: "ทีเด็ดบอลรายวัน",
      text: "แนวทางบอลครบทุกลีกดัง อัปเดตทุกวัน",
    },
    {
      icon: Shield,
      title: "โปร่งใส 100%",
      text: "ไม่มีการปั่นราคา หรือขายฝัน เปิดเผยสถิติจริง",
    },
    {
      icon: Users,
      title: "ชุมชนนักเดิมพัน",
      text: "แลกเปลี่ยนประสบการณ์กับสมาชิกหลักพัน",
    },
    {
      icon: BarChart3,
      title: "รายงานผลงานจริง",
      text: "ดูสถิติย้อนหลังของเซียนแต่ละคนได้ฟรี",
    },
  ];

  const steps = [
    {
      step: "1",
      title: "สมัครสมาชิก",
      description: "สร้างบัญชีกับ GOAL365 ง่าย ๆ ใน 2 นาที",
    },
    {
      step: "2",
      title: "เลือกเซียนบอล",
      description: "เลือกติดตามเซียนที่เหมาะกับสไตล์การเดิมพันของคุณ",
    },
    {
      step: "3",
      title: "รับทีเด็ดทุกวัน",
      description: "รับแนวทางผ่านเว็บไซต์หรือไลน์ออฟฟิเชียล",
    },
  ];

  return (
    <>
      {/* SEO Head */}
      <div className="sr-only">
        <h1>GOAL365 ทีเด็ดบอลวันนี้ ทีเด็ดล้มโต๊ะ</h1>
        <meta
          name="description"
          content="GOAL365 ทีเด็ดบอลวันนี้ ผลบอลสด ราคาบอลวันนี้ เว็บรวมเซียนบอล ทีเด็ดบอลล้มโต๊ะ แทงบอลฟรี 1000"
        />
      </div>

      <div className="min-h-screen bg-gradient-to-br from-slate-900 via-blue-900 to-slate-900">
        {/* Hero Section */}
        <div className="relative overflow-hidden">
          <div className="absolute inset-0 bg-gradient-to-br from-yellow-500/10 to-orange-500/10"></div>
          <div className="relative max-w-7xl mx-auto px-4 py-20 sm:px-6 lg:px-8">
            <div className="text-center">
              <div className="flex justify-center mb-6">
                <div className="bg-gradient-to-r from-yellow-400 to-orange-500 text-black px-6 py-2 rounded-full font-bold text-lg animate-pulse">
                  🏆 GOAL365 - เซียนบอลอันดับ 1
                </div>
              </div>

              <h2 className="text-4xl md:text-7xl font-bold text-white mb-8 leading-tight">
                เดินตามเซียนบอล
                <br />
                <span className="text-transparent bg-clip-text bg-gradient-to-r from-yellow-400 to-orange-500">
                  วินเรทสูงกว่า 80%
                </span>
              </h2>

              <p className="text-xl md:text-2xl text-gray-300 mb-12 max-w-4xl mx-auto">
                วิเคราะห์แม่นทุกแมตช์ โดยเซียนตัวจริง ไม่ใช่บอท
                <br />
                <strong className="text-yellow-400">
                  รับทีเด็ดบอลวันนี้
                </strong>{" "}
                ครอบคลุมทุกลีกดัง
              </p>

              <div className="flex flex-col sm:flex-row gap-4 justify-center mb-16">
                <button className="bg-gradient-to-r from-yellow-400 to-orange-500 text-black px-8 py-4 rounded-lg font-bold text-lg hover:scale-105 transition-all duration-300 shadow-2xl">
                  เริ่มเดินตามเซียนวันนี้
                </button>
                <button className="border-2 border-yellow-400 text-yellow-400 px-8 py-4 rounded-lg font-bold text-lg hover:bg-yellow-400 hover:text-black transition-all duration-300">
                  ดูสถิติย้อนหลังฟรี
                </button>
              </div>

              {/* Features Grid */}
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-6xl mx-auto">
                {features.map((feature, index) => (
                  <div
                    key={index}
                    className="bg-white/5 backdrop-blur-lg rounded-2xl p-6 border border-white/10 hover:border-yellow-400/50 transition-all duration-300 hover:scale-105"
                  >
                    <feature.icon className="w-10 h-10 text-yellow-400 mb-4 mx-auto" />
                    <h3 className="text-white font-bold text-lg mb-2">
                      {feature.title}
                    </h3>
                    <p className="text-gray-300 text-sm">{feature.text}</p>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>

        {/* Why Choose Section */}
        <div className="max-w-7xl mx-auto px-4 py-16 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl md:text-5xl font-bold text-white mb-6">
              ทำไมต้องเดินตาม<span className="text-yellow-400">เซียนบอล</span>
              กับ GOAL365?
            </h2>
            <p className="text-xl text-gray-300 max-w-3xl mx-auto">
              หากคุณกำลังมองหาแนวทางการแทงบอลที่แม่นยำ
              เดินตามเซียนบอลที่วิเคราะห์เกมอย่างมีระบบ —{" "}
              <strong className="text-yellow-400">GOAL365 คือคำตอบ</strong>
            </p>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div className="space-y-8">
              <div className="flex items-start space-x-4">
                <div className="bg-green-500 rounded-full p-2 flex-shrink-0">
                  <CheckCircle className="w-6 h-6 text-white" />
                </div>
                <div>
                  <h3 className="text-white font-bold text-xl mb-2">
                    วิเคราะห์โดยเซียนจริง ไม่ใช่บอท
                  </h3>
                  <p className="text-gray-300">
                    ทีมงานของเราคัดเฉพาะกูรูบอลตัวจริง
                    วิเคราะห์ด้วยข้อมูลลึกทั้งฟอร์มทีม ตัวผู้เล่น และสภาพสนาม
                  </p>
                </div>
              </div>

              <div className="flex items-start space-x-4">
                <div className="bg-yellow-500 rounded-full p-2 flex-shrink-0">
                  <TrendingUp className="w-6 h-6 text-black" />
                </div>
                <div>
                  <h3 className="text-white font-bold text-xl mb-2">
                    วินเรทสูงกว่า 80%
                  </h3>
                  <p className="text-gray-300">
                    จากสถิติย้อนหลัง เซียนบอลของ GOAL365
                    มีอัตราชนะสูงและสม่ำเสมอ
                  </p>
                </div>
              </div>

              <div className="flex items-start space-x-4">
                <div className="bg-blue-500 rounded-full p-2 flex-shrink-0">
                  <Trophy className="w-6 h-6 text-white" />
                </div>
                <div>
                  <h3 className="text-white font-bold text-xl mb-2">
                    แจกทีเด็ดบอลรายวัน
                  </h3>
                  <p className="text-gray-300">
                    รับแนวทางบอลประจำวัน ครอบคลุมทุกลีกดัง เช่น พรีเมียร์ลีก,
                    ลาลีกา, เซเรียอา และไทยลีก
                  </p>
                </div>
              </div>
            </div>

            <div className="bg-gradient-to-br from-yellow-400/10 to-orange-500/10 rounded-3xl p-8 border border-yellow-400/20">
              <h3 className="text-2xl font-bold text-white mb-6 text-center">
                เหมาะกับใคร?
              </h3>
              <div className="space-y-4">
                <div className="flex items-center space-x-3">
                  <Star className="w-5 h-5 text-yellow-400" />
                  <span className="text-gray-300">
                    นักเดิมพันที่อยากเพิ่มโอกาสชนะ
                  </span>
                </div>
                <div className="flex items-center space-x-3">
                  <Star className="w-5 h-5 text-yellow-400" />
                  <span className="text-gray-300">
                    มือใหม่ที่ยังวิเคราะห์ไม่เก่ง
                  </span>
                </div>
                <div className="flex items-center space-x-3">
                  <Star className="w-5 h-5 text-yellow-400" />
                  <span className="text-gray-300">
                    คนที่ต้องการแนวทางแบบมีเหตุผล ไม่ใช่แค่ความรู้สึก
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Experts Section */}
        <div className="max-w-7xl mx-auto px-4 py-16 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl md:text-5xl font-bold text-white mb-6">
              เซียนบอล<span className="text-yellow-400">ยอดนิยม</span>
            </h2>
            <p className="text-xl text-gray-300 max-w-3xl mx-auto">
              สร้างบัญชี GOAL365 และเริ่มติดตามเซียนบอลคุณภาพ
              วินเรทสูงที่สุดในประเทศ
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {experts.map((expert) => (
              <div
                key={expert.id}
                className={`bg-white/5 backdrop-blur-lg rounded-2xl p-6 border transition-all duration-300 hover:scale-105 ${
                  expert.featured
                    ? "border-yellow-400/50 shadow-2xl shadow-yellow-400/20"
                    : "border-white/10 hover:border-yellow-400/30"
                }`}
              >
                {expert.featured && (
                  <div className="bg-gradient-to-r from-yellow-400 to-orange-500 text-black px-3 py-1 rounded-full text-sm font-bold mb-4 w-fit">
                    ⭐ แนะนำ
                  </div>
                )}

                {/* Header */}
                <div className="flex items-center justify-between mb-6">
                  <div className="flex items-center space-x-3">
                    <div className="w-12 h-12 bg-gradient-to-r from-yellow-400 to-orange-500 rounded-full flex items-center justify-center text-black font-bold text-lg">
                      {expert.avatar}
                    </div>
                    <div>
                      <h3 className="text-white font-bold text-lg">
                        {expert.name}
                      </h3>
                      <p className="text-gray-400 text-sm">{expert.since}</p>
                      <p className="text-yellow-400 text-sm font-semibold">
                        {expert.specialty}
                      </p>
                    </div>
                  </div>
                  {expert.verified && (
                    <div className="bg-green-500 rounded-full p-1">
                      <CheckCircle className="w-4 h-4 text-white" />
                    </div>
                  )}
                </div>

                {/* Stats */}
                <div className="grid grid-cols-2 gap-4 mb-6">
                  <div className="text-center">
                    <div className="text-gray-400 text-xs uppercase tracking-wide mb-1">
                      วินเรท
                    </div>
                    <div className="text-white text-2xl font-bold">
                      {expert.winRate}
                    </div>
                  </div>
                  <div className="text-center">
                    <div className="text-gray-400 text-xs uppercase tracking-wide mb-1">
                      ROI
                    </div>
                    <div className="text-white text-2xl font-bold">
                      {expert.roi}
                    </div>
                  </div>
                </div>

                {/* Monthly Profit */}
                <div className="text-center mb-6">
                  <div className="text-gray-400 text-xs uppercase tracking-wide mb-1">
                    กำไรเฉลี่ย/เดือน
                  </div>
                  <div className="text-white text-3xl font-bold">
                    {expert.avgMonthlyProfit}
                  </div>
                </div>

                {/* Description */}
                <p className="text-gray-400 text-sm mb-6 text-center">
                  {expert.description}
                </p>

                {/* Join Button */}
                <div className="space-y-3">
                  <button className="w-full bg-gradient-to-r from-yellow-400 to-orange-500 text-black py-3 px-4 rounded-lg font-bold transition-all duration-300 hover:scale-105">
                    เริ่มติดตาม {expert.joinPrice}
                  </button>
                  <div className="text-center">
                    <p className="text-gray-400 text-xs">
                      สมาชิกต่ออายุอัตโนมัติ {expert.joinPrice} ต่อเดือน
                    </p>
                    <p className="text-gray-400 text-xs">• ยกเลิกได้ทุกเมื่อ</p>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* How It Works */}
        <div className="max-w-7xl mx-auto px-4 py-16 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl md:text-5xl font-bold text-white mb-6">
              เริ่มต้นง่าย ๆ แค่{" "}
              <span className="text-yellow-400">3 ขั้นตอน</span>
            </h2>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {steps.map((step, index) => (
              <div key={index} className="text-center">
                <div className="bg-gradient-to-r from-yellow-400 to-orange-500 text-black w-16 h-16 rounded-full flex items-center justify-center text-2xl font-bold mx-auto mb-4">
                  {step.step}
                </div>
                <h3 className="text-white font-bold text-xl mb-2">
                  {step.title}
                </h3>
                <p className="text-gray-300">{step.description}</p>
              </div>
            ))}
          </div>
        </div>

        {/* Trust Section */}
        <div className="max-w-7xl mx-auto px-4 py-16 sm:px-6 lg:px-8">
          <div className="bg-gradient-to-br from-green-500/10 to-blue-500/10 rounded-3xl p-8 border border-green-500/20">
            <div className="text-center">
              <Shield className="w-16 h-16 text-green-400 mx-auto mb-6" />
              <h2 className="text-3xl md:text-4xl font-bold text-white mb-4">
                🔒 เชื่อถือได้ 100% ไม่มีหมกเม็ด
              </h2>
              <p className="text-xl text-gray-300 max-w-3xl mx-auto">
                GOAL365 ไม่มีการปั่นราคา หรือขายฝัน เราให้ข้อมูลที่โปร่งใส
                พร้อมเปิดเผยสถิติชัดเจน
              </p>
            </div>
          </div>
        </div>

        {/* CTA Section */}
        <div className="max-w-7xl mx-auto px-4 py-16 sm:px-6 lg:px-8">
          <div className="text-center">
            <div className="bg-gradient-to-br from-yellow-400/10 to-orange-500/10 rounded-3xl p-12 border border-yellow-400/20">
              <h2 className="text-3xl md:text-5xl font-bold text-white mb-6">
                อยากแทงบอลให้แม่น?
              </h2>
              <p className="text-xl text-gray-300 mb-8 max-w-2xl mx-auto">
                <strong className="text-yellow-400">
                  ไม่ต้องวิเคราะห์เองให้ปวดหัว
                </strong>
                <br />
                เดินตามเซียนบอลของ GOAL365 วินเรทสูงที่สุดในวงการ
              </p>

              <div className="flex flex-col sm:flex-row gap-4 justify-center">
                <button className="bg-gradient-to-r from-yellow-400 to-orange-500 text-black px-10 py-4 rounded-lg font-bold text-xl hover:scale-105 transition-all duration-300 shadow-2xl">
                  เริ่มเดินตามเซียนวันนี้
                </button>
                <button className="border-2 border-yellow-400 text-yellow-400 px-10 py-4 rounded-lg font-bold text-xl hover:bg-yellow-400 hover:text-black transition-all duration-300">
                  ดูสถิติฟรี
                </button>
              </div>

              <div className="mt-8 text-sm text-gray-400">
                <p>⚡ สมัครง่าย ใช้เวลาแค่ 2 นาที</p>
                <p>🏆 รับทีเด็ดบอลทันที</p>
                <p>📱 ใช้งานได้ทุกอุปกรณ์</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
