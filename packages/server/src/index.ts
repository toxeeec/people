import dotenv from 'dotenv';
import express from 'express';
import { createConnection } from 'typeorm';

dotenv.config({ path: `.env.${process.env.NODE_ENV}` });

const app = express();
app.use(express.json());

createConnection();

const { PORT, NODE_ENV } = process.env;

if (NODE_ENV === 'prod' || NODE_ENV === 'dev') {
  app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
  });
}

export default app;
