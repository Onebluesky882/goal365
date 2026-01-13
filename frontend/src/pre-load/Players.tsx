"use client";
import { authClient } from "@/lib/auth-client";

const Players = () => {
    const { data } = authClient.useSession();
    
  // get player if no player can create

  if (!data?.session) {
    return "กรุณา login";
  }
    
    
  return <div>{data?.user.id}</div>;
};

export default Players;
