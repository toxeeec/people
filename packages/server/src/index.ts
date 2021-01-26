import { createConnection } from 'typeorm';
import app from './server';

const { PORT } = process.env;

createConnection().then(() => {
  app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
  });
});
