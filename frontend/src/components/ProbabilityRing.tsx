
import React from 'react';

interface ProbabilityRingProps {
  home: string;
  draw: string;
  away: string;
  homeName: string;
  awayName: string;
}

const ProbabilityRing: React.FC<ProbabilityRingProps> = ({ home, draw, away, homeName, awayName }) => {
  return (
    <div className="flex flex-col items-center justify-center p-6 bg-white rounded-xl shadow-sm border border-slate-100">
      <h3 className="text-lg font-bold mb-6 text-slate-700">Win Probability</h3>
      <div className="flex w-full h-8 rounded-full overflow-hidden mb-6">
        <div 
          className="bg-indigo-500 flex items-center justify-center text-white text-xs font-bold transition-all duration-1000" 
          style={{ width: home }}
        >
          {home}
        </div>
        <div 
          className="bg-slate-300 flex items-center justify-center text-slate-600 text-xs font-bold transition-all duration-1000" 
          style={{ width: draw }}
        >
          {draw}
        </div>
        <div 
          className="bg-pink-500 flex items-center justify-center text-white text-xs font-bold transition-all duration-1000" 
          style={{ width: away }}
        >
          {away}
        </div>
      </div>
      <div className="flex justify-between w-full text-sm">
        <div className="flex items-center gap-2">
          <div className="w-3 h-3 bg-indigo-500 rounded-full" />
          <span className="text-slate-600">{homeName}</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-3 h-3 bg-slate-300 rounded-full" />
          <span className="text-slate-600">Draw</span>
        </div>
        <div className="flex items-center gap-2">
          <div className="w-3 h-3 bg-pink-500 rounded-full" />
          <span className="text-slate-600">{awayName}</span>
        </div>
      </div>
    </div>
  );
};

export default ProbabilityRing;
