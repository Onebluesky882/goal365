import React from "react";
import ComparisonChart from "./components/ComparisonChart";
import ProbabilityRing from "./components/ProbabilityRing";
import TeamFormBadge from "./components/TeamFormBadge";
import { MOCK_DATA } from "./data/constants";

const App: React.FC = () => {
  const item = MOCK_DATA.Items[0];
  const { Fixture, Predictions } = item;
  const { predictions, comparison, teams, h2h } = Predictions;

  return (
    <div className="min-h-screen bg-slate-50 pb-12 font-sans text-slate-900">
      {/* Header */}
      <nav className="bg-slate-900 text-white p-4 shadow-lg sticky top-0 z-50">
        <div className="max-w-6xl mx-auto flex justify-between items-center">
          <div className="flex items-center gap-3">
            <div className="bg-white p-1 rounded-lg shadow-inner">
              <img
                src={Fixture.league.logo}
                alt="League"
                className="w-8 h-8 object-contain"
              />
            </div>
            <div>
              <h1 className="font-black text-lg tracking-tighter leading-none">
                FOOTY-ANALYSIS PRO
              </h1>
              <p className="text-[10px] text-slate-400 uppercase font-bold tracking-widest">
                {Fixture.league.name} • {Fixture.league.country}
              </p>
            </div>
          </div>
          <div className="flex items-center gap-4">
            <div className="hidden sm:flex flex-col items-end">
              <span className="text-[10px] text-slate-500 font-bold uppercase tracking-widest">
                Match Status
              </span>
              <span className="text-xs font-black text-green-400">
                {Fixture.fixture.status.long}
              </span>
            </div>
            <img
              src={Fixture.league.flag}
              className="w-6 h-4 rounded shadow-sm"
              alt="flag"
            />
          </div>
        </div>
      </nav>

      <main className="max-w-6xl mx-auto px-4 mt-8">
        {/* Result & Prediction Hero */}
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-6 mb-8">
          {/* Actual Match Result Section */}
          <div className="lg:col-span-8 bg-white rounded-3xl shadow-xl overflow-hidden border border-slate-200">
            <div className="bg-gradient-to-br from-slate-800 via-slate-900 to-black p-8 text-white relative">
              <div className="absolute top-4 left-4 bg-white/10 px-3 py-1 rounded-full text-[10px] font-black tracking-widest uppercase border border-white/10 backdrop-blur-md">
                Final Result Summary
              </div>

              <div className="flex justify-around items-center mt-8">
                <div className="text-center w-1/3 group">
                  <div className="bg-white/5 p-4 rounded-2xl mb-4 group-hover:bg-white/10 transition-colors border border-white/5 inline-block">
                    <img
                      src={Fixture.teams.home.logo}
                      className="w-20 h-20 object-contain drop-shadow-2xl"
                      alt=""
                    />
                  </div>
                  <h2 className="text-xl font-black uppercase tracking-tighter truncate">
                    {Fixture.teams.home.name}
                  </h2>
                  <div className="mt-2 flex justify-center">
                    <TeamFormBadge form={teams.home.league.form} />
                  </div>
                </div>

                <div className="text-center w-1/3 flex flex-col items-center">
                  <div className="text-7xl font-black tracking-tighter flex justify-center items-center gap-3">
                    <span className="text-slate-200">{Fixture.goals.home}</span>
                    <span className="text-4xl text-slate-600 font-light">
                      :
                    </span>
                    <span className="text-green-400">{Fixture.goals.away}</span>
                  </div>
                  <div className="mt-2 bg-slate-800 px-4 py-1 rounded-full border border-slate-700">
                    <span className="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
                      HT Score: 0 - 1
                    </span>
                  </div>
                  <div className="mt-4 text-[10px] font-black text-indigo-400 uppercase tracking-[0.2em]">
                    {Fixture.fixture.venue.name}
                  </div>
                </div>

                <div className="text-center w-1/3 group">
                  <div className="bg-white/5 p-4 rounded-2xl mb-4 group-hover:bg-white/10 transition-colors border border-white/5 inline-block">
                    <img
                      src={Fixture.teams.away.logo}
                      className="w-20 h-20 object-contain drop-shadow-2xl"
                      alt=""
                    />
                  </div>
                  <h2 className="text-xl font-black uppercase tracking-tighter truncate">
                    {Fixture.teams.away.name}
                  </h2>
                  <div className="mt-2 flex justify-center">
                    <TeamFormBadge form={teams.away.league.form} />
                  </div>
                </div>
              </div>
            </div>

            <div className="p-6 bg-slate-50 grid grid-cols-2 md:grid-cols-4 gap-4 text-center">
              <div className="border-r border-slate-200 last:border-0">
                <p className="text-[10px] font-bold text-slate-400 uppercase">
                  Season
                </p>
                <p className="font-black text-slate-700">
                  {Fixture.league.season}
                </p>
              </div>
              <div className="border-r border-slate-200 last:border-0">
                <p className="text-[10px] font-bold text-slate-400 uppercase">
                  League Round
                </p>
                <p className="font-black text-slate-700">Round of 32</p>
              </div>
              <div className="border-r border-slate-200 last:border-0">
                <p className="text-[10px] font-bold text-slate-400 uppercase">
                  City
                </p>
                <p className="font-black text-slate-700">
                  {Fixture.fixture.venue.city}
                </p>
              </div>
              <div className="last:border-0">
                <p className="text-[10px] font-bold text-slate-400 uppercase">
                  Status
                </p>
                <p className="font-black text-green-600 uppercase tracking-tighter">
                  Finished
                </p>
              </div>
            </div>
          </div>

          {/* AI Prediction Sidecar */}
          <div className="lg:col-span-4 bg-indigo-700 rounded-3xl shadow-xl p-8 text-white flex flex-col justify-between border border-indigo-600 relative overflow-hidden group">
            <div className="absolute top-0 right-0 p-8 opacity-5 -mr-8 -mt-8 group-hover:opacity-10 transition-opacity">
              <svg
                className="w-48 h-48"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fillRule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z"
                  clipRule="evenodd"
                />
              </svg>
            </div>

            <div>
              <div className="flex items-center gap-2 mb-4">
                <span className="bg-white/20 text-white text-[10px] font-black uppercase tracking-widest px-3 py-1 rounded-full border border-white/20">
                  AI Insights
                </span>
              </div>
              <h3 className="text-3xl font-black leading-tight tracking-tighter mb-4">
                {predictions.advice}
              </h3>
              <p className="text-indigo-100 text-sm opacity-90 leading-relaxed font-medium">
                Analysis: {predictions.winner.comment}. <br />
                Expected goals: {predictions.goals.home} (Home) /{" "}
                {predictions.goals.away} (Away).
              </p>
            </div>

            <div className="mt-8 space-y-4">
              <div className="flex justify-between items-center text-xs font-bold uppercase tracking-widest text-indigo-300">
                <span>Predicted Over/Under</span>
                <span>
                  {predictions.goals.home} Home / {predictions.goals.away} Away
                </span>
              </div>

              <div className="p-5 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-sm">
                <p className="text-[10px] font-black uppercase tracking-widest text-indigo-300 mb-2">
                  Model Confidence
                </p>
                <div className="flex items-center gap-4">
                  <div className="text-3xl font-black">
                    {predictions.percent.home}
                  </div>
                  <div className="text-xs text-indigo-200 font-medium">
                    Probability score calculated via Poisson distribution and
                    weighted team form analysis.
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Analytics Section */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div className="lg:col-span-2 space-y-8">
            <ComparisonChart
              comparison={comparison}
              homeName={teams.home.name}
              awayName={teams.away.name}
            />

            {/* Head-to-Head Section - SHOWING 5 MATCHES */}
            <div className="bg-white rounded-3xl shadow-sm border border-slate-100 overflow-hidden">
              <div className="p-6 border-b border-slate-100 bg-slate-50/50 flex justify-between items-center">
                <h3 className="font-black text-slate-700 uppercase text-sm tracking-widest flex items-center gap-2">
                  <svg
                    className="w-5 h-5 text-indigo-500"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                    />
                  </svg>
                  Head-to-Head Analysis (Last 5)
                </h3>
                <span className="text-[10px] font-black text-slate-400 uppercase bg-slate-100 px-3 py-1 rounded-full">
                  Historical Data
                </span>
              </div>
              <div className="divide-y divide-slate-100">
                {h2h.map((match: any, idx: number) => (
                  <div
                    key={idx}
                    className="p-4 md:p-6 flex items-center justify-between hover:bg-slate-50/80 transition-all group"
                  >
                    <div className="w-24 hidden md:block">
                      <p className="text-[11px] text-slate-500 font-bold">
                        {new Date(match.fixture.date).toLocaleDateString(
                          "en-GB",
                          { day: "2-digit", month: "short", year: "numeric" }
                        )}
                      </p>
                      <p className="text-[10px] text-slate-300 font-black uppercase tracking-tighter">
                        {match.fixture.venue.city}
                      </p>
                    </div>
                    <div className="flex items-center gap-2 md:gap-6 flex-1 justify-center">
                      <div className="flex items-center gap-2 md:gap-4 justify-end flex-1">
                        <span className="text-xs md:text-sm font-black text-slate-600 group-hover:text-indigo-600 transition-colors truncate">
                          {match.teams.home.name}
                        </span>
                        <img
                          src={match.teams.home.logo}
                          className="w-6 h-6 md:w-8 md:h-8 grayscale group-hover:grayscale-0 transition-all object-contain"
                          alt=""
                        />
                      </div>
                      <div className="flex flex-col items-center">
                        <div className="bg-slate-900 text-white px-3 md:px-5 py-1.5 rounded-xl text-lg md:text-xl font-black tracking-tighter shadow-sm group-hover:bg-indigo-600 transition-all group-hover:scale-110">
                          {match.goals.home} - {match.goals.away}
                        </div>
                      </div>
                      <div className="flex items-center gap-2 md:gap-4 justify-start flex-1">
                        <img
                          src={match.teams.away.logo}
                          className="w-6 h-6 md:w-8 md:h-8 grayscale group-hover:grayscale-0 transition-all object-contain"
                          alt=""
                        />
                        <span className="text-xs md:text-sm font-black text-slate-600 group-hover:text-pink-600 transition-colors truncate">
                          {match.teams.away.name}
                        </span>
                      </div>
                    </div>
                    <div className="w-12 text-right">
                      <span className="text-[10px] font-black text-slate-300 group-hover:text-indigo-500 uppercase">
                        {match.fixture.status.short}
                      </span>
                    </div>
                  </div>
                ))}
              </div>
            </div>

            {/* Team Insights Comparison */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-8 pb-8">
              <div className="bg-white rounded-3xl p-6 shadow-sm border border-slate-100 relative overflow-hidden">
                <div className="absolute top-0 right-0 w-1 h-full bg-indigo-500"></div>
                <div className="flex items-center gap-3 mb-6">
                  <img src={teams.home.logo} className="w-10 h-10" alt="" />
                  <h4 className="font-black text-slate-800 uppercase tracking-tighter">
                    Home: {teams.home.name}
                  </h4>
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="bg-slate-50 p-3 rounded-xl border border-slate-100">
                    <p className="text-[10px] font-bold text-slate-400 uppercase mb-1">
                      Clean Sheets
                    </p>
                    <p className="text-xl font-black text-slate-700">
                      {teams.home.league.clean_sheet.total}
                    </p>
                  </div>
                  <div className="bg-slate-50 p-3 rounded-xl border border-slate-100">
                    <p className="text-[10px] font-bold text-slate-400 uppercase mb-1">
                      Avg Score
                    </p>
                    <p className="text-xl font-black text-indigo-600">
                      {teams.home.last_5.goals.for.average}
                    </p>
                  </div>
                </div>
              </div>

              <div className="bg-white rounded-3xl p-6 shadow-sm border border-slate-100 relative overflow-hidden">
                <div className="absolute top-0 right-0 w-1 h-full bg-pink-500"></div>
                <div className="flex items-center gap-3 mb-6">
                  <img src={teams.away.logo} className="w-10 h-10" alt="" />
                  <h4 className="font-black text-slate-800 uppercase tracking-tighter">
                    Away: {teams.away.name}
                  </h4>
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="bg-slate-50 p-3 rounded-xl border border-slate-100">
                    <p className="text-[10px] font-bold text-slate-400 uppercase mb-1">
                      Clean Sheets
                    </p>
                    <p className="text-xl font-black text-slate-700">
                      {teams.away.league.clean_sheet.total}
                    </p>
                  </div>
                  <div className="bg-slate-50 p-3 rounded-xl border border-slate-100">
                    <p className="text-[10px] font-bold text-slate-400 uppercase mb-1">
                      Avg Score
                    </p>
                    <p className="text-xl font-black text-pink-600">
                      {teams.away.last_5.goals.for.average}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="space-y-8">
            <ProbabilityRing
              home={predictions.percent.home}
              draw={predictions.percent.draw}
              away={predictions.percent.away}
              homeName={teams.home.name}
              awayName={teams.away.name}
            />

            <div className="bg-slate-900 rounded-3xl p-8 text-white shadow-2xl relative overflow-hidden">
              <div className="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-indigo-500 via-pink-500 to-yellow-500"></div>
              <h3 className="font-black text-sm uppercase tracking-widest text-indigo-400 mb-6">
                Deep Data Analytics
              </h3>
              <div className="space-y-4">
                {[
                  {
                    label: "Poisson Expectancy",
                    val: comparison.poisson_distribution.away,
                    team: teams.away.name,
                    color: "text-indigo-400",
                  },
                  {
                    label: "Defensive Strength",
                    val: comparison.def.home,
                    team: teams.home.name,
                    color: "text-green-400",
                  },
                  {
                    label: "Form Stability",
                    val: comparison.form.home,
                    team: teams.home.name,
                    color: "text-yellow-400",
                  },
                  {
                    label: "Total AI Weighted Score",
                    val: comparison.total.home,
                    team: teams.home.name,
                    color: "text-pink-400",
                  },
                ].map((stat, i) => (
                  <div
                    key={i}
                    className="bg-white/5 p-4 rounded-2xl border border-white/5 flex justify-between items-center hover:bg-white/10 transition-colors"
                  >
                    <div>
                      <p className="text-[10px] font-bold text-slate-500 uppercase tracking-tighter mb-1">
                        {stat.label}
                      </p>
                      <p className="text-[9px] text-slate-400 font-medium italic">
                        Advantage: {stat.team}
                      </p>
                    </div>
                    <p className={`text-2xl font-black ${stat.color}`}>
                      {stat.val}
                    </p>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </main>

      <footer className="mt-20 py-16 text-center border-t border-slate-200 bg-white/50 backdrop-blur-sm">
        <div className="max-w-md mx-auto px-4">
          <p className="text-slate-400 text-xs font-black uppercase tracking-[0.3em] mb-4">
            FootyPredict intelligence v2.0
          </p>
        </div>
      </footer>
    </div>
  );
};

export default App;
