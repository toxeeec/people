import dotenv from 'dotenv';
import express from 'express';
import { createConnection } from 'typeorm';
import authRouter from './routes/auth';

const { NODE_ENV } = process.env;

dotenv.config({ path: `.env.${NODE_ENV}` });

const app = express();

createConnection();

app.use(express.json());
const api = express.Router();
app.use('/api', api);

api.use('/', authRouter);

const { PORT } = process.env;

if (NODE_ENV === 'prod' || NODE_ENV === 'dev') {
  app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
  });
}

export default app;
