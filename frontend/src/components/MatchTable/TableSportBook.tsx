// Odds Button Component
export const OddsButton = ({
  market,
  selection,
  odd,
  color = "text-white",
  onClick,
}: any) => (
  <button
    onClick={() => odd !== "-" && onClick(market, selection, odd)}
    className={`w-full h-full font-mono font-bold text-[10px] transition-all active:scale-90 hover:bg-blue-500/20 ${color} ${
      odd === "-" ? "opacity-20 cursor-default" : "cursor-pointer"
    }`}
  >
    {odd}
  </button>
);
