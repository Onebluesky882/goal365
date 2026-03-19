import { TableHead, TableHeader, TableRow } from "../ui/table";

export const TableHeaderSection = () => {
  return (
    <TableHeader>
      {/* Main header row */}
      <TableRow className="hover:bg-gray-800/20">
        <TableHead
          rowSpan={2}
          className="text-center text-sm font-bold border-r border-gray-700 p-2 uppercase text-gray-300"
        >
          Time
        </TableHead>

        <TableHead
          rowSpan={2}
          className="text-left text-sm font-bold border-r border-gray-700 p-2 uppercase text-gray-300"
        >
          Match
        </TableHead>

        <TableHead
          colSpan={6}
          className="text-center text-sm font-bold border-r border-gray-700 p-2 uppercase text-gray-300"
        >
          Full Time
        </TableHead>

        <TableHead
          colSpan={6}
          className="text-center text-sm font-bold border-r border-gray-700 p-2 uppercase text-gray-300"
        >
          Half Time
        </TableHead>
      </TableRow>

      {/* Sub-header row */}
      <TableRow className="hover:bg-gray-800/10">
        {/* Full Time */}
        <TableHead className="text-center p-2 border-r border-gray-700">
          FT 1
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          FT 2
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          FT 3
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          FT 4
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          FT 5
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          FT 6
        </TableHead>

        {/* Half Time */}
        <TableHead className="text-center p-2 border-r border-gray-700">
          HT 1
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          HT 2
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          HT 3
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          HT 4
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          HT 5
        </TableHead>
        <TableHead className="text-center p-2 border-r border-gray-700">
          HT 6
        </TableHead>
      </TableRow>
    </TableHeader>
  );
};
