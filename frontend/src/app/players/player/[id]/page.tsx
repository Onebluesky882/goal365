import { authClient } from "@/lib/auth-client";
import React from "react";

const page = () => {
  const { data } = authClient.useSession();
  // get player if no player can create
  return <div>{data?.user.email}</div>;
};

export default page;
