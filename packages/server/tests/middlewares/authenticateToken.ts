import dotenv from 'dotenv';
import { NextFunction, Request, Response } from 'express';
import faker from 'faker';
import jwt from 'jsonwebtoken';
import ApiError from '../../dist/helpers/ApiError';
import authenticateToken from '../../dist/middlewares/authenticateToken';

dotenv.config({ path: '.env.test' });

describe('authenticateToken middleware', () => {
  const sampleUser = {
    id: faker.random.uuid(),
    name: faker.name.firstName(),
    surname: faker.name.lastName(),
  };
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  const nextFunction: NextFunction = jest.fn();

  beforeEach(() => {
    mockRequest = { headers: {} };
    mockResponse = {};
  });

  it('should call next when valid accessToken is given', () => {
    const accessToken = jwt.sign(sampleUser, process.env.ACCESS_TOKEN_SECRET!, {
      expiresIn: '15m',
    });
    mockRequest.headers!.authorization = `Bearer ${accessToken}`;
    authenticateToken(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenCalledWith();
  });

  it('should call next with 401 ApiError when accessToken is not given', () => {
    authenticateToken(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(ApiError.unauthorized());
  });

  it('should call next with 403 ApiError when wrong accessToken is given', () => {
    const accessToken = 'refreshToken';
    mockRequest.headers!.authorization = `Bearer ${accessToken}`;
    authenticateToken(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(ApiError.forbidden());
  });
});
