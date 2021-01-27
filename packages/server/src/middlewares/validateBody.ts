import { NextFunction, Request, Response } from 'express';
import { ObjectSchema } from 'yup';
import ApiError from '../helpers/ApiError';

const validateBody = (schema: ObjectSchema<any>) => async (
  req: Request,
  res: Response,
  next: NextFunction,
) => {
  try {
    const validatedBody = await schema.validate(req.body);
    req.body = validatedBody;
    next();
  } catch (err) {
    next(new ApiError(400, err.errors[0]));
  }
};
export default validateBody;
