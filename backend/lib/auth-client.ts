import { createAuthClient } from "better-auth/react";
export const authClient = createAuthClient({
  baseURL: `http://${process.env.API_URL}`,
});

// Enable calling `getCloudflareContext()` in `next dev`.
// See https://opennext.js.org/cloudflare/bindings#local-access-to-bindings.
// import { initOpenNextCloudflareForDev } from "@opennextjs/cloudflare";
// import { sources } from "next/dist/compiled/webpack/webpack";

// initOpenNextCloudflareForDev();
