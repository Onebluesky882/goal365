"use client";

import React from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { authClient } from "@/lib/auth-client";
import Link from "next/link";
import { useRouter } from "next/navigation";

const ProfileMenu = ({ name }: { name: string }) => {
  const router = useRouter();
  const handleSignOut = () => {
    authClient.signOut();
    router.push("/");
  };
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <div>{name}</div>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="" align="start">
        <DropdownMenuGroup>
          <DropdownMenuItem>
            <Link href={"/players"}>DashBoard</Link>
          </DropdownMenuItem>
          <DropdownMenuItem>Profile</DropdownMenuItem>
          <DropdownMenuItem>History</DropdownMenuItem>
          <DropdownMenuItem>Settings</DropdownMenuItem>
        </DropdownMenuGroup>

        <DropdownMenuItem>Support</DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={handleSignOut}>Log out</DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default ProfileMenu;
