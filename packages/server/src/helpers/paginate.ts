import { NextFunction } from 'express';
import { FindManyOptions } from 'typeorm';

const paginate = async <T>(
  next: NextFunction,
  entity: T,
  paginationOptions?: {
    queryOptions?: FindManyOptions<T>;
    page?: number;
    limit?: number;
  }
): Promise<{ count: number; data: T[] }> => {
  return { count: 40, data: [] };
};

export default paginate;
