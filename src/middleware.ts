import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export async function middleware(req: NextRequest) {
  // อ่าน session token จาก Better Auth
  const sessionToken = req.cookies.get("better-auth.session_token");

  const protectedRoutes = ["/nawin"];

  const isProtectedRoute = protectedRoutes.some((path) =>
    req.nextUrl.pathname.startsWith(path),
  );

  if (isProtectedRoute && !sessionToken) {
    return NextResponse.redirect(new URL("/", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/nawin/:path*"],
};
