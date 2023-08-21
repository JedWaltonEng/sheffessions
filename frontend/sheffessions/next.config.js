/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  env: {
    ENV: process.env.ENV
  }
}

module.exports = nextConfig
