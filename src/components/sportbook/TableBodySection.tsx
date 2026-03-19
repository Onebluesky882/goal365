import { MdExpandCircleDown } from "react-icons/md";
import { TableBody, TableCell, TableRow } from "../ui/table";
import {
  AsianHandicap,
  AsianHandicapFh,
  OverUnder,
  OverUnderFh,
} from "@/types/sportbook";

type SportBookBodyProps = {
  country: string;
  leagueName: string;
  time: string;
  Home: string;
  Away: string;
  asianHandicap: AsianHandicap[];
  overUnderFullIime: OverUnder[];
  firstHapdicap: AsianHandicapFh[];
  overUnderFistHaft: OverUnderFh[];
};

export const TableBodySection = ({
  country,
  leagueName,
  time,
  Home,
  Away,
  asianHandicap,
  overUnderFullIime,
  firstHapdicap,
  overUnderFistHaft,
}: SportBookBodyProps) => {
  const rowSpan = asianHandicap.length;

  // ✅ helper class
  const cellClass = (val: any, color = "text-gray-400") =>
    `text-center ${color} ${val !== undefined ? "border" : "border-none"}`;

  // ✅ reusable cell
  const Cell = ({ val, color }: { val: any; color?: string }) => (
    <TableCell className={["cursor-pointer", cellClass(val, color)].join(" ")}>
      {val ?? null}
    </TableCell>
  );

  return (
    <TableBody>
      {/* League */}
      <TableRow className="bg-amber-300 border">
        <TableCell
          colSpan={14}
          className="text-left font-semibold text-black hover:bg-amber-100"
        >
          {country} : {leagueName}
        </TableCell>
      </TableRow>

      {asianHandicap.map((item, index) => {
        const homeFavorite = item.favorite === "Home";
        const awayFavorite = item.favorite === "Away";
        return (
          <TableRow key={index} className="hover:bg-gray-700/10">
            {/* Time */}
            {index === 0 && (
              <TableCell
                rowSpan={rowSpan}
                className="text-center align-top py-4"
              >
                {time}
              </TableCell>
            )}

            {/* Match */}
            {index === 0 && (
              <TableCell
                rowSpan={rowSpan}
                className="align-top border cursor-pointer"
              >
                <div className="flex items-center justify-between py-2 gap-3 group">
                  <span className={`${homeFavorite && "text-blue-500"}`}>
                    {Home}
                  </span>
                  vs
                  <span className={`${awayFavorite && "text-blue-500"}`}>
                    {Away}
                  </span>
                  <MdExpandCircleDown className="text-blue-700 group-hover:rotate-180 transition" />
                </div>
              </TableCell>
            )}

            {/* ===== FULL TIME ===== */}
            <Cell val={item.line} color="text-orange-400" />
            <Cell val={item.home_odd} />
            <Cell val={item.away_odd} />

            <Cell
              val={overUnderFullIime[index]?.value}
              color="text-orange-400"
            />
            <Cell val={overUnderFullIime[index]?.over} />
            <Cell val={overUnderFullIime[index]?.under} />

            {/* ===== HALF TIME ===== */}
            <Cell val={firstHapdicap[index]?.line} color="text-orange-400" />
            <Cell val={firstHapdicap[index]?.home_odd} />
            <Cell val={firstHapdicap[index]?.away_odd} />

            <Cell
              val={overUnderFistHaft[index]?.value}
              color="text-orange-400"
            />
            <Cell val={overUnderFistHaft[index]?.over} />
            <Cell val={overUnderFistHaft[index]?.under} />
          </TableRow>
        );
      })}
    </TableBody>
  );
};
