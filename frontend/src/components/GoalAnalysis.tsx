
import type { MinuteStats } from '@/types/fixture-analytis';
import React from 'react';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Cell } from 'recharts';

interface GoalAnalysisProps {
  homeMinutes: MinuteStats;
  awayMinutes: MinuteStats;
  homeName: string;
  awayName: string;
}

const GoalAnalysis: React.FC<GoalAnalysisProps> = ({ homeMinutes, awayMinutes, homeName, awayName }) => {
  const intervals = ['0-15', '16-30', '31-45', '46-60', '61-75', '76-90'];
  
  const data = intervals.map(interval => ({
    time: interval,
    home: homeMinutes[interval]?.total || 0,
    away: awayMinutes[interval]?.total || 0,
  }));

  return (
    <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-100">
      <h3 className="text-lg font-bold mb-6 text-slate-700 flex items-center gap-2">
        <svg className="w-5 h-5 text-indigo-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        Scoring Minutes Distribution
      </h3>
      <div className="h-64 w-full">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={data}>
            <CartesianGrid strokeDasharray="3 3" vertical={false} stroke="#f1f5f9" />
            <XAxis dataKey="time" axisLine={false} tickLine={false} tick={{fill: '#64748b', fontSize: 12}} />
            <YAxis axisLine={false} tickLine={false} tick={{fill: '#64748b', fontSize: 12}} />
            <Tooltip 
              cursor={{fill: '#f8fafc'}}
              contentStyle={{borderRadius: '8px', border: 'none', boxShadow: '0 4px 6px -1px rgb(0 0 0 / 0.1)'}}
            />
            <Bar dataKey="home" name={homeName} fill="#6366f1" radius={[4, 4, 0, 0]} />
            <Bar dataKey="away" name={awayName} fill="#ec4899" radius={[4, 4, 0, 0]} />
          </BarChart>
        </ResponsiveContainer>
      </div>
      <div className="flex justify-center gap-6 mt-4">
        <div className="flex items-center gap-2">
          <div className="w-3 h-3 bg-indigo-500 rounded-sm"></div>
          <span className="text-xs text-slate-500 font-medium">{homeName}</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-3 h-3 bg-pink-500 rounded-sm"></div>
          <span className="text-xs text-slate-500 font-medium">{awayName}</span>
        </div>
      </div>
    </div>
  );
};

export default GoalAnalysis;
