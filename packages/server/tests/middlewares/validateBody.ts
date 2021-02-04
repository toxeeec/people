import { authenticateUserSchema } from '@people/common';
import dotenv from 'dotenv';
import { NextFunction, Request, Response } from 'express';
import faker from 'faker';
import ApiError from '../../dist/helpers/ApiError';
import validateBody from '../../dist/middlewares/validateBody';

dotenv.config({ path: '.env.test' });

describe('validateBody middleware', () => {
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  const nextFunction: NextFunction = jest.fn();

  beforeEach(() => {
    mockRequest = {};
    mockResponse = {};
  });

  it('should call next function when given valid input', async (done) => {
    const middleware = validateBody(authenticateUserSchema);
    const sampleUser = {
      email: faker.internet.email(),
      password: faker.internet.password(),
    };
    mockRequest.body = sampleUser;
    await middleware(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenCalledWith();
    done();
  });

  it('should call next function with 400 ApiError when given input is not valid', async (done) => {
    const middleware = validateBody(authenticateUserSchema);
    const wrongUser = {
      email: 'notAnEmail',
      password: 'wrong',
    };
    mockRequest.body = wrongUser;
    await middleware(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(
      new ApiError(400, 'email must be a valid email')
    );
    done();
  });

  it('should call next function with 400 ApiError when required field is not given', async (done) => {
    const middleware = validateBody(authenticateUserSchema);
    const wrongUser = {
      password: 'notAnEmail',
    };
    mockRequest.body = wrongUser;
    await middleware(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(
      new ApiError(400, 'email is a required field')
    );
    done();
  });
});
