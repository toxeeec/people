import { NextFunction } from 'express';
import { EntityTarget, FindManyOptions, getRepository } from 'typeorm';
import ApiError from './ApiError';

interface PaginationResult<T> {
  count: number;
  data: T[];
  next?: { page: number; limit: number };
  previous?: { page: number; limit: number };
}
interface PaginateFunction {
  <T extends EntityTarget<T>>(
    next: NextFunction,
    entity: T,
    paginationOptions?: {
      queryOptions?: FindManyOptions<T>;
      page?: number;
      limit?: number;
    }
  ): Promise<PaginationResult<T>>;
}

const paginate: PaginateFunction = async (next, entity, paginationOptions?) => {
  const { queryOptions, page, limit } = paginationOptions ?? {};
  try {
    const paginationResult = {} as PaginationResult<typeof entity>;
    let count: number;
    let data: typeof entity[];
    if (typeof page !== 'undefined' && typeof limit !== 'undefined') {
      const skip = (page - 1) * limit;
      [data, count] = await getRepository(entity).findAndCount({
        ...queryOptions,
        skip,
        take: limit,
      });
      if (page > 1) paginationResult.previous = { page: page - 1, limit };
      if (count > page * limit)
        paginationResult.next = { page: page + 1, limit };
    } else {
      [data, count] = await getRepository(entity).findAndCount(queryOptions);
    }
    paginationResult.count = count;
    paginationResult.data = data;
    return paginationResult;
  } catch {
    next(ApiError.internal());
  }
};

export default paginate;
