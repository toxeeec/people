import { idSchema } from '@people/common';
import dotenv from 'dotenv';
import { NextFunction, Request, Response } from 'express';
import faker from 'faker';
import ApiError from '../../src/helpers/ApiError';
import validateParams from '../../src/middlewares/validateParams';

dotenv.config({ path: '.env.test' });

describe('validateId middleware', () => {
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  const nextFunction: NextFunction = jest.fn();

  beforeEach(() => {
    mockRequest = { params: {} };
    mockResponse = {};
  });

  it('should call next function when given valid id', async (done) => {
    const uuid = faker.datatype.uuid();
    mockRequest.params!.id = uuid;
    await validateParams(idSchema)(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenCalledWith();
    done();
  });

  it('should call next function with 400 error when given invalid id', async (done) => {
    const wrongUuid = 'wrongUuid';
    mockRequest.params!.id = wrongUuid;
    await validateParams(idSchema)(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(
      new ApiError(400, 'id must be a valid UUID')
    );
    done();
  });

  it('should call next function with 400 error when id is not given', async (done) => {
    await validateParams(idSchema)(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(
      new ApiError(400, 'id is a required field')
    );
    done();
  });
});
