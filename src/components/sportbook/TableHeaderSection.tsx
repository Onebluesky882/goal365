import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { TableHead, TableHeader, TableRow } from "../ui/table";
import { Display } from "next/dist/compiled/@next/font";

type TableHeaderSectionProps = {
  setScrolled: Dispatch<SetStateAction<boolean>>;
  scrolled: Boolean;
};
export const TableHeaderSection = ({
  scrolled,
  setScrolled,
}: TableHeaderSectionProps) => {
  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 50); // 👉 ปรับ threshold ได้
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);
  console.log("scrolled :", scrolled);
  return (
    <TableHeader
      className={`sticky top-0 z-20 ${
        scrolled ? "bg-gray-900 shadow-md" : "bg-transparent"
      }`}
    >
      {/* Main header row */}
      <TableRow>
        <TableHead
          rowSpan={2}
          className="text-center text-sm font-bold border-r border-gray-700      "
        >
          Time
        </TableHead>

        <TableHead
          rowSpan={2}
          className="  text-sm font-bold border-r border-gray-700 text-center  "
        >
          Match
        </TableHead>

        <TableHead
          colSpan={6}
          className="text-center text-sm font-bold border-r border-gray-700   "
        >
          Full Time
        </TableHead>

        <TableHead
          colSpan={6}
          className="text-center text-sm font-bold border-r border-gray-700  0"
        >
          First Half
        </TableHead>
      </TableRow>

      {/* Sub-header row */}
      <TableRow className="hover:bg-gray-800/10">
        {/* Full Time */}
        <TableHead className="text-center  border border-gray-800">
          HDP
        </TableHead>
        <TableHead className="text-center    border border-gray-800">
          Home
        </TableHead>
        <TableHead className="text-center   border border-gray-800">
          Away
        </TableHead>
        <TableHead className="text-center  border border-gray-800">
          O/U
        </TableHead>
        <TableHead className="text-center  border border-gray-800">
          Over
        </TableHead>
        <TableHead className="text-center  border border-gray-800">
          Under
        </TableHead>

        {/* Half Time */}
        <TableHead className="text-center  border border-gray-800">
          HDP
        </TableHead>
        <TableHead className="text-center    border border-gray-800">
          Home
        </TableHead>
        <TableHead className="text-center    border border-gray-800">
          Away
        </TableHead>

        <TableHead className="text-center  border border-gray-800">
          O/U
        </TableHead>
        <TableHead className="text-center  border border-gray-800">
          Over
        </TableHead>
        <TableHead className="text-center  border border-gray-800">
          Under
        </TableHead>
      </TableRow>
    </TableHeader>
  );
};
