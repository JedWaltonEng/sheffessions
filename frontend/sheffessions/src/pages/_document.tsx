import React from 'react';
import Document, { Html, Head, Main, NextScript } from 'next/document';

class MyDocument extends Document {
  render() {
    return (
      <Html>
          <Head>
            <link rel="icon" type="image/png" href="/logo.png" />
            {/* Optionally, add other favicon formats */}
            {/* <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" /> */}
            {/* <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" /> */}
          </Head>
          <body>
            <Main />
            <NextScript />
          </body>
      </Html>
    );
  }
}

export default MyDocument;

