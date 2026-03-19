import { MdExpandCircleDown } from "react-icons/md";
import { TableBody, TableCell, TableRow } from "../ui/table";

export const TableBodySection = () => {
  return (
    <TableBody>
      <TableRow className="hover:bg-gray-700/10">
        <TableCell className="text-center border-r border-gray-700 p-2">
          12:00
        </TableCell>

        <TableCell className="border-r border-gray-700 p-0">
          <div className="flex items-center justify-between px-2 py-2 cursor-pointer hover:bg-gray-800/40 transition group">
            {/* Match name */}
            <span className="text-[12px] text-white truncate group-hover:text-blue-400">
              Team A vs Team B
            </span>

            {/* Icon */}
            <MdExpandCircleDown className="text-blue-700 transition-transform duration-200 group-hover:text-blue-400 group-hover:rotate-180" />
          </div>
        </TableCell>

        {/* Full Time */}
        <TableCell className="text-center border-r border-gray-700 p-2">
          1-0
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-2">
          2-0
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-2">
          1-1
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-2">
          0-0
        </TableCell>

        {/* Half Time */}
        <TableCell className="text-center border-r border-gray-700 p-2">
          0-0
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-2">
          1-0
        </TableCell>
        <TableCell className="text-center border-r border-gray-700 p-2">
          0-1
        </TableCell>
        <TableCell className="text-center p-2">1-1</TableCell>
      </TableRow>
    </TableBody>
  );
};
