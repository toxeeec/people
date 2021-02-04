import dotenv from 'dotenv';
import { NextFunction, Request, Response } from 'express';
import faker from 'faker';
import ApiError from '../../dist/helpers/ApiError';
import validateId from '../../dist/middlewares/validateId';

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
    const uuid = faker.random.uuid();
    mockRequest.params!.id = uuid;
    await validateId(
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
    await validateId(
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
    await validateId(
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
