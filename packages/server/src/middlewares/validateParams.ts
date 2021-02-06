import { NextFunction, Request, Response } from 'express';
import { ObjectSchema } from 'yup';
import ApiError from '../helpers/ApiError';

const validateParams = (schema: ObjectSchema<any>) => async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    await schema.validate(req.params);
    return next();
  } catch (err) {
    return next(new ApiError(400, err.errors[0]));
  }
};
export default validateParams;
