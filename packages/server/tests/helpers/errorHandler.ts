import { NextFunction, Request, Response } from 'express';
import ApiError from '../../src/helpers/ApiError';
import errorHandler from '../../src/helpers/errorHandler';

describe('errorHandler helper function', () => {
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  const nextFunction: NextFunction = jest.fn();

  beforeEach(() => {
    mockRequest = {};
    mockResponse = {
      status(code: number) {
        this.statusCode = code;
        return this as Response;
      },
      json: jest.fn(),
    };
  });

  it('should return response with 404 status and message given 404 ApiError', () => {
    errorHandler(
      ApiError.notFound(),
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(mockResponse.statusCode).toEqual(404);
    expect(mockResponse.json).toHaveBeenCalledWith({ message: 'Not Found' });
  });

  it('should return response with 500 status and message given error that is not Api Error', () => {
    errorHandler(
      new Error('random error'),
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(mockResponse.statusCode).toEqual(500);
    expect(mockResponse.json).toHaveBeenCalledWith({
      message: 'Internal Server Error',
    });
  });
});
