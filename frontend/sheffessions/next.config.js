/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  env: {
    ENV: process.env.ENV,
    NEXT_PUBLIC_ENV: process.env.NEXT_PUBLIC_ENV
  }
}

module.exports = nextConfig
