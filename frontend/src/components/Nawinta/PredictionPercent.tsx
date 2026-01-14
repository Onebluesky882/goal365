// components/prediction/PredictionPercent.tsx
type Props = {
  percent: {
    home: string;
    draw: string;
    away: string;
  };
};

export default function PredictionPercent({ percent }: Props) {
  return (
    <div className="grid grid-cols-3 gap-2 text-center">
      <Stat label="Home" value={percent.home} />
      <Stat label="Draw" value={percent.draw} />
      <Stat label="Away" value={percent.away} />
    </div>
  );
}

function Stat({ label, value }: { label: string; value: string }) {
  return (
    <div className="rounded-lg bg-white border p-3">
      <p className="text-xs text-gray-500">{label}</p>
      <p className="font-bold">{value}</p>
    </div>
  );
}
