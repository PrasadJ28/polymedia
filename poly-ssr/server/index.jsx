import express from 'express';
import React from 'react';
import { renderToString } from 'react-dom/server';
import App from '../src/App.js'; // adjust extension if needed

const app = express();

app.use((req, res) => {
  const appHtml = renderToString(React.createElement(App));

  const html = `
    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>React 19 SSR with Vite</title>
      </head>
      <body>
        <div id="app">${appHtml}</div>
        <script type="module" src="/src/main.jsx"></script>
      </body>
    </html>
  `;

  res.status(200).send(html);
});

app.listen(51743, () => {
  console.log('✅ Server running on http://localhost:3000');
});
