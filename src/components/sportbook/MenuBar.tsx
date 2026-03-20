"use client";
import React from "react";

import { IoIosPaper } from "react-icons/io";
import { HiUserGroup } from "react-icons/hi2";
import { HiHome } from "react-icons/hi2";
import { useRouter } from "next/navigation";
const MenuBar = () => {
  const router = useRouter();
  return (
    <div className=" flex p-2 gap-2  w-full  ">
      <HiHome onClick={() => router.push("/")} />
      <div className="  justify-end  ml-auto flex items-center gap-2">
        <HiUserGroup />
        <IoIosPaper />
      </div>
    </div>
  );
};

export default MenuBar;
