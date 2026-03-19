"use client";
import { Dispatch, SetStateAction } from "react";
import { IoSearchCircle } from "react-icons/io5";

type SportbookHeaderCardProps = {
  setSearch: Dispatch<SetStateAction<boolean>>;
  search: boolean;
};

export const TableHeaderCard = ({
  setSearch,
  search,
}: SportbookHeaderCardProps) => {
  return (
    <div className="overflow-x-auto w-full">
      <div className="flex border-t-2 p-2 items-center justify-between bg-gray-900/40">
        {/* Search */}
        <div className="flex items-center gap-2 px-2 py-1 rounded-md bg-gray-800/50">
          <IoSearchCircle
            className="text-lg text-gray-400 cursor-pointer hover:text-blue-400 transition"
            onClick={() => setSearch((prev) => !prev)}
          />

          {search && (
            <input
              className="bg-transparent outline-none text-[12px] text-white placeholder:text-gray-400 placeholder:text-center text-center w-[120px]"
              placeholder="Finding match"
            />
          )}
        </div>

        {/* Balance */}
        <p className="text-[12px] text-gray-300">
          ยอดเงิน : <span className="text-green-400 font-bold">0.00</span>
        </p>
      </div>
    </div>
  );
};
