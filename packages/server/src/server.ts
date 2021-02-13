import dotenv from 'dotenv';
import express from 'express';
import ApiError from './helpers/ApiError';
import errorHandler from './helpers/errorHandler';
import authRouter from './routes/auth';
import requestsRouter from './routes/requests';
import tokenRouter from './routes/token';
import usersRouter from './routes/users';

const { NODE_ENV } = process.env;

dotenv.config({ path: `.env.${NODE_ENV}` });

const app = express();

app.use(express.json());
const api = express.Router();
app.use('/api', api);

api.use('/', authRouter);
api.use('/', tokenRouter);
api.use('/users', usersRouter);
api.use('/requests', requestsRouter);

app.use((req, res, next) => {
  next(ApiError.notFound());
});

app.use(errorHandler);

export default app;
