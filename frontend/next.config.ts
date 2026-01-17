import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  images: {
    domains: ["media.api-sports.io"],
  },
  async rewrites() {
    return [
      {
        source: "/api/auth/:path*",
        destination:
          "https://goal365-production.up.railway.app/api/auth/:path*",
      },
    ];
  },
};

export default nextConfig;
