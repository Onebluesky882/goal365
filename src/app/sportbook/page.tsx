import MenuBar from "@/components/sportbook/MenuBar";
import React from "react";
import SportBook from '@/pre-load/sportbook';

const page = () => {
  return (
    <div>
      <MenuBar />
      <SportBook />
    </div>
  );
};

export default page;
