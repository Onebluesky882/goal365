import type { NextConfig } from "next";

console.log("process.env.NEXT_PUBLIC_API_URL", process.env.NEXT_PUBLIC_API_URL);
const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "media.api-sports.io",
        pathname: "/**",
      },
    ],
  },
  async rewrites() {
    const API_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:3009";
    console.log("API_URL", API_URL);
    return [
      {
        source: "/api/auth/:path*",
        destination:
          "https://goal365-production.up.railway.app/api/auth/:path*",
      },
      {
        source: "/api/:path*",
        destination: `${API_URL}/api/:path*`,
      },
    ];
  },
};

export default nextConfig;
