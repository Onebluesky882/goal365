import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { db } from "./db";
import { Hono } from "hono";
import { schema } from "./db/schema";
import { cors } from "hono/cors";

const app = new Hono();

app.use(
  "/api/*",
  cors({
    origin: `${process.env.FRONT_END}`,
    allowMethods: ["GET", "POST", "OPTIONS"],
    allowHeaders: ["Content-Type", "Authorization"],
    credentials: true,
  })
);

export const auth = betterAuth({
  database: drizzleAdapter(db, {
    provider: "pg",
    schema,
  }),
  trustedOrigins: [
    "http://localhost:3001",
    "https://goal365-production.up.railway.app",
    "https://goal365.club",
  ],
  emailAndPassword: {
    enabled: true,
  },
  socialProviders: {
    line: {
      clientId: process.env.LINE_CLIENT_ID!,
      clientSecret: process.env.LINE_CLIENT_SECRET!,
    },
  },
});
// mount better-auth handler
app.all("/api/auth/*", async (c) => {
  return auth.handler(c.req.raw);
});

app.get("/", (c) => c.text("API OK"));
const port = Number(process.env.PORT ?? 3000);

app.post("/api/auth/sign-in", async (c) => {
  const res = await auth.handler(c.req.raw);

  // login success
  if (res.status === 200) {
    const session = await auth.api.getSession({
      headers: c.req.raw.headers,
    });

    const playerId = session?.user?.id;
    if (playerId) {
      await logLoginToGo(c, playerId);
    }
    console.log("Player logged in, playerId:", playerId);
  }

  return res;
});
console.log(`🚀 Server running on http://localhost:${port}`);

export default {
  port,
  fetch: app.fetch,
};

async function logLoginToGo(c: any, playerId: string) {
  const ip =
    c.req.header("cf-connecting-ip") ??
    c.req.header("x-forwarded-for") ??
    "unknown";

  const ua = c.req.header("user-agent") ?? "";

  await fetch(process.env.GO_API_URL + "/internal/player/login-log", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Internal-Secret": process.env.INTERNAL_SECRET!,
      "X-Real-IP": ip,
      "User-Agent": ua,
    },
    body: JSON.stringify({ player_id: playerId }),
  });
}
