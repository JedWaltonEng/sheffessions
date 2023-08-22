/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  env: {
    ENV: process.env.ENV,
    NEXT_PUBLIC_ENV: process.env.NEXT_PUBLIC_ENV
    NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL
  }
}

module.exports = nextConfig
