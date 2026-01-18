import type { NextConfig } from "next";

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
    return [
      {
        source: "/api/auth/:path*",
        destination:
          "https://goal365-production.up.railway.app/api/auth/:path*",
      },
      // {
      //   source: "/api/:path*",
      //   destination: "https://mytipster-production.up.railway.app/api/:path*",
      // },
      {
        source: "/api/:path*",
        destination: "http://localhost:3009/api/:path*",
      },
    ];
  },
};

export default nextConfig;
