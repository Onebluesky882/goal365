
import React from 'react';

interface TeamFormBadgeProps {
  form: string;
}

const TeamFormBadge: React.FC<TeamFormBadgeProps> = ({ form }) => {
  // ดึงข้อมูล 5 นัดล่าสุด (หรือตามจำนวนที่มีใน string)
  const matches = form.split('');
  
  return (
    <div className="flex gap-1.5">
      {matches.map((char, idx) => {
        let bgColor = 'bg-slate-200';
        let textColor = 'text-slate-600';
        let ringColor = 'ring-slate-300';
        
        if (char === 'W') {
          bgColor = 'bg-green-500';
          textColor = 'text-white';
          ringColor = 'ring-green-300';
        } else if (char === 'L') {
          bgColor = 'bg-red-500';
          textColor = 'text-white';
          ringColor = 'ring-red-300';
        } else if (char === 'D') {
          bgColor = 'bg-amber-500';
          textColor = 'text-white';
          ringColor = 'ring-amber-200';
        }

        return (
          <div 
            key={idx} 
            className={`w-7 h-7 flex items-center justify-center rounded-lg text-[11px] font-black shadow-sm ring-1 ring-inset ${ringColor} ${bgColor} ${textColor} transition-transform hover:scale-110 cursor-default`}
            title={char === 'W' ? 'Win' : char === 'L' ? 'Loss' : 'Draw'}
          >
            {char}
          </div>
        );
      })}
    </div>
  );
};

export default TeamFormBadge;
