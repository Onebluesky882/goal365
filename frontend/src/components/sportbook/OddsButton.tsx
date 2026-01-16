export type OddsButtonProps = {
  market: string;
  selection: string;
  odd: string;
  highlight?: boolean;
  onClick: (m: string, s: string, o: string) => void;
};

export const OddsButton = ({
  market,
  selection,
  odd,
  highlight,
  onClick,
}: OddsButtonProps) => {
  return (
    <button
      onClick={() => odd !== "-" && onClick(market, selection, odd)}
      className={`
      w-full h-full font-mono font-bold text-[10px] transition-all
      ${highlight ? "text-[#00acec]" : "text-foreground"}
      ${
        odd === "-"
          ? "opacity-10 cursor-default"
          : "hover:bg-[#00acec]/20 active:scale-90 active:bg-[#00acec]/30"
      }
    `}
    >
      {odd}
    </button>
  );
};
