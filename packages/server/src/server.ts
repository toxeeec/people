import dotenv from 'dotenv';
import express from 'express';
import ApiError from './helpers/ApiError.js';
import errorHandler from './helpers/errorHandler.js';
import authRouter from './routes/auth.js';
import usersRouter from './routes/users.js';

const { NODE_ENV } = process.env;

dotenv.config({ path: `.env.${NODE_ENV}` });

const app = express();

app.use(express.json());
const api = express.Router();
app.use('/api', api);

api.use('/', authRouter);
api.use('/users', usersRouter);

app.use((req, res, next) => {
  next(ApiError.notFound());
});

app.use(errorHandler);

export default app;
