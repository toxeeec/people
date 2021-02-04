import { validateIdSchema } from '@people/common';
import { NextFunction, Request, Response } from 'express';
import ApiError from '../helpers/ApiError';

const validateId = async (req: Request, res: Response, next: NextFunction) => {
  try {
    await validateIdSchema.validate(req.params);
    return next();
  } catch (err) {
    return next(new ApiError(400, err.errors[0]));
  }
};
export default validateId;
