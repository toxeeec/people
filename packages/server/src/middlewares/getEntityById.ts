import { NextFunction, Request, Response } from 'express';
import { EntityTarget, getRepository } from 'typeorm';
import ApiError from '../helpers/ApiError';

const getEntityById = <T extends EntityTarget<T>>(entity: EntityTarget<T>) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    const id = req.params.id!;
    try {
      const result = await getRepository(entity).findOne(id);
      if (!result) return next(ApiError.notFound());
      req.entity = result;
      return next();
    } catch (err) {
      return next(ApiError.internal());
    }
  };
};

export default getEntityById;
