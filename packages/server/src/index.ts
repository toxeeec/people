import { createConnection } from 'typeorm';
import app from './server';

const { PORT, NODE_ENV } = process.env;

createConnection(NODE_ENV).then(() => {
  app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
  });
});
