import { ErrorRequestHandler } from 'express';
import ApiError from './ApiError';

const errorHandler: ErrorRequestHandler = (err, req, res, next) => {
  if (err instanceof ApiError) {
    return res.status(err.status).json({ message: err.message });
  }
  res.status(500).json({ message: 'Internal Server Error' });
};

export default errorHandler;
