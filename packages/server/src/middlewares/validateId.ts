import { NextFunction, Request, Response } from 'express';
import validator from 'validator';
import ApiError from '../helpers/ApiError';

const validateId = (req: Request, res: Response, next: NextFunction) => {
  const { id } = req.params;
  if (!validator.isUUID(id, 4))
    return next(new ApiError(400, 'Id is not in valid format'));
  next();
};
export default validateId;
