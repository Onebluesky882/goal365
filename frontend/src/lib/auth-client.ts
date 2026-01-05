import { createAuthClient } from "better-auth/react";

console.log("process.env.API_URL ", process.env.API_URL);

export const authClient = createAuthClient({
  baseURL: `http://${process.env.API_URL}`,
});
