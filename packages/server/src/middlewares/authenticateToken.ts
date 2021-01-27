import { NextFunction, Request, Response } from 'express';
import jwt from 'jsonwebtoken';
import ApiError from '../helpers/ApiError';

const authenticateToken = (req: Request, res: Response, next: NextFunction) => {
  const authHeader = req.headers.authorization;
  const token = authHeader && authHeader.split(' ')[1];
  if (!token) next(ApiError.unauthorized());

  jwt.verify(token, process.env.ACESS_TOKEN_SECRET, (err, user) => {
    if (err) return next(ApiError.forbidden());
    req.user = user as { id: number };
    next();
  });
};

export default authenticateToken;
